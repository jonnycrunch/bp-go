package bp

import (
	"testing"
	"math/big"
	"fmt"
	"crypto/rand"
)

func TestInnerProductProveLen1(t *testing.T) {
	println("TestInnerProductProve1")
	CP = NewECPrimeGroupKey(1)
	a := make([]*big.Int, 1)
	b := make([]*big.Int, 1)

	a[0] = big.NewInt(2)

	b[0] = big.NewInt(2)

	c := InnerProduct(a, b)

	P := TwoVectorPCommit(a, b)

	ipp := InnerProductProve(a, b, c, P,CP.U, CP.BPG, CP.BPH)

	if InnerProductVerify(c, P, CP.U, CP.BPG, CP.BPH, ipp){
		println("Inner Product Proof correct")
	} else {
		println("Inner Product Proof incorrect")
	}
}

func TestInnerProductProveLen2(t *testing.T) {
	println("TestInnerProductProve2")
	CP = NewECPrimeGroupKey(2)
	a := make([]*big.Int, 2)
	b := make([]*big.Int, 2)

	a[0] = big.NewInt(2)
	a[1] = big.NewInt(3)

	b[0] = big.NewInt(2)
	b[1] = big.NewInt(3)

	c := InnerProduct(a, b)

	P := TwoVectorPCommit(a, b)

	ipp := InnerProductProve(a, b, c, P,CP.U, CP.BPG, CP.BPH)

	if InnerProductVerify(c, P,CP.U,CP.BPG, CP.BPH, ipp){
		println("Inner Product Proof correct")
	} else {
		println("Inner Product Proof incorrect")
	}
}


func TestInnerProductProveLen4(t *testing.T) {
	println("TestInnerProductProve4")
	CP = NewECPrimeGroupKey(4)
	a := make([]*big.Int, 4)
	b := make([]*big.Int, 4)

	a[0] = big.NewInt(1)
	a[1] = big.NewInt(1)
	a[2] = big.NewInt(1)
	a[3] = big.NewInt(1)

	b[0] = big.NewInt(1)
	b[1] = big.NewInt(1)
	b[2] = big.NewInt(1)
	b[3] = big.NewInt(1)

	c := InnerProduct(a, b)

	P := TwoVectorPCommit(a, b)

	ipp := InnerProductProve(a, b, c, P,CP.U, CP.BPG, CP.BPH)

	if InnerProductVerify(c, P,CP.U,CP.BPG, CP.BPH, ipp){
		println("Inner Product Proof correct")
	} else {
		println("Inner Product Proof incorrect")
	}
}

func TestInnerProductProveLen8(t *testing.T) {
	println("TestInnerProductProve8")
	CP = NewECPrimeGroupKey(8)
	a := make([]*big.Int, 8)
	b := make([]*big.Int, 8)

	a[0] = big.NewInt(1)
	a[1] = big.NewInt(1)
	a[2] = big.NewInt(1)
	a[3] = big.NewInt(1)
	a[4] = big.NewInt(1)
	a[5] = big.NewInt(1)
	a[6] = big.NewInt(1)
	a[7] = big.NewInt(1)

	b[0] = big.NewInt(2)
	b[1] = big.NewInt(2)
	b[2] = big.NewInt(2)
	b[3] = big.NewInt(2)
	b[4] = big.NewInt(2)
	b[5] = big.NewInt(2)
	b[6] = big.NewInt(2)
	b[7] = big.NewInt(2)

	c := InnerProduct(a, b)

	P := TwoVectorPCommit(a, b)

	ipp := InnerProductProve(a, b, c, P,CP.U, CP.BPG, CP.BPH)

	if InnerProductVerify(c, P,CP.U, CP.BPG, CP.BPH, ipp){
		println("Inner Product Proof correct")
	} else {
		println("Inner Product Proof incorrect")
	}
}

func TestInnerProductProveLen64Rand(t *testing.T) {
	println("TestInnerProductProveLen64Rand")
	CP = NewECPrimeGroupKey(64)
	a := RandVector(64)
	b := RandVector(64)

	c := InnerProduct(a, b)

	P := TwoVectorPCommit(a, b)

	ipp := InnerProductProve(a, b, c, P,CP.U, CP.BPG, CP.BPH)

	if InnerProductVerify(c, P,CP.U,CP.BPG, CP.BPH, ipp){
		println("Inner Product Proof correct")
	} else {
		println("Inner Product Proof incorrect")
		fmt.Printf("Values Used: \n\ta = %s\n\tb = %s\n", a, b)
	}

}

func TestValueBreakdown(t *testing.T){
	v := big.NewInt(20)
	yes := reverse(StrToBigIntArray(PadLeft(fmt.Sprintf("%b", v), "0", 64)))
	vec2 := PowerVector(64, big.NewInt(2))

	calc := InnerProduct(yes, vec2)

	if v.Cmp(calc) != 0 {
		println("Binary Value Breakdown - Failure :(")
		fmt.Println(yes)
		fmt.Println(vec2)
		fmt.Println(calc)
	} else {
		println("Binary Value Breakdown - Success!")
		//fmt.Println(yes)
		//fmt.Println(vec2)
		//fmt.Println(calc)
	}
}

func TestValueBreakdownRand(t *testing.T){
	v, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2), big.NewInt(64), CP.N))
	check(err)

	yes := reverse(StrToBigIntArray(PadLeft(fmt.Sprintf("%b", v), "0", 64)))
	vec2 := PowerVector(64, big.NewInt(2))

	calc := InnerProduct(yes, vec2)

	if v.Cmp(calc) != 0 {
		println("Binary Value Breakdown - Failure :(")
		fmt.Println(yes)
		fmt.Println(vec2)
		fmt.Println(calc)
	} else {
		println("Binary Value Breakdown - Success!")
	}

}

func TestVectorHadamard(t *testing.T) {
	a := make([]*big.Int, 5)
	a[0] = big.NewInt(1)
	a[1] = big.NewInt(1)
	a[2] = big.NewInt(1)
	a[3] = big.NewInt(1)
	a[4] = big.NewInt(1)

	c := VectorHadamard(a, a)

	success := true

	for i := range c {
		if c[i].Cmp(a[i]) != 0 {
			success = false
		}
	}
	if !success{
		println("Failure in the Hadamard")
	} else {
		println("Hadamard good")
	}
}

func TestRPVerify1(t *testing.T) {
	CP = NewECPrimeGroupKey(64)
	// Testing smallest number in range
	if RPVerify(RPProve(big.NewInt(0))) {
		println("Range Proof Verification works")
	} else {
		println("*****Range Proof FAILURE")
	}
}

func TestRPVerify2(t *testing.T) {
	CP = NewECPrimeGroupKey(64)
	// Testing largest number in range
	if RPVerify(RPProve(new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(63), CP.N), big.NewInt(1)))) {
		println("Range Proof Verification works")
	} else {
		println("*****Range Proof FAILURE")
	}
}


func TestRPVerify3(t *testing.T) {
	CP = NewECPrimeGroupKey(64)
	// Testing the value 3
	if RPVerify(RPProve(big.NewInt(3))) {
		println("Range Proof Verification works")
	} else {
		println("*****Range Proof FAILURE")
	}
}

func TestRPVerify4(t *testing.T) {
	CP = NewECPrimeGroupKey(32)
	// Testing smallest number in range
	if RPVerify(RPProve(big.NewInt(0))) {
		println("Range Proof Verification works")
	} else {
		println("*****Range Proof FAILURE")
	}
}

func TestRPVerifyRand(t *testing.T) {
	CP = NewECPrimeGroupKey(64)

	ran, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2), big.NewInt(64), CP.N))
	check(err)

	// Testing the value 3
	if RPVerify(RPProve(ran)) {
		println("Range Proof Verification works")
	} else {
		println("*****Range Proof FAILURE")
		fmt.Printf("Random Value: %s", ran.String())
	}
}

func TestMultiRPVerify1(t *testing.T) {
	values := []*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	CP = NewECPrimeGroupKey(64 * len(values))
	// Testing smallest number in range
	proof := MRPProve(values)
	proofString := fmt.Sprintf("%s", proof)

	fmt.Println(len(proofString)) // length is good measure of bytes, correct?

	if MRPVerify(proof) {
		println("Multi Range Proof Verification works")
	} else {
		println("***** Multi Range Proof FAILURE")
	}
}

func TestMultiRPVerify2(t *testing.T) {
	values := []*big.Int{big.NewInt(0)}
	CP = NewECPrimeGroupKey(64 * len(values))
	// Testing smallest number in range
	if MRPVerify(MRPProve(values)) {
		println("Multi Range Proof Verification works")
	} else {
		println("***** Multi Range Proof FAILURE")
	}
}

func TestMultiRPVerify3(t *testing.T) {
	values := []*big.Int{big.NewInt(0), big.NewInt(1)}
	CP = NewECPrimeGroupKey(64 * len(values))
	// Testing smallest number in range
	if MRPVerify(MRPProve(values)) {
		println("Multi Range Proof Verification works")
	} else {
		println("***** Multi Range Proof FAILURE")
	}
}

func TestMultiRPVerify4(t *testing.T) {
	for j := 1; j < 33; j=2*j {
		values := make([]*big.Int, j)
		for k := 0; k < j; k++{
			values[k] = big.NewInt(0)
		}

		CP = NewECPrimeGroupKey(64 * len(values))
		// Testing smallest number in range
		proof := MRPProve(values)
		proofString := fmt.Sprintf("%s", proof)

		fmt.Println(len(proofString)) // length is good measure of bytes, correct?

		if MRPVerify(proof) {
			println("Multi Range Proof Verification works")
		} else {
			println("***** Multi Range Proof FAILURE")
		}
	}
}

func TestVectorPCommit3(t *testing.T) {
	println("TestVectorPCommit3")
	CP = NewECPrimeGroupKey(3)

	v := make([]*big.Int, 3)
	for j := range v {
		v[j] = big.NewInt(2)
	}

	output, r := VectorPCommit(v)

	if len(r) != 3 {
		println("Failure - rvalues doesn't match length of values")
	}
	// we will verify correctness by replicating locally and comparing output

	GVal := CP.BPG[0].Mult(v[0]).Add(CP.BPG[1].Mult(v[1]).Add(CP.BPG[2].Mult(v[2])))
	HVal := CP.BPH[0].Mult(r[0]).Add(CP.BPH[1].Mult(r[1]).Add(CP.BPH[2].Mult(r[2])))
	Comm := GVal.Add(HVal)

	if output.Equal(Comm) {
		println("Commitment correct")
	} else {
		println("Commitment failed")
	}
}

func TestInnerProduct(t *testing.T) {
	println("TestInnerProduct")
	a := make([]*big.Int, 4)
	b := make([]*big.Int, 4)

	a[0] = big.NewInt(2)
	a[1] = big.NewInt(2)
	a[2] = big.NewInt(2)
	a[3] = big.NewInt(2)

	b[0] = big.NewInt(2)
	b[1] = big.NewInt(2)
	b[2] = big.NewInt(2)
	b[3] = big.NewInt(2)

	c := InnerProduct(a, b)

	if c.Cmp(big.NewInt(16)) == 0 {
		println("Success - Innerproduct works with 2")
	} else{
		println("Failure - Innerproduct equal to ")
		println(c.String())
	}

}


func BenchmarkMRPVerifySize(b *testing.B) {
	for i := 0; i < b.N; i++{
		for j := 1; j < 257; j*=2 {
			values := make([]*big.Int, j)
			for k := 0; k < j; k++{
				values[k] = big.NewInt(0)
			}

			CP = NewECPrimeGroupKey(64 * len(values))
			// Testing smallest number in range
			proof := MRPProve(values)
			proofString := fmt.Sprintf("%s", proof)
			//fmt.Println(proofString)
			fmt.Printf("Size for %d values: %d bytes\n", j, len(proofString)) // length is good measure of bytes, correct?

			if MRPVerify(proof) {
				println("Multi Range Proof Verification works")
			} else {
				println("***** Multi Range Proof FAILURE")
			}
		}
	}
}

var result MultiRangeProof
var boores bool

func BenchmarkMRPProve16(b *testing.B) {
	j := 16
	values := make([]*big.Int, j)
	for k := 0; k < j; k++{
		values[k] = big.NewInt(0)
	}
	CP = NewECPrimeGroupKey(64 * len(values))
	var r MultiRangeProof
	for i := 0; i < b.N; i++{
		r = MRPProve(values)
	}

	result = r
}

func BenchmarkMRPVerify16(b *testing.B) {
	j := 16
	values := make([]*big.Int, j)
	for k := 0; k < j; k++{
		values[k] = big.NewInt(0)
	}
	CP = NewECPrimeGroupKey(64 * len(values))
	proof := MRPProve(values)

	var r bool
	for i := 0; i < b.N; i++{
		r = MRPVerify(proof)
	}
	boores = r
}

func BenchmarkMRPProve32(b *testing.B) {
	j := 32
	values := make([]*big.Int, j)
	for k := 0; k < j; k++{
		values[k] = big.NewInt(0)
	}
	CP = NewECPrimeGroupKey(64 * len(values))
	var r MultiRangeProof
	for i := 0; i < b.N; i++{
		r = MRPProve(values)
	}
	result = r
}

func BenchmarkMRPVerify32(b *testing.B) {
	j := 32
	values := make([]*big.Int, j)
	for k := 0; k < j; k++{
		values[k] = big.NewInt(0)
	}
	CP = NewECPrimeGroupKey(64 * len(values))
	proof := MRPProve(values)
	var r bool
	for i := 0; i < b.N; i++{
		r = MRPVerify(proof)
	}
	boores = r
}