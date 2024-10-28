package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	restadapter "github.com/quocthinhle/file-manager-api/command-ingress/adapter/in/rest"
	pgoutadapter "github.com/quocthinhle/file-manager-api/command-ingress/adapter/out/pg"
	"github.com/quocthinhle/file-manager-api/command-ingress/application/domain/usecase"
	chassispg "github.com/quocthinhle/file-manager-api/pkg/chassis/pg"
)

const (
	baseURL = "/api/file-manager/v1"
)

func main() {
	pgConnection, err := chassispg.NewPool(context.Background())
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	pgOutputAdapter := pgoutadapter.NewNodeOutputAdapter(pgConnection)
	fetchContentUseCase := usecase.NewFetchContentUseCase(pgOutputAdapter)
	restAdapter := restadapter.NewFileManagerRestAdapter(fetchContentUseCase)
	var middlewares []nethttp.StrictHTTPMiddlewareFunc

	strictHandler := restadapter.NewStrictHandler(restAdapter, middlewares)
	httpHandler := restadapter.HandlerFromMuxWithBaseURL(strictHandler, router, baseURL)

	server := &http.Server{
		Addr:    ":80",
		Handler: httpHandler,
	}

	fmt.Println("Server is running on port 80")

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
