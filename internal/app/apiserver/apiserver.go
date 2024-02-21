package apiserver

import (
	"log"
	"net/http"

	"github.com/pu4mane/NATSOrderViewer/config"
	"github.com/pu4mane/NATSOrderViewer/internal/app/handlers"
	"github.com/pu4mane/NATSOrderViewer/internal/app/nats"
	"github.com/pu4mane/NATSOrderViewer/internal/app/storage"
)

func Start(cfg *config.AppConfig) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandlerIndex)
	mux.HandleFunc("/order", handlers.HandlerOrder)

	db, err := storage.NewDB(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	sc := nats.ConnectToNatsStreaming(cfg.Stan.ClusterID, cfg.Stan.ClientID)
	defer sc.Close()

	go func() {
		err = nats.PublishOrders(db, sc, ".././test")
		if err != nil {
			log.Printf("Error publishing orders: %v", err)
		}
	}()

	return http.ListenAndServe(":8080", mux)
}
