package handler

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"golang_bootstrap_project/errors"
	"golang_bootstrap_project/service"
	"golang_bootstrap_project/utils"
	"net/http"
)

type UserHandler interface {
	User(c echo.Context) error
}

type UserHandlerImpl struct {
	UserService service.UserService
	StringUtils utils.StringUtils
}

func (u UserHandlerImpl) User(c echo.Context) error {
	userId := u.StringUtils.ConvertStringToLowerCase(c.Param("userId"))
	if len(userId) == 0 {
		log.Errorf("no userId was given")
		return c.JSON(http.StatusUnprocessableEntity, errors.TechnicalErrorNoUserIdWasGiven())
	}
	user, userErr := u.UserService.FindActiveUserWithUserId(userId)
	if userErr != nil {
		log.WithError(userErr).Warnf("error while gettting user information for %s", userId)
		return c.JSON(http.StatusOK, errors.TechnicalErrorWasOccurred())
	}
	return c.JSON(http.StatusOK, user)
}
