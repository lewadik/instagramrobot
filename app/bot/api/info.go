package api

import (
	"github.com/pkg/errors"
	"gopkg.in/telebot.v3"
)

// HandlerStart is the entry point for the incoming update
func HandlerStart(c telebot.Context) error {
	// Ignore channels and groups
	if c.Chat().Type != telebot.ChatPrivate {
		return nil
	}

	if err := c.Reply("Hello! I'm instagram keeper, post some instagram public post/reels to me."); err != nil {
		return errors.Wrap(err, "Couldn't sent the start command response")
	}

	return nil
}
