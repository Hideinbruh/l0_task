package nats

import (
	"awesomeProject/cache"
	"awesomeProject/internal/model"
	"awesomeProject/internal/service"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

type Subscriber struct {
	clusterId    string
	clientId     string
	uri          string
	channel      string
	service      *service.Service
	connection   stan.Conn
	subscription stan.Subscription
}

func NewSubscriber(clusterId, clientId, uri, channel string, service *service.Service) *Subscriber {
	return &Subscriber{
		clusterId: clusterId,
		clientId:  clientId,
		uri:       uri,
		channel:   channel,
		service:   service,
	}
}

func (s *Subscriber) ConnectAndSubscribe(cache *cache.Cache, service *service.Service, order *model.Order) {
	// создается соединение к nats-streaming
	sc, err := stan.Connect("test-cluster", "test-client", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatalf("failed to connect to NATS Streaming Server: %v", err)
	}
	defer sc.Close()

	// создание подписки на канал "example"
	sub, err := sc.Subscribe("example", func(msg *stan.Msg) {
		service.Save(msg.Data)
	}, stan.DurableName("example-sub"))
	if err != nil {
		log.Fatalf("failed to subscribe to channel: %v", err)
	}

	defer sub.Close()

	order = &model.Order{
		OrderUid:    "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: model.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: model.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestId:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDt:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: model.Items{
			ChrtId:      9934930,
			TrackNumber: "WBILMTESTTRACK",
			Price:       453,
			Rid:         "ab4219087a764ae0btest",
			Name:        "Mascaras",
			Sale:        30,
			Size:        "0",
			TotalPrice:  317,
			NmId:        2389212,
			Brand:       "Vivienne Sabo",
			Status:      202,
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmId:              99,
		DateCreated:       "time.Time{2021-11-26T06:22:19Z}",
		OofShard:          "1",
	}

	// отправка сообщения в канал "example"
	ord, _ := json.Marshal(order)
	if err := sc.Publish("example", ord); err != nil {
		log.Fatalf("failed to publish message: %v", err)
	}
	cache.Put(order.OrderUid, order)

	// ожидание получения сообщений
	time.Sleep(10 * time.Second) // дайте подписчику время на обработку сообщений
}
