package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/pu4mane/NATSOrderViewer/internal/app/model"
)

type Cache struct {
	data map[string][]byte
	mu   sync.RWMutex
}

var cache *Cache

func init() {
	cache = &Cache{
		data: make(map[string][]byte),
	}
}

func SaveToCache(order *model.Order) error {
	if err := saveOrderToCache(order); err != nil {
		return err
	}
	if err := saveItemsToCache(order.Items); err != nil {
		return err
	}
	if err := saveDeliveryToCache(order.Delivery); err != nil {
		return err
	}
	if err := savePaymentToCache(order.Payment); err != nil {
		return err
	}

	return nil
}

func saveOrderToCache(order *model.Order) error {
	orderData, err := json.Marshal(order)
	if err != nil {
		return err
	}

	orderKey := fmt.Sprintf("order:%d", order.ID)
	return setData(orderKey, orderData)
}

func saveItemsToCache(items []model.Items) error {
	for _, item := range items {
		itemKey := fmt.Sprintf("item:%d", item.ID)
		if err := setData(itemKey, []byte(strconv.Itoa(item.ChrtID))); err != nil {
			return err
		}
	}
	return nil
}

func saveDeliveryToCache(delivery model.Delivery) error {
	deliveryKey := fmt.Sprintf("delivery:%d", delivery.ID)
	return setData(deliveryKey, []byte(delivery.Name))
}

func savePaymentToCache(payment model.Payment) error {
	paymentKey := fmt.Sprintf("payment:%d", payment.ID)
	return setData(paymentKey, []byte(payment.Transaction))
}

func setData(key string, value []byte) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.data[key] = value
	return nil
}

func Get(key string) ([]byte, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	value, ok := cache.data[key]
	if !ok {
		return nil, errors.New("key not found in cache")
	}

	return value, nil
}
