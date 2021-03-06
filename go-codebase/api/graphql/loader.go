package graphql

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// LoadSchema graphql from file
func LoadSchema() string {
	var schema strings.Builder
	here := fmt.Sprintf("%s/api/graphql/", os.Getenv("APP_PATH"))

	// load main schema
	s, err := ioutil.ReadFile(here + "schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	schema.Write(s)

	// load auth schema
	s, err = ioutil.ReadFile(here + "auth.graphql")
	if err != nil {
		log.Fatal(err)
	}
	schema.Write(s)

	return schema.String()
}
