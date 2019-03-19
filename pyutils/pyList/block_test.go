package pylist

import (
	"reflect"
	"testing"
)

func TestKeywords_Get(t *testing.T) {
	type args struct {
		key          string
		defaultValue interface{}
	}
	tests := []struct {
		name string
		kw   *Keywords
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kw.Get(tt.args.key, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keywords.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
