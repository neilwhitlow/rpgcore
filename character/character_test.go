package character_test

import (
	"testing"

	"github.com/neilwhitlow/rpgcore/character"
	"github.com/neilwhitlow/rpgcore/stats"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCharacter(t *testing.T) {
	c := character.GenerateCharacter()
	t.Logf("generated a character %s", c.Name)
}

func TestGetReadOnlyPrimeAbilities(t *testing.T) {
	c := character.GenerateCharacter()
	pa := c.GetReadOnlyAbilities(c.PrimeAbilities)
	expected := stats.Ability{}
	actual := pa.First().Value
	assert.IsType(t, expected, actual)
	t.Logf("number of prime abilities %d", pa.Len())
}
