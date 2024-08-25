package generator

import (
	"fmt"
	"math/rand"
	"time"
)

var spaceTerms = []string{
	"galaxy",
	"planet",
	"ufo",
	"asteroid",
	"comet",
	"nebula",
	"star",
	"moon",
	"cosmos",
	"blackhole",
}

func GenerateRandomProjectName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	term1 := spaceTerms[r.Intn(len(spaceTerms))]
	term2 := spaceTerms[r.Intn(len(spaceTerms))]

	projectName := fmt.Sprintf("%s %s", term1, term2)
	return projectName
}
