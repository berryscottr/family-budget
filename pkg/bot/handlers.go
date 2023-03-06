package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

// MessageHandler for interpreting which function to launch from message contents
func (bot Data) MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == bot.User.ID {
		return
	}
	if strings.Contains(strings.ToLower(m.Content), "!summary") {
		bot.HandleSummary(s, m)
	}
}


// HandleSummary for handling the summary of finances
func (bot Data) HandleSummary(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Info().Msg("handling finance summary post creation")
	bot = bot.CreatePlaidClient()
	log.Info().Msg("Plaid client created")
	ctx := context.Background()
	accounts := bot.Plaid.Client.PlaidApi.AccountsBalanceGet(ctx)
	log.Info().Msg("Plaid accounts retrieved")
	message := discordgo.MessageSend{
		Content: fmt.Sprintf("Hello, %s. Your financial summary: %v", m.Author.Mention(), accounts),
	}
	_, bot.Err = s.ChannelMessageSendComplex(m.ChannelID, &message)
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to post message")
		return
	}
	log.Info().Msgf("finance summary posted to Discord channel %s", m.ChannelID)
}
