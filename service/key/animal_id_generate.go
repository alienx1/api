package key

import (
	"math/rand"
	"time"
)

func GenerateKeyAnimal(Type string, Gender string, Color string, Birthday time.Time) string {
	key_set := `1234567890`
	seededRand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	result := make([]byte, 4)
	for i := range result {
		result[i] = key_set[seededRand.Intn(len(key_set))]
	}
	rand := string(result)
	birtday := Birthday.Format("20060102")
	return Type + `-` + Gender + `-` + Color + `-` + birtday + rand
}
