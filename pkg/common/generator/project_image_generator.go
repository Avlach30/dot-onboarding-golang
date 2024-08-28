package generator

import (
	"math/rand"
	"time"
)

var projectImageUrl = []string{
	"https://res.cloudinary.com/deafomwc7/image/upload/v1710229364/codespace/images/WhyChooseUs/Screenshot_2024-03-12_at_14.41.16_zcei8q.png",
}

func GenerateRandomProjectImage() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return projectImageUrl[r.Intn(len(projectImageUrl))]
}
