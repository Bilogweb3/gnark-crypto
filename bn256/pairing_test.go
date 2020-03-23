// Code generated by internal/pairing DO NOT EDIT
package bn256

import (
	"testing"

	"github.com/consensys/gurvy/bn256/fp"
	"github.com/consensys/gurvy/bn256/fr"
)

func TestPairingLineEval(t *testing.T) {

	G := G2Jac{}
	G.X.SetString("9632395095185379999533066141603122875949398208988555443965873841297402654342",
		"19450252117195313925432134971824010362990079611665366649991932746671645587645")
	G.Y.SetString("13239408568242169453745727884443206250074028527951082640891183851714572387720",
		"2520488583727600099537136301403866807710407226542998911254294118959203578262")
	G.Z.SetString("1",
		"0")

	H := G2Jac{}
	H.X.SetString("4530090633702281412192028985619138850948035436035867170944613703864893134673",
		"11763830287135508066855029300189693921150081474419219185678083042857354676820")
	H.Y.SetString("13655518303190975472485738431514725885449587406831405681539688568601822661241",
		"6670653009080699707124366082582497740207251583656225538428663533545504823643")
	H.Z.SetString("1",
		"0")

	var a, b, c fp.Element
	a.SetString("19515237996314166984214782036479482283774444027267835315583249221845302746118")
	b.SetString("532769920444494182800832779540002510383749869333290362263178053098069316746")
	c.SetString("1")
	P := G1Jac{}
	P.X = a
	P.Y = b
	P.Z = c

	var Paff G1Affine
	P.ToAffineFromJac(&Paff)

	lRes := &lineEvalRes{}
	lineEvalJac(G, H, &Paff, lRes)

	var expectedA, expectedB, expectedC e2
	expectedA.SetString("20931116379304672891639612952659718019832871436773720309370367703766193898366",
		"1227691887220329021277611524657925318059159105924776637288096476583462700785")
	expectedB.SetString("11075701809971638713510823345421988501476980830739507328194508302903924813857",
		"12185932250924285590604781728248318914343702835207456116739570443012761818757")
	expectedC.SetString("20296547859951799284675362603476655855080142116855216859667285955222541979778",
		"3309163358421889473325056160526496300609105424816625582436053461638362255177")

	if !lRes.r1.Equal(&expectedA) {
		t.Fatal("Error A coeff")
	}
	if !lRes.r0.Equal(&expectedB) {
		t.Fatal("Error A coeff")
	}
	if !lRes.r2.Equal(&expectedC) {
		t.Fatal("Error A coeff")
	}
}

func TestMagicPairing(t *testing.T) {

	curve := BN256()

	var r1, r2 e12

	r1.SetRandom()
	r2.SetRandom()

	res1 := curve.FinalExponentiation(&r1)
	res2 := curve.FinalExponentiation(&r2)

	if res1.Equal(&res2) {
		t.Fatal("TestMagicPairing failed")
	}
}

func TestComputePairing(t *testing.T) {

	curve := BN256()

	G := curve.g2Gen.Clone()
	P := curve.g1Gen.Clone()
	sG := G.Clone()
	sP := P.Clone()

	var Gaff, sGaff G2Affine
	var Paff, sPaff G1Affine

	// checking bilinearity

	// check 1
	scalar := fr.Element{123}
	sG.ScalarMul(curve, sG, scalar)
	sP.ScalarMul(curve, sP, scalar)

	var mRes1, mRes2, mRes3 e12

	P.ToAffineFromJac(&Paff)
	sP.ToAffineFromJac(&sPaff)
	G.ToAffineFromJac(&Gaff)
	sG.ToAffineFromJac(&sGaff)

	res1 := curve.FinalExponentiation(curve.MillerLoop(Paff, sGaff, &mRes1))
	res2 := curve.FinalExponentiation(curve.MillerLoop(sPaff, Gaff, &mRes2))

	if !res1.Equal(&res2) {
		t.Fatal("pairing failed")
	}

	// check 2
	s1G := G.Clone()
	s2G := G.Clone()
	s3G := G.Clone()
	s1 := fr.Element{29372983}
	s2 := fr.Element{209302420904}
	var s3 fr.Element
	s3.Add(&s1, &s2)
	s1G.ScalarMul(curve, s1G, s1)
	s2G.ScalarMul(curve, s2G, s2)
	s3G.ScalarMul(curve, s3G, s3)

	var s1Gaff, s2Gaff, s3Gaff G2Affine
	s1G.ToAffineFromJac(&s1Gaff)
	s2G.ToAffineFromJac(&s2Gaff)
	s3G.ToAffineFromJac(&s3Gaff)

	rs1 := curve.FinalExponentiation(curve.MillerLoop(Paff, s1Gaff, &mRes1))
	rs2 := curve.FinalExponentiation(curve.MillerLoop(Paff, s2Gaff, &mRes2))
	rs3 := curve.FinalExponentiation(curve.MillerLoop(Paff, s3Gaff, &mRes3))
	rs1.Mul(&rs2, &rs1)
	if !rs3.Equal(&rs1) {
		t.Fatal("pairing failed2")
	}

}

//--------------------//
//     benches		  //
//--------------------//

func BenchmarkLineEval(b *testing.B) {

	curve := BN256()

	H := G2Jac{}
	H.ScalarMul(curve, &curve.g2Gen, fr.Element{1213})

	lRes := &lineEvalRes{}
	var g1GenAff G1Affine
	curve.g1Gen.ToAffineFromJac(&g1GenAff)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lineEvalJac(curve.g2Gen, H, &g1GenAff, lRes)
	}

}

func BenchmarkPairing(b *testing.B) {

	curve := BN256()

	var mRes e12

	var g1GenAff G1Affine
	var g2GenAff G2Affine

	curve.g1Gen.ToAffineFromJac(&g1GenAff)
	curve.g2Gen.ToAffineFromJac(&g2GenAff)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.FinalExponentiation(curve.MillerLoop(g1GenAff, g2GenAff, &mRes))
	}
}

func BenchmarkFinalExponentiation(b *testing.B) {

	var a e12

	curve := BN256()

	a.SetString(
		"1382424129690940106527336948935335363935127549146605398842626667204683483408227749",
		"0121296909401065273369489353353639351275491466053988426266672046834834082277499690",
		"7336948129690940106527336948935335363935127549146605398842626667204683483408227749",
		"6393512129690940106527336948935335363935127549146605398842626667204683483408227749",
		"2581296909401065273369489353353639351275491466053988426266672046834834082277496644",
		"5331296909401065273369489353353639351275491466053988426266672046834834082277495363",
		"1296909401065273369489353353639351275491466053988426266672046834834082277491382424",
		"0129612969094010652733694893533536393512754914660539884262666720468348340822774990",
		"7336948129690940106527336948935335363935127549146605398842626667204683483408227749",
		"6393129690940106527336948935335363935127549146605398842626667204683483408227749512",
		"2586641296909401065273369489353353639351275491466053988426266672046834834082277494",
		"5312969094010652733694893533536393512754914660539884262666720468348340822774935363")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.FinalExponentiation(&a)
	}

}
