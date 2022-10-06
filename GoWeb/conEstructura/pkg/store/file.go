package store

import (
	"encoding/json"
	"os"
)

//package que se va a encargar de manipular el archivo json

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string //un alias

const(
	FileType Type = "file"
	MongoType Type = "mongo"
)

func NewStore(store Type, fileName string) Store {
	switch store{
	case FileType: 
		return &fileStore{fileName}
	}

	return nil
}

type fileStore struct{
	FilePath string
}

func (fs *fileStore) Write(data interface{}) error {
	//identamos
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}