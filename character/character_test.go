package character_test

import (
	"testing"

	"github.com/neilwhitlow/rpgcore/character"
)

func TestGenerateCharacter(t *testing.T) {
	c := character.GenerateCharacter()
	t.Logf("generated a character %s", c.Name)
}
