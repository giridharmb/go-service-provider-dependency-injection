package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func GetEpochTime() int64 {
	now := time.Now()
	epochTime := now.Unix()
	return epochTime
}

func RandStr() string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	n := 6
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func PrettyPrintData(data interface{}) {
	dataBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Printf("error : could not MarshalIndent json : %v", err.Error())
		return
	}
	fmt.Printf("\n%v\n\n", string(dataBytes))
}
