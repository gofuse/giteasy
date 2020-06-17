package git

import (
	"giteasy/internal/model"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func Clone() error {
	_, err := git.PlainClone(model.CurrentProfile.LocalRepo, false, &git.CloneOptions{
		URL: model.CurrentProfile.RemoteRepo,
		Auth: &http.BasicAuth{
			Username: model.CurrentProfile.UserName,
			Password: model.CurrentProfile.Password,
		},
		Progress: os.Stdout,
	})
	return err
}
