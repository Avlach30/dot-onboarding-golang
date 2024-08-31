package generator

import (
	"math/rand"
	"time"
)

var projectImageUrl = []string{
	"https://minio-cloud.codespace.id/backoffice/project-thumbnail/project-1.png",
	"https://minio-cloud.codespace.id/backoffice/project-thumbnail/project-2.png",
	"https://minio-cloud.codespace.id/backoffice/project-thumbnail/project-3.png",
}

func GenerateRandomProjectImage() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return projectImageUrl[r.Intn(len(projectImageUrl))]
}
