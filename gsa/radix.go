package gsa

// In this implementation we use the rotation trick to index the columns
// of the suffix/rotation matrix. Unlike the padding case, with rotations
// the letters are therefore always the same in each column, and thus so is
// the initial bucket pointers, but we do need the sentinel to make it work.

// Add the sentinel to x.
func addSentinel(x string) []byte {
	return append([]byte(x), 0)
}

func rotIdx(x []byte, i int) byte {
	return x[i%len(x)]
}

// We assume here that we are counting bytes and not unicode points.
// That means that we can assume that we have 256 different characters.
type counts [256]int
type buckets [256]int

func countLetters(x []byte) counts {
	c := counts{}
	for _, b := range x {
		c[b]++
	}
	return c
}

func cumsum(c counts) buckets {
	sum, b := 0, buckets{}
	for i := 0; i < 256; i++ {
		b[i] = sum
		sum += c[i]
	}
	return b
}

func (b *buckets) next(a byte) int {
	idx := b[a]
	b[a]++
	return idx
}

// Sort the string x using a count sort
func CountSort(x string) string {
	c := countLetters([]byte(x))
	y := make([]byte, len(x))
	i, j := 0, 0
	for ; j < 256; j++ {
		for k := 0; k < c[j]; k++ {
			y[i] = byte(j)
			i++
		}
	}
	return string(y)
}

// Sort the indices in idx according to the letters in x
// using a bucket sort
func BucketSort(x string, idx []int) []int {
	b := cumsum(countLetters([]byte(x)))
	out := make([]int, len(idx))
	for _, i := range idx {
		out[b.next(x[i])] = i
	}
	return out
}

// Sort a single column.
//
// the buckets, b, passed by value so we don't modify the original. We can
// reuse the already calculated buckets when we are using rotations for our
// sort.
func bucketColumn(x []byte, sufs []int, col int, b buckets) []int {
	out := make([]int, len(sufs))
	for _, s := range sufs {
		out[b.next(rotIdx(x, s+col))] = s
	}
	return out
}

// Compute the suffix array for x using a least-significant digits
// first radix sort
func LsdRadixSort(x string) []int {
	y := addSentinel(x)
	b := cumsum(countLetters(y))

	idx := make([]int, len(y))
	for i := 0; i < len(idx); i++ {
		idx[i] = i
	}

	for col := len(idx) - 1; col >= 0; col-- {
		idx = bucketColumn(y, idx, col, b)
	}

	return idx
}

// Count the letters in the slice of suffixes in sufs at the rotation colum col
func countLettersSlice(x []byte, sufs []int, col int) counts {
	c := counts{}
	for _, i := range sufs {
		c[rotIdx(x, i+col)]++
	}
	return c
}

type StackFrame struct {
	col   int
	slice []int
}
type Stack []StackFrame

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(frame StackFrame) {
	*s = append(*s, frame)
}

func (s *Stack) pop() StackFrame {
	if s.isEmpty() {
		panic("Don't pop an empty stack")
	}
	i := len(*s) - 1
	v := (*s)[i]
	*s = (*s)[:i]
	return v
}

func recSubSlices(sufs []int, col int, b buckets, s *Stack) {
	prev := 0
	for i := 1; i < 256; i++ {
		if b[i] == b[prev] {
			continue
		}
		if b[prev]+1 < b[i] {
			s.push(StackFrame{col: col + 1, slice: sufs[b[prev]:b[i]]})
		}
		prev = i
	}
	// last slice to stack...
	if b[prev]+1 < b[len(b)-1] {
		s.push(StackFrame{col: col + 1, slice: sufs[b[prev]:]})
	}
}

func sortSlice(x []byte, sufs []int, col int, s *Stack) {
	b := cumsum(countLettersSlice(x, sufs, col))
	recSubSlices(sufs, col, b, s)
	res := make([]int, len(sufs))
	for _, s := range sufs {
		res[b.next(rotIdx(x, s+col))] = s
	}
	copy(sufs, res)
}

// Compute the suffix array for x using a most-significant digits
// first radix sort
func MsdRadixSort(x string) []int {
	y := addSentinel(x)

	idx := make([]int, len(y))
	for i := 0; i < len(idx); i++ {
		idx[i] = i
	}

	s := Stack{StackFrame{col: 0, slice: idx}}
	for !s.isEmpty() {
		frame := s.pop()
		sortSlice(y, frame.slice, frame.col, &s)
	}

	return idx
}
