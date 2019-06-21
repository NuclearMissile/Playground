package main

import (
	"fmt"
	"math/big"
)

func main() {
	toBig := func(x int64) *big.Int { return big.NewInt(x) }

	empty := func() *big.Int { return new(big.Int) }

	Np := func(r, n *big.Int) *big.Int {
		res := toBig(0)
		t := toBig(0)
		i := toBig(1)
		rr := new(big.Int).Set(r)
		for rr.Cmp(toBig(1)) > 0 {
			if t.Bit(0) == 0 {
				t.Add(t, n)
				res.Add(res, i)
			}
			t.Rsh(t, 1)
			rr.Rsh(rr, 1)
			i.Lsh(i, 1)
		}
		return res
	}

	mont := func(a, b, R, N, Np *big.Int) *big.Int {
		A := empty().Mul(a, R)
		B := empty().Mul(b, R)
		T := empty().Mul(A, B)
		t := empty().Div(empty().Add(T, empty().Mul(empty().Mod(empty().Mul(T, Np), R), N)), R)
		if t.Cmp(N) >= 0 {
			return empty().Sub(t, N)
		} else {
			return t
		}
	}

	leftBin := func(x, y *big.Int) *big.Int {
		flag := empty().Lsh(toBig(1), uint(y.BitLen()-1))
		res := toBig(1)
		for flag.Cmp(empty()) != 0 {
			res.Mul(res, res)
			if empty().And(y, flag).Cmp(empty()) != 0 {
				res.Mul(res, x)
			}
			flag.Rsh(flag, 1)
		}
		return res
	}

	leftBinMod := func(x, y, m *big.Int) *big.Int {
		flag := empty().Lsh(toBig(1), uint(y.BitLen()-1))
		res := toBig(1)
		for flag.Cmp(empty()) != 0 {
			res = empty().Mod(empty().Mul(res, res), m)
			if empty().And(y, flag).Cmp(empty()) != 0 {
				res = empty().Mod(empty().Mul(res, x), m)
			}
			flag.Rsh(flag, 1)
		}
		return res
	}

	u := empty().Neg(empty().Add(empty().Add(empty().Exp(toBig(2), big.NewInt(62), nil),
		empty().Exp(toBig(2), big.NewInt(55), nil)), big.NewInt(1)))
	N := empty().Mul(empty().Exp(u, toBig(4), nil), toBig(36))
	N = empty().Add(N, empty().Mul(empty().Exp(u, toBig(3), nil), toBig(36)))
	N = empty().Add(N, empty().Mul(empty().Exp(u, toBig(2), nil), toBig(24)))
	N = empty().Add(N, empty().Mul(u, toBig(6)))
	N = empty().Add(N, toBig(1))
	R := empty().Lsh(toBig(2), 255)
	np := Np(R, N)

	fmt.Println("leftBin(2, 3):")
	fmt.Println(leftBin(toBig(2), toBig(3)))
	fmt.Println("leftBinMod(2, 3, 3):")
	fmt.Println(leftBinMod(toBig(2), toBig(3), toBig(3)))
	fmt.Printf("N:\n%d\n", N)
	fmt.Printf("N':\n%x\n", np)
	fmt.Println("2 mont:")
	fmt.Printf("%x\n", mont(toBig(2), empty().Mod(empty().Mul(R, R), N), R, N, np))
	fmt.Printf("%x\n", empty().Mod(empty().Mul(toBig(2), R), N))
	fmt.Println("3 mont:")
	fmt.Printf("%x\n", empty().Mod(empty().Mul(toBig(3), R), N))
	fmt.Println("2 * 3 mont:")
	fmt.Printf("%x\n", empty().Mod(empty().Mul(toBig(2*3), R), N))
}
