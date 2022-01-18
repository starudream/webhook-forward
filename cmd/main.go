package main

import (
	"github.com/go-sdk/lib/app"
	"github.com/go-sdk/lib/log"

	webhookForward "github.com/starudream/webhook-forward"
)

const name = "webhook-forward"

func main() {
	a := app.New(name)
	defer a.Recover()

	a.Add(webhookForward.Start)

	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
