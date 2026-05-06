package interfaces

import "github.com/nickchervov/go-library-system/internal/models"

type Storage interface {
	SaveToFile(filename string) error
	LoadFromFile(filename string) (*models.Library, error)
}
