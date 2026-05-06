package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nickchervov/go-library-system/internal/models"
)

func SaveToFile(filename string) error {
	data, err := json.MarshalIndent(models.Lib, "", "    ")
	if err != nil {
		return fmt.Errorf("ошибка при сериализации библиотеки: %v", err)
	}
	if err := os.WriteFile(filename, data, 0755); err != nil {
		return fmt.Errorf("ошибка при сохранении библиотеки в файл: %v", err)
	}
	return nil
}

func LoadFromFile(filename string) (*models.Library, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0755)
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузки данных из файла: %v", err)
	}
	defer file.Close()

	jsonData := make([]byte, 0)
	file.Read(jsonData)

	if len(jsonData) == 0 {
		return models.NewLibrary(), nil
	}

	var decodedLibrary *models.Library
	if err := json.Unmarshal(jsonData, &decodedLibrary); err != nil {
		return nil, fmt.Errorf("ошибка при десериализации данных: %v", err)
	}
	return decodedLibrary, nil
}
