package mockhdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"session_token": "nkdnje4h981u4901312i4n114u194m1ke283ur901i41e",
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil logout",
	})
}

func ChartData(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"proper":   78,
		"improper": 1823,
		"no":       1238,
	})
}

func PaginateData(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": []map[string]interface{}{
			{
				"id":         1,
				"type":       "proper",
				"created_at": "14 Maret 2022 08:30:00",
				"image":      "ndjak898e4j4i1j4901u341oj4914jo1j1", // byte type
			},
			{
				"id":         2,
				"type":       "proper",
				"created_at": "14 Maret 2022 08:30:00",
				"image":      "ndjak898e4j4i1j4901u341oj4914jo1j1", // byte type
			},
			{
				"id":         3,
				"type":       "proper",
				"created_at": "14 Maret 2022 08:30:00",
				"image":      "ndjak898e4j4i1j4901u341oj4914jo1j1", // byte type
			},
		},
		"meta": map[string]interface{}{
			"page":       1,
			"total_page": 10,
			"total_data": 100,
		},
	})
}
