package common

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

type (
	GitRepository struct {
		*git.Repository
	}

	Log struct {
		Author  string
		Message string
	}
	Logs []Log
)

func NewGitRepository(url string) (*GitRepository, error) {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: url})
	if err != nil {
		return nil, err
	}
	return &GitRepository{Repository: r}, nil
}

func (gr GitRepository) FetchLogs() (Logs, error) {
	ref, err := gr.Head()
	if err != nil {
		return nil, err
	}

	cIter, err := gr.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, err
	}

	var logs Logs
	err = cIter.ForEach(func(c *object.Commit) error {
		logs = append(logs, Log{
			Author:  c.Author.String(),
			Message: c.Message,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return logs, nil
}
