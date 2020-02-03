package service

import (
	"net/http"

	"GMOps/src/proto"
	"GMOps/src/service/bi/model"

	"github.com/gin-gonic/gin"
)

func (s *Service) getUserRepoRefsAction(c *gin.Context) {
	repo, ok := c.Keys["Repo"].(*model.Repo)
	if !ok {
		c.ProtoBuf(http.StatusNotFound, &proto.Error{
			ErrorCode: 404,
			Reason:    "repo not exist",
		})
		return
	}

	refs := s.lgc.GetRepoRefs(repo)
	ret := proto.RepoRefs{
		Repo: &proto.Repo{
			Id:       repo.Id.Hex(),
			Name:     repo.Name,
			Descript: repo.Descript,
		},
	}
	protoRefs := make([]*proto.RepoRef, 0)

	for _, ref := range refs {
		protoRefs = append(protoRefs, &proto.RepoRef{
			TargetId: ref.Target().String(),
			IsBranch: ref.IsBranch(),
			IsNote:   ref.IsNote(),
			IsRemote: ref.IsRemote(),
			IsTag:    ref.IsTag(),
			Name:     ref.Name(),
		})
		ref.Free()
	}
	ret.Refs = protoRefs

	c.ProtoBuf(http.StatusOK, &ret)
}
