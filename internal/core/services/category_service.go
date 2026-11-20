package services

import (
	"context"
	"errors"
	"portal-berita/internal/adapters/repositories"
	"portal-berita/internal/core/domains/entity"
	"portal-berita/libs/conv"

	"github.com/gofiber/fiber/v3/log"
)

var ErrInvalidId = errors.New("invalid id")

type CategoryService interface {
	CreateCategory(ctx context.Context, req entity.CategoryEntity) (*entity.CategoryEntity, error)
	GetCategories(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryById(ctx context.Context, id int64) (*entity.CategoryEntity, error)
	UpdateCategoryById(ctx context.Context, id int64, req entity.CategoryEntity) error
	DeleteCategoryById(ctx context.Context, id int64) error
}

type categoryService struct {
	categoryrepo repositories.CategoryRepository
}

// CreateCategory implements [CategoryService].
func (c *categoryService) CreateCategory(ctx context.Context, req entity.CategoryEntity) (*entity.CategoryEntity, error) {
	// [PSEUDOCODE]
	var err error
	var code string

	if req.CreatedById <= 0 {
		code = "[SERVICE] CreateCategory - 1"
		log.Errorw(code, "error", err)
		return nil, ErrInvalidId
	}

	if req.Slug == "" {
		req.Slug = conv.GenerateSlug(req.Title)
	}

	category, err := c.categoryrepo.CreateCategory(ctx, req)

	if err != nil {
		code = "[SERVICE] CreateCategory - 2"
		log.Errorw(code, "error", err)
		return nil, err
	}

	return category, nil
}

// DeleteCategoryById implements [CategoryService].
func (c *categoryService) DeleteCategoryById(ctx context.Context, id int64) error {
	// [PSEUDOCODE]
	// 1. Validasi ID kategori:
	//    if id <= 0 -> return errors.New("invalid category id")
	// 2. Cek apakah kategori ada di database sebelum dihapus:
	//    _, err := c.categoryrepo.GetCategoryById(ctx, id)
	//    if err != nil -> return err (misal ErrCategoryNotFound)
	// 3. Eksekusi penghapusan di repository:
	//    err = c.categoryrepo.DeleteCategoryById(ctx, id)
	//    if err != nil -> catat log error, return err
	// 4. Return nil jika berhasil.
	//
	// [TODO]
	// - Cek aturan bisnis: periksa apakah kategori ini sedang memiliki artikel aktif sebelum dihapus.
	//   Hal ini untuk menghindari terhapusnya konten berita penting secara otomatis via database CASCADE.
	panic("unimplemented")
}

// GetCategories implements [CategoryService].
func (c *categoryService) GetCategories(ctx context.Context) ([]entity.CategoryEntity, error) {
	// [PSEUDOCODE]
	// 1. Ambil seluruh data kategori dari repository:
	//    categories, err := c.categoryrepo.GetCategories(ctx)
	//    if err != nil -> catat log error, return nil, err
	// 2. Terapkan aturan bisnis atau transformasi data jika diperlukan.
	// 3. Return categories, nil.
	//
	// [TODO]
	// - Pertimbangkan implementasi caching (misal via Redis atau in-memory cache) karena data kategori jarang berubah tapi sering dibaca.
	// - Jika di masa depan kategori berjumlah ratusan, tambahkan parameter pagination (page, perPage).
	panic("unimplemented")
}

// GetCategoryById implements [CategoryService].
func (c *categoryService) GetCategoryById(ctx context.Context, id int64) (*entity.CategoryEntity, error) {
	var err error
	var code string

	if id <= 0 {
		code = "[SERVICE] GetCategoryById - 1"
		log.Errorw(code, "error", ErrInvalidId)
		return nil, ErrInvalidId
	}

	category, err := c.categoryrepo.GetCategoryById(ctx, id)
	if err != nil {
		if errors.Is(err, repositories.ErrCategoryNotFound) {
			return nil, repositories.ErrCategoryNotFound
		}
		code = "[SERVICE] GetCategoryById - 2"
		log.Errorw(code, "error", err)
		return nil, err
	}

	return category, nil
}

// UpdateCategoryById implements [CategoryService].
func (c *categoryService) UpdateCategoryById(ctx context.Context, id int64, req entity.CategoryEntity) error {
	// [PSEUDOCODE]
	// 1. Validasi parameter ID dan field wajib:
	//    if id <= 0 -> return errors.New("invalid category id")
	//    if req.Title == "" -> return errors.New("title is required")
	// 2. Cek apakah record kategori yang ingin diperbarui benar-benar ada:
	//    existing, err := c.categoryrepo.GetCategoryById(ctx, id)
	//    if err != nil -> return err
	// 3. Auto-update slug jika judul berubah atau slug kosong:
	//    if req.Slug == "" || req.Title != existing.Title {
	//        req.Slug = conv.GenerateSlug(req.Title)
	//    }
	// 4. Panggil repository untuk menyimpan perubahan:
	//    err = c.categoryrepo.UpdateCategoryById(ctx, id, req)
	//    if err != nil -> catat log error, return err
	// 5. Return nil jika sukses.
	//
	// [TODO]
	// - Pastikan autorisasi user: cek apakah user memiliki hak akses (role Admin atau author pemilik kategori).
	// - Hindari slug bentrok jika judul baru menghasilkan slug yang sudah dimiliki oleh kategori lain.
	panic("unimplemented")
}

func NewCategoryService(categoryRepo repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryrepo: categoryRepo}
}
