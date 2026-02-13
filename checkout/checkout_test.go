package checkout

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestCloneOrUpdateClonesWhenDirectoryMissing(t *testing.T) {
	remoteDir, _ := setupRemoteRepo(t)
	localDir := filepath.Join(t.TempDir(), "local")

	cloneOrUpdate(localDir, remoteDir)

	if _, err := os.Stat(localDir); err != nil {
		t.Fatalf("expected local repository directory to exist: %v", err)
	}

	if _, err := git.PlainOpen(localDir); err != nil {
		t.Fatalf("expected local directory to be a git repository: %v", err)
	}
}

func TestCloneOrUpdatePullsNewCommit(t *testing.T) {
	remoteDir, seedDir := setupRemoteRepo(t)
	localDir := filepath.Join(t.TempDir(), "local")

	cloneOrUpdate(localDir, remoteDir)
	before := repoHeadHash(t, localDir)

	commitAndPush(t, seedDir, "second.txt", "second commit", "add second commit")
	cloneOrUpdate(localDir, remoteDir)
	after := repoHeadHash(t, localDir)

	if before == after {
		t.Fatalf("expected head hash to change after pull, before=%s after=%s", before, after)
	}
}

func TestCloneOrUpdateNoopPullDoesNotFail(t *testing.T) {
	remoteDir, _ := setupRemoteRepo(t)
	localDir := filepath.Join(t.TempDir(), "local")

	cloneOrUpdate(localDir, remoteDir)
	before := repoHeadHash(t, localDir)

	cloneOrUpdate(localDir, remoteDir)
	after := repoHeadHash(t, localDir)

	if before != after {
		t.Fatalf("expected head hash to remain unchanged on no-op pull, before=%s after=%s", before, after)
	}
}

func setupRemoteRepo(t *testing.T) (string, string) {
	t.Helper()

	baseDir := t.TempDir()
	seedDir := filepath.Join(baseDir, "seed")
	remoteDir := filepath.Join(baseDir, "remote.git")

	seedRepo, err := git.PlainInit(seedDir, false)
	if err != nil {
		t.Fatalf("init seed repo: %v", err)
	}

	if _, err := git.PlainInit(remoteDir, true); err != nil {
		t.Fatalf("init remote repo: %v", err)
	}

	_, err = seedRepo.CreateRemote(&config.RemoteConfig{Name: "origin", URLs: []string{remoteDir}})
	if err != nil {
		t.Fatalf("create remote: %v", err)
	}

	commitAndPush(t, seedDir, "first.txt", "first commit", "add first commit")

	return remoteDir, seedDir
}

func commitAndPush(t *testing.T, repoDir, fileName, content, message string) {
	t.Helper()

	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		t.Fatalf("open repo: %v", err)
	}

	filePath := filepath.Join(repoDir, fileName)
	if err := os.WriteFile(filePath, []byte(content), 0o600); err != nil {
		t.Fatalf("write file %s: %v", fileName, err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		t.Fatalf("worktree: %v", err)
	}

	if _, err := worktree.Add(fileName); err != nil {
		t.Fatalf("git add %s: %v", fileName, err)
	}

	_, err = worktree.Commit(message, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "test",
			Email: "test@example.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		t.Fatalf("git commit: %v", err)
	}

	err = repo.Push(&git.PushOptions{RemoteName: "origin"})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		t.Fatalf("git push: %v", err)
	}
}

func repoHeadHash(t *testing.T, repoDir string) string {
	t.Helper()

	repo, err := git.PlainOpen(repoDir)
	if err != nil {
		t.Fatalf("open repo: %v", err)
	}

	head, err := repo.Head()
	if err != nil {
		t.Fatalf("head: %v", err)
	}

	return head.Hash().String()
}
