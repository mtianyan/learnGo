package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	order := createOrder()
	fmt.Printf("%+v\n", order)
	//&{ID:1234 Item:{ID:iterm_6 Name:Golang Price:66} Description: Quantity:6 TotalPrice:666}

	bytes, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", bytes)
	//without tag: {"ID":"1234","Name":"Golang","Quantity":6,"TotalPrice":666}
	//with tag: {"id":"1234","item":{"id":"iterm_6","name":"Golang","price":66},"quantity":6,"total_price":666}

	//unmarshalExample()
	//unmarshalMapExample()
	unmarshalCustomStructExample()
}

func unmarshalCustomStructExample() {
	s := `{"id":"1234","item":{"id":"iterm_6","name":"Golang","price":66},"quantity":6,"total_price":666}`
	item := struct {
		Item struct {
			ID    string  `json:"id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		} `json:"item"`
	}{}
	err := json.Unmarshal([]byte(s), &item)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", item.Item)
}

func unmarshalMapExample() {
	s := `{"id":"1234","item":{"id":"iterm_6","name":"Golang","price":66},"quantity":6,"total_price":666}`
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", m)
	//map[id:1234 item:map[id:iterm_6 name:Golang price:66] quantity:6 total_price:666]

	fmt.Printf("%v\n", m["item"].(map[string]interface{})["id"])
	//iterm_6
}

func unmarshalExample() {
	s := `{"id":"1234","item":{"id":"iterm_6","name":"Golang","price":66},"quantity":6,"total_price":666}`
	var order Order
	err := json.Unmarshal([]byte(s), &order)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", order)
	//{ID:1234 Item:{ID:iterm_6 Name:Golang Price:66} Description: Quantity:6 TotalPrice:666}
}

func createOrder() *Order {
	order := &Order{
		ID:         "1234",
		Quantity:   6,
		TotalPrice: 666,
		Item: OrderItem{
			ID:    "iterm_6",
			Name:  "Golang",
			Price: 66,
		},
	}
	return order
}

type Order struct {
	// 1. tag for json parse
	// 2. tag for skipping empty fields
	ID          string    `json:"id"`
	Item        OrderItem `json:"item"`
	Description string    `json:"description,omitempty"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
}

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
