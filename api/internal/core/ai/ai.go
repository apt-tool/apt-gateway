package ai

import (
	"fmt"
	"math/rand"

	"github.com/automated-pen-testing/api/pkg/models"
	"github.com/automated-pen-testing/api/pkg/models/instruction"
)

type AI struct {
	Models *models.Interface
}

func (a AI) GetAttacks() ([]*instruction.Instruction, error) {
	list, err := a.Models.Instructions.Get()
	if err != nil {
		return nil, fmt.Errorf("[ai.Get] failed to get instructions error=%w", err)
	}

	records := make([]*instruction.Instruction, 0)

	// logic goes here (now it's random)
	for _, item := range list {
		if rand.Intn(10) > 7 {
			records = append(records, item)
		}
	}

	return records, nil
}
