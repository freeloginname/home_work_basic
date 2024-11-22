package main

import (
	"context"
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

func GetData(url string) (string, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса", err)
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return "", err
	}
	return string(body), nil
}

func PostData(url string, body []byte) (string, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*40))
	defer cancel()
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("Ошибка создания запроса", err)
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Ошибка HTTP-ответа: %d\n", resp.StatusCode)
		return "", err
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return "", err
	}
	return string(body), nil
}

func main() {
	url := flag.String("url", "http://127.0.0.1:8080", "server url")
	method := flag.String("method", "GET", "method")
	path := flag.String("path", "get_user", "path")
	flag.Parse()

	switch *method {
	case "GET":
		resp, err := GetData(*url + "/v1/" + *path)
		if err != nil {
			fmt.Println("Ошибка запроса", err)
			return
		}
		fmt.Println("Response body: ", resp)
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
			return
		}
		resp, err := PostData(*url+"/v1/"+*path, body)
		if err != nil {
			fmt.Println("Ошибка запроса", err)
			return
		}
		fmt.Println("Response body: ", resp)
	default:
		fmt.Println("unknown method")
		return
	}
}
