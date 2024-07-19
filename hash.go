package fttl

import (
	"encoding/hex"

	"github.com/dchest/siphash"
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
