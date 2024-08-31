package book

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// List returns all books.
func (r *Repository) List() (Books, error) {
	books := make([]*Book, 0)
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// Create creates a new book.
func (r *Repository) Create(book *Book) (*Book, error) {
	if err := r.db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// Read returns a book by the given id.
func (r *Repository) Read(id uuid.UUID) (*Book, error) {
	book := &Book{}
	if err := r.db.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *Repository) Update(book *Book) (int64, error) {
	result := r.db.
		Model(&Book{}).
		Select("Title", "Author", "PublishedDate", "ImageURL", "Description").
		Where("id = ?", book.ID).
		Updates(map[string]interface{}{
			"Title":         book.Title,
			"Author":        book.Author,
			"PublishedDate": book.PublishedDate,
			"ImageURL":      book.ImageURL,
			"Description":   book.Description,
		})

	return result.RowsAffected, result.Error
}

// Delete deletes a book by the given id and returns the number of deleted rows and a potential error.
func (r *Repository) Delete(id uuid.UUID) (int64, error) {
	result := r.db.Unscoped().Where("id = ?", id).Delete(&Book{})

	return result.RowsAffected, r.db.Error
}
