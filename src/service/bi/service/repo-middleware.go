package service

import (
	"net/http"

	"GMOps/src/proto"

	"github.com/gin-gonic/gin"
)

func (s *Service) repoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		repoId := c.Param("repoId")
		repo, err := s.lgc.GetRepo(repoId)
		if err != nil {
			c.ProtoBuf(http.StatusNotFound, &proto.Error{
				ErrorCode: 404,
				Reason:    "repo not exist",
			})
			return
		}
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["Repo"] = repo
		c.Next()
	}
}
