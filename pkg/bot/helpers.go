package bot

import (
	"os"

	"github.com/plaid/plaid-go/plaid"
	"github.com/rs/zerolog/log"
)

// CreatePlaidClient for Plaid API
func (bot Data) CreatePlaidClient() Data {
	log.Info().Msg("handling creating Plaid API client")
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", os.Getenv("PLAID_CLIENT_ID"))
	configuration.AddDefaultHeader("PLAID-SECRET", os.Getenv("PLAID_SANDBOX_KEY"))
	configuration.UseEnvironment(plaid.Sandbox)

	bot.Plaid = plaid.NewAPIClient(configuration)
	return bot
}