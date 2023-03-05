package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"strings"
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
	message := discordgo.MessageSend{
		Content: "Finance Summary",
	}
	_, bot.Err = s.ChannelMessageSendComplex(m.ChannelID, &message)
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to post message")
		return
	}
	log.Info().Msgf("finance summary posted to Discord channel %s", m.ChannelID)
}
