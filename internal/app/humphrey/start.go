package humphrey

import (
	"github.com/cloudlena/humphrey/internal/app/humphrey/img"
	"github.com/cloudlena/humphrey/pkg/game"
)

// Start is the first scene of the game.
var Start = &game.Scene{
	Image: img.Title,
	Body:  "Das isch d Gschicht vomne chliine Gschöpf namens Humphrey.",
	Actions: map[string]game.Renderer{
		"Los geits! 🎬":                  intro,
		"Wi funktioniert dä Scheiss? 💩": manual,
	},
}
