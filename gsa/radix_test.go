package gsa

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestCountSort(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{"abaab"},
			"aaabb",
		},
		{
			"Test 2",
			args{"mississippi"},
			"iiiimppssss",
		},
		{
			"Test 3",
			args{""},
			"",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSort(tt.args.x); got != tt.want {
				t.Errorf("CountSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucketSort(t *testing.T) {
	type args struct {
		x   string
		idx []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test 1",
			args{"abaab", []int{0, 1, 2, 3, 4}},
			[]int{0, 2, 3, 1, 4},
		},
		{
			"Test 2",
			args{"abaab", []int{4, 3, 2, 1, 0}},
			[]int{3, 2, 0, 4, 1},
		},
		{
			"Test 3",
			args{"mississippi", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			[]int{1, 4, 7, 10, 0, 8, 9, 2, 3, 5, 6},
		},
		{
			"Test 4",
			args{"mississippi", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
			[]int{10, 7, 4, 1, 0, 9, 8, 6, 5, 3, 2},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BucketSort(tt.args.x, tt.args.idx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BucketSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// checkSAIndices checks that the suffix array sa has all the
// indices in x (plus one for the sentinel if len(sa) == len(x) + 1).
// Reports an error to t otherwise
func checkSAIndices(t *testing.T, x string, sa []int) bool {
	t.Helper()

	if len(sa) != len(x) && len(sa) != len(x)+1 {
		t.Errorf("Suffix %v has an invalid length: %d. "+
			"It should be %d without sentinel or %d with.",
			sa, len(sa), len(x), len(x)+1)
	}

	indices := make([]int, len(sa))
	for i, j := range sa {
		indices[i] = int(j)
	}

	sort.Ints(indices)

	for i, j := range indices {
		if j < 0 || j > len(x) {
			t.Errorf("Index %d is not valid for a suffix array over a string of length %d.",
				j, len(x))
		}

		if i < j {
			t.Errorf("Index %d is missing from the suffix array.",
				i)
			return false
		}
	}

	return true
}

// checkSASorted checks if a suffix array sa actually
// represents the sorted suffix in the string x. Reports
// errors to t.
func checkSASorted(t *testing.T, x string, sa []int) bool {
	t.Helper()

	result := true

	for i := 1; i < len(sa); i++ {
		if x[sa[i-1]:] >= x[sa[i]:] {
			t.Errorf("Suffix array is not sorted! %q >= %q",
				x[sa[i-1]:], x[sa[i]:])

			result = false
		}
	}

	return result
}

// checkSuffixArray runs all the consistency checks for
// suffix array sa over string x, reporting errors to t.
func checkSuffixArray(t *testing.T, x string, sa []int) bool {
	t.Helper()

	result := true
	result = result && checkSAIndices(t, x, sa)
	result = result && checkSASorted(t, x, sa)

	return result
}

// newRandomSeed creates a new random number generator
func newRandomSeed(tb testing.TB) *rand.Rand {
	tb.Helper()

	seed := time.Now().UTC().UnixNano()
	return rand.New(rand.NewSource(seed))
}

// randomStringN constructs a random string of length in n, over the alphabet alpha.
func randomStringN(n int, alpha string, rng *rand.Rand) string {
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = alpha[rng.Intn(len(alpha))]
	}

	return string(bytes)
}

func TestLsdRadixSort(t *testing.T) {
	rng := newRandomSeed(t)
	for i := 0; i < 10; i++ {
		x := randomStringN(10, "acgt", rng)
		sa := LsdRadixSort(x)
		checkSuffixArray(t, x, sa)
	}
}

func TestMsdRadixSort(t *testing.T) {
	rng := newRandomSeed(t)
	for i := 0; i < 10; i++ {
		x := randomStringN(10, "acgt", rng)
		sa := MsdRadixSort(x)
		checkSuffixArray(t, x, sa)
	}

}
