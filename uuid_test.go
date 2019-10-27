package shortuuid

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUUIDString(t *testing.T) {
	expected := "704389bc-244c-4d1a-bbfd-e4f03cceff02"
	require.Equal(t, expected, NewFromString(expected).String())
	require.Equal(t, expected, NewFromString(strings.ToUpper(expected)).String())
}

func TestBadUUID(t *testing.T) {
	require.Nil(t, NewFromString(""))
	require.Nil(t, NewFromString("123"))
	require.Nil(t, NewFromString("704389bc-244c-4d1a-bbfd-e4f03cceff0211"))
	require.Nil(t, NewFromString("704389by-244c-4d1a-bbfd-e4f03cceff02"))
	require.Nil(t, NewFromString("704-389bc244c4d-1a-bbfd-e4f03cceff02"))
	require.Equal(t, "", NewFromString("").String())
}

type str string

func (s str) String() string {
	return string(s)
}

func TestUUIDStringer(t *testing.T) {
	expected := str("704389bc-244c-4d1a-bbfd-e4f03cceff02")
	require.Equal(t, string(expected), NewFromStringer(expected).String())
}

/*
	BenchmarkUUIDBytes-8   	1000000000	         2.70 ns/op	       0 B/op	       0 allocs/op
*/
func BenchmarkUUIDBytes(b *testing.B) {
	u := []byte("949ea5ff-0bfd-4c4f-8f2c-af3d3edd0641")

	for i := 0; i < b.N; i++ {
		encodeUUID(u)
	}
}

/*
	BenchmarkUUID-8        	20000000	        95.6 ns/op	      48 B/op	       1 allocs/op
*/
func BenchmarkUUID(b *testing.B) {
	u := "949ea5ff-0bfd-4c4f-8f2c-af3d3edd0641"

	for i := 0; i < b.N; i++ {
		NewFromString(u)
	}
}
