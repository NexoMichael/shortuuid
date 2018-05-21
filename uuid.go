package shortuuid

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// UUID representation
type UUID []byte

// String is implementation of standart Stringer interface
func (u UUID) String() string {
	if u == nil {
		return ""
	}
	return decodeUUID(u)
}

// NewFromString returns new uuid value converted from string representation of UUID
func NewFromString(in string) UUID {
	return encodeUUID([]byte(in))
}

// NewFromStringer returns uuid value converted from any stringable representation of proper length and format
func NewFromStringer(in fmt.Stringer) UUID {
	return NewFromString(in.String())
}

func decodeUUID(in []byte) string {
	str := hex.EncodeToString(in)
	return strings.Join([]string{str[8:16], str[4:8], str[:4], str[16:20], str[20:]}, "-")
}

/*
	Encode UUID according to the following scheme:
		AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE (36 bytes)
		CCCCBBBBAAAAAAAADDDDEEEEEEEEEEEE (32 bytes)
		(hex decoded string) (16 bytes)

	ref https://www.percona.com/blog/2014/12/19/store-uuid-optimized-way/
*/
func encodeUUID(u []byte) []byte {
	if len(u) != 36 {
		return nil
	}
	// validate only hyphens, in case if uuid contain wrong symbols, hex.DecodeString will fail
	if u[8] != '-' || u[13] != '-' || u[18] != '-' || u[23] != '-' {
		return nil
	}
	var b1 [8]byte
	copy(b1[:], u[:8])
	copy(u[0:], u[14:18])
	copy(u[4:], u[9:13])
	copy(u[8:], b1[:])
	copy(u[16:], u[19:23])
	copy(u[20:], u[24:])
	u = u[:32]

	n, err := hex.Decode(u, u)
	if err != nil {
		return nil
	}
	return u[:n]
}
