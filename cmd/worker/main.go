package main

import (
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/lstrgiang/gitscan/internal/usecase/tasks"
)

func main() {
	redisConnection := asynq.RedisClientOpt{
		Addr: os.Getenv("REDIS_LOCATION"), // Redis server address
	}

	// Create and configuring Asynq worker server.
	worker := asynq.NewServer(redisConnection, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: 10,
		// Specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 10, // processed 60% of the time
		},
	})

	// Create a new task's mux instance.
	mux := asynq.NewServeMux()

	mux.HandleFunc(
		tasks.TypeNewScanTask,
		tasks.HandleScan,
	)

	// Run worker server.
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}
