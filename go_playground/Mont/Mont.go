package Mont

import "math/big"

type mont struct {
	n  uint     // m.BitLen()
	m  *big.Int // m.Bit(0) must equal to 1
	r2 *big.Int // (1 << 2n) % m
}

func (m *mont) redc(t *big.Int) *big.Int {
	temp := new(big.Int).Set(t)
	for i := uint(0); i < m.n; i++ {
		if temp.Bit(0) == 1 {
			temp.Add(temp, m.m)
		}
		temp.Rsh(temp, 1)
	}
	if temp.Cmp(m.m) >= 0 {
		temp.Sub(temp, m.m)
	}
	return temp
}

func newMont(m *big.Int) *mont {
	if m.Bit(0) != 1 {
		return nil
	}
	n := uint(m.BitLen())
	one := big.NewInt(1)
	temp := one.Lsh(one, 2*n)
	return &mont{n, m, temp.Mod(temp, m)}
}

func Exp(x, y, m *big.Int) *big.Int {
	mont := newMont(m)
	t1 := new(big.Int).Mul(x, mont.r2)
	prod := mont.redc(mont.r2)
	base := mont.redc(t1.Mul(x, mont.r2))
	exp := new(big.Int).Set(y)
	for exp.BitLen() > 0 {
		if exp.Bit(0) == 1 {
			prod = mont.redc(prod.Mul(prod, base))
		}
		exp.Rsh(exp, 1)
		base = mont.redc(base.Mul(base, base))
	}
	return mont.redc(prod)
}
