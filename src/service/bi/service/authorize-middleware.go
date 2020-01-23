package service

import (
	"net/http"

	"GMOps/src/proto"

	"github.com/gin-gonic/gin"
)

func (s *Service) authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("GMOps-Username")
		s.logger.Info("valid user authorized: ", username)
		if len(username) == 0 {
			c.ProtoBuf(http.StatusUnauthorized, &proto.Error{
				ErrorCode: 401,
				Reason:    "Unauthorized",
			})
			return
		}
		user, err := s.lgc.LoadUser(username)
		if err != nil {
			c.ProtoBuf(http.StatusUnauthorized, &proto.Error{
				ErrorCode: 401,
				Reason:    "Unauthorized",
			})
			return
		}
		if !user.Available {
			c.ProtoBuf(http.StatusUnauthorized, &proto.Error{
				ErrorCode: 401,
				Reason:    "Unauthorized",
			})
			return
		}
		c.Keys = map[string]interface{}{"User": user}
		c.Next()
	}
}
