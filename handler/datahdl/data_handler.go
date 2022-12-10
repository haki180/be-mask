package datahdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/platformsh/template-golang/domain"
	"github.com/platformsh/template-golang/service"
)

type dataHandler struct {
	dataService service.DataService
	userService service.UserService
}

func NewDataHandler(dataService service.DataService, userService service.UserService) *dataHandler {
	return &dataHandler{dataService: dataService, userService: userService}
}

func (instance dataHandler) Create(c *gin.Context) {
	// check authorization header
	if c.Request.Header.Get("Authorization") != "jbdae8dy8293je2jnrke23r92rj2knkfxr239ujx2j90j22r8h28xrh289hnxhrc298rh2" {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "UNAUTHORIZED",
			"message": "you are not authorized to access this service",
		})
		return
	} else {
		in := new(domain.CreateDataRequest)
		if err := c.ShouldBindJSON(&in); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATOR",
					"message": err.Error(),
				})
			return
		} else {
			if err := instance.dataService.Create(c, *in); err != nil {
				domain.NewCreateDataResponse(c, http.StatusInternalServerError, err)
			} else {
				domain.NewCreateDataResponse(c, http.StatusOK, "success")
				return
			}
		}
	}
}

func (instance dataHandler) Delete(c *gin.Context) {
	// check authorization header
	token := c.Request.Header.Get("Authorization")
	if token != "" {
		// users
		if isTokenExist, err := instance.userService.CheckToken(c, token); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error":   "INTERNAL SERVER ERROR",
				"message": err.Error(),
			})
			return
		} else {
			if !isTokenExist {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "UNAUTHORIZED",
					"message": "you are not authorized to access this service",
				})
				return
			} else {
				uuid := c.Param("uuid")
				if uuid == "" {
					c.JSON(http.StatusBadRequest, map[string]interface{}{
						"error":   "BAD REQUEST",
						"message": "uuid params bad request",
					})
				} else {
					if err := instance.dataService.Delete(c, uuid); err != nil {
						if err.Error() == "data not found" {
							c.JSON(http.StatusNotFound, map[string]interface{}{
								"error":   "NOT FOUND",
								"message": err.Error(),
							})
						} else {
							c.JSON(http.StatusInternalServerError, map[string]interface{}{
								"error":   "INTERNAL SERVER ERROR",
								"message": err.Error(),
							})
						}
					} else {
						c.JSON(http.StatusOK, map[string]interface{}{
							"message": "success",
						})
					}
				}
			}
		}
	} else {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "UNAUTHORIZED",
			"message": "missing token",
		})
		return
	}
}

func (instance dataHandler) ChartData(c *gin.Context) {
	// check authorization header
	token := c.Request.Header.Get("Authorization")
	if token != "" {
		// users
		if isTokenExist, err := instance.userService.CheckToken(c, token); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error":   "INTERNAL SERVER ERROR",
				"message": err.Error(),
			})
			return
		} else {
			if !isTokenExist {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "UNAUTHORIZED",
					"message": "you are not authorized to access this service",
				})
				return
			} else {
				in := new(domain.GetDataChartParams)
				in.SetStart(c.Query("start"))
				in.SetEnd(c.Query("end"))
				datas, err := instance.dataService.GetTotalDataPerType(c, *in)
				if err != nil {
					c.JSON(http.StatusUnauthorized, map[string]interface{}{
						"error":   "INTERNAL SERVER ERROR",
						"message": err.Error(),
					})
				} else {
					c.JSON(http.StatusOK, datas)
				}
			}
		}
	} else {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "UNAUTHORIZED",
			"message": "missing token",
		})
		return
	}
}

func (instance dataHandler) PaginateData(c *gin.Context) {
	// check authorization header
	token := c.Request.Header.Get("Authorization")
	if token != "" {
		// users
		if isTokenExist, err := instance.userService.CheckToken(c, token); err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error":   "INTERNAL SERVER ERROR",
				"message": err.Error(),
			})
			return
		} else {
			if !isTokenExist {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "UNAUTHORIZED",
					"message": "you are not authorized to access this service",
				})
				return
			} else {
				in := new(domain.GetDatasPaginateParams)
				in.SetLimit(c.Query("limit"))
				in.SetPage(c.Query("page"))
				in.SetType(c.Query("type"))
				in.SetStart(c.Query("start"))
				in.SetEnd(c.Query("end"))

				datas, err := instance.dataService.GetDatasPaginate(c, *in)
				if err != nil {
					c.JSON(http.StatusUnauthorized, map[string]interface{}{
						"error":   "INTERNAL SERVER ERROR",
						"message": err.Error(),
					})
				} else {
					c.JSON(http.StatusOK, datas)
				}
			}
		}
	} else {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "UNAUTHORIZED",
			"message": "missing token",
		})
		return
	}
}
