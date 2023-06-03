package registry

import (
	"os"
	"reflect"

	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
)

var typeRegistry = make(map[string]reflect.Type)

func RegisterType(typeString string, structType reflect.Type) {
	log.SetHandler(json.New(os.Stdout))
	log.Infof("Adding %s to registry", typeString)
	typeRegistry[typeString] = structType
}

func GetType(typeString string) reflect.Type {
	return typeRegistry[typeString]
}

func PrintTypeRegistry() {
	for key, value := range typeRegistry {
		log.Infof("Key %s Value %s", key, value.String())
	}
}
