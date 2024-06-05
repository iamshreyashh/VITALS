package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		insuranceList: []insurance{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in insurance
	insuranceIdMap := make(map[uint64]bool)
	insuranceCount := gs.GetinsuranceCount()
	for _, elem := range gs.insuranceList {
		if _, ok := insuranceIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for insurance")
		}
		if elem.Id >= insuranceCount {
			return fmt.Errorf("insurance id should be lower or equal than the last id")
		}
		insuranceIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
