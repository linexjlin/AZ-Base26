package base26

import (
	bc "github.com/chtison/baseconverter"
	"math/big"
)

type Base26 struct {
	base26str string
	dec       *big.Int
}

func NewBase26(in string) Base26 {
	var b26 Base26
	b26.base26str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if d, e := bc.BaseToDecimal(in, b26.base26str); e != nil {
		b26.dec = &big.Int{}
	} else {
		b26.dec = d
	}
	return b26
}

func (b *Base26) Add(i int64) Base26 {
	b.dec = big.NewInt(0).Add(b.dec, big.NewInt(i))
	return *b
}

func (b *Base26) String() string {
	res, _ := bc.DecimalToBase(b.dec, b.base26str)
	return res
}
