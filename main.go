package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	url := "https://www.cbr-xml-daily.ru/latest.js"

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите количество денег: ")
	valueStr, _ := reader.ReadString('\n')
	valueStr = strings.TrimSpace(valueStr)

	fmt.Print("Введите валюту, в которую вы хотите перевести: ")
	// currencyStr, _ := reader.ReadString('\n')  // Пока не используется

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err) // Выводим ошибку с помощью log.Fatal
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err) // Выводим ошибку с помощью log.Fatal
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err) // Выводим ошибку с помощью log.Fatal
	}

	rates, ok := data["rates"].(map[string]interface{})
	if !ok {
		log.Fatal("Неверный формат данных: 'rates' не является map")
	}

	course, ok := rates["USD"].(float64)
	if !ok {
		log.Fatal("Неверный формат данных: курс USD не найден")
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Fatal(err) // Выводим ошибку с помощью log.Fatal
	}

	result := course * value
	fmt.Println(result)
}