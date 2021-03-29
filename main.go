package main

import (
	"github.com/jhallat/udt-devtools/cors"
	"github.com/jhallat/udt-devtools/filesys"
	"net/http"
)

func formHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Forms not yet implemented"))
}

func main() {
	http.Handle("/filesystem", cors.MiddlewareHandler(http.HandlerFunc(filesys.Handler)))
	http.Handle("/form", cors.MiddlewareHandler(http.HandlerFunc(formHandler)))
	http.ListenAndServe(":5000", nil)
}

/* func buildForm() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input Type: (T)ext, (P)assword")
	inputTypeReq, _ := reader.ReadString('\n')
	inputTypeCommands := strings.Split(strings.Trim(inputTypeReq, "\n")," ")
	inputTypeCommands = clearSpaces(inputTypeCommands)
	if len(inputTypeCommands) == 1 {

	}
	if len(inputTypeCommands) == 2 {

	}
	for index, command := range inputTypeCommands {
		fmt.Println(index, command)
	}
	formBuilder := form.NewFormBuilder(form.Reactive)
	formHtml := formBuilder.Build()
	fmt.Println(formHtml)
}

func clearSpaces(values []string) []string {
	cleared := []string{}
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			cleared = append(cleared, value)
		}
	}
	return cleared
} */