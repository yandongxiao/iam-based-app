// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import "iam-based-app/internal/apiserver/store"

//go:generate mockgen -self_package=iam-based-app/internal/apiserver/service/v1 -destination mock_service.go -package v1 iam-based-app/internal/apiserver/service/v1 Service,UserSrv

// Service defines functions used to return resource interface.
type Service interface {
	Users() UserSrv
}

type service struct {
	store store.Factory
}

// NewService returns Service interface.
func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}
