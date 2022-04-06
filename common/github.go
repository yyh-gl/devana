package common

import (
	"context"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	"path"
	"strings"
	"time"
)

const perPage = 100

type (
	GitHubClient struct {
		*github.Client

		Owner        string
		Repository   string
		IsEnterprise bool
	}

	PR struct {
		Name      string
		CreatedAt time.Time
		ClosedAt  time.Time
	}
	PRs []PR
)

func NewGitHubClient(ctx context.Context, url, token string, isEnterprise bool) (*GitHubClient, error) {
	urlParts := strings.Split(url, "/")
	owner, repo := urlParts[len(urlParts)-2], urlParts[len(urlParts)-1]

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	if isEnterprise {
		// TODO: urlパッケージ使ってbaseURL作る
		gitHubBaseURL := "https://" + path.Join(strings.Join(urlParts[2:len(urlParts)-2], ""), "/api/v3")
		c, err := github.NewEnterpriseClient(gitHubBaseURL, "", tc)
		if err != nil {
			return nil, err
		}

		return &GitHubClient{
			Client:       c,
			Owner:        owner,
			Repository:   repo,
			IsEnterprise: true,
		}, nil
	}

	c := github.NewClient(tc)
	return &GitHubClient{
		Client:       c,
		Owner:        owner,
		Repository:   repo,
		IsEnterprise: false,
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
