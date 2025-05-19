package httputil

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/constraints"

	pkgapperror "highload-architect/pkg/apperror"
	pkgstrings "highload-architect/pkg/strings"
)

func ParamString(c *gin.Context, name string, typ requiredType) (string, error) {
	param := c.Param(name)

	if param == "" && typ == Required {
		return "", pkgapperror.ErrIncorrectParameter.WithDetails(fmt.Sprintf("path parameter %s is required", name))
	}

	return param, nil
}

func ParamInt[T constraints.Signed](c *gin.Context, name string, typ requiredType) (T, error) {
	strParam, err := ParamString(c, name, typ)
	if err != nil {
		return 0, err
	}

	res, err := pkgstrings.ParseInt[T](strParam)
	if err != nil {
		return 0, pkgapperror.ErrIncorrectParameter.WithDetails(fmt.Sprintf("path parameter %s in invalid: %v", name, err))
	}

	return res, nil
}

func ParamUint[T constraints.Unsigned](c *gin.Context, name string, typ requiredType) (T, error) {
	strParam, err := ParamString(c, name, typ)
	if err != nil {
		return 0, err
	}

	res, err := pkgstrings.ParseUint[T](strParam)
	if err != nil {
		return 0, pkgapperror.ErrIncorrectParameter.WithDetails(fmt.Sprintf("path parameter %s in invalid: %v", name, err))
	}

	return res, nil
}
