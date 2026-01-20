package sequencemanipulation_test

import (
	"reflect"
	"slices"
	"testing"

	"github.com/palebluedot4/quark/cookbook/sequencemanipulation"
)

func TestReverseString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "single character",
			input: "a",
			want:  "a",
		},
		{
			name:  "ASCII string with even length",
			input: "abcd",
			want:  "dcba",
		},
		{
			name:  "ASCII string with odd length",
			input: "abcde",
			want:  "edcba",
		},
		{
			name:  "Unicode string",
			input: "你好，世界",
			want:  "界世，好你",
		},
		{
			name:  "string with spaces",
			input: "hello world",
			want:  "dlrow olleh",
		},
		{
			name:  "string with emoji",
			input: "Gopher ✅",
			want:  "✅ rehpoG",
		},
		{
			name:  "palindrome string",
			input: "madam",
			want:  "madam",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := sequencemanipulation.ReverseString(tt.input)
			if got != tt.want {
				t.Errorf("ReverseString(%q) got %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestReverseSlice(t *testing.T) {
	t.Parallel()
	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name  string
			input []int
			want  []int
		}{
			{
				name:  "nil slice",
				input: nil,
				want:  nil,
			},
			{
				name:  "empty slice",
				input: []int{},
				want:  []int{},
			},
			{
				name:  "single element slice",
				input: []int{1},
				want:  []int{1},
			},
			{
				name:  "even number of elements",
				input: []int{1, 2, 3, 4},
				want:  []int{4, 3, 2, 1},
			},
			{
				name:  "odd number of elements",
				input: []int{1, 2, 3, 4, 5},
				want:  []int{5, 4, 3, 2, 1},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				var clone []int
				if tt.input != nil {
					clone = make([]int, len(tt.input))
					copy(clone, tt.input)
				}
				got := sequencemanipulation.ReverseSlice(tt.input)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ReverseSlice() got %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(tt.input, clone) {
					t.Errorf("ReverseSlice() modified the input, got %v, want %v", tt.input, clone)
				}
				if tt.input != nil && cap(tt.input) > 0 {
					if reflect.ValueOf(got).Pointer() == reflect.ValueOf(tt.input).Pointer() {
						t.Errorf("ReverseSlice() should have returned a new slice, but it shares the same memory address as the input")
					}
				}
			})
		}
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name  string
			input []string
			want  []string
		}{
			{
				name:  "nil slice",
				input: nil,
				want:  nil,
			},
			{
				name:  "empty slice",
				input: []string{},
				want:  []string{},
			},
			{
				name:  "single element slice",
				input: []string{"gopher"},
				want:  []string{"gopher"},
			},
			{
				name:  "even number of elements",
				input: []string{"hello", "world"},
				want:  []string{"world", "hello"},
			},
			{
				name:  "odd number of elements",
				input: []string{"hello", "world", "!"},
				want:  []string{"!", "world", "hello"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				var clone []string
				if tt.input != nil {
					clone = make([]string, len(tt.input))
					copy(clone, tt.input)
				}
				got := sequencemanipulation.ReverseSlice(tt.input)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ReverseSlice() got %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(tt.input, clone) {
					t.Errorf("ReverseSlice() modified the input, got %v, want %v", tt.input, clone)
				}
				if tt.input != nil && cap(tt.input) > 0 {
					if reflect.ValueOf(got).Pointer() == reflect.ValueOf(tt.input).Pointer() {
						t.Errorf("ReverseSlice() should have returned a new slice, but it shares the same memory address as the input")
					}
				}
			})
		}
	})
}

func TestReverseSliceInPlace(t *testing.T) {
	t.Parallel()
	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name  string
			input []int
			want  []int
		}{
			{
				name:  "nil slice",
				input: nil,
				want:  nil,
			},
			{
				name:  "empty slice",
				input: []int{},
				want:  []int{},
			},
			{
				name:  "single element slice",
				input: []int{1},
				want:  []int{1},
			},
			{
				name:  "even number of elements",
				input: []int{1, 2, 3, 4},
				want:  []int{4, 3, 2, 1},
			},
			{
				name:  "odd number of elements",
				input: []int{1, 2, 3, 4, 5},
				want:  []int{5, 4, 3, 2, 1},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := slices.Clone(tt.input)
				sequencemanipulation.ReverseSliceInPlace(got)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ReverseSliceInPlace() modified the input, got %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name  string
			input []string
			want  []string
		}{
			{
				name:  "nil slice",
				input: nil,
				want:  nil,
			},
			{
				name:  "empty slice",
				input: []string{},
				want:  []string{},
			},
			{
				name:  "single element slice",
				input: []string{"gopher"},
				want:  []string{"gopher"},
			},
			{
				name:  "even number of elements",
				input: []string{"hello", "world"},
				want:  []string{"world", "hello"},
			},
			{
				name:  "odd number of elements",
				input: []string{"hello", "world", "!"},
				want:  []string{"!", "world", "hello"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := slices.Clone(tt.input)
				sequencemanipulation.ReverseSliceInPlace(got)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ReverseSliceInPlace() modified the input, got %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestReverseSliceWithIter(t *testing.T) {
	t.Parallel()
	t.Run("int slice", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name  string
			input []int
			want  []int
		}{
			{
				name:  "nil slice",
				input: nil,
				want:  nil,
			},
			{
				name:  "empty slice",
				input: []int{},
				want:  []int{},
			},
			{
				name:  "single element slice",
				input: []int{1},
				want:  []int{1},
			},
			{
				name:  "even number of elements",
				input: []int{1, 2, 3, 4},
				want:  []int{4, 3, 2, 1},
			},
			{
				name:  "odd number of elements",
				input: []int{1, 2, 3, 4, 5},
				want:  []int{5, 4, 3, 2, 1},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				var clone []int
				if tt.input != nil {
					clone = make([]int, len(tt.input))
					copy(clone, tt.input)
				}
				got := sequencemanipulation.ReverseSliceWithIter(tt.input)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ReverseSliceWithIter() got %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(tt.input, clone) {
					t.Errorf("ReverseSliceWithIter() modified the input, got %v, want %v", tt.input, clone)
				}
				if tt.input != nil && cap(tt.input) > 0 {
					if reflect.ValueOf(got).Pointer() == reflect.ValueOf(tt.input).Pointer() {
						t.Errorf("ReverseSliceWithIter() should have returned a new slice, but it shares the same memory address as the input")
					}
				}
			})
		}
	})

	t.Run("string slice", func(t *testing.T) {
		t.Parallel()
		tests := []struct {
			name  string
			input []string
			want  []string
		}{
			{
				name:  "nil slice",
				input: nil,
				want:  nil,
			},
			{
				name:  "empty slice",
				input: []string{},
				want:  []string{},
			},
			{
				name:  "single element slice",
				input: []string{"gopher"},
				want:  []string{"gopher"},
			},
			{
				name:  "even number of elements",
				input: []string{"hello", "world"},
				want:  []string{"world", "hello"},
			},
			{
				name:  "odd number of elements",
				input: []string{"hello", "world", "!"},
				want:  []string{"!", "world", "hello"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				var clone []string
				if tt.input != nil {
					clone = make([]string, len(tt.input))
					copy(clone, tt.input)
				}
				got := sequencemanipulation.ReverseSliceWithIter(tt.input)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ReverseSliceWithIter() got %v, want %v", got, tt.want)
				}
				if !reflect.DeepEqual(tt.input, clone) {
					t.Errorf("ReverseSliceWithIter() modified the input, got %v, want %v", tt.input, clone)
				}
				if tt.input != nil && cap(tt.input) > 0 {
					if reflect.ValueOf(got).Pointer() == reflect.ValueOf(tt.input).Pointer() {
						t.Errorf("ReverseSliceWithIter() should have returned a new slice, but it shares the same memory address as the input")
					}
				}
			})
		}
	})
}
