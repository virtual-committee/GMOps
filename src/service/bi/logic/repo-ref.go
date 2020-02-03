package logic

import (
	"GMOps/src/service/bi/model"

	git "github.com/libgit2/git2go"
)

func (lgc *Logic) GetRepoRefs(repo *model.Repo) []*git.Reference {
	ret := make([]*git.Reference, 0)
	gitRepo, err := repo.OpenGitRepo(GMOPS_REPO_BASE_PATH, lgc.logger)
	if err != nil {
		lgc.logger.Error("Get Repo Refs error open git repo failed, ", err)
		return ret
	}

	refIter, err := gitRepo.NewReferenceIterator()
	if err != nil {
		lgc.logger.Error("Get Repo Refs error NewReferenceIterator failed, ", err)
		return ret
	}
	defer refIter.Free()

	for {
		ref, err := refIter.Next()
		if err != nil {
			lgc.logger.Error("Get Repo Refs error, refIter.Next failed, ", err)
			return ret
		}
		ret = append(ret, ref)
	}

	lgc.logger.Info("GetRepoRefs result count: ", len(ret))
	return ret
}
