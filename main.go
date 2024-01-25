package main

import (
	"context"
	"fmt"
	"joust/app"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.New().Start(ctx); err != nil {
		fmt.Println("failed to start: ", err)
	}
}
