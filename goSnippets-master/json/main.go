package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"item"`
	Quantity   int         `json:"quantity"`
	TotalPrice float64     `json:"total_price"`
}

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	parseNLP()
}

func marshal() {
	o := Order{
		ID: "1234",
		// Name:       "learn go",
		Quantity:   3,
		TotalPrice: 30,
		Items: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_2",
				Name:  "interview",
				Price: 20,
			},
		},
	}

	marshal, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", marshal)
}

func unmarshal() {
	s := `{"id":"1234","item":[{"id":"item_1","name":"learn go","price":15},{"id":"item_2","name":"interview","price":20}],"quantity":3,"total_price":30}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func parseNLP() {
	res := `
{
"data": [
    {
        "synonym":"",
        "weight":"0.6",
        "word": "真丝",
        "tag":"材质"
    },
    {
        "synonym":"",
        "weight":"0.8",
        "word": "韩都衣舍",
        "tag":"品牌"
    },
    {
        "synonym":"连身裙;联衣裙",
        "weight":"1.0",
        "word": "连衣裙",
        "tag":"品类"
    }
]
}`
	dict := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Tag     string `json:"tag"`
		} `json:"data"`
	}{}
	err := json.Unmarshal([]byte(res), &dict)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", dict.Data[2].Synonym, dict.Data[2].Tag)
}
