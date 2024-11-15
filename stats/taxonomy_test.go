package stats_test

import (
	"testing"

	"github.com/neilwhitlow/rpgcore/stats"
	"github.com/stretchr/testify/assert"
)

func TestGetTaxonomyDefinitions(t *testing.T) {
	taxonomies := stats.GetTaxonomyDefinitions()
	assert.NotNil(t, taxonomies)
}
