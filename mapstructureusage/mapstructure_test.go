package mapstructureusage

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	FirstName   string `mapstructure:"givenName"`
	Surname     string `mapstructure:"surname"`
	City        string `mapstructure:"city"`
	YearOfBirth int    `mapstructure:"yearOfBirth"`
}

func TestMapStructureDecoding(t *testing.T) {
	input := `{
	"givenName": "Frank",
	"surname": "Sinatra",
	"city": "Hoboken",
	"yearOfBirth": 1915,
	"topTenAlbums": 42
	}`

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(input), &parsed); err != nil {
		t.Error("parse:", err)
	}

	var errorAccum *multierror.Error
	var result Person
	var metadata mapstructure.Metadata

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: &metadata,
		Result:   &result,
	})
	if err != nil {
		t.Error("Failed constructing Decoder")
	}
	if err := decoder.Decode(parsed); err != nil {
		errorAccum = multierror.Append(errorAccum, err.(*mapstructure.Error).WrappedErrors()...)
	}

	fieldsPresent := make(map[string]struct{}, len(metadata.Keys))
	var present struct{}
	for _, fieldName := range metadata.Keys {
		fieldsPresent[fieldName] = present
	}

	for _, fieldName := range []string{"givenName", "surname", "yearOfBirth", "city"} {
		if _, ok := fieldsPresent[fieldName]; !ok {
			errorAccum = multierror.Append(errorAccum, fmt.Errorf("'%s' was not specified", fieldName))
		}
	}

	if errorAccum.ErrorOrNil() != nil {
		t.Error(errorAccum.Error())
	} else {
		log.Printf("%#v", result)
	}
}
