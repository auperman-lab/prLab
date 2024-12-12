package middleware

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"regexp"
)

func LeaderCheckerMiddleware(isLeaderFunc func() bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !isLeaderFunc() {
				http.Error(w, "Access restricted: this node is not the leader", http.StatusForbidden)
				return
			}
			fmt.Printf("\033[31m node passed the leader check \033[0m\n")

			next.ServeHTTP(w, r)
		})
	}
}

func RaftReplicationMiddleware(sendLogToRaft func(logEntry []byte) error) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}
			r.Body.Close()

			data := make([]byte, 0)

			if matched, _ := regexp.MatchString(`/products/\d+/upload`, r.URL.Path); matched {
				data = replicateImage(r)
			} else {
				data = body
			}

			if err := sendLogToRaft(data); err != nil {
				http.Error(w, "Failed to replicate log entry", http.StatusInternalServerError)
				return
			}

			r.Body = io.NopCloser(bytes.NewBuffer(body))
			fmt.Printf("\033[31m node passed the replication check \033[0m\n")

			next.ServeHTTP(w, r)
		})
	}
}

func replicateImage(r *http.Request) []byte {
	var imageName string
	if r.MultipartForm != nil {
		imageName = r.FormValue("image")
	} else {
		vars := mux.Vars(r)
		id := vars["id"]
		imageName = fmt.Sprintf("%s_image", id)
	}

	return []byte(imageName)
}
