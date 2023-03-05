package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"path/filepath"
)

// SetDir for setting the directory for the bot
func (bot Data) SetDir() Data {
	bot.Dir, bot.Err = filepath.Abs(".")
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to set foxley-berry-bot directory")
	}
	return bot
}

// Start the Discord bot listener
func (bot Data) Start() {
	bot.GoBot, bot.Err = discordgo.New("Bot " + bot.Token)
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to instantiate foxley-berry-bot bot")
		return
	}
	bot.User, bot.Err = bot.GoBot.User(UserID)
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to set foxley-berry-bot user id")
		return
	}
	bot.GoBot.AddHandler(bot.MessageHandler)
	bot.Err = bot.GoBot.Open()
	if bot.Err != nil {
		log.Err(bot.Err).Msg("failed to start foxley-berry-bot listener")
		return
	}
	log.Info().Msg("foxley-berry-bot listening")
}
