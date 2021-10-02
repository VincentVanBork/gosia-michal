package models

import (
	"github.com/gofiber/fiber/v2"
)

func ValidateGuestToken(ctx *fiber.Ctx, s string) (bool, error) {
	if s == "michal" {
		return true, nil
	}
	return false, nil
}
