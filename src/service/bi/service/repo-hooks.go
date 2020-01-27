package service

import (
	"net/http"

	"GMOps/src/proto"

	"github.com/gin-gonic/gin"
)

type createHookMsg struct {
	Type   string `form:"type"`
	Name   string `form:"name"`
	Source string `form:"source"`
}

func (s *Service) getRepoPreReceiveHooksAction(c *gin.Context) {
	repoId := c.Param("repoId")
	repo, err := s.lgc.GetRepo(repoId)
	if err != nil {
		c.ProtoBuf(http.StatusNotFound, &proto.Error{
			ErrorCode: 404,
			Reason:    "repo not exist",
		})
		return
	}
	hooks, err := s.lgc.GetRepoPreReceiveHooks(repo)
	if err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}
	sources := make([]*proto.Hook, 0)
	ret := &proto.Hooks{Type: "pre-receive"}
	for _, hook := range hooks {
		sources = append(sources, &proto.Hook{
			Id:     hook.Id.Hex(),
			Name:   hook.Name,
			Source: hook.LuaSource,
		})

	}
	ret.Hooks = sources

	c.ProtoBuf(http.StatusOK, ret)
}

func (s *Service) useRepoHookAction(c *gin.Context) {
	repoId := c.Param("repoId")
	hookId := c.Param("hookId")

	repo, err := s.lgc.GetRepo(repoId)
	if err != nil {
		c.ProtoBuf(http.StatusNotFound, &proto.Error{
			ErrorCode: 404,
			Reason:    "repo not exist",
		})
		return
	}
	hook, err := s.lgc.GetHook(hookId)
	if err != nil {
		c.ProtoBuf(http.StatusNotFound, &proto.Error{
			ErrorCode: 404,
			Reason:    "repo not exist",
		})
		return
	}

	repoHookId, err := s.lgc.AddRepoHook(repo, hook)
	if err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}

	c.ProtoBuf(http.StatusCreated, &proto.Created{Id: repoHookId})
}

func (s *Service) addGitHookAction(c *gin.Context) {
	req := createHookMsg{}
	if err := c.ShouldBind(&req); err != nil {
		s.logger.Error("BI Service addGitHookAction bind req failed: ", err)
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
		return
	}
	s.logger.Info(req)
	if len(req.Type) == 0 || len(req.Name) == 0 || len(req.Source) == 0 {
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
		return
	}
	if req.Type != "pre-receive" && req.Type != "update" && req.Type != "post-receive" && req.Type != "post-update" {
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "type error",
		})
		return
	}

	hookId, err := s.lgc.AddHook(req.Type, req.Name, req.Source)
	if err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}

	c.ProtoBuf(http.StatusCreated, &proto.Created{Id: hookId})
}
