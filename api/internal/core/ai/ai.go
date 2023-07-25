package ai

import "math/rand"

type AI struct{}

func (a AI) GetAttacks(list []string) []string {
	records := make([]string, 0)

	// logic goes here (now it's random)
	for _, item := range list {
		if rand.Intn(10) > 7 {
			records = append(records, item)
		}
	}

	return records
}
