package service

import (
	"net/http"

	"GMOps/src/proto"
	"GMOps/src/service/bi/model"

	"github.com/gin-gonic/gin"
)

type createUserRepoMsg struct {
	Name     string `form:"name"`
	Descript string `form:"descript"`
}

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
		})
	}
	ret.Repos = retRepos
	c.ProtoBuf(http.StatusOK, ret)
}

func (s *Service) getUserRepoAction(c *gin.Context) {
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
	name := c.Param("name")
	for _, repo := range repos {
		if repo.Name == name {
			c.ProtoBuf(http.StatusOK, &proto.Repo{
				Id:       repo.Id.Hex(),
				Name:     repo.Name,
				Descript: repo.Descript,
			})
			return
		}
	}
	c.ProtoBuf(http.StatusNotFound, &proto.Error{
		ErrorCode: 404,
		Reason:    "repo not exist",
	})
}

func (s *Service) createUserRepoAction(c *gin.Context) {
	user, ok := c.Keys["User"].(*model.User)
	if !ok {
		c.ProtoBuf(http.StatusUnauthorized, &proto.Error{
			ErrorCode: 401,
			Reason:    "Unauthorized",
		})
		return
	}
	req := createUserRepoMsg{}
	if err := c.ShouldBind(&req); err != nil {
		s.logger.Error("BI Service createUserRepoAction bind req failed: ", err)
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
		return
	}
	if len(req.Name) == 0 {
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
		return
	}
	exist, err := s.lgc.ExistUserRepo(user, req.Name)
	if err != nil {
		s.logger.Error("BI Service createUserRepoAction ExistUserRepo failed: ", err)
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}
	if exist {
		c.ProtoBuf(http.StatusConflict, &proto.Error{
			ErrorCode: 409,
			Reason:    "repo already existed",
		})
		return
	}
	createdId, err := s.lgc.CreateUserRepo(user, req.Name, req.Descript)
	if err != nil {
		s.logger.Error("BI Service createUserRepoAction ExistUserRepo failed: ", err)
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}

	c.ProtoBuf(http.StatusCreated, &proto.Created{Id: createdId})
}
