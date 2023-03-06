package main

import (
	"family-budget/pkg/bot"
	"os"
)

func main() {
	gobot := bot.Data{
		Token: os.Getenv("BOT_TOKEN"),
		Plaid: bot.PlaidAPI{
			ClientID: os.Getenv("PLAID_CLIENT_ID"),
			Secret:   os.Getenv("PLAID_SANDBOX_KEY"),
		},
	}
	if gobot.Token == "" {
		panic("BOT_TOKEN not set")
	}
	if gobot.Plaid.ClientID == "" {
		panic("PLAID_CLIENT_ID not set")
	}
	if gobot.Plaid.ClientID == "" {
		panic("PLAID_SANDBOX_KEY not set")
	}
	gobot = gobot.SetDir()
	gobot.Start()
	<-make(chan struct{})
	return
}
