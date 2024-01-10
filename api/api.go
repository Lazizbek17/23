package api

import (
  "app/config"
  "app/storage"

  //"database/sql"

  "github.com/gofiber/fiber/v2"
)

type Api struct {
  f   *fiber.App
  stg *storage.Storage
}

func NewApi(stg *storage.Storage) *Api {

  f := fiber.New()
  a := &Api{
    f:   f,
    stg: stg,
  }
 
  {
	l := f.Group("like")
	l.Get("/",a.GetLikeList)
	l.Post("/",a.CreateLike)
	l.Delete("/",a.DeleteLike)
	l.Put("/:id",a.UpdateLike)

}
return a
}
func (a *Api ) RUN(){
	if err := a.f.Listen(config.Port);err != nil{
		panic(err)
	}
}