package api

import "bins/config"

type Api struct {
	Config *config.Config
}

func NewApi(cfg *config.Config) *Api {
	return&Api{Config: cfg}
}