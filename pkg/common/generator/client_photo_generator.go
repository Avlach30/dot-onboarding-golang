package generator

import (
	"math/rand"
	"time"
)

var photoURLs = []string{
	"https://res.cloudinary.com/deafomwc7/image/upload/v1724774465/codespace/codespace-x/uscl63wsbbxbyujbie6l.png",
	"https://res.cloudinary.com/deafomwc7/image/upload/v1724774544/codespace/codespace-x/hroz1ipqaupc2cwmjnsl.jpg",
}

func GenerateRandomPhotoProfile() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return photoURLs[r.Intn(len(photoURLs))]
}
