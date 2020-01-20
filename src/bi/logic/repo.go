package logic

import (
	"GMOps/src/bi/model"
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
