package routes

import (
	"context"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/markwallsgrove/muzz_devops/src/database"
	"github.com/markwallsgrove/muzz_devops/src/models/security"
	"go.uber.org/zap"
)

// LoginController login functionality
type LoginController struct {
	Ctx      context.Context
	Database database.Database
	Logger   *zap.Logger
	Secret   string
}

type payload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login login with a email and password. The response is a JWT token that can be used
// on the APIs.
func (i *LoginController) Login(c echo.Context) error {
	var payload payload
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		i.Logger.Error("cannot marshall body", zap.Error(err))
		return c.String(http.StatusBadRequest, "invalid payload")
	}

	user, err := i.Database.GetUserByEmail(i.Ctx, payload.Email)
	if err == database.ErrNotFound {
		return c.String(http.StatusForbidden, "unknown email address or incorrect password")
	}
	if err != nil {
		i.Logger.Error("cannot retrieve password hash", zap.Error(err))
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	err = security.VerifyPasswordHash(payload.Password, user.PasswordHash)
	if err == security.ErrInvalidPassword {
		return c.String(http.StatusForbidden, "unknown email address or incorrect password")
	}
	if err != nil {
		i.Logger.Error("cannot verify password hash", zap.Error(err))
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	token, err := security.CreateToken(user.ID, i.Secret)
	if err != nil {
		i.Logger.Error("cannot create token", zap.Error(err))
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, security.LoginResponse{Result: token})
}
