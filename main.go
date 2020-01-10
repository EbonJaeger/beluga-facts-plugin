package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/EbonJaeger/beluga"
	"github.com/EbonJaeger/beluga/config"
	"github.com/bwmarrin/discordgo"
)

// Handle handles the "!hunter2" command
func Handle(s *discordgo.Session, c beluga.Command) {
	// Send the command to the right function
	switch c.Command {
	// Send fact command
	case "fact":
		sendFact(s, c)
		break
	// We don't handle this command
	default:
		return
	}
}

func sendFact(s *discordgo.Session, c beluga.Command) {
	// Get the list of facts
	facts := config.Conf.Facts

	// Make sure we have facts to choose from
	if len(facts) == 0 {
		_, _ = s.ChannelMessageSend(c.ChannelID, fmt.Sprintf("%s, I don't appear to know any facts! This is embarassing...", c.Sender.Mention()))
		return
	}

	var fact string

	// Check if we have any args
	if len(c.MessageNoCmd) > 0 {
		// Make sure that it is a number
		if i, err := strconv.Atoi(c.MessageNoCmd); err == nil {
			// Check if the index is in bounds
			if i < 0 || i >= len(facts) {
				_, _ = s.ChannelMessageSend(c.ChannelID, "There is no fact at that index! Please try again.")
				return
			}
			// Get the fact at that index
			fact = facts[i]
		} else {
			_, _ = s.ChannelMessageSend(c.ChannelID, "That is not a valid fact index. Just who do you take me for?")
			return
		}
	} else {
		// Choose a random fact
		i := rand.Intn(len(facts))
		fact = facts[i]
	}

	// Send the fact to the channel
	_, _ = s.ChannelMessageSend(c.ChannelID, fact)
}
