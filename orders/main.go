package main

func main() {
	store := NewStore()
	svc := NewService()

	svc.CreateOrder(context.Background())
}