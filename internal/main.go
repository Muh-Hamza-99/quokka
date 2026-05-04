package main

import (
	"context"
	"fmt"

	"github.com/Muh-Hamza-99/quokka"
)

func main() {
	ctx := context.Background()

	client := quokka.NewClient(nil)

	video, err := client.GetVideo(ctx, "xhgt47nvZUQ")
	if err != nil {
		panic(err)
	}

	fmt.Println("+%v\n", video)
}
