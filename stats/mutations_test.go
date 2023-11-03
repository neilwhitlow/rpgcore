package stats_test

import (
	"testing"

	"github.com/neilwhitlow/rpgcore/stats"
	"github.com/stretchr/testify/assert"
)

func TestGetMutationDefinitions(t *testing.T) {
	mutations := stats.GetMutationDefinitions()
	assert.NotNil(t, mutations)
}

func TestGetRandomPhysicalMutation(t *testing.T) {
	mutations := stats.GetMutationDefinitions()
	assert.NotNil(t, mutations)
	mut := mutations.GetRandomPhysicalMutation()
	assert.NotNil(t, mut)
	assert.NotEmpty(t, mut.Name)
}

func TestGetPhysicalMutationSuccess(t *testing.T) {
	mutations := stats.GetMutationDefinitions()
	assert.NotNil(t, mutations)
	mut := mutations.GetPhysicalMutation("Wings")
	assert.NotNil(t, mut)
	assert.NotEmpty(t, mut.Name)
}

func TestGetPhysicalMutationEmpty(t *testing.T) {
	mutations := stats.GetMutationDefinitions()
	assert.NotNil(t, mutations)
	mut := mutations.GetPhysicalMutation("NoWayThereIsAMutationWithThisNameInTheData")
	assert.Equal(t, stats.Mutation{}, mut)
}

func TestGetRandomPsychicMutation(t *testing.T) {
	mutations := stats.GetMutationDefinitions()
	assert.NotNil(t, mutations)
	mut := mutations.GetRandomPsychicMutation()
	assert.NotNil(t, mut)
	assert.NotEmpty(t, mut.Name)
}

func TestGetPsychicMutationSuccess(t *testing.T) {
	mutations := stats.GetMutationDefinitions()
	assert.NotNil(t, mutations)
	mut := mutations.GetPsychicMutation("Enhanced Charisma")
	assert.NotNil(t, mut)
	assert.NotEmpty(t, mut.Name)
}

func TestGetPsychicMutationEmpty(t *testing.T) {
	mutations := stats.GetMutationDefinitions()
	assert.NotNil(t, mutations)
	mut := mutations.GetPsychicMutation("NoWayThereIsAMutationWithThisNameInTheData")
	assert.Equal(t, stats.Mutation{}, mut)
}

func TestRollMutationStrength(t *testing.T) {
	mutations := stats.GetMutationDefinitions()
	assert.NotNil(t, mutations)
	mut := getTestMutationWithMutationStrength()
	actual := mut.RollMutationStrength().MutationStrength
	assert.GreaterOrEqual(t, actual, 3)
	assert.LessOrEqual(t, actual, 18)
}

func TestGetMutationStrengthModifier(t *testing.T) {
	mutations := stats.GetMutationDefinitions()
	assert.NotNil(t, mutations)
	mut := getTestMutationWithMutationStrength()
	actual := mut.RollMutationStrength().GetMutationStrengthModifier()
	assert.GreaterOrEqual(t, actual, -3)
	assert.LessOrEqual(t, actual, 3)
}

func getTestMutationWithMutationStrength() stats.Mutation {
	adjustments := []stats.Adjustments{{Type: "CalculatedAbilityBonus", AbilityKey: "HUAPPEAR", ScoreBonus: -2}}
	return stats.Mutation{
		Name:                "Wings",
		RollChance:          stats.RollChance{From: 32, To: 33},
		HasMutationStrength: true,
		Adjustments:         adjustments,
	}
}
