package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Grains struct {
	ItemId   string
	ItemName string
	Quantity int
	Price    string
}
type Vegetables struct {
	ProductId   string
	ProductName string
	Quantity    int
	Price       string
}
type Fruits struct {
	Id       string
	Name     string
	Quantity int
	Price    string
}

const fruits_Url = "https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b"
const grains_Url = "https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148"
const vegetable_Url = "https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c"

var wg sync.WaitGroup

var Fruit []Fruits
var Grain []Grains
var Vegetable []Vegetables

var summarydetails []interface{}

func GetByItemName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get  Name")
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)

	response, err := http.Get(fruits_Url)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err1 := json.Unmarshal(content, &Fruit)
	if err1 != nil {
		panic(err1)
	}

	for _, p := range Fruit {

		if p.Name == par["item"] {
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	response1, er := http.Get(vegetable_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err2 := json.Unmarshal(content1, &Vegetable)
	if err2 != nil {
		panic(err2)
	}

	response2, er := http.Get(grains_Url)
	if er != nil {
		panic(er)
	}

	for _, c := range Vegetable {
		if c.ProductName == par["item"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err3 := json.Unmarshal(content2, &Grain)
	if err3 != nil {
		panic(err3)
	}

	for _, g := range Grain {
		if g.ItemName == par["item"] {
			json.NewEncoder(w).Encode(g)
			return
		}
	}

	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!! Item Not Found")

}

func getFruits() {
	response, err := http.Get(fruits_Url)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err1 := json.Unmarshal(content, &Fruit)
	if err1 != nil {
		panic(err1)
	}
	wg.Done()
	defer response.Body.Close()
}
func getVegitables() {
	response1, er := http.Get(vegetable_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err2 := json.Unmarshal(content1, &Vegetable)
	if err2 != nil {
		panic(err2)
	}
	wg.Done()
	defer response1.Body.Close()
}
func getGrains() {
	response2, er := http.Get(grains_Url)
	if er != nil {
		panic(er)
	}
	content2, _ := ioutil.ReadAll(response2.Body)

	err3 := json.Unmarshal(content2, &Grain)
	if err3 != nil {
		panic(err3)
	}
	wg.Done()
	defer response2.Body.Close()
}
func GetByItemNameFast(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get  Name")
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)
	go getFruits()
	go getVegitables()
	go getGrains()
	wg.Add(3)

	for _, p := range Fruit {

		if p.Name == par["item"] {
			json.NewEncoder(w).Encode(p)
			return
		}

	}

	for _, c := range Vegetable {
		if c.ProductName == par["item"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	for _, g := range Grain {
		if g.ItemName == par["item"] {
			json.NewEncoder(w).Encode(g)
			return
		}
	}

	json.NewEncoder(w).Encode("!! Item Not Found")

}

func GetByItemNameAndQuantity(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBy Name And Quanity")
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)
	j, _ := strconv.Atoi(par["quantity"])
	response, err := http.Get(fruits_Url)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err1 := json.Unmarshal(content, &Fruit)
	if err1 != nil {
		panic(err1)
	}

	for _, p := range Fruit {

		if p.Quantity >= j && p.Name == par["item"] {
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	response1, er := http.Get(vegetable_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err2 := json.Unmarshal(content1, &Vegetable)
	if err2 != nil {
		panic(err2)
	}

	response2, er := http.Get(grains_Url)
	if er != nil {
		panic(er)
	}

	for _, c := range Vegetable {
		if c.Quantity >= j && c.ProductName == par["item"] {
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err3 := json.Unmarshal(content2, &Grain)
	if err3 != nil {
		panic(err3)
	}

	for _, g := range Grain {
		if g.Quantity >= j && g.ItemName == par["item"] {
			json.NewEncoder(w).Encode(g)
			return
		}
	}

	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!! Item Not Found")

}

func GetByItemNameQuantityAndPrice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetBy Name Quanity And Price")
	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)
	str := "$" + par["price"]
	j, _ := strconv.Atoi(par["quantity"])
	response, err := http.Get(fruits_Url)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err1 := json.Unmarshal(content, &Fruit)
	if err1 != nil {
		panic(err1)
	}

	for _, p := range Fruit {

		if p.Quantity >= j && p.Name == par["item"] && p.Price == str {
			summarydetails = append(summarydetails, p)
			json.NewEncoder(w).Encode(p)
			return
		}

	}
	response1, er := http.Get(vegetable_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err2 := json.Unmarshal(content1, &Vegetable)
	if err2 != nil {
		panic(err2)
	}

	response2, er := http.Get(grains_Url)
	if er != nil {
		panic(er)
	}

	for _, c := range Vegetable {
		if c.Quantity >= j && c.ProductName == par["item"] && c.Price == str {
			summarydetails = append(summarydetails, c)
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err3 := json.Unmarshal(content2, &Grain)
	if err3 != nil {
		panic(err3)
	}

	for _, g := range Grain {
		if g.Quantity >= j && g.ItemName == par["item"] && g.Price == str {
			summarydetails = append(summarydetails, g)
			json.NewEncoder(w).Encode(g)
			return
		}
	}

	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!! Item Not Found")
}

func summary_details(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(summarydetails)
}
func main() {
	fmt.Println("Started.........")
	rou := mux.NewRouter()

	rou.HandleFunc("/buy-item/{item}", GetByItemName).Methods("GET")
	rou.HandleFunc("/buy-item-qty/{item}/{quantity}", GetByItemNameAndQuantity).Methods("GET")
	rou.HandleFunc("/buy-item-qty-price/{item}/{quantity}/{price}", GetByItemNameQuantityAndPrice).Methods("GET")
	rou.HandleFunc("/show-summary/", summary_details).Methods("GET")
	rou.HandleFunc("/fastbuy-item/{item}", GetByItemNameFast).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", rou))
	wg.Wait()
}
