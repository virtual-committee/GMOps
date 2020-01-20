package service

import (
	"net/http"

	"GMOps/src/bi/model"
	"GMOps/src/proto"

	"github.com/gin-gonic/gin"
)

func (s *Service) getUserReposAction(c *gin.Context) {
	user, ok := c.Keys["User"].(*model.User)
	if !ok {
		c.ProtoBuf(http.StatusUnauthorized, &proto.Error{
			ErrorCode: 401,
			Reason:    "Unauthorized",
		})
		return
	}
	repos, err := s.lgc.GetUserRepos(user)
	if err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}

	ret := &proto.UserRepos{UserId: user.Id.Hex()}
	retRepos := make([]*proto.Repo, 0)
	for _, repo := range repos {
		retRepos = append(retRepos, &proto.Repo{
			Id:       repo.Id.Hex(),
			Name:     repo.Name,
			Descript: repo.Descript,
			Attr:     repo.Attr,
		})
	}
	ret.Repos = retRepos
	c.ProtoBuf(http.StatusOK, ret)
}
