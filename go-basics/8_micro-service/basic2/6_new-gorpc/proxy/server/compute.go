package server

import (
	"go-basics/8_micro-service/basic2/6_new-gorpc/proxy/service"
	"net/rpc"
)

type ComputeService interface {
	Add(req service.RequestBody, res *service.ResponseBody) error
	Minus(req service.RequestBody, res *service.ResponseBody) error
	Multiply(req service.RequestBody, res *service.ResponseBody) error
	Div(req service.RequestBody, res *service.ResponseBody) error
}

func RegisterComputeService(svc ComputeService) error {
	return rpc.RegisterName(service.ComputeName, svc)
}
