package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/davecgh/go-spew/spew"

	"github.com/hibiken/asynq"
)

var lock = sync.RWMutex{}
var tasksProcessed = map[int]int{}

func main() {
	r := &asynq.RedisClientOpt{
		Addr:     "redis:6379",
		Password: "kjkjfoip",
	}

	fmt.Printf("Creating redis worker\n")
	bg := asynq.NewBackground(r, &asynq.Config{
		Concurrency: 1,
	})

	bg.Run(asynq.HandlerFunc(handler))
}

func handler(ctx context.Context, t *asynq.Task) error {
	tasknum , err := t.Payload.GetInt("tasknum")
	if err != nil {
		return err
	}
	
	fmt.Printf("Got task: %s\n", t.Type)
	spew.Dump(t)

	func() {
		lock.Lock()
		defer lock.Unlock()
		tasksProcessed[tasknum]++
		fmt.Printf("Tasks processed\n")
		spew.Dump(tasksProcessed)
	}()


	return nil
}
