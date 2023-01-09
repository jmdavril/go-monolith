package domain

import (
	"github.com/jmdavril/pubsub"
	"github.com/jmdavril/template/shop/context"
	"github.com/rs/zerolog/log"
)

func StartSubscriptions(sub pubsub.Sub, pss *ProductSalesService) {
	eventHandler := func(e pubsub.Event) {
		pce, _ := e.(context.ProductCreatedEvent)
		err := pss.InitProductSales(pce.Sku)
		log.Err(err) //TODO
		return
	}

	sub.Subscribe(context.PRODUCT_CREATED_TYPE, eventHandler)
}
