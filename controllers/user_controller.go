package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rischoo/project-management/models"
	"github.com/rischoo/project-management/services"
	"github.com/rischoo/project-management/utils"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data!", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Gagal registrasi user!", err.Error())
	}

	return utils.Created(ctx, "User berhasil didaftarkan!", user)
}
