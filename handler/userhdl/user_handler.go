package userhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platformsh/template-golang/domain"
	"github.com/platformsh/template-golang/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (instance userHandler) Create(c *gin.Context) {
	// check authorization header
	if c.Request.Header.Get("Authorization") != "jbdae8dy8293je2jnrke23r92rj2knkfxr239ujx2j90j22r8h28xrh289hnxhrc298rh2" {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "UNAUTHORIZED",
			"message": "you are not authorized to access this service",
		})
		return
	} else {
		in := new(domain.CreteUserRequest)
		if err := c.ShouldBindJSON(&in); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATOR",
					"message": err.Error(),
				})
			return
		} else {
			if err := instance.userService.CreateNewUser(c, *in); err != nil {
				domain.CreateNewUserResponseError(c, err)
			} else {
				domain.CreateNewUserResponseSuccess(c)
				return
			}
		}
	}
}

func (instance userHandler) Login(c *gin.Context) {
	in := new(domain.LoginRequest)
	if err := c.ShouldBindJSON(&in); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATOR",
				"message": err.Error(),
			})
		return
	} else {
		sessionToken, err := instance.userService.Login(c, *in)
		if err != nil {
			if err.Error() == "invalid username" || err.Error() == "invalid password" {
				domain.LoginResponseError(c, http.StatusBadRequest, "BAD REQUEST", err)
			} else {
				domain.LoginResponseError(c, http.StatusInternalServerError, "INTERNAL SERVER ERROR", err)
			}
		} else {
			domain.LoginResponseSuccess(c, sessionToken)
			return
		}
	}
}

func (instance userHandler) Logout(c *gin.Context) {
	in := new(domain.LogoutRequest)
	if err := c.ShouldBindJSON(&in); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATOR",
				"message": err.Error(),
			})
		return
	} else {
		if err := instance.userService.Logout(c, *in); err != nil {
			if err.Error() == "invalid session token" {
				domain.LogoutResponseError(c, http.StatusBadRequest, "BAD REQUEST", err)
			} else {
				domain.LogoutResponseError(c, http.StatusInternalServerError, "INTERNAL SERVER ERROR", err)
			}
		} else {
			domain.LogoutResponseSuccess(c)
			return
		}
	}
}
