package repository

import (
	"awesomeProject/cache"
	"awesomeProject/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(order *model.Order) error {
	query := `
	INSERT INTO myorder (order_uid, track_number, entry, delivery_name, delivery_phone,
	                     delivery_zip, delivery_city, delivery_address, delivery_region,
	                     delivery_email, payment_transaction, payment_request_id,
	                     payment_currency, payment_provider, payment_amount, payment_dt,
	                     payment_bank, payment_delivery_cost, payment_goods_total,
	                     payment_custom_fee, items_chrt_id, items_track_number, items_price,
	                     items_rid, items_name, items_sale, items_size, items_total_price,
	                     items_nm_id, items_brand, items_status, locale, internal_signature,
	                     customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19,
	        $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36,
	        $37, $38, $39)
`
	var id int64
	if err := r.db.QueryRow(query, order.OrderUid, order.TrackNumber, order.Entry, order.Delivery.Name,
		order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region,
		order.Delivery.Email, order.Payment.Transaction, order.Payment.RequestId, order.Payment.Currency, order.Payment.Provider,
		order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal,
		order.Payment.CustomFee, order.Items.ChrtId, order.Items.TrackNumber, order.Items.Price, order.Items.Rid, order.Items.Name,
		order.Items.Sale, order.Items.Size, order.Items.TotalPrice, order.Items.NmId, order.Items.Brand, order.Items.Status,
		order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.ShardKey, order.SmId, order.DateCreated, order.OofShard).Scan(&id); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetDataById(orderUid string) (*model.Order, error) {
	var result *model.Order
	query := `
		SELECT * FROM myorder
		WHERE order_uid = $1
`
	err := r.db.QueryRow(query, orderUid).Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) LoadDataToCache(order *model.Order, cache *cache.Cache) error {
	query := `
		SELECT order_uid, track_number, entry, delivery_name, delivery_phone,
	                     delivery_zip, delivery_city, delivery_address, delivery_region,
	                     delivery_email, payment_transaction, payment_request_id,
	                     payment_currency, payment_provider, payment_amount, payment_dt,
	                     payment_bank, payment_delivery_cost, payment_goods_total,
	                     payment_custom_fee, items_chrt_id, items_track_number, items_price,
	                     items_rid, items_name, items_sale, items_size, items_total_price,
	                     items_nm_id, items_brand, items_status, locale, internal_signature,
	                     customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard  
		FROM myorder
`
	rows, err := r.db.Query(query)
	if err != nil {
		return err
	}
	var data *model.Order
	for rows.Next() {
		rows.Scan(&order.TrackNumber, &order.Entry, &order.Delivery.Name,
			&order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region,
			&order.Delivery.Email, &order.Payment.Transaction, &order.Payment.RequestId, &order.Payment.Currency, &order.Payment.Provider,
			&order.Payment.Amount, &order.Payment.PaymentDt, &order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal,
			&order.Payment.CustomFee, &order.Items.ChrtId, &order.Items.TrackNumber, &order.Items.Price, &order.Items.Rid, &order.Items.Name,
			&order.Items.Sale, &order.Items.Size, &order.Items.TotalPrice, &order.Items.NmId, &order.Items.Brand, &order.Items.Status,
			&order.Locale, &order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.ShardKey, &order.SmId, &order.DateCreated, &order.OofShard)
	}
	cache.Put(order.OrderUid, data)
	return nil
}
