package infrastructure

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MuxRouter struct{}

var (
	MuxRouterInstance = mux.NewRouter()
)

func NewMuxRouter() RouterIF {
	return &MuxRouter{}
}

type RouterIF interface {
	Post(uri string, f func(w http.ResponseWriter, r *http.Request))
	ServeEndpoint(port string)
}

func (m MuxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	MuxRouterInstance.HandleFunc(uri, f).Methods("POST")
}

func (m MuxRouter) ServeEndpoint(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	_ = http.ListenAndServe(port, MuxRouterInstance)
}
