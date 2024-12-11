package middleware

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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

			if err := sendLogToRaft(body); err != nil {
				http.Error(w, "Failed to replicate log entry", http.StatusInternalServerError)
				return
			}

			r.Body = io.NopCloser(bytes.NewBuffer(body))
			fmt.Printf("\033[31m node passed the replication check \033[0m\n")

			next.ServeHTTP(w, r)
		})
	}
}
