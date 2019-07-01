package common

import (
	"math/big"
)

// modInt is a *big.Int that performs all of its arithmetic with modular reduction.
type modInt big.Int

func ModInt(mod *big.Int) *modInt {
	return (*modInt)(new(big.Int).Set(mod))
}

func (mi *modInt) Add(x, y *big.Int) *big.Int {
	i := new(big.Int).Add(x, y)
	return new(big.Int).Mod(i, (*big.Int)(mi))
}

func (mi *modInt) Sub(x, y *big.Int) *big.Int {
	i := new(big.Int).Sub(x, y)
	return new(big.Int).Mod(i, (*big.Int)(mi))
}

func (mi *modInt) Mul(x, y *big.Int) *big.Int {
	i := new(big.Int).Mul(x, y)
	return new(big.Int).Mod(i, (*big.Int)(mi))
}

func (mi *modInt) Exp(x, y *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, (*big.Int)(mi))
}
