package nats

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"

	"github.com/pu4mane/NATSOrderViewer/internal/app/cache"
	"github.com/pu4mane/NATSOrderViewer/internal/app/model"
	"github.com/pu4mane/NATSOrderViewer/internal/app/storage"
	"github.com/pu4mane/NATSOrderViewer/pkg"
)

func PublishOrders(db *sql.DB, sc stan.Conn, directory string) error {
	orders, err := pkg.ReadJSONFilesFromDirectory(directory)
	if err != nil {
		return err
	}

	for _, orderBytes := range orders {
		order := &model.Order{}
		err = json.Unmarshal(orderBytes, order)
		if err != nil {
			log.Printf("Error unmarshalling order: %v", err)
			continue
		}

		err = storage.InsertOrderToDB(db, order)
		if err != nil {
			log.Printf("Error inserting order to DB: %v", err)
			continue
		}

		err = cache.SaveToCache(order)
		if err != nil {
			log.Printf("Error saving order to cache: %v", err)
			continue
		}

		PublishOrder(sc, order)
	}

	return nil
}
