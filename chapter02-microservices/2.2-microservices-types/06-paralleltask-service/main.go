// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"math/rand"
	"net/http"
	"time"
)

func main() {
	ms := NewMicroservice()

	cacheServer := "localhost:6379"
	mqServers := "localhost:9094"

	// 1. Start PTask endpoint
	ms.PTask("/citizen/batch", cacheServer, mqServers)

	// 2. Start 3 workers
	for i := 0; i < 3; i++ {
		ms.PTaskWorker("/citizen/batch", cacheServer, mqServers, func(ctx IContext) error {
			ctx.Log(ctx.ReadInput())
			res := map[string]interface{}{
				"id": "123",
			}
			n := rand.Intn(5)
			time.Sleep(time.Duration(n) * time.Second)
			ctx.Response(http.StatusOK, res)
			return nil
		})
	}

	defer ms.Cleanup()
	ms.Start()
}
