package api
import "github.com/gofiber/fiber/v2"

func handlerResponse(c *fiber.Ctx,code int, response any ) error{
	return c.JSON(response)
}