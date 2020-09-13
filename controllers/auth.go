package controllers

import (
	"github.com/gofiber/fiber"
)

// Login authenticate users
func Login(ctx *fiber.Ctx) {
	ctx.Send("Login")
}

// Register user register here
func Register(ctx *fiber.Ctx) {
	ctx.Send("Register")
}
