package main

import (
	"fmt"
	"math/big"
)

func main() {
	toBig := func(x int64) *big.Int { return big.NewInt(x) }

	empty := func() *big.Int { return new(big.Int) }

	np := func(r, n *big.Int) *big.Int {
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

	// montTest := func(a, R, N *big.Int) *big.Int { return empty().Mod(empty().Mul(a, R), N) }

	leftBinExp := func(x, y *big.Int) *big.Int {
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
	} // x^y

	leftBinExpMod := func(x, y, m *big.Int) *big.Int {
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
	} // x^y%m

	montREDC := func(T, NP, R, N *big.Int) *big.Int {
		m := leftBinExpMod(empty().Mul(leftBinExpMod(T, toBig(1), R), NP), toBig(1), R)
		res := empty().Rsh(empty().Add(T, empty().Mul(m, N)), 256)
		if res.Cmp(N) >= 0 {
			return empty().Sub(res, N)
		} else {
			return res
		}
	} // T/R%N

	/*U := empty().Neg(empty().Add(empty().Add(empty().Exp(toBig(2), big.NewInt(62), nil),
		empty().Exp(toBig(2), big.NewInt(55), nil)), big.NewInt(1)))
	tempN := empty().Mul(empty().Exp(U, toBig(4), nil), toBig(36))
	tempN = empty().Add(tempN, empty().Mul(empty().Exp(U, toBig(3), nil), toBig(36)))
	tempN = empty().Add(tempN, empty().Mul(empty().Exp(U, toBig(2), nil), toBig(24)))
	tempN = empty().Add(tempN, empty().Mul(U, toBig(6)))
	N := empty().Add(tempN, toBig(1))*/
	N, _ := empty().SetString("16798108731015832284940804142231733909889187121439069848933715426072753864723", 10)
	R := empty().Exp(toBig(2), toBig(256), nil)
	NP := np(R, N)
	R2ModN := leftBinExpMod(R, toBig(2), N)
	MONT2 := montREDC(empty().Mul(toBig(2), R2ModN), NP, R, N)
	MONT3 := montREDC(empty().Mul(toBig(3), R2ModN), NP, R, N)
	MONT6 := montREDC(empty().Mul(toBig(6), R2ModN), NP, R, N)

	fmt.Printf("N:\n%x\n", N)
	fmt.Printf("N':\n%x\n", NP)
	fmt.Println("R^2ModN:")
	fmt.Printf("%x\n", R2ModN)

	fmt.Println("2 mont:")
	fmt.Printf("%x\n", MONT2)
	fmt.Println("convert mont2 to int:")
	fmt.Printf("%d\n", montREDC(MONT2, NP, R, N))
	fmt.Println("3 mont:")
	fmt.Printf("%x\n", MONT3)
	fmt.Println("convert mont3 to int:")
	fmt.Printf("%d\n", montREDC(MONT3, NP, R, N))
	fmt.Println("6 mont:")
	fmt.Printf("%x\n", MONT6)
	fmt.Println("3 + 3 mont:")
	fmt.Printf("%x\n", leftBinExpMod(empty().Add(MONT3, MONT3), toBig(1), N))
	fmt.Println("2 * 3 mont:")
	fmt.Printf("%x\n", leftBinExpMod(empty().Mul(MONT2, toBig(3)), toBig(1), N))
	fmt.Println("convert mont6 to int:")
	fmt.Printf("%d\n", montREDC(MONT6, NP, R, N))
	fmt.Println("======================================")
	fmt.Println("leftBin(2, 3):")
	fmt.Println(leftBinExp(toBig(2), toBig(3)))
	fmt.Println("leftBinExpMod(2, 1, 3):")
	fmt.Println(leftBinExpMod(toBig(2), toBig(1), toBig(3)))
	fmt.Println("leftBinExpMod(2, 40710, N):")
	fmt.Printf("%d\n", leftBinExpMod(toBig(2), toBig(40710), N))
}
