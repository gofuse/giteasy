package git

import (
	"errors"
	"giteasy/internal/logger"
	"giteasy/internal/model"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func Push() error {
	if model.CurrentProfile.UserName == "" || model.CurrentProfile.Password == "" {
		return errors.New("Username/Password not set")
	}
	logger.Debug("Pushing local changes to remote origin...")
	r, err := git.PlainOpen(model.CurrentProfile.LocalRepo)
	if err != nil {
		logger.Error(err)
	}
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: model.CurrentProfile.UserName,
			Password: model.CurrentProfile.Password,
		},
		Progress: os.Stdout,
	})
	if err != nil {
		logger.Error(err)
	}
	logger.Debug("Done")
	Status()
	return err
}
