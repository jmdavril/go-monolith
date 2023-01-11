package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmdavril/template/shop/app"
	"github.com/stretchr/testify/assert"
)

type ProductDto struct {
	Sku string `json:"Sku"`
}

type ProductResponse struct {
	Data ProductDto `json:"data"`
}

func TestCreateProduct1(t *testing.T) {
	t.Setenv("PG_CONN", "host=localhost user=user_name dbname=shop port=5432 sslmode=disable")

	server := app.Run()

	w := httptest.NewRecorder()

	jsonBody := []byte(`{"sku": "SKUPROD", "name": "Product New", "price": 39.99}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := "http://localhost:3000/products"
	req1, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	req1.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("error: %v", err)
	}

	server.ServeHTTP(w, req1)

	assert.Equal(t, http.StatusOK, w.Code)

	var respBody1 map[string]string
	err = json.Unmarshal([]byte(w.Body.String()), &respBody1)

	productId := respBody1["productId"]

	w.Body.Reset()

	req2, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s", productId), nil)
	server.ServeHTTP(w, req2)

	var respBody2 ProductResponse
	err = json.Unmarshal([]byte(w.Body.String()), &respBody2)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "SKUPROD", respBody2.Data.Sku)
}
