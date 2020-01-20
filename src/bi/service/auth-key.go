package service

import (
	"net/http"

	"GMOps/src/bi/model"
	"GMOps/src/proto"

	"github.com/gin-gonic/gin"
)

type addUserAuthKeyMsg struct {
	Title string `form:"title"`
	Key   string `form:"key"`
}

func (s *Service) getUserAuthKeysAction(c *gin.Context) {
	user, ok := c.Keys["User"].(*model.User)
	if !ok {
		c.ProtoBuf(http.StatusUnauthorized, &proto.Error{
			ErrorCode: 401,
			Reason:    "Unauthorized",
		})
		return
	}

	authKeys, err := s.lgc.GetUserAuthKeys(user)
	if err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}

	ret := &proto.UserAuthKeys{UserId: user.Id.Hex()}
	retKeys := make([]*proto.AuthKey, 0)
	for _, key := range authKeys {
		retKey := &proto.AuthKey{
			Id:        key.Id.Hex(),
			Title:     key.Title,
			Key:       key.AuthKey,
			Available: key.Writed && key.Available,
		}
		retKeys = append(retKeys, retKey)
	}
	ret.Keys = retKeys
	c.ProtoBuf(http.StatusOK, ret)
}

func (s *Service) addUserAuthKeyAction(c *gin.Context) {
	user, ok := c.Keys["User"].(*model.User)
	if !ok {
		c.ProtoBuf(http.StatusUnauthorized, &proto.Error{
			ErrorCode: 401,
			Reason:    "Unauthorized",
		})
		return
	}
	req := addUserAuthKeyMsg{}
	if err := c.ShouldBind(&req); err != nil {
		s.logger.Error("BI Service addUserAuthKeyAction bind req failed: ", err)
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
		return
	}
	s.logger.Info(req)
	if len(req.Title) == 0 || len(req.Key) == 0 {
		c.ProtoBuf(http.StatusBadRequest, &proto.Error{
			ErrorCode: 400,
			Reason:    "req body cannot bind",
		})
		return
	}

	createdId, err := s.lgc.AddUserAuthKey(user, req.Title, req.Key)
	if err != nil {
		c.ProtoBuf(http.StatusConflict, &proto.Error{
			ErrorCode: 409,
			Reason:    "already existed auth_key",
		})
		return
	}
	c.ProtoBuf(http.StatusCreated, &proto.Created{Id: createdId})
}

func (s *Service) validUserAuthKeyAction(c *gin.Context) {
	user, ok := c.Keys["User"].(*model.User)
	if !ok {
		c.ProtoBuf(http.StatusUnauthorized, &proto.Error{
			ErrorCode: 401,
			Reason:    "Unauthorized",
		})
		return
	}
	id := c.Param("id")
	authKey, err := s.lgc.GetUserAuthKeyByID(id)
	if err != nil {
		c.ProtoBuf(http.StatusNotFound, &proto.Error{
			ErrorCode: 404,
			Reason:    "Not found auth_key",
		})
		return
	}
	if !authKey.Writed {
		c.ProtoBuf(http.StatusForbidden, &proto.Error{
			ErrorCode: 403,
			Reason:    "the auth_key not writed yet",
		})
		return
	}
	if !authKey.Available {
		c.ProtoBuf(http.StatusForbidden, &proto.Error{
			ErrorCode: 403,
			Reason:    "the auth_key disabled",
		})
		return
	}
	if authKey.User != user.Id {
		c.ProtoBuf(http.StatusForbidden, &proto.Error{
			ErrorCode: 403,
			Reason:    "the auth_key do not belong this user",
		})
		return
	}
	c.String(http.StatusOK, "pass")
}

func (s *Service) applyUserAuthKeyAction(c *gin.Context) {
	id := c.Param("id")
	if err := s.lgc.ApplyUserAuthKey(id); err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}
	c.String(http.StatusOK, "pass")
}

func (s *Service) cancelUserAuthKeyAction(c *gin.Context) {
	id := c.Param("id")
	if err := s.lgc.CancelUserAuthKey(id); err != nil {
		c.ProtoBuf(http.StatusInternalServerError, &proto.Error{
			ErrorCode: 500,
			Reason:    "server internal error",
		})
		return
	}
	c.String(http.StatusOK, "pass")
}
