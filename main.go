package main

import (
	"core/config"
	"core/routes"
	"master/resource"
)

func main() {
	cf := resource.ResourceConfig{
		IsEnablePostgres: true,
	}
	r, err := resource.Init(cf)
	if err != nil {
		return
	}
	defer r.Close()

	app := routes.GetEngine(r)
	app.Run(config.AppHost + ":" + config.AppPort)
}
