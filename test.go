package foobar

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func f(s string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		resp, _ := http.Get("http://192.168.14.145:30358/")
		if resp.StatusCode == 200 {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			log.Printf("%s:  %s", s, body)

		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f(fmt.Sprintf("%s %d", "goroutine", i), &wg)
	}

	wg.Wait()
}
