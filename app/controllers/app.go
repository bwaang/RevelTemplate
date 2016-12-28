package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Health() revel.Result {
  return c.RenderJson(map[string]interface{}{
    "Status": "Ok",
  })
}

func (c App) Version() revel.Result {
  return c.RenderJson(map[string]interface{} {
    "Version": "0.0.0",
  })
}
