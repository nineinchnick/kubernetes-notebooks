package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
)

func main() {
	data, err := ioutil.ReadFile("/etc/config-printer/config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-ctx.Done()
}
