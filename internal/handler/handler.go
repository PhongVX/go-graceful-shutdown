package handler

import (
	"log"
	"net/http"
	"time"

	"graceful-shutdown/pkg/utils/response"

	"github.com/gorilla/mux"
)

func testGracefulShutDown(res http.ResponseWriter, req *http.Request) {
	time.Sleep(10 * time.Second)
	log.Println("testGracefulShutdown job completed")
	response.ResponseWithJSON(res, 200, map[string]interface{}{"status": "completed"})
}

func New(r *mux.Router) {
	r.HandleFunc("/test-graceful-shutdown", testGracefulShutDown).Methods(http.MethodPost)
}
