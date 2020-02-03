package service

import (
	"net/http"
	"time"

	"GMOps/src/proto"
	"GMOps/src/service/bi/model"

	"github.com/gin-gonic/gin"
)

func (s *Service) getCommitInfoAction(c *gin.Context) {
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

func (s *Service) getTreeInfoAction(c *gin.Context) {
	repo, ok := c.Keys["Repo"].(*model.Repo)
	if !ok {
		c.ProtoBuf(http.StatusNotFound, &proto.Error{
			ErrorCode: 404,
			Reason:    "repo not exist",
		})
		return
	}
	treeId := c.Param("treeId")
	tree, err := s.lgc.GetRepoTree(repo, treeId)
	if err != nil {
		c.ProtoBuf(http.StatusNotFound, &proto.Error{
			ErrorCode: 404,
			Reason:    "tree not exist",
		})
		return
	}
	entries := make([]*proto.GitTreeEntry, tree.EntryCount())
	for i := uint64(0); i < tree.EntryCount(); i++ {
		entry := tree.EntryByIndex(i)
		if entry == nil {
			continue
		}
		entries[i] = &proto.GitTreeEntry{
			Id:       entry.Id.String(),
			Name:     entry.Name,
			Type:     entry.Type.String(),
			Filemode: int32(entry.Filemode),
		}
	}
	ret := proto.GitTree{
		Id:      treeId,
		Entries: entries,
	}
	c.ProtoBuf(http.StatusOK, &ret)
}
