package main

type store struct {
	//add here our mongo db
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) {
	return nil
}