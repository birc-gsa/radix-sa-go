package gsa

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

// Compute the suffix array for x using a least-significant digits
// first radix sort
func LsdRadixSort(x string) []int {
	return []int{}
}

// Compute the suffix array for x using a most-significant digits
// first radix sort
func MsdRadixSort(x string) []int {
	return []int{}
}
