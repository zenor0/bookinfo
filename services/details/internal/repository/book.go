package repository

import (
	"github.com/zenor0/bookinfo/services/details/internal/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepository) GetByID(id uint) (*model.Book, error) {
	var book model.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) GetByISBN(isbn string) (*model.Book, error) {
	var book model.Book
	err := r.db.Where("isbn = ?", isbn).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) List(page, pageSize int) ([]model.Book, int64, error) {
	var books []model.Book
	var total int64

	err := r.db.Model(&model.Book{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.Offset(offset).Limit(pageSize).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

func (r *BookRepository) Update(book *model.Book) error {
	return r.db.Save(book).Error
}

func (r *BookRepository) Delete(id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}

func (r *BookRepository) Search(query string) ([]model.Book, error) {
	var books []model.Book
	err := r.db.Where("title ILIKE ? OR author ILIKE ? OR isbn LIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%").
		Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
} 