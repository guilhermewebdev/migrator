package lib_test

import (
	"testing"

	"github.com/guilhermewebdev/migrator/src/lib"
)

func TestSnakeCase(t *testing.T) {
	expect_mapping := map[string]string{
		"Testing snake case": "testing_snake_case",
		"TestingSnakeCase":   "testing_snake_case",
		"testingSnakeCase":   "testing_snake_case",
		"TESTING_SNAKE_CASE": "testing_snake_case",
		"TESTING SNAKE CASE": "testing_snake_case",
	}
	for text, expecting := range expect_mapping {
		exemple := lib.SnakeCase(text)
		if exemple != expecting {
			t.Log(exemple, " is not ", expecting, " no exemplo ", text)
			t.Fail()
		}
	}
}
