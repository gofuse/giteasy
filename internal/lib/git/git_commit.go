package git

import (
	"giteasy/internal/logger"
	"giteasy/internal/model"
	"os/user"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func Commit(message string) {
	r, err := git.PlainOpen(model.CurrentProfile.LocalRepo)
	w, err := r.Worktree()
	_, err = w.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name: func() string {
				user, _ := user.Current()
				return user.Username
			}(),
			When: time.Now(),
		},
	})
	if err != nil {
		logger.Error(err)
	}
	Status()
}
