package golib

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
)

type entity struct {
	name   string
	prefix string
}

var entities []entity

func init() {
	entitiesJSON, err := os.ReadFile("../shared/config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(entitiesJSON, &entities)

	if err != nil {
		fmt.Println("Error parsing entities JSON")
	}
}

func GetEntities() []entity {
	return entities
}
