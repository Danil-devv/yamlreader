package yamlreader

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type reader struct {
	data []byte
}

func (c *reader) LoadYAML(path string) error {
	if filepath.Ext(path) != ".yaml" {
		return fmt.Errorf("%s is not path to yaml file", path)
	}

	var err error
	c.data, err = os.ReadFile(path)
	return err
}

func (c *reader) ReadYAML(s any) error {
	return yaml.Unmarshal(c.data, s)
}

type LoadReader interface {
	// LoadYAML загружает данные из файла .yaml по пути path
	// во внутренне хранилище.
	//
	// Этот метод не преобразует данные в структуру go
	LoadYAML(string) error

	// ReadYAML преобразует данные из внутреннего хранилища
	// в структуру go
	ReadYAML(any) error
}

// New возвращает LoadReader
//
// Файл находящийся по пути path сразу будет считан при
// вызове данной функции. Поэтому можно не вызывать
// метод LoadYAML и сразу преобразовывать данные в структуру
// go при помощи ReadYAML
func New(path string) (LoadReader, error) {
	c := &reader{}
	err := c.LoadYAML(path)
	if err != nil {
		return nil, err
	}
	return c, nil
}
