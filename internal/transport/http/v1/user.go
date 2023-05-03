package v1

import (
	"net/http"
	"vk-api/internal/usecase"

	"github.com/gin-gonic/gin"
)

const (
	ErrUserNotFound  = "err: user not found"
	ErrUserIsPrivate = "err: user is private"
	ErrInvalidUserID = "err: invalid user_id"
)

type logger interface {
	Info(msg string)
	Warning(msg string)
	Error(msg string)
	Fatal(msg string)
}

type UserHandler struct {
	uc *usecase.UserUsecase
	logger
}

func NewUserHandler(userUsecase *usecase.UserUsecase, l logger) *UserHandler {
	return &UserHandler{userUsecase, l}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.uc.GetUser(userID)
	if err != nil {
		if err.Error() == ErrUserNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			h.logger.Error(err.Error())
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			h.logger.Error(err.Error())
			return
		}
	}

	c.JSON(200, user)
}

func (h *UserHandler) GetUserFriends(c *gin.Context) {
	userID := c.Param("id")

	friends, err := h.uc.GetUserFriendsByID(userID)
	if err != nil {
		if err.Error() == ErrUserIsPrivate {
			c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
			h.logger.Error(err.Error())
			return
		} else if err.Error() == ErrInvalidUserID {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			h.logger.Error(err.Error())
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			h.logger.Error(err.Error())
			return
		}
	}

	c.JSON(200, friends)
}
