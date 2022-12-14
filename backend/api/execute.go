package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/gorilla/websocket"
)

func (c *Controller) enqueueExecution(w http.ResponseWriter, r *http.Request) {
	// CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	if r.Method == http.MethodOptions {
		return
	}

	// decode submission data
	var submission models.ExecSubmission
	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// respond to submission
	resp, err := json.Marshal(submission)
	if err != nil {
		log.Printf("HTTP %s", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Write(resp) // echo back the submission
	c.scheduler.EnqueueExecution(submission)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *Controller) communicateWithExec(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte("labdien"))
	conn.Close()
}
