package utils

import "testing"
import "github.com/stretchr/testify/assert"

func TestSliceInsert(t *testing.T) {
	type input struct {
		slice interface{}
		index int
		item  interface{}
	}
	tests := []struct {
		input input
		want  interface{}
	}{
		{input: input{slice: &[]string{"a", "b", "c"}, index: 0, item: "0"}, want: []string{"0", "a", "b", "c"}},
		{input: input{slice: &[]string{"a", "b", "c"}, index: 1, item: "1"}, want: []string{"a", "1", "b", "c"}},
		{input: input{slice: &[]string{"a", "b", "c"}, index: 2, item: "2"}, want: []string{"a", "b", "2", "c"}},

		{input: input{slice: &[]int32{0, 1, 2}, index: 0, item: int32(-1)}, want: []int32{-1, 0, 1, 2}},
		{input: input{slice: &[]int32{0, 1, 2}, index: 1, item: int32(-1)}, want: []int32{0, -1, 1, 2}},
		{input: input{slice: &[]int32{0, 1, 2}, index: 2, item: int32(-1)}, want: []int32{0, 1, -1, 2}},
	}

	for _, ts := range tests {
		err := SliceInsert(ts.input.slice, ts.input.index, ts.input.item)
		assert.NoError(t, err)
		t.Log(ts.input.slice, ts.want)
		switch got := ts.input.slice.(type) {
		case *[]string:
			want := ts.want.([]string)
			assert.Equal(t, len(*got), len(want))
			for i := 0; i < len(want); i++ {
				assert.Equal(t, (*got)[i], want[i])
			}
		case *[]int32:
			want := ts.want.([]int32)
			assert.Equal(t, len(*got), len(want))
			for i := 0; i < len(want); i++ {
				assert.Equal(t, (*got)[i], want[i])
			}
		}
	}
}
