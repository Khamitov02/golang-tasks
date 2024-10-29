package fridge_test

import (
	"context"
	"testing"

	"sandbox/internal/fridge"
	"sandbox/internal/fridge/memory"

	"github.com/google/go-cmp/cmp"
)

func TestAppService_Products(t *testing.T) {
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

	// Invoke the service's Products method
	got, err := service.Products(context.Background())
	if err != nil {
		t.Fatalf("Products() returned error: %v", err)
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
		t.Errorf("Products() mismatch (-want +got):\n%s", diff)
	}
}

func TestAppService_Place(t *testing.T) {
	// Initialize the in-memory storage and service
	storage := memory.NewStorage()
	service := fridge.NewAppService(storage)

	// Test product to place
	product := fridge.Product{
		Name:  "Milk",
		Count: 2,
	}

	// Invoke the service's Place method
	id, err := service.Place(context.Background(), product)
	if err != nil {
		t.Fatalf("Place() returned error: %v", err)
	}

	// Verify that the product was saved correctly
	gotProducts, err := storage.LoadProducts(context.Background())
	if err != nil {
		t.Fatalf("LoadProducts() returned error: %v", err)
	}

	if len(gotProducts) != 1 {
		t.Fatalf("expected 1 product, got %d", len(gotProducts))
	}

	gotProduct := gotProducts[0]
	wantProduct := fridge.Product{
		ID:    id,
		Name:  "Milk",
		Count: 2,
	}

	if diff := cmp.Diff(wantProduct, gotProduct); diff != "" {
		t.Errorf("Stored product mismatch (-want +got):\n%s", diff)
	}
}
