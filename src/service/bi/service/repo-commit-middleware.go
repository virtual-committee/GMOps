package service

import (
	"net/http"

	"GMOps/src/proto"
	"GMOps/src/service/bi/model"

	"github.com/gin-gonic/gin"
)

func (s *Service) repoCommitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		repo, ok := c.Keys["Repo"].(*model.Repo)
		if !ok {
			c.ProtoBuf(http.StatusNotFound, &proto.Error{
				ErrorCode: 404,
				Reason:    "repo not exist",
			})
			return
		}
		commitId := c.Param("commitId")
		commit, err := s.lgc.GetRepoCommit(repo, commitId)
		if err != nil {
			c.ProtoBuf(http.StatusNotFound, &proto.Error{
				ErrorCode: 404,
				Reason:    "commit not exist",
			})
			return
		}
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["Commit"] = commit

		c.Next()
	}
}
