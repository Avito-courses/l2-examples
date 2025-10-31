package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go-course-l2-db/connections"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	pool := connections.InitPool(ctx)
	defer pool.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		queryLoop(ctx, pool)
	}()

	<-ctx.Done()
	log.Println("shutdown signal received")
	wg.Wait() // ждём пока закончится функция queryLoop и затем закрываем пул
}

func queryLoop(ctx context.Context, _ *pgxpool.Pool) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("делаем запрос через пул")
		case <-ctx.Done():
			time.Sleep(5 * time.Second)
			fmt.Println("bye")
			return
		}
	}
}
