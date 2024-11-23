package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/freeloginname/home_work_basic/hw13_http/client"
	"github.com/freeloginname/home_work_basic/hw13_http/internal"
	"github.com/freeloginname/home_work_basic/hw13_http/server"
	"github.com/pkg/errors"
)

const (
	GetUser = `{"id":1,"name":"John Doe","email":"xk0e5@example.com"}
`
	GetOrder = `{"id":1,"user_id":"1","order_date":"2022-01-01","total_amount":100}
`
	GetProduct = `{"id":1,"name":"Product 1","price":10}
`
)

// server tests.
func TestGetUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/get_user", nil)
	w := httptest.NewRecorder()
	server.GetUser(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != GetUser {
		t.Errorf("expected %s got %v", GetUser, string(data))
	}
}

func TestCreateUser(t *testing.T) {
	user := internal.User{
		ID:    1,
		Name:  "John Doe",
		Email: "xk0e5@example.com",
	}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/v1/create_user", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	server.CreateUser(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != GetUser {
		t.Errorf("expected %s got %v", GetUser, string(data))
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected %d got %v", http.StatusCreated, res.StatusCode)
	}
}

func TestGetOrder(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/get_order", nil)
	w := httptest.NewRecorder()
	server.GetOrder(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != GetOrder {
		t.Errorf("expected %s got %v", GetOrder, string(data))
	}
}

func TestCreateOrder(t *testing.T) {
	order := internal.Order{
		ID:          1,
		UserID:      "1",
		OrderDate:   "2022-01-01",
		TotalAmount: 100.0,
	}
	body, _ := json.Marshal(order)

	req := httptest.NewRequest(http.MethodPost, "/v1/create_order", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	server.CreateOrder(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != GetOrder {
		t.Errorf("expected %s got %v", GetOrder, string(data))
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected %d got %v", http.StatusCreated, res.StatusCode)
	}
}

func TestGetProduct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/v1/get_product", nil)
	w := httptest.NewRecorder()
	server.GetProduct(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != GetProduct {
		t.Errorf("expected %s got %v", GetProduct, string(data))
	}
}

func TestCreateProduct(t *testing.T) {
	product := internal.Product{
		ID:    1,
		Name:  "Product 1",
		Price: 10.0,
	}
	body, _ := json.Marshal(product)

	req := httptest.NewRequest(http.MethodPost, "/v1/create_product", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	server.CreateProduct(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != GetProduct {
		t.Errorf("expected %s got %v", GetProduct, string(data))
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected %d got %v", http.StatusCreated, res.StatusCode)
	}
}

// client tests.

func TestClientGetUser(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, GetUser)
	}))
	defer svr.Close()
	res, err := client.GetData(svr.URL + "/v1/get_user")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	if res != GetUser {
		t.Errorf("expected res to be %s got %s", GetUser, res)
	}
}

func TestClientGetOrder(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, GetOrder)
	}))
	defer svr.Close()
	res, err := client.GetData(svr.URL + "/v1/get_order")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	if res != GetOrder {
		t.Errorf("expected res to be %s got %s", GetOrder, res)
	}
}

func TestClientGetProduct(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, GetProduct)
	}))
	defer svr.Close()
	res, err := client.GetData(svr.URL + "/v1/get_product")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	if res != GetProduct {
		t.Errorf("expected res to be %s got %s", GetProduct, res)
	}
}

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c Client) PostData(path string, body []byte) (string, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*40))
	defer cancel()
	res, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url+path, strings.NewReader(string(body)))
	// res, err := http.Post(c.url+path, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return "", errors.Wrap(err, "unable to complete Post request")
	}
	defer res.Body.Close()
	out, err := io.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "unable to read response data")
	}

	return string(out) + "\n", nil
}

func TestClientCreateUser(t *testing.T) {
	user := internal.User{
		ID:    1,
		Name:  "John Doe",
		Email: "xk0e5@example.com",
	}
	body, _ := json.Marshal(user)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(GetUser))
	}))
	defer svr.Close()
	c := NewClient(svr.URL)
	res, err := c.PostData("/v1/create_user", body)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	if res != GetUser {
		t.Errorf("expected res to be %s got %s", GetUser, res)
	}
}

func TestClientCreateOrder(t *testing.T) {
	order := internal.Order{
		ID:          1,
		UserID:      "1",
		OrderDate:   "2022-01-01",
		TotalAmount: 100.0,
	}
	body, _ := json.Marshal(order)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(GetOrder))
	}))
	defer svr.Close()
	c := NewClient(svr.URL)
	res, err := c.PostData("/v1/create_order", body)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	if res != GetOrder {
		t.Errorf("expected res to be %s got %s", GetOrder, res)
	}
}

func TestClientCreateProduct(t *testing.T) {
	product := internal.Product{
		ID:    1,
		Name:  "Product 1",
		Price: 10.0,
	}
	body, _ := json.Marshal(product)
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(GetProduct))
	}))
	defer svr.Close()
	c := NewClient(svr.URL)
	res, err := c.PostData("/v1/create_product", body)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	if res != GetProduct {
		t.Errorf("expected res to be %s got %s", GetProduct, res)
	}
}
