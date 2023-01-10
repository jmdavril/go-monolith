package domain

import (
	"github.com/jmdavril/pubsub"
	"github.com/jmdavril/template/shop/context"
)

func StartSubscriptions(sub pubsub.Sub, pss *ProductSalesService) {
	eventHandler := func(e pubsub.Event) {
		pce, _ := e.(context.ProductCreatedEvent)

		if err := pss.InitProductSales(pce.Sku); err != nil {
			logger.Error().
				Str("Op", "productCreate").
				Str("ProductSKU", pce.Sku).
				Msg("Could not create product sales entity for new product")

			return
		}

		logger.Info().
			Str("Op", "productCreate").
			Str("ProductSKU", pce.Sku).
			Msg("Created product sales entity for new product")
		
		return
	}

	sub.Subscribe(context.PRODUCT_CREATED_TYPE, eventHandler)
}
