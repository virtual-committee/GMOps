package logic

import (
	"GMOps/src/service/bi/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (lgc *Logic) GetRepoHooks(repo *model.Repo, hookType string) ([]*model.GitHook, error) {
	ret := make([]*model.GitHook, 0)
	hooks, err := model.LoadGitHooksByRepoId(repo, lgc.db, lgc.logger)
	if err != nil {
		return nil, err
	}
	for _, hook := range hooks {
		if hook.Type == hookType {
			ret = append(ret, hook)
		}
	}
	return ret, nil
}

func (lgc *Logic) GetHook(id string) (*model.GitHook, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return model.LoadGitHookById(oid, lgc.db, lgc.logger)
}

func (lgc *Logic) AddRepoHook(repo *model.Repo, hook *model.GitHook) (string, error) {
	repoHook := model.NewRepoGitHook()
	repoHook.RepoId = repo.Id
	repoHook.GitHookId = hook.Id
	if err := repoHook.Save(lgc.db, lgc.logger); err != nil {
		return "", err
	}
	return repoHook.Id.Hex(), nil
}

func (lgc *Logic) AddHook(hookType string, name string, source string, spec bool) (string, error) {
	hook := model.NewGitHook()
	hook.Type = hookType
	hook.Name = name
	hook.LuaSource = source
	hook.Spec = spec

	if err := hook.Save(lgc.db, lgc.logger); err != nil {
		return "", err
	}
	return hook.Id.Hex(), nil
}
