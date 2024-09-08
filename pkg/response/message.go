package response

import "net/http"

type Message interface {
}

var msg = map[int]string {
	http.StatusOK: "Success",
	http.StatusBadRequest: "Error Happened",
}