package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Order struct {
	ID          int     `json:"id"`
	UserID      string  `json:"user_id"`
	OrderDate   string  `json:"order_date"`
	TotalAmount float64 `json:"total_amount"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	url := flag.String("url", "http://127.0.0.1:8080", "server url")
	method := flag.String("method", "GET", "method")
	path := flag.String("path", "get_user", "path")
	flag.Parse()
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	switch *method {
	case "GET":
		resp, err := netClient.Get(
			*url + "/v1/" + *path,
		)
		if err != nil {
			fmt.Println("Ошибка запроса", err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка чтения", err)
			return
		}

		fmt.Println("Response body: ", string(body))
	case "POST":
		var body []byte
		switch *path {
		case "create_user":
			user := User{
				ID:    1,
				Name:  "John Doe",
				Email: "xk0e5@example.com",
			}
			body, _ = json.Marshal(user)
		case "create_order":
			order := Order{
				ID:          1,
				UserID:      "1",
				OrderDate:   "2023-01-01",
				TotalAmount: 10.0,
			}
			body, _ = json.Marshal(order)
		case "create_product":
			product := Product{
				ID:    1,
				Name:  "Product 1",
				Price: 10.0,
			}
			body, _ = json.Marshal(product)
		default:
			fmt.Println("unknown path")
		}
		resp, err := netClient.Post(
			*url+"/v1/"+*path,
			"application/json",
			strings.NewReader(string(body)),
		)
		if err != nil {
			fmt.Println("Ошибка запроса", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
			return
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка чтения", err)
			return
		}

		fmt.Println("Response body: ", string(body))
	default:
		fmt.Println("unknown method")
	}
}
