package api

import "blog-go/service"

type ApiGroup struct {
	BaseApi
}

var ApiGroupApp = &ApiGroup{}

var baseService = service.ServiceGroupApp.BaseService
