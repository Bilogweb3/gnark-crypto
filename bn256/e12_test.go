// Code generated by internal/tower DO NOT EDIT
package bn256

import (
	"reflect"
	"testing"
)

type e12TestPoint struct {
	in  [2]e12
	out [17]e12
}

var e12TestPoints []e12TestPoint

// TODO this method is the same everywhere. move it someplace central and call it "compare"
func e12compare(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Fatal("\nexpect:\t", want, "\ngot:\t", got)
	}
}

func e12check(t *testing.T, f func(*e12, *e12, *e12) *e12, m int) {

	if len(e12TestPoints) < 1 {
		t.Log("no tests to run")
	}

	for i := range e12TestPoints {
		var receiver e12
		var out *e12
		var inCopies [len(e12TestPoints[i].in)]e12

		for j := range inCopies {
			inCopies[j].Set(&e12TestPoints[i].in[j])
		}

		// receiver, return value both set to result
		out = f(&receiver, &inCopies[0], &inCopies[1])

		e12compare(t, receiver, e12TestPoints[i].out[m]) // receiver correct
		e12compare(t, *out, e12TestPoints[i].out[m])     // return value correct
		for j := range inCopies {
			e12compare(t, inCopies[j], e12TestPoints[i].in[j]) // inputs unchanged
		}

		// receiver == one of the inputs
		for j := range inCopies {
			out = f(&inCopies[j], &inCopies[0], &inCopies[1])

			e12compare(t, inCopies[j], e12TestPoints[i].out[m]) // receiver correct
			e12compare(t, *out, e12TestPoints[i].out[m])        // return value correct
			for k := range inCopies {
				if k == j {
					continue
				}
				e12compare(t, inCopies[k], e12TestPoints[i].in[k]) // other inputs unchanged
			}
			inCopies[j].Set(&e12TestPoints[i].in[j]) // reset input for next tests
		}
	}
}

//--------------------//
//     tests		  //
//--------------------//

func TestE12Add(t *testing.T) {
	e12check(t, (*e12).Add, 0)
}

func TestE12Sub(t *testing.T) {
	e12check(t, (*e12).Sub, 1)
}

func TestE12Mul(t *testing.T) {
	e12check(t, (*e12).Mul, 2)
}

func TestE12MulByV(t *testing.T) {
	e12check(t, (*e12).MulByVBinary, 3)
}

func TestE12MulByVW(t *testing.T) {
	e12check(t, (*e12).MulByVWBinary, 4)
}

func TestE12MulByV2W(t *testing.T) {
	e12check(t, (*e12).MulByV2WBinary, 5)
}

func TestE12MulByV2NRInv(t *testing.T) {
	e12check(t, (*e12).MulByV2NRInvBinary, 6)
}

func TestE12MulByVWNRInv(t *testing.T) {
	e12check(t, (*e12).MulByVWNRInvBinary, 7)
}

func TestE12MulByWNRInv(t *testing.T) {
	e12check(t, (*e12).MulByWNRInvBinary, 8)
}

func TestE12Square(t *testing.T) {
	e12check(t, (*e12).SquareBinary, 9)
}

func TestE12Inverse(t *testing.T) {
	e12check(t, (*e12).InverseBinary, 10)
}

func TestE12Conjugate(t *testing.T) {
	e12check(t, (*e12).ConjugateBinary, 11)
}

func TestE12Frobenius(t *testing.T) {
	e12check(t, (*e12).FrobeniusBinary, 12)
}

func TestE12FrobeniusSquare(t *testing.T) {
	e12check(t, (*e12).FrobeniusSquareBinary, 13)
}

func TestE12FrobeniusCube(t *testing.T) {
	e12check(t, (*e12).FrobeniusCubeBinary, 14)
}

func TestE12Expt(t *testing.T) {
	e12check(t, (*e12).ExptBinary, 15)
}

func TestE12FinalExponentiation(t *testing.T) {
	e12check(t, (*e12).FinalExponentiationBinary, 16)
}

//--------------------//
//     benches		  //
//--------------------//

var e12BenchIn1, e12BenchIn2, e12BenchOut e12

func BenchmarkE12Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.Add(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12Sub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.Sub(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12Mul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.Mul(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12MulByV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.MulByVBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12MulByVW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.MulByVWBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12MulByV2W(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.MulByV2WBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12MulByV2NRInv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.MulByV2NRInvBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12MulByVWNRInv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.MulByVWNRInvBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12MulByWNRInv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.MulByWNRInvBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12Square(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.SquareBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12Inverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.InverseBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12Conjugate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.ConjugateBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12Frobenius(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.FrobeniusBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12FrobeniusSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.FrobeniusSquareBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12FrobeniusCube(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.FrobeniusCubeBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12Expt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.ExptBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

func BenchmarkE12FinalExponentiation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e12BenchOut.FinalExponentiationBinary(&e12BenchIn1, &e12BenchIn2)
	}
}

//-------------------------------------//
// unary helpers for e12 methods
//-------------------------------------//

// SquareBinary a binary wrapper for Square
func (z *e12) SquareBinary(x, y *e12) *e12 {
	return z.Square(x)
}

// InverseBinary a binary wrapper for Inverse
func (z *e12) InverseBinary(x, y *e12) *e12 {
	return z.Inverse(x)
}

// ConjugateBinary a binary wrapper for Conjugate
func (z *e12) ConjugateBinary(x, y *e12) *e12 {
	return z.Conjugate(x)
}

// FrobeniusBinary a binary wrapper for Frobenius
func (z *e12) FrobeniusBinary(x, y *e12) *e12 {
	return z.Frobenius(x)
}

// FrobeniusSquareBinary a binary wrapper for FrobeniusSquare
func (z *e12) FrobeniusSquareBinary(x, y *e12) *e12 {
	return z.FrobeniusSquare(x)
}

// FrobeniusCubeBinary a binary wrapper for FrobeniusCube
func (z *e12) FrobeniusCubeBinary(x, y *e12) *e12 {
	return z.FrobeniusCube(x)
}

// FinalExponentiationBinary a binary wrapper for FinalExponentiation
func (z *e12) FinalExponentiationBinary(x, y *e12) *e12 {
	return z.FinalExponentiation(x)
}

//-------------------------------------//
// custom helpers for e12 methods
//-------------------------------------//

// ExptBinary a binary wrapper for Expt
func (z *e12) ExptBinary(x, y *e12) *e12 {
	z.Expt(x)

	// if tAbsVal is negative then need to undo the conjugation in order to match the test point

	return z
}

// MulByVBinary a binary wrapper for MulByV
func (z *e12) MulByVBinary(x, y *e12) *e12 {
	yCopy := y.C0.B1
	z.MulByV(x, &yCopy)
	return z
}

// MulByVWBinary a binary wrapper for MulByVW
func (z *e12) MulByVWBinary(x, y *e12) *e12 {
	yCopy := y.C1.B1
	z.MulByVW(x, &yCopy)
	return z
}

// MulByV2WBinary a binary wrapper for MulByV2W
func (z *e12) MulByV2WBinary(x, y *e12) *e12 {
	yCopy := y.C1.B2
	z.MulByV2W(x, &yCopy)
	return z
}

// MulByV2NRInvBinary a binary wrapper for MulByV2NRInv
func (z *e12) MulByV2NRInvBinary(x, y *e12) *e12 {
	yCopy := y.C0.B2
	z.MulByV2NRInv(x, &yCopy)
	return z
}

// MulByVWNRInvBinary a binary wrapper for MulByVWNRInv
func (z *e12) MulByVWNRInvBinary(x, y *e12) *e12 {
	yCopy := y.C1.B1
	z.MulByVWNRInv(x, &yCopy)
	return z
}

// MulByWNRInvBinary a binary wrapper for MulByWNRInv
func (z *e12) MulByWNRInvBinary(x, y *e12) *e12 {
	yCopy := y.C1.B0
	z.MulByWNRInv(x, &yCopy)
	return z
}
