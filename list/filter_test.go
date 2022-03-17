package list

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	type args struct {
		list []int
		fn   func(item int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "can filter out simple numbers",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(item int) bool {
					return item > 5
				},
			},
			want: []int{6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Filter(tt.args.list, tt.args.fn))
		})
	}
}

func TestAll(t *testing.T) {
	type args struct {
		list []int
		fn   func(item int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true if all items match",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(item int) bool {
					return item < 10
				},
			},
			want: true,
		},
		{
			name: "returns false if one of the items does not match",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				fn: func(item int) bool {
					return item < 10
				},
			},
			want: false,
		},
		{
			name: "returns false if none of the items match",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				fn: func(item int) bool {
					return item > 10
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, All(tt.args.list, tt.args.fn), "All(%v, %v)", tt.args.list, tt.args.fn)
		})
	}
}

func TestNone(t *testing.T) {
	type args struct {
		list []int
		fn   func(item int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true if none of the items match",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(item int) bool {
					return item > 10
				},
			},
			want: true,
		},
		{
			name: "returns false if one of the items matches",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				fn: func(item int) bool {
					return item < 10
				},
			},
			want: false,
		},
		{
			name: "returns false if all of the items are matching",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				fn: func(item int) bool {
					return item < 10
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, None(tt.args.list, tt.args.fn), "None(%v, %v)", tt.args.list, tt.args.fn)
		})
	}
}

func TestSome(t *testing.T) {
	type args struct {
		list []int
		fn   func(item int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns true if more than one item matches the function",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6},
				fn: func(item int) bool {
					return item > 3
				},
			},
			want: true,
		},
		{
			name: "returns true if only one item matches the function",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6},
				fn: func(item int) bool {
					return item == 3
				},
			},
			want: true,
		},
		{
			name: "returns false if none match the function",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6},
				fn: func(item int) bool {
					return item == 9
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Some(tt.args.list, tt.args.fn), "Some(%v, %v)", tt.args.list, tt.args.fn)
		})
	}
}

func TestFilterWithDifferentTypes(t *testing.T) {
	floats := []float64{1.0, 2.3, 5.4, 6.7}
	floatResult := Filter(floats, func(item float64) bool {
		return item > 5.6
	})
	assert.Equal(t, []float64{6.7}, floatResult)

	s := []string{"string1", "string11", "string111"}
	stringResult := Filter(s, func(item string) bool {
		return len(item) > 7
	})
	assert.Equal(t, []string{"string11", "string111"}, stringResult)
}

func TestNoneWithDifferentTypes(t *testing.T) {
	floats := []float64{1.0, 2.3, 5.4, 6.7}
	floatResult := None(floats, func(item float64) bool {
		return item > 11.6
	})
	assert.True(t, floatResult)

	s := []string{"string1", "string11", "string111"}
	stringResult := None(s, func(item string) bool {
		return item == "nope"
	})
	assert.True(t, stringResult)
}

func TestAllWithDifferentTypes(t *testing.T) {
	floats := []float64{1.0, 2.3, 5.4, 6.7}
	floatResult := All(floats, func(item float64) bool {
		return item < 10.2
	})
	assert.True(t, floatResult)

	s := []string{"string1", "string11", "string111"}
	stringResult := All(s, func(item string) bool {
		return strings.Contains(item, "string")
	})
	assert.True(t, stringResult)
}

func TestSomeWithDifferentTypes(t *testing.T) {
	floats := []float64{1.0, 2.3, 5.4, 6.7}
	floatResult := Some(floats, func(item float64) bool {
		return item > 5.6
	})
	assert.True(t, floatResult)

	s := []string{"string1", "string11", "string111"}
	stringResult := Some(s, func(item string) bool {
		return len(item) > 7
	})
	assert.True(t, stringResult)
}
