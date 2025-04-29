package main

type service struct {
	store OrdersService
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context){
	return nil
}