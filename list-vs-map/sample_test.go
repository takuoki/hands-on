package sample_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	sample "github.com/takuoki/hands-on/list-vs-map"
)

func TestSliceSample(t *testing.T) {
	testSample(t, sample.SliceSample)
}

func TestMapSample(t *testing.T) {
	testSample(t, sample.MapSample)
}

func testSample(t *testing.T, fn func(d1, d2 [][]string) ([][]string, error)) {
	t.Helper()

	d1 := [][]string{
		{"1", "foo1", "bar1"},
		{"2", "foo2", "bar2"},
		{"3", "foo3", "bar3"},
	}
	d2 := [][]string{
		{"2", "abc2", "xyz2"},
		{"3", "abc3", "xyz3"},
		{"1", "abc1", "xyz1"},
	}
	want := [][]string{
		{"1", "foo1", "bar1", "abc1", "xyz1"},
		{"2", "foo2", "bar2", "abc2", "xyz2"},
		{"3", "foo3", "bar3", "abc3", "xyz3"},
	}

	r, err := fn(d1, d2)
	if err != nil {
		t.Fatalf("error occurred: %v", err)
	}

	assert.Equal(t, want, r)
}

func BenchmarkSample(b *testing.B) {
	benchmarkSample(b, 1000, sample.MapSample)
}

func BenchmarkSliceSample10(b *testing.B) {
	benchmarkSample(b, 10, sample.SliceSample)
}

func BenchmarkMapSample10(b *testing.B) {
	benchmarkSample(b, 10, sample.MapSample)
}

func BenchmarkSliceSample100(b *testing.B) {
	benchmarkSample(b, 100, sample.SliceSample)
}

func BenchmarkMapSample100(b *testing.B) {
	benchmarkSample(b, 100, sample.MapSample)
}

func BenchmarkSliceSample1000(b *testing.B) {
	benchmarkSample(b, 1000, sample.SliceSample)
}

func BenchmarkMapSample1000(b *testing.B) {
	benchmarkSample(b, 1000, sample.MapSample)
}

func benchmarkSample(b *testing.B, length int, fn func(d1, d2 [][]string) ([][]string, error)) {
	b.Helper()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(createData(length), createData(length))
	}
}

func createData(length int) [][]string {
	r := make([][]string, length)
	for i := 0; i < length; i++ {
		r[i] = []string{fmt.Sprint(i), "a", "b"}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(length, func(i, j int) { r[i], r[j] = r[j], r[i] })
	return r
}
