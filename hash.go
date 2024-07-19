package fttl

import (
	"encoding/hex"

	"github.com/dchest/siphash"
	"golang.org/x/exp/constraints"
)

const (
	_key0       = 506097522914230528
	_key1       = 1084818905618843912
	_eight      = 8
	_sixteen    = 16
	_twentyFour = 24
	_thirtyTwo  = 32
	_forty      = 40
	_fortyEight = 48
	_fiftySix   = 56
)

func Path(num uint64) (string, string) {
	return hex.EncodeToString([]byte{
			byte(num),
			byte(num >> _eight),
		}),
		hex.EncodeToString([]byte{
			byte(num >> _sixteen),
			byte(num >> _twentyFour),
			byte(num >> _thirtyTwo),
			byte(num >> _forty),
			byte(num >> _fortyEight),
			byte(num >> _fiftySix),
		})
}

func Hash(key []byte) uint64 {
	return siphash.Hash(_key0, _key1, key)
}

func IntHash[I constraints.Integer](num I) uint64 {
	num64 := uint64(num)

	return Hash([]byte{
		byte(num64),
		byte(num64 >> _eight),
		byte(num64 >> _sixteen),
		byte(num64 >> _twentyFour),
		byte(num64 >> _thirtyTwo),
		byte(num64 >> _forty),
		byte(num64 >> _fortyEight),
		byte(num64 >> _fiftySix),
	})
}
