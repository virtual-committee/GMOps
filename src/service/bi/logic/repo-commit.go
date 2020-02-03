package logic

import (
	"GMOps/src/service/bi/model"

	git "github.com/libgit2/git2go"
)

func (lgc *Logic) GetRepoCommit(repo *model.Repo, commitId string) (*git.Commit, error) {
	commitOId, err := git.NewOid(commitId)
	if err != nil {
		lgc.logger.Error("GetRepoCommit NewOid failed: ", err)
		return nil, err
	}
	gitRepo, err := repo.OpenGitRepo(GMOPS_REPO_BASE_PATH, lgc.logger)
	if err != nil {
		return nil, err
	}
	commit, err := gitRepo.LookupCommit(commitOId)
	if err != nil {
		lgc.logger.Error("GetRepoCommit LookupCommit failed: ", err)
		return nil, err
	}
	return commit, err
}
