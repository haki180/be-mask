package domain

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uint64
	UUID         string
	UserName     string
	Password     string
	LoginAt      *time.Time
	LogoutAt     *time.Time
	SessionToken string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (u *User) IsEmpty() bool {
	return u == nil
}

func (u User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u User) GenerateSession() string {
	return base64.StdEncoding.EncodeToString([]byte(u.UUID))
}

func (u *User) SetSessionTokenToBeEmpty() {
	u.SessionToken = ""
}

func (u *User) SetLogin() {
	loginAt := time.Now()
	u.LoginAt = &loginAt
}

func (u *User) SetLogout() {
	logoutAt := time.Now()
	u.LogoutAt = &logoutAt
}

func (u *User) SetSessionToken(sessionToken string) {
	u.SessionToken = sessionToken
}

type CreteUserRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l CreteUserRequest) GetUserName() string {
	return l.UserName
}

func (l CreteUserRequest) GetPassword() string {
	return l.Password
}

func (l CreteUserRequest) GeneratePassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(l.Password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (l CreteUserRequest) ToUser(hashedPassword string) *User {
	return &User{
		UUID:     uuid.New().String(),
		UserName: l.UserName,
		Password: hashedPassword,
	}
}

type CreateNewUserResponse struct {
	SessionToken string `json:"session_token"`
}

func CreateNewUserResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func CreateNewUserResponseError(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
		"message": err.Error(),
	})
}

type LoginRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l LoginRequest) GetUserName() string {
	return l.UserName
}

func (l LoginRequest) GetPassword() string {
	return l.Password
}

type LoginResponse struct {
	SessionToken string `json:"session_token"`
}

func LoginResponseSuccess(c *gin.Context, session_token string) {
	c.JSON(http.StatusOK, LoginResponse{
		SessionToken: session_token,
	})
}

func LoginResponseError(c *gin.Context, statusCode int, errDesc string, err error) {
	c.JSON(statusCode, map[string]interface{}{
		"error":   errDesc,
		"message": err.Error(),
	})
}

type LogoutRequest struct {
	SessionToken string `json:"session_token"`
}

func (l LogoutRequest) GetSessionToken() string {
	return l.SessionToken
}

func LogoutResponseSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func LogoutResponseError(c *gin.Context, statusCode int, errDesc string, err error) {
	c.JSON(statusCode, map[string]interface{}{
		"error":   errDesc,
		"message": err.Error(),
	})
}
