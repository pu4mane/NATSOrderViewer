package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pu4mane/NATSOrderViewer/config"
	"github.com/pu4mane/NATSOrderViewer/internal/app/model"
)

func NewDB(config *config.AppConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return db, nil
}

func InsertItem(db *sql.DB, item *model.Items) error {
	err := db.QueryRow(
		`SELECT id FROM items WHERE chrt_id = $1 AND track_number = $2 AND price = $3`,
		item.ChrtID,
		item.TrackNumber,
		item.Price).Scan(&item.ID)

	if err == sql.ErrNoRows {
		err = db.QueryRow(
			`INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status).Scan(&item.ID)
	}
	return err
}

func InsertDelivery(conn *sql.DB, delivery *model.Delivery) error {
	err := conn.QueryRow(
		`INSERT INTO delivery (name, phone, zip, city, address, region, email)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email).Scan(&delivery.ID)
	return err
}

func InsertPayment(conn *sql.DB, payment *model.Payment) error {
	err := conn.QueryRow(
		`INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		payment.Transaction,
		payment.ReqestID,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDT,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee).Scan(&payment.ID)
	return err
}

func InsertOrder(conn *sql.DB, order *model.Order) error {
	err := conn.QueryRow(
		`INSERT INTO orders (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`,
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		order.Delivery.ID,
		order.Payment.ID,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard).Scan(&order.ID)

	if err != nil {
		return err
	}
	return nil
}

func InsertOrderToDB(db *sql.DB, order *model.Order) error {
	err := InsertDelivery(db, &order.Delivery)
	if err != nil {
		return fmt.Errorf("error inserting delivery: %v", err)
	}

	err = InsertPayment(db, &order.Payment)
	if err != nil {
		return fmt.Errorf("error inserting payment: %v", err)
	}

	for _, item := range order.Items {
		err = InsertItem(db, &item)
		if err != nil {
			return fmt.Errorf("error inserting item: %v", err)
		}
	}

	err = InsertOrder(db, order)
	if err != nil {
		return fmt.Errorf("error inserting order: %v", err)
	}

	return nil
}
