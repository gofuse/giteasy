package git

import (
	"fmt"
	"giteasy/internal/logger"
	"giteasy/internal/model"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

func Stage() {
	r, _ := git.PlainOpen(model.CurrentProfile.LocalRepo)
	w, _ := r.Worktree()
	s, _ := w.Status()
	// staged := make(map[string]string)
	for path, status := range s {
		//skip gitli
		if strings.Contains(os.Args[0], filepath.Base(path)) {
			continue
		}
		if status.Staging == git.Renamed {
			path = fmt.Sprintf("%s -> %s", path, status.Extra)
		}

		switch status.Worktree {
		case git.Unmodified:
			continue
		case git.Added:
		case git.Untracked:
			logger.Info(string(status.Worktree), path)
			// staged[path] = string(status.Worktree)
		case git.Modified:
			logger.Warn(string(status.Worktree), path)
			// staged[path] = string(status.Worktree)
		case git.Deleted:
			logger.Error(string(status.Worktree), path)
			// staged[path] = string(status.Worktree)
		default:
			logger.Debug(string(status.Worktree), path)
			// staged[path] = string(status.Worktree)
		}
		w.Add(path)
	}
	// model.Set(staged, constants.STAGED)
	Status()
}
