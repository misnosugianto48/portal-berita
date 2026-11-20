package repositories

import (
	"context"
	"errors"
	"portal-berita/internal/core/domains/entity"
	"portal-berita/internal/core/domains/models"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/gorm"
)

var ErrCategoryNotFound = errors.New("category not found")
var ErrDuplicated = errors.New("duplicate same key")

type CategoryRepository interface {
	CreateCategory(ctx context.Context, req entity.CategoryEntity) (*entity.CategoryEntity, error)
	GetCategories(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryById(ctx context.Context, id int64) (*entity.CategoryEntity, error)
	UpdateCategoryById(ctx context.Context, id int64, req entity.CategoryEntity) error
	DeleteCategoryById(ctx context.Context, id int64) error
}

type categoryRepository struct {
	db *gorm.DB
}

// CreateCategory implements [CategoryRepository].
func (c *categoryRepository) CreateCategory(ctx context.Context, req entity.CategoryEntity) (*entity.CategoryEntity, error) {

	category := models.Category{
		Title:       req.Title,
		Slug:        req.Slug,
		CreatedById: req.CreatedById,
	}

	if err := c.db.WithContext(ctx).Create(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ErrDuplicated
		}

		code := "[REPOSITORY] CreateCategory - 1"
		log.Errorw(code, err)
		return nil, err
	}

	return &entity.CategoryEntity{
		Id:          category.Id,
		Title:       category.Title,
		Slug:        category.Slug,
		CreatedById: category.CreatedById,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}, nil
}

// DeleteCategoryById implements [CategoryRepository].
func (c *categoryRepository) DeleteCategoryById(ctx context.Context, id int64) error {
	// [PSEUDOCODE]
	// 1. Eksekusi query delete berdasarkan ID:
	//    res := c.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Category{})
	// 2. Validasi error query:
	//    if res.Error != nil -> catat log error, return res.Error
	// 3. Validasi apakah data benar-benar ada dan terhapus:
	//    if res.RowsAffected == 0 -> return ErrCategoryNotFound
	// 4. Jika sukses -> return nil.
	//
	// [TODO]
	// - Perhatikan efek relasi ON DELETE CASCADE di database (tabel contents berelasi ke categories).
	//   Menghapus kategori akan otomatis menghapus seluruh artikel terkait di database.
	panic("unimplemented")
}

// GetCategories implements [CategoryRepository].
func (c *categoryRepository) GetCategories(ctx context.Context) ([]entity.CategoryEntity, error) {
	// [PSEUDOCODE]
	// 1. Siapkan slice penampung model:
	//    var modelCategories []models.Category
	// 2. Ambil data dari DB dengan relasi pembuat (User) dan urutkan berdasarkan waktu terbaru:
	//    err := c.db.WithContext(ctx).Preload("User").Order("created_at DESC").Find(&modelCategories).Error
	//    if err != nil -> catat log error, return nil, err
	// 3. Petakan slice model ke slice domain entity:
	//    categories := make([]entity.CategoryEntity, 0, len(modelCategories))
	//    for _, item := range modelCategories:
	//        categories = append(categories, entity.CategoryEntity{
	//            Id:          item.Id,
	//            Title:       item.Title,
	//            Slug:        item.Slug,
	//            CreatedById: item.CreatedById,
	//            User: entity.UserEntity{
	//                Id:    item.User.Id,
	//                Name:  item.User.Name,
	//                Email: item.User.Email,
	//            },
	//            CreatedAt: item.CreatedAt,
	//            UpdatedAt: item.UpdatedAt,
	//        })
	// 4. Return categories, nil.
	//
	// [TODO]
	// - Tambahkan dukungan parameter pagination (limit & offset) jika daftar kategori bertambah banyak.
	// - Pastikan selalu mengembalikan empty slice [] bukan nil jika record kosong agar serialize JSON aman.
	panic("unimplemented")
}

// GetCategoryById implements [CategoryRepository].
func (c *categoryRepository) GetCategoryById(ctx context.Context, id int64) (*entity.CategoryEntity, error) {
	var modelCategory models.Category

	if err := c.db.WithContext(ctx).Preload("User").Where("id = ?", id).First(&modelCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCategoryNotFound
		}

		code := "[REPOSITORY] GetCategoryById - 1"
		log.Errorw(code, "error", err)
		return nil, err
	}

	response := entity.CategoryEntity{
		Id:          modelCategory.Id,
		Title:       modelCategory.Title,
		Slug:        modelCategory.Slug,
		CreatedById: modelCategory.CreatedById,
		User: entity.UserEntity{
			Id:   modelCategory.User.Id,
			Name: modelCategory.User.Name,
		},
		CreatedAt: modelCategory.CreatedAt,
		UpdatedAt: modelCategory.UpdatedAt,
	}

	return &response, nil
}

// UpdateCategoryById implements [CategoryRepository].
func (c *categoryRepository) UpdateCategoryById(ctx context.Context, id int64, req entity.CategoryEntity) error {
	// [PSEUDOCODE]
	// 1. Tentukan kolom yang diizinkan untuk di-update (hindari menimpa CreatedById dan CreatedAt):
	//    updateData := map[string]interface{}{
	//        "title":      req.Title,
	//        "slug":       req.Slug,
	//        "updated_at": time.Now(),
	//    }
	// 2. Eksekusi query update:
	//    res := c.db.WithContext(ctx).Model(&models.Category{}).Where("id = ?", id).Updates(updateData)
	// 3. Validasi error query:
	//    if res.Error != nil -> catat log error, return res.Error
	// 4. Validasi apakah ada baris yang ter-update:
	//    if res.RowsAffected == 0 -> return ErrCategoryNotFound
	// 5. Jika sukses -> return nil.
	//
	// [TODO]
	// - Tangani potensi duplicate key error jika slug baru sudah digunakan kategori lain.
	// - Gunakan .Select("title", "slug") atau map eksplisit agar field sensitif tidak tertimpa.
	panic("unimplemented")
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
