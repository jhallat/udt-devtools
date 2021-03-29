package filesys

import (
	"os"
)

type FileSystem struct {
	Directory bool         `json:"isDirectory"`
	Name      string       `json:"name"`
	Children  []FileSystem `json:"children"`
}

func GetRoot() (FileSystem, error) {
	root, err := os.UserHomeDir()
	if err != nil {
		return FileSystem{}, err
	}
	return FileSystem{
		Directory: true,
		Name: root,
		Children: findChildren(root),
	}, nil
}

func GetPath(path string) (FileSystem, error) {
	if path == "" {
		return GetRoot()
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return FileSystem{}, err
	}
	if !fileInfo.IsDir() {
		return FileSystem{
			Directory: false,
			Name: fileInfo.Name(),
			Children: make([]FileSystem, 0),
		}, nil
	}
	return FileSystem{
		Directory: true,
		Name: fileInfo.Name(),
		Children: findChildren(path),
	}, nil
}

func findChildren(root string) []FileSystem {
	var children []FileSystem
	file, err := os.Open(root)
	if err != nil {
		return children
	}
	fileInfo, err := file.Readdir(-1)
	file.Close()
	if err != nil {
		return children
	}
	for _, info := range fileInfo {
		children = append(children,
			FileSystem{
				Directory: info.IsDir(),
				Name: info.Name(),
				Children: make([]FileSystem, 0)})
	}
	return children
}