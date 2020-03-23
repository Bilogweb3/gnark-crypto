package fp2

import "github.com/consensys/gurvy/internal/tower"

// CodeSource is the aggregated source code
var CodeSource []string

// CodeTest is the aggregated test code
var CodeTest []string

// CodeTestPoints is the aggregated test points code
var CodeTestPoints []string

func init() {
	CodeSource = []string{
		base,
		mul,
		Inline,
	}

	CodeTest = []string{
		tower.Tests,
		customTests,
	}

	CodeTestPoints = []string{
		tower.TestPoints,
	}
}

const base = `
// Code generated by internal/fp2 DO NOT EDIT 

package {{.PackageName}}

import (
	"github.com/consensys/gurvy/{{.PackageName}}/fp"
)

// {{.Name}} is a degree-two finite field extension of fp.Element:
// A0 + A1u where u^2 == {{.Fp2NonResidue}} is a quadratic nonresidue in fp

type {{.Name}} struct {
	A0, A1 fp.Element
}

// SetString sets a {{.Name}} element from strings
func (z *{{.Name}}) SetString(s1, s2 string) *{{.Name}} {
	z.A0.SetString(s1)
	z.A1.SetString(s2)
	return z
}

func (z *{{.Name}}) SetZero() *{{.Name}} {
	z.A0.SetZero()
	z.A1.SetZero()
	return z
}

// Clone returns a copy of self
func (z *{{.Name}}) Clone() *{{.Name}} {
	return &{{.Name}}{
		A0: z.A0,
		A1: z.A1,
	}
}

// Set sets an {{.Name}} from x
func (z *{{.Name}}) Set(x *{{.Name}}) *{{.Name}} {
	z.A0.Set(&x.A0)
	z.A1.Set(&x.A1)
	return z
}

// Set sets z to 1
func (z *{{.Name}}) SetOne() *{{.Name}} {
	z.A0.SetOne()
	z.A1.SetZero()
	return z
}

// SetRandom sets a0 and a1 to random values
func (z *{{.Name}}) SetRandom() *{{.Name}} {
	z.A0.SetRandom()
	z.A1.SetRandom()
	return z
}

// Equal returns true if the two elements are equal, fasle otherwise
func (z *{{.Name}}) Equal(x *{{.Name}}) bool {
	return z.A0.Equal(&x.A0) && z.A1.Equal(&x.A1)
}

// Equal returns true if the two elements are equal, fasle otherwise
func (z *{{.Name}}) IsZero() bool {
	return z.A0.IsZero() && z.A1.IsZero()
}

// Neg negates an {{.Name}} element
func (z *{{.Name}}) Neg(x *{{.Name}}) *{{.Name}} {
	z.A0.Neg(&x.A0)
	z.A1.Neg(&x.A1)
	return z
}

// String implements Stringer interface for fancy printing
func (z *{{.Name}}) String() string {
	return (z.A0.String() + "+" + z.A1.String() + "*u")
}

// ToMont converts to mont form
func (z *{{.Name}}) ToMont() *{{.Name}} {
	z.A0.ToMont()
	z.A1.ToMont()
	return z
}

// FromMont converts from mont form
func (z *{{.Name}}) FromMont() *{{.Name}} {
	z.A0.FromMont()
	z.A1.FromMont()
	return z
}

// Add adds two elements of {{.Name}}
func (z *{{.Name}}) Add(x, y *{{.Name}}) *{{.Name}} {
	z.A0.Add(&x.A0, &y.A0)
	z.A1.Add(&x.A1, &y.A1)
	return z
}

// AddAssign adds x to z
func (z *{{.Name}}) AddAssign(x *{{.Name}}) *{{.Name}} {
	z.A0.AddAssign(&x.A0)
	z.A1.AddAssign(&x.A1)
	return z
}

// Sub two elements of {{.Name}}
func (z *{{.Name}}) Sub(x, y *{{.Name}}) *{{.Name}} {
	z.A0.Sub(&x.A0, &y.A0)
	z.A1.Sub(&x.A1, &y.A1)
	return z
}

// SubAssign subs x from z
func (z *{{.Name}}) SubAssign(x *{{.Name}}) *{{.Name}} {
	z.A0.SubAssign(&x.A0)
	z.A1.SubAssign(&x.A1)
	return z
}

// Double doubles an {{.Name}} element
func (z *{{.Name}}) Double(x *{{.Name}}) *{{.Name}} {
	z.A0.Double(&x.A0)
	z.A1.Double(&x.A1)
	return z
}
`
