package tests

import (
	"net/http"
	"sync"
	"testing"
)

func TestStressAPI(t *testing.T) {
	url := "http://localhost:8080/api/stats"
	const concurrentUsers = 1000
	var wg sync.WaitGroup

	for range concurrentUsers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil || resp.StatusCode != 200 {
				t.Errorf("Sistema colapsó bajo carga")
			}
		}()
	}
	wg.Wait()
}