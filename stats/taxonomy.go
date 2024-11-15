package stats

import (
	_ "embed"
	"encoding/json"

	lhm "github.com/neilwhitlow/collections/linkedhashmap"
)

type TaxAdjustments struct {
	Type                   string `json:"type,omitempty"`
	AbilityKey             string `json:"abilityKey,omitempty"`
	ScoreBonus             int    `json:"scoreBonus,omitempty"`
	BaseInitialScore       int    `json:"baseInitialScore,omitempty"`
	AddMutationStrengthMod bool   `json:"addMutationStrengthMod,omitempty"`
}

type Taxonomy struct {
	Name           string           `json:"name"`
	Abbreviation   string           `json:"abbreviation"`
	BaseWalkSpeed  int              `json:"baseWalkSpeed,omitempty"`
	BaseSwimSpeed  int              `json:"baseSwimSpeed,omitempty"`
	BaseArmorClass int              `json:"baseArmorClass,omitempty"`
	Adjustments    []TaxAdjustments `json:"adjustments,omitempty"`
}

//go:embed data/taxonomies.json
var taxonomyDefinitions []byte

func GetTaxonomyDefinitions() *lhm.LinkedHashMap[string, Taxonomy] {
	taxonomies := lhm.New[string, Taxonomy]()

	a := make([]Taxonomy, 8)
	_ = json.Unmarshal(taxonomyDefinitions, &a)
	for _, item := range a {
		taxonomies.Put(item.Abbreviation, item)
	}
	return taxonomies
}
