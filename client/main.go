package main

import (
	"log"
	"time"

	"github.com/nafeem-evatix/asynq/tasks"

	"github.com/hibiken/asynq"
)

func main() {
	// Create a new Redis connection for the client.
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379", // Redis server address
	}

	// Create a new Asynq client.
	client := asynq.NewClient(redisConnection)
	defer client.Close()

	task1 := tasks.NewSendEmailTask("a@abc.com", "hello, a")
	task2 := tasks.NewSendEmailTask("b@abc.com", "hello, b")

	if _, err := client.Enqueue(
		task1,                   // task payload
		asynq.Queue("critical"), // set queue for task
	); err != nil {
		log.Fatal(err)
	}

	// Process the task 2 minutes later in low queue.
	if _, err := client.Enqueue(
		task2,                           // task payload
		asynq.Queue("low"),              // set queue for task
		asynq.ProcessIn(time.Second*30), // set time to process task
	); err != nil {
		log.Fatal(err)
	}

}
