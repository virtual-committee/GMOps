package logic

import (
	"fmt"

	"GMOps/src/bi/model"

	git "github.com/libgit2/git2go"
)

const (
	GMOPS_REPO_BASE_PATH = "/opt/GMOps/data/"
)

func (lgc *Logic) GetUserRepos(user *model.User) ([]*model.Repo, error) {
	repos := make([]*model.Repo, 0)
	userRepos, err := model.LoadUserReposByUser(user, lgc.db, lgc.logger)
	if err != nil {
		return nil, err
	}
	for _, userRepo := range userRepos {
		repo, err := model.LoadRepoById(userRepo.Repo, lgc.db, lgc.logger)
		if err != nil {
			return nil, err
		}
		repos = append(repos, repo)
	}

	return repos, nil
}

func (lgc *Logic) CreateUserRepo(user *model.User, name, descript string) (string, error) {
	repo := model.NewRepo()
	repo.Name = name
	repo.Descript = descript

	if err := repo.Save(lgc.db, lgc.logger); err != nil {
		return "", err
	}
	userRepo := model.NewUserRepo()
	userRepo.Repo = repo.Id
	userRepo.User = user.Id
	if err := userRepo.Save(lgc.db, lgc.logger); err != nil {
		return "", err
	}

	if _, err := git.InitRepository(fmt.Sprintf("%s%s", GMOPS_REPO_BASE_PATH, repo.Id.Hex()), true); err != nil {
		return "", err
	}
	return repo.Id.Hex(), nil
}

func (lgc *Logic) ExistUserRepo(user *model.User, name string) (bool, error) {
	repos, err := lgc.GetUserRepos(user)
	if err != nil {
		return false, err
	}
	for _, repo := range repos {
		if name == repo.Name {
			return true, nil
		}
	}
	return false, nil
}
