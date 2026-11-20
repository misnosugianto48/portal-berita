package handlers

import (
	"errors"
	"portal-berita/internal/adapters/handlers/responses"
	"portal-berita/internal/adapters/repositories"
	"portal-berita/internal/core/domains/entity"
	"portal-berita/internal/core/services"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

type CategoryHandlerInterface interface {
	CreateCategory(c fiber.Ctx) error
	GetCategories(c fiber.Ctx) error
	GetCategoryById(c fiber.Ctx) error
	UpdateCategoryById(c fiber.Ctx) error
	DeleteCategoryById(c fiber.Ctx) error
}

type categoryHandler struct {
	categoryServ services.CategoryService
}

// CreateCategory implements [CategoryHandlerInterface].
func (c *categoryHandler) CreateCategory(ctx fiber.Ctx) error {
	// [PSEUDOCODE - Alur Handler HTTP Fiber v3]
	// 1. Inisialisasi request DTO:
	//    var req requests.CreateCategoryRequest
	// 2. Parsing body JSON dan jalankan auto-validasi:
	//    if err := c.Bind().JSON(&req); err != nil {
	//        return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponseDefault{
	//            Meta: responses.Meta{Status: false, Message: err.Error()},
	//        })
	//    }
	// 3. Ambil data user yang sedang login dari context JWT middleware:
	//    userClaims := c.Locals("user").(*entity.JWTClaims)
	//    userId := int64(userClaims.UserId)
	// 4. Petakan DTO ke domain entity:
	//    categoryEntity := entity.CategoryEntity{
	//        Title:       req.Title,
	//        CreatedById: userId,
	//    }
	// 5. Panggil service:
	//    if err := c.categoryServ.CreateCategory(c.Context(), categoryEntity); err != nil {
	//        return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponseDefault{...})
	//    }
	// 6. Kembalikan respons sukses HTTP 201 Created:
	//    return c.Status(fiber.StatusCreated).JSON(responses.DefaultSuccessResponse{...})
	//
	// [TODO]
	// - Buat request struct `CreateCategoryRequest` di `internal/adapters/handlers/requests/category_request.go`
	//   dengan tag validasi seperti `validate:"required,min=3,max=200"`.
	// - Pasang middleware auth `middlewares.CheckToken()` pada rute `POST /categories`.
	panic("unimplemented")
}

// DeleteCategoryById implements [CategoryHandlerInterface].
func (c *categoryHandler) DeleteCategoryById(ctx fiber.Ctx) error {
	// [PSEUDOCODE - Alur Handler HTTP Fiber v3]
	// 1. Ekstrak parameter ID dari URL path:
	//    idParam := c.Params("id")
	//    id, err := strconv.ParseInt(idParam, 10, 64)
	//    if err != nil {
	//        return c.Status(fiber.StatusBadRequest).JSON(...)
	//    }
	// 2. Panggil service untuk menghapus data:
	//    if err := c.categoryServ.DeleteCategoryById(c.Context(), id); err != nil {
	//        if errors.Is(err, repositories.ErrCategoryNotFound) {
	//            return c.Status(fiber.StatusNotFound).JSON(...)
	//        }
	//        return c.Status(fiber.StatusInternalServerError).JSON(...)
	//    }
	// 3. Kembalikan respons sukses HTTP 200 OK:
	//    return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{Message: "Category deleted successfully"})
	//
	// [TODO]
	// - Proteksi endpoint ini dengan middleware auth `CheckToken()`.
	panic("unimplemented")
}

// GetCategories implements [CategoryHandlerInterface].
func (c *categoryHandler) GetCategories(ctx fiber.Ctx) error {
	// [PSEUDOCODE - Alur Handler HTTP Fiber v3]
	// 1. Panggil service untuk mengambil daftar kategori:
	//    categories, err := c.categoryServ.GetCategories(c.Context())
	//    if err != nil {
	//        return c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponseDefault{...})
	//    }
	// 2. Petakan slice entity ke slice response DTO:
	//    resData := make([]responses.CategoryResponse, 0, len(categories))
	//    for _, item := range categories {
	//        resData = append(resData, mapToCategoryResponse(item))
	//    }
	// 3. Kembalikan respons HTTP 200 OK dengan format standar:
	//    return c.Status(fiber.StatusOK).JSON(responses.CategoryListResponse{
	//        Meta: responses.Meta{Status: true, Message: "Success get categories"},
	//        Data: resData,
	//    })
	//
	// [TODO]
	// - Endpoint ini bisa dibuka sebagai public endpoint (tanpa middleware auth).
	panic("unimplemented")
}

// GetCategoryById implements [CategoryHandlerInterface].
func (c *categoryHandler) GetCategoryById(ctx fiber.Ctx) error {
	// TODO: Endpoint kategori ini diakses oleh Admin menggunakan token JWT.
	// 1. Pastikan rute ini diproteksi oleh middleware auth: `middlewares.CheckToken()` (di app.go).
	// 2. Data admin login dapat diambil dari context locals:
	//    userClaims := ctx.Locals("user").(*entity.JWTClaims)
	//    adminId := int64(userClaims.UserId)
	var code string
	var errorResponse responses.ErrorResponseDefault
	res := responses.CategoryDetailResponse{}

	userClaims, ok := ctx.Locals("user").(*entity.JWTClaims)

	if !ok || userClaims == nil || userClaims.UserId == 0 {
		errorResponse.Meta.Status = false
		errorResponse.Meta.Message = "Unauthorized access"
		return ctx.Status(fiber.StatusUnauthorized).JSON(errorResponse)
	}

	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		code = "[HANDLER] GetCategory - 1"
		log.Error().Err(err).Msg(code)
		errorResponse.Meta.Status = false
		errorResponse.Meta.Message = "Invalid Id"
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse)

	}

	result, err := c.categoryServ.GetCategoryById(ctx.Context(), id)
	if err != nil {
		code = "[HANDLER] GetCategory - 2"
		log.Error().Err(err).Msg(code)
		errorResponse.Meta.Status = false

		if errors.Is(err, repositories.ErrCategoryNotFound) {
			errorResponse.Meta.Message = "Category not found"
			return ctx.Status(fiber.StatusNotFound).JSON(errorResponse)
		}

		errorResponse.Meta.Message = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse)
	}

	res.Meta.Status = true
	res.Meta.Message = "Successfully retrieved"
	res.Data.Id = result.Id
	res.Data.Title = result.Title
	res.Data.Slug = result.Slug
	res.Data.User.Id = result.User.Id
	res.Data.User.Name = result.User.Name
	res.Data.CreatedAt = result.CreatedAt
	res.Data.UpdatedAt = result.UpdatedAt

	return ctx.JSON(res)

}

// UpdateCategoryById implements [CategoryHandlerInterface].
func (c *categoryHandler) UpdateCategoryById(ctx fiber.Ctx) error {
	// [PSEUDOCODE - Alur Handler HTTP Fiber v3]
	// 1. Ambil ID dari URL path:
	//    id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	//    if err != nil { return c.Status(fiber.StatusBadRequest).JSON(...) }
	// 2. Binding dan validasi request body JSON:
	//    var req requests.UpdateCategoryRequest
	//    if err := c.Bind().JSON(&req); err != nil {
	//        return c.Status(fiber.StatusBadRequest).JSON(...)
	//    }
	// 3. Panggil service untuk melakukan update:
	//    categoryEntity := entity.CategoryEntity{ Title: req.Title }
	//    if err := c.categoryServ.UpdateCategoryById(c.Context(), id, categoryEntity); err != nil {
	//        if errors.Is(err, repositories.ErrCategoryNotFound) {
	//            return c.Status(fiber.StatusNotFound).JSON(...)
	//        }
	//        return c.Status(fiber.StatusInternalServerError).JSON(...)
	//    }
	// 4. Kembalikan respons sukses HTTP 200 OK.
	//
	// [TODO]
	// - Proteksi endpoint dengan middleware auth `CheckToken()`.
	// - Buat struct `UpdateCategoryRequest` di package requests.
	panic("unimplemented")
}

func NewCategoryHandler(categoryServ services.CategoryService) CategoryHandlerInterface {
	return &categoryHandler{categoryServ: categoryServ}
}
