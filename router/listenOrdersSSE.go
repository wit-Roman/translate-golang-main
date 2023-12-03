package router

import (
	"encoding/json"
	"net/http"
	"strings"

	"translate/types"
)

var messageChannel = make(chan types.ISayHandlerMessage)

//var viewersActiveCount int64 = 0

func ListenHandler(w http.ResponseWriter, r *http.Request) {
	addHeaders(&w, r)
	query := r.URL.Query()

	viewer := query.Get("v")
	if viewer == "" {
		http.Error(w, "not valid", http.StatusBadRequest)
		return
	}

	addHeadersSSE(&w)

	for {
		select {
		case _msg := <-messageChannel:
			//fmt.Println(viewer, _msg)
			//if viewer != _msg.Viewer {
			jsonStructure, err := json.Marshal(_msg)
			if err != nil {
				return
			}

			w.Write(formatSSE("message", string(jsonStructure)))
			w.(http.Flusher).Flush()
			//}

		case <-r.Context().Done():
			// TODO удалить канал?
			return
		}
	}
}

func formatSSE(event string, data string) []byte {
	eventPayload := "event: " + event + "\n"
	dataLines := strings.Split(data, "\n")
	for _, line := range dataLines {
		eventPayload = eventPayload + "data: " + line + "\n"
	}
	return []byte(eventPayload + "\n")
}

func sayHandler(mesType string, session types.ITokenData, data any) {
	go func() {
		messageChannel <- types.ISayHandlerMessage{mesType, session.Viewer, data}
	}()
}
