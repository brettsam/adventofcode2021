package polymer

import (
	"fmt"
)

type Polymer struct {
	Inserts map[string]string
}

func (p *Polymer) Generate(input string, g int) map[string]uint64 {
	cache := make(map[string]uint64)
	// split into pairs
	for i := 0; i < len(input)-1; i++ {
		pair := input[i : i+2]
		cache[pair]++
	}

	// Generate pair counts
	for i := 0; i < g; i++ {
		cacheSnapshot := make(map[string]uint64, len(cache))
		for k, v := range cache {
			cacheSnapshot[k] = v
		}

		for key, value := range cacheSnapshot {
			if value == 0 {
				continue
			}
			insert := p.Inserts[key]
			nodes := generateOne(key, insert)
			cache[key] -= value
			cache[nodes[0]] += value
			cache[nodes[1]] += value
		}
		fmt.Print()
	}

	counts := make(map[string]uint64, 10)

	for key, value := range cache {
		counts[key[0:1]] += value
	}
	// add the last letter
	counts[input[len(input)-1:]]++
	return counts
}

func generateOne(pair string, insert string) [2]string {
	first := pair[0:1] + insert
	second := insert + pair[1:2]
	return [2]string{first, second}
}
