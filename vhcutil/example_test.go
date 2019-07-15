package vhcutil_test

import (
	"fmt"
	"math"

	"github.com/valhallacoin/vhcd/vhcutil"
)

func ExampleAmount() {

	a := vhcutil.Amount(0)
	fmt.Println("Zero Atom:", a)

	a = vhcutil.Amount(1e8)
	fmt.Println("100,000,000 Atoms:", a)

	a = vhcutil.Amount(1e5)
	fmt.Println("100,000 Atoms:", a)
	// Output:
	// Zero Atom: 0 VHC
	// 100,000,000 Atoms: 1 VHC
	// 100,000 Atoms: 0.001 VHC
}

func ExampleNewAmount() {
	amountOne, err := vhcutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := vhcutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := vhcutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := vhcutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 VHC
	// 0.01234567 VHC
	// 0 VHC
	// invalid coin amount
}

func ExampleAmount_unitConversions() {
	amount := vhcutil.Amount(44433322211100)

	fmt.Println("Atom to kCoin:", amount.Format(vhcutil.AmountKiloCoin))
	fmt.Println("Atom to Coin:", amount)
	fmt.Println("Atom to MilliCoin:", amount.Format(vhcutil.AmountMilliCoin))
	fmt.Println("Atom to MicroCoin:", amount.Format(vhcutil.AmountMicroCoin))
	fmt.Println("Atom to Atom:", amount.Format(vhcutil.AmountAtom))

	// Output:
	// Atom to kCoin: 444.333222111 kVHC
	// Atom to Coin: 444333.222111 VHC
	// Atom to MilliCoin: 444333222.111 mVHC
	// Atom to MicroCoin: 444333222111 μVHC
	// Atom to Atom: 44433322211100 Atom
}
