package main

import (
	"goAdmin/config"
	"goAdmin/models"
	"goAdmin/controllers"
)

type Router struct {
	Prefix  string
	Method  string
	Handler controller.EndPointFun
}

func GenerateRoutes(tables map[string]models.GlobalTable) (router map[string]Router) {

	router = make(map[string]Router, 0)
	for k, _ := range tables {
		router["/"+k+"/info"] = Router{
			Prefix:  k,
			Method:  "GET",
			Handler: controller.ShowInfo,
		}
		router["/"+k+"/info/edit"] = Router{
			Prefix:  k,
			Method:  "GET",
			Handler: controller.ShowForm,
		}
		router["/"+k+"/edit"] = Router{
			Prefix:  k,
			Method:  "POST",
			Handler: controller.EditForm,
		}
		router["/"+k+"/info/new"] = Router{
			Prefix:  k,
			Method:  "GET",
			Handler: controller.ShowNewForm,
		}
		router["/"+k+"/new"] = Router{
			Prefix:  k,
			Method:  "POST",
			Handler: controller.NewForm,
		}
		router["/"+k+"/delete"] = Router{
			Prefix:  k,
			Method:  "POST",
			Handler: controller.DeleteData,
		}
	}

	return
}

var GlobalRouter = GenerateRoutes(config.GlobalTableList)
