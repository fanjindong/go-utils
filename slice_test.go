package utils

import (
	"testing"
)
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
		{input: input{slice: &[]string{"a", "b", "c"}, index: 3, item: "3"}, want: []string{"a", "b", "c", "3"}},

		{input: input{slice: &[]int32{0, 1, 2}, index: 0, item: int32(-1)}, want: []int32{-1, 0, 1, 2}},
		{input: input{slice: &[]int32{0, 1, 2}, index: 1, item: int32(-1)}, want: []int32{0, -1, 1, 2}},
		{input: input{slice: &[]int32{0, 1, 2}, index: 2, item: int32(-1)}, want: []int32{0, 1, -1, 2}},
		{input: input{slice: &[]int32{0, 1, 2}, index: 3, item: int32(-1)}, want: []int32{0, 1, 2, -1}},
	}

	for _, ts := range tests {
		err := SliceInsert(ts.input.slice, ts.input.index, ts.input.item)
		assert.NoError(t, err)
		//t.Log(ts.input.slice, ts.want)
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

// BenchmarkSliceInsert-4   	  2000000	    594 ns/op
func BenchmarkSliceInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "g"}
		index := i % len(s)
		_ = SliceInsert(&s, index, "0")
	}
}

func TestSliceShuffle(t *testing.T) {
	type args struct {
		slice interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{slice: []int{1, 2, 3, 4, 5}}},
		{name: "2", args: args{slice: []string{"a", "b", "c", "d", "e"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.slice)
			SliceShuffle(tt.args.slice)
			t.Log(tt.args.slice)
		})
	}
}

func TestSliceContain(t *testing.T) {
	type args struct {
		slice interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{slice: []int{1, 2, 3, 4, 5}, value: 3}, want: true},
		{name: "2", args: args{slice: []int{1, 2, 3, 4, 5}, value: 5}, want: true},
		{name: "3", args: args{slice: []int{1, 2, 3, 4, 5}, value: 6}, want: false},
		{name: "4", args: args{slice: []int{1, 2, 3, 4, 5}, value: 0}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContain(tt.args.slice, tt.args.value); got != tt.want {
				t.Errorf("SliceContain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceContainsAny(t *testing.T) {
	type args struct {
		slice  interface{}
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{slice: []int{1, 2, 3, 4, 5}, values: []interface{}{1, 2}}, want: true},
		{name: "2", args: args{slice: []int{1, 2, 3, 4, 5}, values: []interface{}{4, 2}}, want: true},
		{name: "3", args: args{slice: []int{1, 2, 3, 4, 5}, values: []interface{}{0, 3}}, want: true},
		{name: "4", args: args{slice: []int{1, 2, 3, 4, 5}, values: []interface{}{0, 6}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContainsAny(tt.args.slice, tt.args.values...); got != tt.want {
				t.Errorf("SliceContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceContainsAll(t *testing.T) {
	type args struct {
		slice  interface{}
		values []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{slice: []int{1, 2, 3, 4, 5}, values: []interface{}{1, 2, 3}}, want: true},
		{name: "2", args: args{slice: []int{1, 2, 3, 4, 5}, values: []interface{}{1, 3}}, want: true},
		{name: "3", args: args{slice: []int{1, 2, 3, 4, 5}, values: []interface{}{1}}, want: true},
		{name: "4", args: args{slice: []int{1, 2, 3, 4, 5}, values: []interface{}{0, 2, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContainsAll(tt.args.slice, tt.args.values...); got != tt.want {
				t.Errorf("SliceContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
