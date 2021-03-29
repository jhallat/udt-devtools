package cors

import "net/http"

func MiddlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// TODO refine the CORS logic
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		handler.ServeHTTP(writer, request)
	})
}
