package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/plaid/plaid-go/plaid"
)

const (
	// UserID assigns ID to the bot
	UserID = "@me"
)

// Data for the bot to track along a request
type Data struct {
	// Err for error tracking
	Err error
	// User for the bot to track
	User *discordgo.User
	// GoBot for the bot to track
	GoBot *discordgo.Session
	// Token for the bot to track
	Token string
	// Dir for the bot to track
	Dir string
	// Plaid client
	Plaid PlaidAPI
}

type PlaidAPI struct {
	// Plaid client ID
	ClientID string
	// Plaid secret
	Secret string
	// Plaid client
	Client *plaid.APIClient
}

// Methods for the bot to use
type Methods interface {
	// SetDir for the bot to use
	SetDir() Data
	// Start the Discord bot listener
	Start()
	// MessageHandler for interpreting which function to launch from message contents
	MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate)
	// HandleSummary for handling the summary of finances
	HandleSummary(s *discordgo.Session, m *discordgo.MessageCreate)
	// CreatePlaidClient for Plaid API
	CreatePlaidClient() Data
}
