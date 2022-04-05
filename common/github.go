package common

import (
	"context"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	"strings"
	"time"
)

const perPage = 100

type (
	GitHubClient struct {
		*github.Client

		Owner      string
		Repository string
	}

	PR struct {
		Name      string
		CreatedAt time.Time
		ClosedAt  time.Time
	}
	PRs []PR
)

func NewGitHubClient(ctx context.Context, url, token string) (*GitHubClient, error) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	c := github.NewClient(tc)

	strs := strings.Split(url, "/")
	owner, repo := strs[len(strs)-2], strs[len(strs)-1]

	return &GitHubClient{
		Client:     c,
		Owner:      owner,
		Repository: repo,
	}, nil
}

func (c GitHubClient) FetchPRs(ctx context.Context, since, until time.Time) (PRs, error) {
	var pullRequests PRs

	for i := 0; ; i++ {
		prs, _, err := c.PullRequests.List(ctx, c.Owner, c.Repository, &github.PullRequestListOptions{
			State:       "closed",
			Sort:        "created",
			Direction:   "desc",
			ListOptions: github.ListOptions{Page: i, PerPage: perPage},
		})
		if err != nil {
			if _, ok := err.(*github.RateLimitError); ok {
				i--
				time.Sleep(60 * time.Second)
			} else {
				return nil, err
			}
		}

		for _, pr := range prs {
			// CreatedAtより後にsinceがある場合、それ以降のPRは集計期間外のPRであることが確定
			if since.After(*pr.CreatedAt) {
				goto Done
			}

			if until.Before(*pr.CreatedAt) {
				continue
			}

			pullRequests = append(pullRequests, PR{
				Name:      *pr.Title,
				CreatedAt: *pr.CreatedAt,
				ClosedAt:  *pr.ClosedAt,
			})
		}

		if len(prs) < perPage {
			break
		}

		time.Sleep(2 * time.Second)
	}

Done:
	return pullRequests, nil
}
