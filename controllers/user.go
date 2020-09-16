package controllers

import (
	"fmt"

	"github.com/ezedinff/appzoid/models"
	"github.com/gofiber/fiber"
)

// FinaAllUsers retrun all users
func FinaAllUsers(ctx *fiber.Ctx) {
	fmt.Println(models.GetUsers())
	ctx.Send("Users retrie....")
}
