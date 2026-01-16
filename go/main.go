package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "https://mood-tracker.25775lkdbeg1.eu-de.codeengine.appdomain.cloud" 
	start := time.Now()

	for i := 0; i < 100000; i++ { // 100000 Testanfragen
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Anfrage %d fehlgeschlagen\n", id)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("Anfrage %d: Status %s\n", id, resp.Status)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Dauer für 100000 Anfragen: %s\n", time.Since(start))
}