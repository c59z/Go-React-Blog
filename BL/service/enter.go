package service

type ServiceGroup struct {
	EsService
	BaseService
}

var ServiceGroupApp = &ServiceGroup{}
