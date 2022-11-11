// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import (
	"github.com/gin-gonic/gin"
	_ "github.com/marmotedu/iam/pkg/validator"

	"iam-based-app/internal/apiserver/controller/v1/user"
	"iam-based-app/internal/apiserver/store/mysql"
)

func initRouter(g *gin.Engine) {
	installController(g)
}

func installController(g *gin.Engine) *gin.Engine {
	storeIns, _ := mysql.GetMySQLFactoryOr(nil)
	v1 := g.Group("/v1")
	{
		// user RESTful resource
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)

			userv1.POST("", userController.Create)
		}
	}
	return g
}
