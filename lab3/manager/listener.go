package manager

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Listener struct {
	server   *http.Server
	consumer *Consumer
}

type ChangeLeaderRequest struct {
	NewLeaderURL string `json:"new_leader_url"`
}

func NewListener(port string, consumer *Consumer) *Listener {
	l := &Listener{}
	mux := http.NewServeMux()
	mux.HandleFunc("/changeLeader", l.changeLeaderHandler)
	l.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
	l.consumer = consumer

	return l
}

func (l *Listener) Start() error {

	slog.Info("Listening on", "addr", l.server.Addr)

	return l.server.ListenAndServe()
}

func (l *Listener) changeLeaderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ChangeLeaderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.NewLeaderURL == "" {
		http.Error(w, "New leader URL is required", http.StatusBadRequest)
		return
	}

	(*l.consumer).UpdateURL(req.NewLeaderURL)

	w.WriteHeader(http.StatusOK)
	fmt.Printf("Leader URL updated successfully : %s\n", req.NewLeaderURL)
	w.Write([]byte("Leader URL updated successfully"))
}
