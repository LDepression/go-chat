/**
 * @Author: lenovo
 * @Description:
 * @File:  base
 * @Version: 1.0.0
 * @Date: 2023/03/25 13:36
 */

package base

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-chat/internal/global"
	"net/http"
	"strings"
)

func RemoveTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": RemoveTopStruct(errs.Translate(global.Trans)),
	})
}
