package httputil

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/constraints"

	pkgapperror "highload-architect/pkg/apperror"
	pkgstrings "highload-architect/pkg/strings"
)

func QueryStruct(c *gin.Context, obj any) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return pkgapperror.ErrIncorrectParameter.WithInternal(err)
	}

	return nil
}

func QueryString(c *gin.Context, name string, typ requiredType) (string, error) {
	param := c.Query(name)

	if param == "" && typ == Required {
		return "", pkgapperror.ErrIncorrectParameter.WithDetails(fmt.Sprintf("path parameter %s is required", name))
	}

	return param, nil
}

func QueryInt[T constraints.Signed](c *gin.Context, name string, typ requiredType) (T, error) {
	strParam, err := QueryString(c, name, typ)
	if err != nil {
		return 0, err
	}

	res, err := pkgstrings.ParseInt[T](strParam)
	if err != nil {
		return 0, pkgapperror.ErrIncorrectParameter.WithDetails(fmt.Sprintf("path parameter %s in invalid: %v", name, err))
	}

	return res, nil
}

func QueryUint[T constraints.Unsigned](c *gin.Context, name string, typ requiredType) (T, error) {
	strParam, err := QueryString(c, name, typ)
	if err != nil {
		return 0, err
	}

	res, err := pkgstrings.ParseUint[T](strParam)
	if err != nil {
		return 0, pkgapperror.ErrIncorrectParameter.WithDetails(fmt.Sprintf("path parameter %s in invalid: %v", name, err))
	}

	return res, nil
}
