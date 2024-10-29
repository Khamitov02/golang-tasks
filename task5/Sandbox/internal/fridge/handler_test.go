package fridge_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sandbox/internal/fridge"
	"sandbox/internal/fridge/memory"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/go-cmp/cmp"
)

func TestHandler_getProducts(t *testing.T) {
	// Initialize the in-memory storage and service
	storage := memory.NewStorage()
	service := fridge.NewAppService(storage)

	// Add test products to storage
	product1 := fridge.Product{
		Name:  "Apple",
		Count: 5,
	}
	product2 := fridge.Product{
		Name:  "Banana",
		Count: 10,
	}

	_, err := storage.SaveProduct(context.Background(), product1)
	if err != nil {
		t.Fatalf("failed to save product1: %v", err)
	}

	_, err = storage.SaveProduct(context.Background(), product2)
	if err != nil {
		t.Fatalf("failed to save product2: %v", err)
	}

	// Setup router and handler
	router := chi.NewRouter()
	h := fridge.NewHandler(router, service)
	h.Register()

	// Create a request to the handler
	req, err := http.NewRequest(http.MethodGet, "/api/v1/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: want %d, got %d", http.StatusOK, rr.Code)
	}

	// Parse the response body
	var got []fridge.Product
	err = json.NewDecoder(rr.Body).Decode(&got)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	// Expected products
	want := []fridge.Product{
		{
			ID:    "1",
			Name:  "Apple",
			Count: 5,
		},
		{
			ID:    "2",
			Name:  "Banana",
			Count: 10,
		},
	}

	// Compare the expected and actual products
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GET /api/v1/products mismatch (-want +got):\n%s", diff)
	}
}
