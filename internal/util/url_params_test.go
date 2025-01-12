package util_test

import (
	"testing"

	"github.com/boolean-software/aikido-http-client/internal/util"
)

func TestBuildURLParams(t *testing.T) {

	type Filters struct {
		Page int `url:"page"`
		Size int `url:"size"`
	}

	params, err := util.BuildURLParams(Filters{Page: 10, Size: 12})
	if err != nil {
		t.Error(err)
	}

	expected := "page=10&size=12"

	if params != expected {
		t.Errorf("`%s` does not match expected `%s`", params, expected)
	}
}
