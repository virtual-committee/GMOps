package service

import (
	"net/http"

	"GMOps/src/proto"

	"github.com/gin-gonic/gin"
)

type createUserMsg struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

func (s *Service) createUserAction(c *gin.Context) {
	req := createUserMsg{}
	if err := c.ShouldBind(&req); err != nil {
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
	}

	// TODO create User

	c.ProtoBuf(http.StatusCreated, &proto.Created{Id: "fuck"})
}
