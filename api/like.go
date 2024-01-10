package api

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateLike(c *fiber.Ctx) error {
	var u models.Like

	err := c.BodyParser(&u)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parser error: "+err.Error())
	}
	err = a.stg.Like.Create(&u)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusCreated, u)
}

func (a *Api) GetLikeList(c *fiber.Ctx) error {
	var req = models.Like{
		UserId: c.Query("user_id"),
	}

	m, err := a.stg.Like.GetList(&req)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusOK, m)
}

func (a *Api) UpdateLike(c *fiber.Ctx) error {
	var m models.Like
	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parser error during update: "+err.Error())
	}
	id := c.Params("id")
	err = a.stg.Like.Update(&m, &id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, err.Error())
	}
	return handlerResponse(c, http.StatusAccepted, "SUCCESS, UPDATED")
}

func (a *Api) DeleteLike(c *fiber.Ctx) error {
	id := c.Params("id")
	err := a.stg.Like.DeleteLike(&id)
	if err != nil {
		fmt.Println("Error in DeleteLike function: ", err.Error())
	}

	return handlerResponse(c, http.StatusAccepted, "SUCCESS, DELETED")
}
