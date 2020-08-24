package bot

import (
	"fmt"
	"strings"

	"github.com/Tony-Moon/CasinoCord/config"
	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session

// Start begins a new discord session and adds handlers
func Start() {
	// Start a new discord session
	fmt.Println("Creating new session...")
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Failed to create Discord session:", err.Error())
		return
	}
	fmt.Println("Success")

	// Add Message handlers
	dg.AddHandler(listenText)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open the session.
	fmt.Println("Opening discord session.")
	err = dg.Open()
	if err != nil {
		fmt.Println("Failed to open connection:", err.Error())
		return
	}
	defer dg.Close()
}

// listenText Handler for handling Ping! Pong! response
func listenText(s *discordgo.Session, m *discordgo.MessageCreate) {
	// The bot will not respond to bots or replies to incorrect prefixes
	if !strings.HasPrefix(m.Content, config.BotPrefix) {
		return
	}
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Printf("Command call from %v: %v\n", m.Author.Username, m.Content)
	cmd := strings.ToLower(m.Content[1:len(m.Content)])
	switch cmd {
	case "ping":
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	case "pong":
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
