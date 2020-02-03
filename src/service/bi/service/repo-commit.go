package service

import (
	"net/http"
	"time"

	"GMOps/src/proto"

	"github.com/gin-gonic/gin"
	git "github.com/libgit2/git2go"
)

func (s *Service) getCommitInfoAction(c *gin.Context) {
	commit, ok := c.Keys["Commit"].(*git.Commit)
	if !ok {
		c.ProtoBuf(http.StatusNotFound, &proto.Error{
			ErrorCode: 404,
			Reason:    "commit not exist",
		})
		return
	}
	ret := proto.GitCommit{
		Message: commit.Message(),
		Author: &proto.GitSignature{
			Name:  commit.Author().Name,
			Email: commit.Author().Email,
			Time:  commit.Author().When.Format(time.UnixDate),
		},
		Committer: &proto.GitSignature{
			Name:  commit.Committer().Name,
			Email: commit.Committer().Email,
			Time:  commit.Committer().When.Format(time.UnixDate),
		},
		TreeId: commit.TreeId().String(),
	}
	parentId := make([]string, commit.ParentCount())
	for i := uint(0); i < commit.ParentCount(); i++ {
		parentId[i] = commit.ParentId(i).String()
	}
	ret.ParentId = parentId

	c.ProtoBuf(http.StatusOK, &ret)
}
