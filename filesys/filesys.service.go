package filesys

import (
	"encoding/json"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		var fileSystem FileSystem
		var err error
		path, ok := request.URL.Query()["path"]
		if ok && len(path) > 0 {
			fileSystem, err = GetPath(path[0])
		} else {
			fileSystem, err = GetRoot()
		}
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		fileSystemJson, err := json.Marshal(fileSystem)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(fileSystemJson)
	/*case http.MethodPost:
	var child filesys.FileSystem
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyBytes, &child)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fileSystem.Children = append(fileSystem.Children, &child)
	fmt.Printf("Children = %v\n", len(fileSystem.Children))
	writer.WriteHeader(http.StatusCreated) */
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}

}
