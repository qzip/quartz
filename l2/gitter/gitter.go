package gitter

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Repo struct {
	git  *git.Repository
	user *Author
}

func Open(path string) (g *git.Repository, err error) {
	g, err = git.PlainOpen(path)
	return
}

func NewRepo(path string, user *Author) (repo *Repo, err error) {
	repo = &Repo{user: user}
	repo.git, err = git.PlainOpen(path)
	return
}

func (r *Repo) Hash() (hash string, err error) {

	ref, erx := r.git.Head()
	if erx != nil {
		err = erx
		return
	}
	if ref.Hash().IsZero() {
		hash = ""
	} else {
		hash = ref.Hash().String()
	}
	return
}

func (r *Repo) Commit() (hash string, err error) {
	w, er := r.git.Worktree()
	if er != nil {
		err = er
		return
	}
	status, erx := w.Status()

	if erx != nil {
		err = erx
		return
	}
	if status.IsClean() {
		return r.Hash()
	}
	var tm = time.Now()
	var treeHash plumbing.Hash
	treeHash, err = w.Commit(tm.String(), &git.CommitOptions{
		Author: &object.Signature{
			Name:  r.user.Name,
			Email: r.user.Email,
			When:  tm,
		}})
	if err == nil {
		hash = treeHash.String()
	}
	return
}
