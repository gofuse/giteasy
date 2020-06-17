package git

import (
	"fmt"
	"giteasy/internal/constants"
	"giteasy/internal/logger"
	"giteasy/internal/model"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func Status() {
	findUnstaged()
	findStaged()
	findCommited()
}

func findUnstaged() {
	logger.Debug("\nUnstaged changes")
	r, _ := git.PlainOpen(model.CurrentProfile.LocalRepo)
	w, _ := r.Worktree()
	s, _ := w.Status()
	unstaged := make(map[string]string)
	for path, status := range s {
		if status.Staging == git.Renamed {
			path = fmt.Sprintf("%s -> %s", path, status.Extra)
		}

		switch status.Worktree {
		case git.Unmodified:
			continue
		case git.Added:
		case git.Untracked:
			logger.Info(string(status.Worktree), path)
			unstaged[path] = string(status.Worktree)
		case git.Modified:
			logger.Warn(string(status.Worktree), path)
			unstaged[path] = string(status.Worktree)
		case git.Deleted:
			logger.Error(string(status.Worktree), path)
			unstaged[path] = string(status.Worktree)
		default:
			logger.Debug(string(status.Worktree), path)
		}
	}
	model.Set(unstaged, constants.UNSTAGED)
}

func findStaged() {
	logger.Debug("\nStaged changes")
	r, _ := git.PlainOpen(model.CurrentProfile.LocalRepo)
	w, _ := r.Worktree()
	s, _ := w.Status()
	staged := make(map[string]string)
	for path, status := range s {
		if status.Staging == git.Renamed {
			path = fmt.Sprintf("%s -> %s", path, status.Extra)
		}
		switch status.Staging {
		case git.Unmodified:
			continue
		case git.Added:
			logger.Info(string(status.Staging), path)
			staged[path] = string(status.Staging)
		case git.Modified:
			logger.Info(string(status.Staging), path)
			staged[path] = string(status.Staging)
		case git.Deleted:
			logger.Error(string(status.Staging), path)
			staged[path] = string(status.Staging)
		default:
			logger.Debug(string(status.Staging), path)
		}
	}
	model.Set(staged, constants.STAGED)
}

func findCommited() {
	r, _ := git.PlainOpen(model.CurrentProfile.LocalRepo)
	pl, _ := r.Head()
	revHash, _ := r.ResolveRevision(plumbing.Revision("origin/" + pl.Name().Short()))

	revCommit, _ := r.CommitObject(*revHash)

	revTree, _ := revCommit.Tree()
	logger.Debug("REV TREE")
	// revTree.Files().ForEach(func(f *object.File) error {
	// 	fmt.Printf("%s    %s\n", f.Hash, f.Name)
	// 	return nil
	// })
	logger.Debug("REV HASH_1: %s", revHash)
	logger.Debug("REV HASH_2: %s", revTree.Hash)

	headRef, _ := r.Head()

	// ... retrieving the commit object
	headCommit, _ := r.CommitObject(headRef.Hash())

	headTree, _ := headCommit.Tree()
	logger.Debug("HEAD TREE")
	// headTree.Files().ForEach(func(f *object.File) error {
	// 	fmt.Printf("%s    %s\n", f.Hash, f.Name)
	// 	return nil
	// })

	logger.Debug("HEAD HASH_1: %s", headRef.Hash())
	logger.Debug("HEAD HASH_2: %s", headTree.Hash)

	// commitIter, _ := r.Log(&git.LogOptions{})
	// commitIter.ForEach(func(c *object.Commit) error {
	// 	fmt.Print(c.Hash.String())
	// 	fmt.Print(" " + c.Message)
	// 	fmt.Println(" " + revHash.String())
	// 	if c.Hash.String() == revHash.String() {

	// 		fmt.Println("Should break")
	// 		return nil
	// 	}
	// fileIterator, _ := c.Files()
	// fileIterator.ForEach(func(f *object.File) error {
	// 	fmt.Println("\t" + f.Name)
	// 	return nil
	// })
	// 	return nil
	// })

	isAncestor, _ := headCommit.IsAncestor(revCommit)
	commited := make(map[string]string)
	if !isAncestor {
		commited["Local commits present"] = "A"
	}
	model.Set(commited, constants.COMMITED)
}
