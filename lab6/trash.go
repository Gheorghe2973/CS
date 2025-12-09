package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

// Mesajul din Laboratorul 2
const message = `the first one was waiting in the wings. this was poly-alphabeticsubstitution, in the form
of the straight-alphabet vigenere with shortrepeating keyword. the old objections to its
use, which boiled down tothe impossibility of correcting a garbled dispatch quickly
enough,vanished with the telegraph. it fulfilled the requirements of noncom-promisability
of  the  general  system  and  of  ease  of  key  changes.  moreover,it  had  the  reputation  of  being
unbreakable—which, if its cryptogramswere not divided into words, it largely was. the military
adopted  it  atonce.then,  in  1863,  a  retired  prussian  infantry  major  discovered  thegeneral
solution for the periodic polyalpha-betic substitution. at onestroke he demolished the
only impregnable structure in cryptography.signal officers, compelled to provide secure
communications,  huntedfrantically  for  new  field  ciphers.  they  found  many  good  ideas  in
thewritings  of  the  dilettante  cryptographers  who  had  proposed  ciphers  forthe  protection
of private messages. soon some of these systems wereserving in the various armies of europe
and the americas. more ideascame from army officers who had studied cryptography in the courses
insignal communication that the national military academies, such as st.cyr, had added in the
mid-1800s. inevitably, cryptanalysts—who wereeither amateurs or soldiers with a professional
interest, for fullprofessionals there were none—replied with new techniques for breakingthe
new ciphers. from the slow crawl of nomenclator days, when theintroduction of a special group
meaning disregard the preceding groupwould constitute a remarkable technical advance, the race
betweenoffense  and  defense  in  cryptology  acclerated  to  its  modern  pace.the  history  of
cryptology from the decade that saw both the death ofthe black chambers and the birth of the
telegraph  to  world  war  i  is  thusa  story  of  internal  development.  without  rossignols  or
willeses, andwithout major wars or diplomatic struggles, cryptology could notinfluence world
events,  and,  efcept  for  one  or  two  unusual  cases,  it  didnot.  the  telegraph  launched  this
evolution  of  cryptology.  it  broke  the  monopoly  of  the  nomenclator.  thenomenclator  had
reigned for 450 years as a general, all-purpose system,but it could not meet the new requirements
either of high-leveldiplomatic or military communications or of low-level
signalcommunications,  which  the  telegraph  had  engendered.  each  called  for  itsown  kind  of
cryptosystem, a specialized one. signal officers ranked thesesystems in a hierarchy, rising from
the simple and flefible and easilysolved to the eftensive and hard to solve. the telegraph
thus stimulatedthe invention of many new ciphers and, by reaction, many new methodsof crypt-
analysis,  and  compelled  their  arrangement  in  a  scale  ofcomplefity.many  of  these  ciphers  and
techniques  have  become  classic  and  are  inuse  today.  moreover,  cryptography  still  functions
through a hierarchyand employs a multitude of special systems. the telegraph therebyfurnished
cryptography with the structure and the content that it stillhas. it made it what it is
today. all these things have antecedents, and just as the telegraph itself did,so were there
precursors  of  the  cryptographic  systems  that  itengendered.one  cipher  system  invented
before  the  telegraph  was  so  far  ahead  ofits  time,  and  so  much  in  the  spirit  of  the  later
inventions, that it deservesto be classed with them. indeed, it deserves the front rank
among them,for this system was beyond doubt the most remarkable of all. so wellconceived was it
that today, more than a century and a half of rapidtechnological progress after its invention,
it remains in active use.but then it was invented by a remarkable man, a well-known
writer,agriculturalist, bibliophile, architect, diplomat, gadgeteer, and statesmannamed
thomas jefferson. he called it his "wheel cypher," and it seemslikely that he invented it
either during 1790 to 1793 or during 1797 to1800.`

// Hash-uri pre-calculate
const (
	// RIPEMD-128 pentru RSA (Sarcina 2)
	hashRIPEMD128Hex = "29f4c31d4911700655bcac70b51bcd36"

	// SHA3-512 pentru ElGamal (Sarcina 3)
	hashSHA3_512Hex = "bcc0a3f43a0991cfa3af12ec3480b05e3563922d46bb8145c4642bab5dbcb37c8064717a4020bef8d5aa633aecfb0d97d0b63035e0965437396ba0eb9330b406"
)

// ====================================================================================
// PARTEA 1: UTILITĂȚI GENERALE
// ====================================================================================

// hashToDecimal convertește hash-ul hexazecimal în reprezentare zecimală
func hashToDecimal(hashHex string) *big.Int {
	hashBytes, _ := hex.DecodeString(hashHex)
	return new(big.Int).SetBytes(hashBytes)
}

// printSeparator afișează un separator vizual
func printSeparator(title string) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 80))
}

// ====================================================================================
// PARTEA 2: RSA DIGITAL SIGNATURE (Sarcina 2)
// ====================================================================================

// RSAKeys structura pentru cheile RSA
type RSAKeys struct {
	N *big.Int // modulul n = p * q
	E *big.Int // exponentul public e
	D *big.Int // exponentul privat d
	P *big.Int // numărul prim p
	Q *big.Int // numărul prim q
}

// generateRSAKeys generează perechi de chei RSA cu n >= 3072 biți
func generateRSAKeys(bits int) (*RSAKeys, error) {
	fmt.Printf("\n[RSA] Generare chei RSA cu n >= %d biți...\n", bits)

	// Pentru n de 3072 biți, fiecare prim trebuie să fie de ~1536 biți
	primeBits := bits / 2

	fmt.Printf("[RSA] Generare număr prim p (%d biți)...\n", primeBits)
	p, err := rand.Prime(rand.Reader, primeBits)
	if err != nil {
		return nil, fmt.Errorf("eroare la generarea lui p: %v", err)
	}
	fmt.Printf("[RSA] p generat: %d biți\n", p.BitLen())

	fmt.Printf("[RSA] Generare număr prim q (%d biți)...\n", primeBits)
	q, err := rand.Prime(rand.Reader, primeBits)
	if err != nil {
		return nil, fmt.Errorf("eroare la generarea lui q: %v", err)
	}
	fmt.Printf("[RSA] q generat: %d biți\n", q.BitLen())

	// Calculăm n = p * q
	n := new(big.Int).Mul(p, q)
	fmt.Printf("[RSA] n = p × q calculat: %d biți\n", n.BitLen())

	// Calculăm φ(n) = (p-1) * (q-1)
	pMinus1 := new(big.Int).Sub(p, big.NewInt(1))
	qMinus1 := new(big.Int).Sub(q, big.NewInt(1))
	phi := new(big.Int).Mul(pMinus1, qMinus1)
	fmt.Printf("[RSA] φ(n) = (p-1) × (q-1) calculat\n")

	// Alegem e = 65537 (valoare standard)
	e := big.NewInt(65537)
	fmt.Printf("[RSA] Exponent public e = %d\n", e.Int64())

	// Calculăm d = e^(-1) mod φ(n)
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		return nil, fmt.Errorf("nu s-a putut calcula inversul modular")
	}
	fmt.Printf("[RSA] Exponent privat d calculat: %d biți\n", d.BitLen())

	return &RSAKeys{
		N: n,
		E: e,
		D: d,
		P: p,
		Q: q,
	}, nil
}

// rsaSign semnează un hash folosind cheia privată RSA
func rsaSign(hash *big.Int, keys *RSAKeys) *big.Int {
	fmt.Println("\n[RSA] Semnare mesaj...")
	fmt.Printf("[RSA] Hash-ul mesajului (decimal): %s\n", hash.String())

	// Semnătura: s = hash^d mod n
	signature := new(big.Int).Exp(hash, keys.D, keys.N)
	fmt.Printf("[RSA] Semnătură generată: %d biți\n", signature.BitLen())

	return signature
}

// rsaVerify verifică o semnătură RSA folosind cheia publică
func rsaVerify(hash *big.Int, signature *big.Int, keys *RSAKeys) bool {
	fmt.Println("\n[RSA] Verificare semnătură...")

	// Decriptare semnătură: hash_verificat = s^e mod n
	verifiedHash := new(big.Int).Exp(signature, keys.E, keys.N)
	fmt.Printf("[RSA] Hash verificat: %s\n", verifiedHash.String())
	fmt.Printf("[RSA] Hash original: %s\n", hash.String())

	// Comparăm hash-urile
	isValid := hash.Cmp(verifiedHash) == 0
	if isValid {
		fmt.Println("[RSA] ✓ Semnătura este VALIDĂ!")
	} else {
		fmt.Println("[RSA] ✗ Semnătura este INVALIDĂ!")
	}

	return isValid
}

// ====================================================================================
// PARTEA 3: ELGAMAL DIGITAL SIGNATURE (Sarcina 3)
// ====================================================================================

// ElGamalKeys structura pentru cheile ElGamal
type ElGamalKeys struct {
	P *big.Int // numărul prim p (dat în cerință)
	G *big.Int // generatorul g = 2
	X *big.Int // cheia privată x
	Y *big.Int // cheia publică y = g^x mod p
}

// ElGamalSignature structura pentru semnătura ElGamal
type ElGamalSignature struct {
	R *big.Int // r = g^k mod p
	S *big.Int // s = (hash - x*r) * k^(-1) mod (p-1)
}

// generateElGamalKeys generează chei ElGamal cu p și g dați
func generateElGamalKeys() (*ElGamalKeys, error) {
	fmt.Println("\n[ElGamal] Generare chei ElGamal...")

	// p este dat în cerință (2048 biți)
	pStr := "323170060713110073001535134778251633624880571334890751745884" +
		"34139269806834136210002792056362640164685458556357935330816928" +
		"82902308057347262527355474246124574102620252791657297286270630" +
		"03252634282131457669314142236542209411113486299916574782680342" +
		"30553086349050635557712219187890332729569696129743856241741236" +
		"23722519734640269185579776797682301462539793305801522685873076" +
		"11975324364674758554607150438968449403661304976978128542959586" +
		"59597567051283852132784468522925504568272879113720098931873959" +
		"14337417583782600027803497319855206060753323412260325468408812" +
		"0031105907484281003994966956119696956248629032338072839127039"

	p := new(big.Int)
	p.SetString(pStr, 10)
	fmt.Printf("[ElGamal] p (din cerință): %d biți\n", p.BitLen())

	// g este dat în cerință: g = 2
	g := big.NewInt(2)
	fmt.Printf("[ElGamal] g (din cerință): %d\n", g.Int64())

	// Generăm cheia privată x (aleatoriu, 1 < x < p-1)
	pMinus2 := new(big.Int).Sub(p, big.NewInt(2))
	x, err := rand.Int(rand.Reader, pMinus2)
	if err != nil {
		return nil, fmt.Errorf("eroare la generarea cheii private: %v", err)
	}
	x.Add(x, big.NewInt(1))
	fmt.Printf("[ElGamal] Cheia privată x generată: %d biți\n", x.BitLen())

	// Calculăm cheia publică y = g^x mod p
	y := new(big.Int).Exp(g, x, p)
	fmt.Printf("[ElGamal] Cheia publică y = g^x mod p calculată: %d biți\n", y.BitLen())

	return &ElGamalKeys{
		P: p,
		G: g,
		X: x,
		Y: y,
	}, nil
}

// elgamalSign semnează un hash folosind cheia privată ElGamal
func elgamalSign(hash *big.Int, keys *ElGamalKeys) (*ElGamalSignature, error) {
	fmt.Println("\n[ElGamal] Semnare mesaj...")
	fmt.Printf("[ElGamal] Hash-ul mesajului (decimal): %s\n", hash.String())

	pMinus1 := new(big.Int).Sub(keys.P, big.NewInt(1))

	var k, r, s *big.Int
	var err error

	// Încercăm să găsim un k valid (gcd(k, p-1) = 1)
	for {
		// Generăm k aleatoriu: 1 < k < p-1
		kMax := new(big.Int).Sub(keys.P, big.NewInt(2))
		k, err = rand.Int(rand.Reader, kMax)
		if err != nil {
			return nil, fmt.Errorf("eroare la generarea lui k: %v", err)
		}
		k.Add(k, big.NewInt(1))

		// Verificăm gcd(k, p-1) = 1
		gcd := new(big.Int).GCD(nil, nil, k, pMinus1)
		if gcd.Cmp(big.NewInt(1)) == 0 {
			break
		}
	}

	fmt.Printf("[ElGamal] k generat: %d biți (gcd(k, p-1) = 1)\n", k.BitLen())

	// Calculăm r = g^k mod p
	r = new(big.Int).Exp(keys.G, k, keys.P)
	fmt.Printf("[ElGamal] r = g^k mod p calculat: %d biți\n", r.BitLen())

	// Calculăm s = (hash - x*r) * k^(-1) mod (p-1)
	kInv := new(big.Int).ModInverse(k, pMinus1)
	if kInv == nil {
		return nil, fmt.Errorf("nu s-a putut calcula k^(-1)")
	}

	xr := new(big.Int).Mul(keys.X, r)
	xr.Mod(xr, pMinus1)

	hashMinusXr := new(big.Int).Sub(hash, xr)
	hashMinusXr.Mod(hashMinusXr, pMinus1)

	s = new(big.Int).Mul(hashMinusXr, kInv)
	s.Mod(s, pMinus1)

	fmt.Printf("[ElGamal] s = (hash - x*r) * k^(-1) mod (p-1) calculat\n")
	fmt.Printf("[ElGamal] Semnătură generată: (r=%d biți, s=%d biți)\n", r.BitLen(), s.BitLen())

	return &ElGamalSignature{R: r, S: s}, nil
}

// elgamalVerify verifică o semnătură ElGamal folosind cheia publică
func elgamalVerify(hash *big.Int, sig *ElGamalSignature, keys *ElGamalKeys) bool {
	fmt.Println("\n[ElGamal] Verificare semnătură...")

	// Verificăm condiții preliminare: 0 < r < p și 0 < s < p-1
	if sig.R.Cmp(big.NewInt(0)) <= 0 || sig.R.Cmp(keys.P) >= 0 {
		fmt.Println("[ElGamal] ✗ r nu este în intervalul valid (0 < r < p)")
		return false
	}

	pMinus1 := new(big.Int).Sub(keys.P, big.NewInt(1))
	if sig.S.Cmp(big.NewInt(0)) <= 0 || sig.S.Cmp(pMinus1) >= 0 {
		fmt.Println("[ElGamal] ✗ s nu este în intervalul valid (0 < s < p-1)")
		return false
	}

	// Verificăm ecuația: g^hash ≡ y^r * r^s (mod p)
	// Partea stângă: g^hash mod p
	left := new(big.Int).Exp(keys.G, hash, keys.P)
	fmt.Printf("[ElGamal] g^hash mod p calculat\n")

	// Partea dreaptă: y^r * r^s mod p
	yr := new(big.Int).Exp(keys.Y, sig.R, keys.P)
	rs := new(big.Int).Exp(sig.R, sig.S, keys.P)
	right := new(big.Int).Mul(yr, rs)
	right.Mod(right, keys.P)
	fmt.Printf("[ElGamal] y^r * r^s mod p calculat\n")

	// Comparăm
	isValid := left.Cmp(right) == 0
	if isValid {
		fmt.Println("[ElGamal] ✓ Semnătura este VALIDĂ!")
		fmt.Println("[ElGamal] g^hash ≡ y^r * r^s (mod p)")
	} else {
		fmt.Println("[ElGamal] ✗ Semnătura este INVALIDĂ!")
		fmt.Println("[ElGamal] g^hash ≢ y^r * r^s (mod p)")
	}

	return isValid
}

// ====================================================================================
// PARTEA 4: FUNCȚIA MAIN
// ====================================================================================

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║         LUCRARE DE LABORATOR NR. 6 - SEMNĂTURI DIGITALE                   ║")
	fmt.Println("║         Funcții Hash și Semnături Digitale (RSA & ElGamal)                ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════╝")

	// ================================================================================
	// ETAPA 0: Afișarea hash-urilor pre-calculate
	// ================================================================================
	printSeparator("ETAPA 0: Hash-uri Pre-calculate")

	fmt.Println("\n[Info] Mesajul din Laboratorul 2:")
	fmt.Printf("[Info] Lungime mesaj: %d caractere\n", len(message))
	fmt.Println("\n[Info] Număr de ordine în catalog: 15")
	fmt.Println("[Info] Formula: i = (15 mod 24) + 1 = 16")

	// Hash pentru RSA (RIPEMD-128)
	fmt.Println("\n[Hash] === Pentru RSA (Sarcina 2): RIPEMD-128 (poziția 16) ===")
	hashDecimalRSA := hashToDecimal(hashRIPEMD128Hex)
	fmt.Printf("[Hash] Hash RIPEMD-128 (hexazecimal): %s\n", hashRIPEMD128Hex)
	fmt.Printf("[Hash] Hash RIPEMD-128 (decimal): %s\n", hashDecimalRSA.String())
	fmt.Printf("[Hash] Hash RIPEMD-128 (lungime în biți): %d\n", hashDecimalRSA.BitLen())

	// Hash pentru ElGamal (SHA3-512)
	fmt.Println("\n[Hash] === Pentru ElGamal (Sarcina 3): SHA3-512 (poziția 16) ===")
	hashDecimalElGamal := hashToDecimal(hashSHA3_512Hex)
	fmt.Printf("[Hash] Hash SHA3-512 (hexazecimal): %s\n", hashSHA3_512Hex)
	fmt.Printf("[Hash] Hash SHA3-512 (decimal): %s\n", hashDecimalElGamal.String())
	fmt.Printf("[Hash] Hash SHA3-512 (lungime în biți): %d\n", hashDecimalElGamal.BitLen())

	// ================================================================================
	// SARCINA 2: RSA DIGITAL SIGNATURE
	// ================================================================================
	printSeparator("SARCINA 2: RSA DIGITAL SIGNATURE (RIPEMD-128)")

	// Generare chei RSA
	rsaKeys, err := generateRSAKeys(3072)
	if err != nil {
		fmt.Printf("Eroare la generarea cheilor RSA: %v\n", err)
		return
	}

	fmt.Println("\n[RSA] Cheile RSA generate:")
	fmt.Printf("[RSA] Cheia publică (n, e):\n")
	fmt.Printf("      n (modulul): %d biți\n", rsaKeys.N.BitLen())
	fmt.Printf("      e (exponent): %d\n", rsaKeys.E.Int64())
	fmt.Printf("\n[RSA] Cheia privată (d):\n")
	fmt.Printf("      d: %d biți\n", rsaKeys.D.BitLen())

	// Semnare mesaj
	rsaSignature := rsaSign(hashDecimalRSA, rsaKeys)

	// Verificare semnătură
	rsaValid := rsaVerify(hashDecimalRSA, rsaSignature, rsaKeys)

	fmt.Println("\n[RSA] Rezultat final:")
	if rsaValid {
		fmt.Println("[RSA] ✓✓✓ Semnătura RSA a fost generată și verificată cu succes!")
	} else {
		fmt.Println("[RSA] ✗✗✗ Eroare: Semnătura RSA nu este validă!")
	}

	// ================================================================================
	// SARCINA 3: ELGAMAL DIGITAL SIGNATURE
	// ================================================================================
	printSeparator("SARCINA 3: ELGAMAL DIGITAL SIGNATURE (SHA3-512)")

	// Generare chei ElGamal
	elgamalKeys, err := generateElGamalKeys()
	if err != nil {
		fmt.Printf("Eroare la generarea cheilor ElGamal: %v\n", err)
		return
	}

	fmt.Println("\n[ElGamal] Cheile ElGamal generate:")
	fmt.Printf("[ElGamal] Cheia publică (p, g, y):\n")
	fmt.Printf("          p: %d biți (dat în cerință)\n", elgamalKeys.P.BitLen())
	fmt.Printf("          g: %d (dat în cerință)\n", elgamalKeys.G.Int64())
	fmt.Printf("          y: %d biți\n", elgamalKeys.Y.BitLen())
	fmt.Printf("\n[ElGamal] Cheia privată (x):\n")
	fmt.Printf("          x: %d biți\n", elgamalKeys.X.BitLen())

	// Semnare mesaj
	elgamalSig, err := elgamalSign(hashDecimalElGamal, elgamalKeys)
	if err != nil {
		fmt.Printf("Eroare la semnarea ElGamal: %v\n", err)
		return
	}

	// Verificare semnătură
	elgamalValid := elgamalVerify(hashDecimalElGamal, elgamalSig, elgamalKeys)

	fmt.Println("\n[ElGamal] Rezultat final:")
	if elgamalValid {
		fmt.Println("[ElGamal] ✓✓✓ Semnătura ElGamal a fost generată și verificată cu succes!")
	} else {
		fmt.Println("[ElGamal] ✗✗✗ Eroare: Semnătura ElGamal nu este validă!")
	}

	// ================================================================================
	// REZUMAT FINAL
	// ================================================================================
	printSeparator("REZUMAT FINAL - LABORATOR 6")

	fmt.Println("\n✓ Număr de ordine în catalog: 15")
	fmt.Println("  Formula: i = (15 mod 24) + 1 = 16")

	fmt.Println("\n✓ RSA - Hash Function: RIPEMD-128 (poziția 16)")
	fmt.Printf("  Hash (hex): %s\n", hashRIPEMD128Hex)
	fmt.Printf("  Hash (dec): %s\n", hashDecimalRSA.String())

	fmt.Println("\n✓ ElGamal - Hash Function: SHA3-512 (poziția 16)")
	fmt.Printf("  Hash (hex): %s\n", hashSHA3_512Hex)
	fmt.Printf("  Hash (dec): %s\n", hashDecimalElGamal.String())

	fmt.Println("\n✓ RSA Digital Signature:")
	fmt.Printf("  n: %d biți (cerință: >= 3072 biți) ✓\n", rsaKeys.N.BitLen())
	fmt.Printf("  Semnătură validă: %v\n", rsaValid)

	fmt.Println("\n✓ ElGamal Digital Signature:")
	fmt.Printf("  p: %d biți (din cerință: 2048 biți) ✓\n", elgamalKeys.P.BitLen())
	fmt.Printf("  Semnătură validă: %v\n", elgamalValid)

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("Laborator finalizat cu succes!")
	fmt.Println(strings.Repeat("=", 80))
}
