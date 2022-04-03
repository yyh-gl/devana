package common

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

type (
	GitRepository struct {
		*git.Repository
	}

	Tag struct {
		Name     string
		Datetime time.Time
	}
	Tags []Tag
)

func NewGitRepository(url string, token string) (*GitRepository, error) {
	opt := git.CloneOptions{
		URL:   url,
		Depth: 1,
		// TODO: -vオプション指定時は出力
		//Progress: os.Stdout,
	}
	if token != "" {
		opt.Auth = &http.BasicAuth{
			// 空文字以外であればOK
			Username: "ninja",
			Password: token,
		}
	}

	r, err := git.Clone(memory.NewStorage(), nil, &opt)
	if err != nil {
		return nil, err
	}

	return &GitRepository{Repository: r}, nil
}

func (gr GitRepository) FetchTags(since, until time.Time) (Tags, error) {
	iter, err := gr.Tags()
	if err != nil {
		return nil, err
	}

	var tags Tags
	err = iter.ForEach(func(t *plumbing.Reference) error {
		obj, err := gr.CommitObject(t.Hash())
		if err != nil {
			return err
		}

		when := obj.Author.When
		if since.After(when) || until.Before(when) {
			return nil
		}

		tags = append(tags, Tag{
			Name:     t.Name().String(),
			Datetime: when,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tags, nil
}
