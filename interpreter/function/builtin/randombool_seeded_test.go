// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of randombool_seeded
// Arguments may be:
// - INTEGER, INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randombool-seeded/
func Test_Randombool_seeded(t *testing.T) {
	tests := []struct {
		n int64
		d int64
		s int64
	}{
		{n: 1, d: 10, s: 1000000},
		{n: 3, d: 4, s: 1111111},
		{n: 5, d: 10, s: 2222222},
	}

	for i, tt := range tests {
		ret, err := Randombool_seeded(
			&context.Context{},
			&value.Integer{Value: tt.n},
			&value.Integer{Value: tt.d},
			&value.Integer{Value: tt.s},
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if ret.Type() != value.BooleanType {
			t.Errorf("[%d] Unexpected return type, expect=STRING, got=%s", i, ret.Type())
		}
	}
}
