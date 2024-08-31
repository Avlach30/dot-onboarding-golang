package generator

import (
	"fmt"
	"math/rand"
	"time"
)

var spaceTerms = []string{
	"Galaxy",
	"Planet",
	"Ufo",
	"Asteroid",
	"Comet",
	"Nebula",
	"Star",
	"Moon",
	"Cosmos",
	"Blackhole",
}

func GenerateRandomProjectName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	term1 := spaceTerms[r.Intn(len(spaceTerms))]
	term2 := spaceTerms[r.Intn(len(spaceTerms))]

	projectName := fmt.Sprintf("%s %s", term1, term2)
	return projectName
}
