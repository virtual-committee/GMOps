package service

import (
	"net/http"

	"GMOps/src/proto"
	"GMOps/src/service/bi/model"

	"github.com/gin-gonic/gin"
)

type createUserMsg struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

type validUserMsg struct {
	Password string `form:"password"`
}

func (s *Service) createUserAction(c *gin.Context) {
	req := createUserMsg{}
	if err := c.ShouldBind(&req); err != nil {
		s.logger.Error("BI Service createUserAction bind req failed: ", err)
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
		return
	}
	if len(req.Name) == 0 || len(req.Password) == 0 {
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "missing fields",
		})
		return
	}

	existed, err := s.lgc.ExistsUser(req.Name)
	if err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}
	if existed {
		c.ProtoBuf(http.StatusConflict, &proto.Error{
			ErrorCode: 409,
			Reason:    "existed same name user",
		})
		return
	}

	createdId, err := s.lgc.CreateUser(req.Name, req.Password)
	if err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}

	c.ProtoBuf(http.StatusCreated, &proto.Created{Id: createdId})
}

func (s *Service) validUserAction(c *gin.Context) {
	req := validUserMsg{}
	if err := c.ShouldBind(&req); err != nil {
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
		return
	}
	user, ok := c.Keys["User"].(*model.User)
	if !ok {
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
		return
	}
	if len(req.Password) != 0 && req.Password != user.Password {
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "password wrong",
		})
		return
	}
	c.String(http.StatusOK, user.Id.Hex())
}
