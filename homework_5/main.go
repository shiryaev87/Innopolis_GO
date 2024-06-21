// 12 . Горутины и синхронизация 03.06.2024
// исправить код
// исходники https://raw.githubusercontent.com/bogatyr285/go-course/master/concurrency/dz.go

package main

import (
	"fmt"
	"sync"
)

func RunProcessor(wg *sync.WaitGroup, mu *sync.RWMutex, prices []map[string]float64) {
	go func() {
		defer wg.Done()

		for _, price := range prices {
			mu.RLock()
			for key, value := range price {
				price[key] = value + 1
			}
			mu.RUnlock()
			fmt.Println(price)
		}
	}()
}

func RunWriter() <-chan map[string]float64 {
	var prices = make(chan map[string]float64)
	// 1 Добавил мьютекс для защиты доступа к currentprice
	//var mu sync.RWMutex

	go func() {
		var currentPrice = map[string]float64{
			"inst1": 1.1,
			"inst2": 2.1,
			"inst3": 3.1,
			"inst4": 4.1,
		}
		//
		for i := 1; i < 5; i++ {
			//newPrice := make(map[string]float64)
			//mu.Lock()
			newPrice := make(map[string]float64)

			for key, value := range currentPrice {
				newPrice[key] = value + 1
			}
			for key := range currentPrice {
				currentPrice[key] = newPrice[key]
			}
			//mu.Unlock()

			//fmt.Println(currentPrice)
			prices <- newPrice

			//time.Sleep(time.Second)
		}
		//

		close(prices)
	}()
	return prices

}
func main() {
	//done := make(chan bool, 1)
	var mu sync.RWMutex
	p := RunWriter()
	var prices []map[string]float64

	//go func() {
	for price := range p {
		mu.Lock()
		prices = append(prices, price)
		mu.Unlock()
	}
	//	done <- true
	//}()
	//<-done
	for _, price := range prices {
		fmt.Println(price)
	}

	// 2 wg  передается по ссылке
	wg := sync.WaitGroup{}
	wg.Add(3)
	RunProcessor(&wg, &mu, prices)
	RunProcessor(&wg, &mu, prices)
	RunProcessor(&wg, &mu, prices)
	wg.Wait()

}
