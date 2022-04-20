package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID         string      `json:"id"`
	Items      []OrderItem `json:"item"`
	Quantity   int         `json:"quantity"`
	TotalPrice float64     `json:"total_price"`
}

func main() {
	//marshal()
	//unmarshal()
	parseNLP()
}

func marshal() {
	o := Order{
		ID:         "1234",
		Quantity:   3,
		TotalPrice: 20,
		Items: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_1",
				Name:  "interview",
				Price: 10,
			},
		},
	}

	//fmt.Printf("%+v\n", o)

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
}

func unmarshal() {
	s := `{"id":"1234","item":[{"id":"item_1","name":"learn go","price":15},{"id":"item_1","name":"interview","price":10}],"quantity":3,"total_price":20}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", o)

}

type result struct {
	synonym string `json:"synonym"`
	weight  string ` json:"wight"`
	tag     string `json:"tag"`
	word    string `json:"word"`
}

func parseNLP() {
	res := `{"RequestId":"FA53D08F-37D1-4D81-BEE7-41F24E825F60","Data":"{\"result\":[{\"synonym\":\"\",\"weight\":\"0.100000\",\"tag\":\"普通词\",\"word\":\"请\"},{\"synonym\":\"\",\"weight\":\"0.100000\",\"tag\":\"普通词\",\"word\":\"输入\"},{\"synonym\":\"\",\"weight\":\"1.000000\",\"tag\":\"品类\",\"word\":\"文本\"}],\"success\":true}"}`

	m := struct {
		RequestId string `json:"id"`
		Data      struct {
			Result  []result `json:"result"`
			Success string   `json:"success"`
		} `json:"data"`
	}{}

	//m := make(map[string]interface{})
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", m)

	// Type assertion

	//fmt.Printf("%+v\n", m["Data"])
}
