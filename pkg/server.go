package bare_go

import (
    "github.com/go-chi/chi/v5"
    "net/http"
    "github.com/ruby-network/bare-go/internal/routes"
)

func HandleBareChi(directory string, chiRouter *chi.Mux) *chi.Mux {
    router := routes.ChiRouter(directory, chiRouter)
    return router
}

func HandleBareHttp(directory string, router *http.ServeMux) http.Handler {
    router = routes.NetHttp(directory, router)
    return router
}