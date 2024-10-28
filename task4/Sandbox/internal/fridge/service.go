package fridge

import "context"

type AppService struct {
	store Store
}

func NewAppService(s Store) *AppService {
	return &AppService{
		store: s,
	}
}

func (s *AppService) Products(ctx context.Context) ([]Product, error) {
	products, err := s.store.LoadProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *AppService) Place(ctx context.Context, product Product) (id string, err error) {
	id, err = s.store.SaveProduct(ctx, product)
	if err != nil {
		return "", err
	}

	return id, nil
}
