package main

import (
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	r := &asynq.RedisClientOpt{
		Addr:     "redis:6379",
		Password: "kjkjfoip",
	}

	fmt.Printf("Creating redis client\n")
	client := asynq.NewClient(r)

	n := 1
	for i := -10; i < 10; i++ {
		name := fmt.Sprintf("task %d", n)
		t := asynq.NewTask(name, map[string]interface{}{"i": i, "tasknum": n})
		err := client.Schedule(t, time.Now().Add(time.Duration(i)*time.Minute))
		if err != nil {
			fmt.Printf("Error scheduling\n")
			panic(err)
		}

		n++
	}

	fmt.Printf("Complete\n")
}
