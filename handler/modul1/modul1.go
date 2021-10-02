package modul1

import (
	"fmt"
	"net/http"
	"strconv"

	ctrl "simple-crud-go/controller/modul1"
	mdl "simple-crud-go/models/modul1"

	"github.com/gin-gonic/gin"
)

const modul = "modul1"

func DataHandler(c *gin.Context) {
	var param mdl.Request
	c.BindJSON(&param)

	if param.Id >= 0 {
		result, err := ctrl.GetData(param)
		if err == nil {
			result.Code = 200
			result.Message = "success retrieve data"
			c.JSON(http.StatusOK, result)
		} else {
			result.Code = 500
			result.Message = "internal server error"
			c.JSON(http.StatusInternalServerError, result)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": fmt.Sprintf("kode %s must be a string", modul)})
	}
}

func CreateHandler(c *gin.Context) {
	var param mdl.Request
	c.BindJSON(&param)

	if _, err := strconv.Atoi(param.Nama); err != nil {
		err = ctrl.CreateData(param)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success create data"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "this account doesn't have access to create data"})
	}
}

func UpdateHandler(c *gin.Context) {
	var param mdl.Request
	c.BindJSON(&param)

	if param.Id >= 0 {
		err := ctrl.UpdateData(param)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success update data"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "this account doesn't have access to update data"})
	}
}

func DeleteHandler(c *gin.Context) {
	var param mdl.Request
	c.BindJSON(&param)

	if param.Id >= 0 {
		err := ctrl.DeleteData(param)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success delete data"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "this account doesn't have access to delete data"})
	}
}
