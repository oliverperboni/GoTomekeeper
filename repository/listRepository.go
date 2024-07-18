package repository

// TODO Tests
import (
	"time"

	"github.com/oliverperboni/GoApi/schemas"
	"gorm.io/gorm"
)

type ListRepository struct {
	DB *gorm.DB
}

func CreateListRepository(db *gorm.DB) ListRepository {
	return ListRepository{DB: db}
}

func (l *ListRepository) CreateList(list *schemas.List) error {
	return l.DB.Create(list).Error
}

func (l *ListRepository) AddBookToList(bookID uint, listID uint) error {
	var listbook schemas.ListBook
	listbook.BookID = bookID
	listbook.ListID = listID
	listbook.CreatedAt = time.Now()
	return l.DB.Create(listbook).Error
}

func (l *ListRepository) RemoveBookToList(bookId uint) error {
	var listbook schemas.ListBook
	l.DB.Find(&listbook, "BookID = ?", bookId)
	return l.DB.Delete(listbook).Error
}

func (l *ListRepository) SeachBookToList(bookId uint) (schemas.Book, error) {
	var listbook schemas.ListBook
	var book schemas.Book
	err := l.DB.Find(&listbook, "Id = ?", bookId).Error
	if err != nil {
		return book, err
	}

	err = l.DB.Find(&book, "Id = ?", listbook.BookID).Error

	if err != nil {
		return book, err
	}

	return book, err
}

func (l *ListRepository) DeleteList(list *schemas.List) error {
	return l.DB.Delete(list).Error
}
