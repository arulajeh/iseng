package utils

import (
	"pertama_go/types"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetLimit(ctx *fiber.Ctx) int {
	limit := ctx.Get("limit")
	if limit == "" {
		return 10
	}
	return ConvertStringToInt(limit)
}

func GetPage(ctx *fiber.Ctx) int {
	page := ctx.Get("page")
	if page == "" {
		return 1
	}
	return ConvertStringToInt(page)
}

func ConvertStringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func GetPagination(ctx *fiber.Ctx) *types.Pagination {
	return &types.Pagination{
		Page:  GetPage(ctx),
		Limit: GetLimit(ctx),
	}
}

func GetOffset(pagination *types.Pagination) int {
	return (pagination.Page - 1) * pagination.Limit
}
