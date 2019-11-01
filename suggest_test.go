package main

import (
	"strings"
	"testing"

	"github.com/Masterminds/sprig/v3"
)

func TestAllFunctionsExist(t *testing.T) {
	funcMap := sprig.GenericFuncMap()
	for _, sug := range suggest {
		if !strings.HasPrefix(sug.Text, "/") {
			if _, exist := funcMap[sug.Text]; exist {
				delete(funcMap, sug.Text)
			} else {
				t.Errorf("Function %s not found in Sprig", sug.Text)
			}
		}
	}
	if len(funcMap) != 0 {
		keys := ""
		for k := range funcMap {
			//keys += k + "\n"
			keys += `{Text: "` + k + `", Description: ""},` + "\n"
		}
		t.Errorf("There are missing functions suggestions:\n%s", keys)
	}
}

func TestEmptyDescription(t *testing.T) {
	for _, sug := range suggest {
		if sug.Description == "" {
			t.Errorf("Function %s has empty Description", sug.Text)
		}
	}
}
