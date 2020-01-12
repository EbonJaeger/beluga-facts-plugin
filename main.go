package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

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
	case "listfacts":
		listFacts(s, c)
		break
	// We don't handle this command
	default:
		return
	}
}

func listFacts(s *discordgo.Session, c beluga.Command) {
	// Get the list of facts
	facts := config.Conf.Facts

	// Make sure we have facts to choose from
	if len(facts) == 0 {
		_, _ = s.ChannelMessageSend(c.ChannelID, fmt.Sprintf("%s, I don't appear to know any facts! This is embarassing...", c.Sender.Mention()))
		return
	}

	// Calculate number of pages
	pages := (len(facts) + 10 - 1) / 10
	pageNum := 1

	// Check for args
	if len(c.MessageNoCmd) > 0 {
		// Make sure that it is a number
		if i, err := strconv.Atoi(c.MessageNoCmd); err == nil {
			// Check that the given page number is in bounds
			if i < 1 || i > pages {
				_, _ = s.ChannelMessageSend(c.ChannelID, "That page doesn't exist! Please try again.")
				return
			}
			pageNum = i
		} else {
			_, _ = s.ChannelMessageSend(c.ChannelID, "That is not a valid page number. Just who do you take me for?")
			return
		}
	}

	// Split facts into a page
	start := (pageNum - 1) * 10
	end := start + 10
	// Clamp to number of facts available
	if end >= len(facts) {
		end = len(facts)
	}
	// Slice the facts we want to send in this response
	entries := facts[start:end]
	// Get the Guild name
	guild, _ := s.Guild(c.GuildID)
	// Make sure the message was sent in a Guild
	if guild == nil {
		s.ChannelMessageSend(c.ChannelID, "I don't know what you want. Please try again in your Discord server.")
		return
	}
	// Build our response
	var b strings.Builder
	b.WriteString(fmt.Sprintf(" **Facts for %s (Page: %d / %d):**\n", guild.Name, pageNum, pages))
	curr := start
	for _, e := range entries {
		b.WriteString(fmt.Sprintf("> `%d` - %s\n", curr, e))
		curr++
	}
	b.WriteString("To send a specific fact, type `!fact <index>` in the chat.")
	// Create a DM channel
	dm, _ := s.UserChannelCreate(c.Sender.ID)
	// Send the response
	_, err := s.ChannelMessageSend(dm.ID, b.String())
	if err != nil {
		s.ChannelMessageSend(c.ChannelID, fmt.Sprintf("%s I can't slide into your DM's, are they open?", c.Sender.Mention()))
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
