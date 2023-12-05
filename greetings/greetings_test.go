package greetings

import (
	"fmt"
	"pgregory.net/rapid"
	"strings"
	"testing"
	"testing/quick"
)

func TestNotEmptyName(t *testing.T) {

	personeName := "Pippo"
	_, err := Hello(personeName)
	if err != nil {
		t.Errorf("Hello() error = %v expecting a string containing %v", err, personeName)
		return

	}
}

func TestNEmptyName(t *testing.T) {

	personeName := ""
	got, _ := Hello(personeName)
	if got != "" {
		t.Errorf("Hello() expecting error got %v", got)
		return

	}
}

func TestExampleBased(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "Mario test", args: args{"Mario"}, want: "Mario", wantErr: false},
		{name: "Luigi test", args: args{"Luigi"}, want: "Luigi", wantErr: false},
		{name: "Empty name test", args: args{""}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hello(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !strings.Contains(got, tt.want) {
				t.Errorf("Hello() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPropertyBased_with_quick(t *testing.T) {

	property := func(name string) bool {

		got, err := Hello(name)

		if name != "" {
			return strings.Contains(got, name)
		} else {
			return err != nil
		}

	}
	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}

}

func TestPropertyBased_with_rapid(t *testing.T) {

	prop := func(t *rapid.T) {
		stringGenerator := rapid.String()

		for i := 0; i < 100; i++ {
			input := stringGenerator.Example(i)

			result, err := Hello(input)

			if input != "" {
				checkProperty(strings.Contains(result, input), fmt.Sprintf("string |%v| does not contains |%v|", result, input), t)
			} else {
				checkProperty(err != nil, "expected an error if empty input", t)
			}
		}
	}

	rapid.Check(t, prop)
}

func checkProperty(condition bool, errorMessage string, t *rapid.T) {
	if !condition {
		t.Fatalf(errorMessage)
	}
}
