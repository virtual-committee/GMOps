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
	return commit, nil
}

func (lgc *Logic) GetRepoTree(repo *model.Repo, treeId string) (*git.Tree, error) {
	treeOId, err := git.NewOid(treeId)
	if err != nil {
		lgc.logger.Error("GetRepoTree NewOid failed: ", err)
		return nil, err
	}
	gitRepo, err := repo.OpenGitRepo(GMOPS_REPO_BASE_PATH, lgc.logger)
	if err != nil {
		return nil, err
	}
	tree, err := gitRepo.LookupTree(treeOId)
	if err != nil {
		lgc.logger.Error("GetRepoTree LookupTree failed: ", err)
		return nil, err
	}
	return tree, nil
}
