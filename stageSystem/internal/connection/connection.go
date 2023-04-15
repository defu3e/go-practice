package connection

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stageSystem/internal/constants"
	"stageSystem/internal/functions"
	"stageSystem/internal/result"

	"github.com/gorilla/mux"
)


func handleConnection (w http.ResponseWriter, r *http.Request)  {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "OK")
}

func handleApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	data := result.GetResultData()
	res  := result.PrepareResponce(data)

	json.NewEncoder(w).Encode(res)
}

func ListenAndServeHTTP(port string) {
	r := mux.NewRouter()
	
	r.HandleFunc("/", handleConnection)
	r.HandleFunc("/api", handleApi).Methods("GET", "OPTIONS")
	
	log.Printf("Stage system starting on %s port...\n", port)
	
	err := http.ListenAndServe(":"+port, r)
	functions.CheckErr(err, constants.ERR_INFO_MODE)
}