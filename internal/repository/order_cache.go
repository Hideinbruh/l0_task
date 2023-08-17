package repository

import (
	"awesomeProject/cache"
	"awesomeProject/internal/model"
)

type OrderCache struct {
	cache *cache.Cache
}

func NewOrderCache(cache *cache.Cache) *OrderCache {
	return &OrderCache{cache: cache}
}

func (oc *OrderCache) GetOrderCache(key string) (*model.Order, error) {
	val, err := oc.cache.GetByIdFromCache(key)
	if err != nil {
		return nil, err
	}
	orderCache := &model.Order{
		OrderUid:    val.OrderUid,
		TrackNumber: val.TrackNumber,
		Entry:       val.TrackNumber,
		Delivery: model.Delivery{
			Name:    val.Delivery.Name,
			Phone:   val.Delivery.Phone,
			Zip:     val.Delivery.Zip,
			City:    val.Delivery.City,
			Address: val.Delivery.City,
			Region:  val.Delivery.Region,
			Email:   val.Email,
		},
		Payment: model.Payment{
			Transaction:  val.Payment.Transaction,
			RequestId:    val.Payment.RequestId,
			Currency:     val.Payment.Currency,
			Provider:     val.Payment.Provider,
			Amount:       val.Payment.Amount,
			PaymentDt:    val.Payment.PaymentDt,
			Bank:         val.Payment.Bank,
			DeliveryCost: val.Payment.DeliveryCost,
			GoodsTotal:   val.Payment.GoodsTotal,
			CustomFee:    val.Payment.CustomFee,
		},
		Items: model.Items{
			ChrtId:      val.Items.ChrtId,
			TrackNumber: val.Items.TrackNumber,
			Price:       val.Items.Price,
			Rid:         val.Items.Rid,
			Name:        val.Items.Name,
			Sale:        val.Items.Sale,
			Size:        val.Items.Size,
			TotalPrice:  val.Items.TotalPrice,
			NmId:        val.Items.NmId,
			Brand:       val.Items.Brand,
			Status:      val.Items.Status,
		},
		Locale:            val.Locale,
		InternalSignature: val.InternalSignature,
		CustomerId:        val.CustomerId,
		DeliveryService:   val.DeliveryService,
		ShardKey:          val.ShardKey,
		SmId:              val.SmId,
		DateCreated:       val.DateCreated,
		OofShard:          val.OofShard,
	}
	return orderCache, nil
}

func (oc *OrderCache) CreateOrderCache(order *model.Order) error {
	oc.cache.Put(order.OrderUid, order)
	return nil
}
