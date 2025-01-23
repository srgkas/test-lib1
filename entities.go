package test_lib1

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed shared-registry/entities.json
var entitiesJSON []byte

type entity struct {
	name   string
	prefix string
}

var entities []entity

func init() {
	err := json.Unmarshal(entitiesJSON, &entities)

	if err != nil {
		fmt.Println("Error parsing entities JSON")
	}
}

func GetEntities() []entity {
	return entities
}
