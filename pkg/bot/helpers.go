package bot

import (
	"github.com/plaid/plaid-go/plaid"
	"github.com/rs/zerolog/log"
)

// CreatePlaidClient for Plaid API
func (bot Data) CreatePlaidClient() Data {
	log.Info().Msg("handling creating Plaid API client")
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", bot.Plaid.ClientID)
	configuration.AddDefaultHeader("PLAID-SECRET", bot.Plaid.Secret)
	configuration.UseEnvironment(plaid.Sandbox)

	bot.Plaid.Client = plaid.NewAPIClient(configuration)
	return bot
}