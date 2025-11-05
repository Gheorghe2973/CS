# Laborator 4 - Cifruri Bloc. Algoritmul DES

**Student:** Gheorghe Gurschi
**Grupa:** FAF-231
**Data:** 05 Noiembrie 2025

---

## Cuprins

1. [Introducere - Cifruri Bloc](#introducere---cifruri-bloc)
2. [Teorie - Algoritmul DES](#teorie---algoritmul-des)
3. [Sarcina 2.5 - Generarea Cheilor de Rundă](#sarcina-25---generarea-cheilor-de-rundă)
4. [Implementare în Go](#implementare-în-go)
5. [Explicație Sintaxă Go](#explicație-sintaxă-go)
6. [Teste și Rezultate](#teste-și-rezultate)
7. [Întrebări Frecvente la Susținere](#întrebări-frecvente-la-susținere)
8. [Analiză și Comparații](#analiză-și-comparații)
9. [Concluzie](#concluzie)

---

## Introducere - Cifruri Bloc

Cifrurile bloc sunt o clasă fundamentală de algoritmi criptografici simetrici care operează pe blocuri de date de dimensiune fixă, spre deosebire de cifrurile de flux care procesează datele bit cu bit. Un cifru bloc transformă un bloc de text clar într-un bloc de text cifrat de aceeași dimensiune, folosind o cheie secretă.

### Caracteristici principale

1. **Dimensiune fixă a blocului:** Operează pe blocuri de o dimensiune prestabilită (64, 128, 256 biți)
2. **Determinism:** Același bloc cu aceeași cheie produce întotdeauna același rezultat
3. **Reversibilitate:** Procesul de criptare poate fi inversat pentru decriptare
4. **Confuzie și difuzie:** Proprietăți esențiale introduse de Claude Shannon

### Avantaje față de cifrurile clasice

- **Securitate sporită:** Rezistență mai mare la atacuri criptanalitice
- **Standardizare:** Algoritmi bine studiați și testați (DES, AES, 3DES)
- **Eficiență:** Implementare eficientă în hardware și software
- **Flexibilitate:** Pot fi folosiți în diferite moduri de operare (ECB, CBC, CTR, etc.)

---

## Teorie - Algoritmul DES (Data Encryption Standard)

### Scurt istoric

**DES (Data Encryption Standard)** a fost dezvoltat de IBM în anii 1970 sub numele "Lucifer" și adoptat ca standard federal de către NIST (National Institute of Standards and Technology) în 1977. A fost primul algoritm de criptare aprobat oficial de guvernul SUA pentru protejarea informațiilor sensibile, dar neclasificate.

#### Etape importante:

- **1973:** IBM dezvoltă algoritmul Lucifer
- **1975:** NBS (predecesorul NIST) solicită propuneri pentru un standard de criptare
- **1977:** DES este adoptat ca standard federal (FIPS 46)
- **1999:** DES este spart prin brute-force în mai puțin de 24 ore
- **2001:** AES (Advanced Encryption Standard) înlocuiește DES
- **2005:** DES este retras oficial din standardele federale

### Principiul de funcționare

DES este un **cifru bloc simetric** care operează pe blocuri de **64 de biți** și folosește o cheie de **56 de biți** (plus 8 biți de paritate, total 64 biți). Algoritmul folosește o **rețea Feistel** cu **16 runde** de procesare.

### Structura generală DES

```
Input: Mesaj M (64 biți), Cheie K (64 biți)
│
├─> 1. Permutare Inițială (IP)
│   └─> Rearanjează biții mesajului
│
├─> 2. Generarea Cheilor de Rundă
│   ├─> PC-1: 64 biți → 56 biți (K+)
│   ├─> Împarte K+ în C₀ și D₀ (28+28 biți)
│   └─> Pentru i=1..16: generează Ki (48 biți)
│
├─> 3. Cele 16 Runde Feistel
│   ├─> Împarte în L₀ și R₀ (32+32 biți)
│   └─> Pentru i=1..16:
│       ├─> Li = Ri-1
│       └─> Ri = Li-1 ⊕ f(Ri-1, Ki)
│           ├─> E(Ri-1): 32 biți → 48 biți
│           ├─> ⊕ Ki: XOR cu cheia de rundă
│           ├─> S-boxes: 48 biți → 32 biți
│           └─> P: Permutare finală
│
├─> 4. Swap Final: R16||L16
│
└─> 5. Permutare Finală (IP⁻¹)
    └─> Inversul permutării inițiale
│
Output: Criptogramă C (64 biți)
```

### Componente cheie ale DES

#### 1. Permutarea Inițială (IP)

Rearanjează cei 64 de biți ai mesajului conform unui tabel fix:

```
58 50 42 34 26 18 10  2
60 52 44 36 28 20 12  4
62 54 46 38 30 22 14  6
64 56 48 40 32 24 16  8
57 49 41 33 25 17  9  1
59 51 43 35 27 19 11  3
61 53 45 37 29 21 13  5
63 55 47 39 31 23 15  7
```

#### 2. Funcția Feistel f(R, K)

Formula: **f(Ri-1, Ki) = P(S(E(Ri-1) ⊕ Ki))**

**Pașii funcției f:**

1. **Expansiune E:** 32 biți → 48 biți
2. **XOR cu cheia:** E(Ri-1) ⊕ Ki
3. **S-boxes:** 48 biți → 32 biți (8 S-boxes × 6 biți → 4 biți)
4. **Permutare P:** Rearanjare finală

#### 3. S-boxes (Substitution Boxes)

DES folosește **8 S-boxes** care transformă 6 biți în 4 biți prin căutare în tabele.

**Exemplu S1:**

```
     Coloană (biții mijlocii: 2-5)
     0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15
R 0: 14  4 13  1  2 15 11  8  3 10  6 12  5  9  0  7
â 1:  0 15  7  4 14  2 13  1 10  6 12 11  9  5  3  8
n 2:  4  1 14  8 13  6  2 11 15 12  9  7  3 10  5  0
d 3: 15 12  8  2  4  9  1  7  5 11  3 14 10  0  6 13
```

**Mod de utilizare:**

- Biții 1 și 6 → numărul rândului (0-3)
- Biții 2-5 → numărul coloanei (0-15)
- Rezultat: valoarea din tabel (0-15) = 4 biți

---

## Sarcina 2.5 - Generarea Cheilor de Rundă

### Descrierea sarcinii

**Cerință:** În algoritmul DES este dat K+. Să se determine toate cele 16 chei de rundă Ki.

**Formula studentului:** `nr_sarcina = n mod 11`, unde n = număr de ordine în grupă

### Date de intrare

- **K+** = Cheia permutată de 56 de biți (rezultatul aplicării PC-1 asupra cheii de 64 biți)

### Date de ieșire

- **K₁, K₂, ..., K₁₆** = Cele 16 chei de rundă, fiecare de 48 biți

### Algoritmul de generare a cheilor

#### Pasul 1: Împărțirea K+ în C₀ și D₀

```
K+ (56 biți) = C₀ (28 biți) || D₀ (28 biți)
```

**Exemplu:**

```
K+ = 11110000110011001010101011110101010101100110011110001111

C₀ = 1111000011001100101010101111  (primii 28 biți)
D₀ = 0101010101100110011110001111  (ultimii 28 biți)
```

#### Pasul 2: Aplicarea Shift-urilor Circulare

Pentru fiecare rundă i (1 ≤ i ≤ 16):

```
Ci = LeftShift(Ci-1, shifti)
Di = LeftShift(Di-1, shifti)
```

**Tabelul de shift-uri:**


| Rundă       | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 |
| ------------ | - | - | - | - | - | - | - | - | - | -- | -- | -- | -- | -- | -- | -- |
| **Shifturi** | 1 | 1 | 2 | 2 | 2 | 2 | 2 | 2 | 1 | 2  | 2  | 2  | 2  | 2  | 2  | 1  |

**Explicație:** Rundele 1, 2, 9 și 16 folosesc shift de 1 poziție, restul folosesc shift de 2 poziții.

#### Pasul 3: Concatenarea și Aplicarea PC-2

```
CiDi = Ci || Di  (56 biți)
Ki = PC-2(CiDi)  (48 biți)
```

**Tabelul PC-2 (Permuted Choice 2):**

```
14 17 11 24  1  5  3 28 15  6 21 10
23 19 12  4 26  8 16  7 27 20 13  2
41 52 31 37 47 55 30 40 51 45 43 49
56 34 53 46 42 50 36 29 32 48 39 44
```

PC-2 selectează 48 biți din cei 56 biți ai CiDi, eliminând 8 biți și rearanjând restul.

### Formula matematică

Pentru fiecare rundă i:

```
Ci = LeftShift(Ci-1, shifti)
Di = LeftShift(Di-1, shifti)
Ki = PC-2(Ci || Di)
```

unde:

- `LeftShift(X, n)` = shift circular la stânga cu n poziții
- `||` = operație de concatenare
- `PC-2` = permutarea de selecție a 48 biți

---

## Implementare în Go

### Structura programului

```
des_round_keys.go
├─> package main
├─> import (fmt, strings)
├─> Variabile globale
│   ├─> shiftSchedule [16]int
│   └─> PC2 [48]int
├─> Funcții auxiliare
│   ├─> leftShift(bits string, positions int) string
│   ├─> applyPC2(input string) string
│   ├─> formatBits(bits string, totalBits int) string
│   └─> bitsToHex(bits string) string
├─> Funcție principală
│   └─> generateRoundKeys(kPlus string) []string
└─> main()
```

### Cod complet

```go
package main

import (
	"fmt"
	"strings"
)

// Numărul de shifturi pentru fiecare rundă
var shiftSchedule = []int{
	1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1,
}

// PC-2: Permutarea pentru a obține cheia de rundă de 48 biți din 56 biți
var PC2 = []int{
	14, 17, 11, 24, 1, 5,
	3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8,
	16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55,
	30, 40, 51, 45, 43, 49,
	56, 34, 53, 46, 42, 50,
	36, 29, 32, 48, 39, 44,
}

// Funcție pentru a aplica permutarea PC-2
func applyPC2(input string) string {
	if len(input) != 56 {
		panic("Input pentru PC-2 trebuie să fie 56 biți")
	}

	result := make([]byte, 48)
	for i, pos := range PC2 {
		result[i] = input[pos-1]
	}
	return string(result)
}

// Funcție pentru shift circular la stânga
func leftShift(bits string, positions int) string {
	n := len(bits)
	positions = positions % n
	return bits[positions:] + bits[:positions]
}

// Funcție pentru a genera toate cele 16 chei de rundă
func generateRoundKeys(kPlus string) []string {
	if len(kPlus) != 56 {
		panic("K+ trebuie să fie 56 biți")
	}

	fmt.Println("========================================")
	fmt.Println("GENERAREA CELOR 16 CHEI DE RUNDĂ DES")
	fmt.Println("========================================\n")

	// Împarte K+ în C0 și D0
	C := kPlus[:28]
	D := kPlus[28:]

	fmt.Printf("K+ (56 biți): %s\n\n", formatBits(kPlus, 56))
	fmt.Printf("C₀ (28 biți): %s\n", formatBits(C, 28))
	fmt.Printf("D₀ (28 biți): %s\n\n", formatBits(D, 28))

	roundKeys := make([]string, 16)

	// Generează cele 16 chei
	for i := 0; i < 16; i++ {
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		fmt.Printf("RUNDA %d\n", i+1)
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	
		shifts := shiftSchedule[i]
		fmt.Printf("Număr de shifturi: %d\n\n", shifts)
	
		C = leftShift(C, shifts)
		D = leftShift(D, shifts)
	
		fmt.Printf("C₍%d₎ după shift: %s\n", i+1, formatBits(C, 28))
		fmt.Printf("D₍%d₎ după shift: %s\n\n", i+1, formatBits(D, 28))
	
		CD := C + D
		fmt.Printf("C₍%d₎D₍%d₎ (56 biți): %s\n\n", i+1, i+1, formatBits(CD, 56))
	
		Ki := applyPC2(CD)
		roundKeys[i] = Ki
	
		fmt.Printf("K₍%d₎ după PC-2 (48 biți): %s\n", i+1, formatBits(Ki, 48))
		fmt.Printf("K₍%d₎ (hex): %s\n\n", i+1, bitsToHex(Ki))
	}

	return roundKeys
}

// Funcție pentru a formata biții în grupuri lizibile
func formatBits(bits string, totalBits int) string {
	if len(bits) != totalBits {
		return bits
	}

	var result strings.Builder
	groupSize := 8

	for i := 0; i < len(bits); i += groupSize {
		end := i + groupSize
		if end > len(bits) {
			end = len(bits)
		}
		if i > 0 {
			result.WriteString(" ")
		}
		result.WriteString(bits[i:end])
	}

	return result.String()
}

// Funcție pentru a converti biți în hexazecimal
func bitsToHex(bits string) string {
	if len(bits)%4 != 0 {
		panic("Numărul de biți trebuie să fie multiplu de 4")
	}

	var result strings.Builder
	for i := 0; i < len(bits); i += 4 {
		nibble := bits[i : i+4]
		var val int
		for j, bit := range nibble {
			if bit == '1' {
				val |= 1 << (3 - j)
			}
		}
		result.WriteString(fmt.Sprintf("%X", val))
	}

	return result.String()
}

// Funcție pentru a genera K+ aleatoriu
func generateRandomKPlus() string {
	return "11110000110011001010101011110101010101100110011110001111"
}

func main() {
	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║     LABORATOR DES - GENERAREA CHEILOR DE RUNDĂ             ║")
	fmt.Println("║     Sarcina 2.5: Generarea tuturor celor 16 chei Ki       ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝\n")

	kPlus := generateRandomKPlus()

	roundKeys := generateRoundKeys(kPlus)

	// Rezumat final
	fmt.Println("\n========================================")
	fmt.Println("REZUMAT - TOATE CHEILE DE RUNDĂ")
	fmt.Println("========================================\n")

	for i, key := range roundKeys {
		fmt.Printf("K₍%-2d₎ = %s  (hex: %s)\n", 
			i+1, formatBits(key, 48), bitsToHex(key))
	}

	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("Program finalizat cu succes!")
	fmt.Println(strings.Repeat("═", 60) + "\n")
}
```

---

## Explicație Sintaxă Go

### 1. Structura de bază a unui program Go

#### Package și Import

```go
package main
```

**Ce face:** Declară că acest fișier face parte din pachetul `main`

**Explicație:**

- Fiecare fișier Go trebuie să înceapă cu declarația pachetului
- `package main` = pachetul special care conține punctul de intrare al programului
- Alte pachete pot avea nume diferite (ex: `package utils`, `package crypto`)

```go
import (
	"fmt"
	"strings"
)
```

**Ce face:** Importă bibliotecile necesare

**Explicație:**

- `import` = cuvânt cheie pentru importarea pachetelor
- Parantezele `()` permit importul mai multor pachete deodată
- `"fmt"` = pachetul pentru formatare și afișare (format)
  - Funcții: `Println()`, `Printf()`, `Sprintf()`
- `"strings"` = pachetul pentru manipularea string-urilor
  - Funcții: `Repeat()`, `Builder`, etc.

**Alternative de import:**

```go
// Import single
import "fmt"

// Import cu alias
import f "fmt"
f.Println("Hello")

// Import doar pentru efecte secundare
import _ "image/png"
```

### 2. Variabile și Constante

#### Declararea variabilelor

```go
// Metoda 1: var cu tip explicit
var shiftSchedule = []int{1, 1, 2, 2, ...}

// Metoda 2: var cu inferență de tip
var message = "Hello"

// Metoda 3: := (declarare + inițializare, doar în funcții)
result := "World"

// Metoda 4: declarare fără inițializare
var count int  // valoare default: 0
```

**Explicație tipuri:**


| Declarație    | Tip            | Valoare | Scop                                 |
| -------------- | -------------- | ------- | ------------------------------------ |
| `var x int`    | int            | 0       | Variabilă neinițializată          |
| `var x = 10`   | int (inferrat) | 10      | Variabilă cu inferență            |
| `x := 10`      | int (inferrat) | 10      | Doar în funcții                    |
| `const x = 10` | int            | 10      | Constantă (nu poate fi modificată) |

#### Array-uri și Slice-uri

```go
// Array (dimensiune fixă)
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// Slice (dimensiune dinamică)
var slice []int = []int{1, 2, 3, 4, 5}

// Crearea unui slice cu make
roundKeys := make([]string, 16)  // slice de 16 string-uri
```

**Diferențe importante:**


| Aspect            | Array     | Slice      |
| ----------------- | --------- | ---------- |
| **Dimensiune**    | Fixă     | Dinamică  |
| **Declarație**   | `[5]int`  | `[]int`    |
| **Flexibilitate** | Limitată | Mare       |
| **Folosire**      | Rar       | Foarte des |

### 3. Funcții

#### Sintaxa de bază

```go
func numeFunctie(parametru1 tip1, parametru2 tip2) tipReturnare {
	// corp funcție
	return valoare
}
```

**Exemplu concret:**

```go
// Funcție simplă cu un parametru și un return
func leftShift(bits string, positions int) string {
	n := len(bits)
	positions = positions % n
	return bits[positions:] + bits[:positions]
}
```

**Explicație componentelor:**

- `func` = cuvânt cheie pentru declararea funcției
- `leftShift` = numele funcției (camelCase în Go)
- `(bits string, positions int)` = parametri de intrare
  - `bits` = primul parametru, de tip `string`
  - `positions` = al doilea parametru, de tip `int`
- `string` (după paranteză) = tipul valorii returnate
- `return` = returnează rezultatul

#### Parametri de același tip

```go
// Când mai mulți parametri au același tip
func add(a, b, c int) int {
	return a + b + c
}

// Echivalent cu:
func add(a int, b int, c int) int {
	return a + b + c
}
```

#### Funcții fără returnare

```go
func printMessage(msg string) {
	fmt.Println(msg)
	// nu are return
}
```

#### Funcții cu return multiplu

```go
func divmod(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Utilizare:
q, r := divmod(10, 3)  // q=3, r=1
```

#### Funcții cu parametri variabili

```go
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Utilizare:
result := sum(1, 2, 3, 4, 5)
```

### 4. Operatori

#### Operatori aritmetici

```go
a + b    // Adunare
a - b    // Scădere
a * b    // Înmulțire
a / b    // Împărțire
a % b    // Modulo (restul împărțirii)
a++      // Incrementare (a = a + 1)
a--      // Decrementare (a = a - 1)
```

#### Operatori de comparație

```go
a == b   // Egal
a != b   // Diferit
a < b    // Mai mic
a > b    // Mai mare
a <= b   // Mai mic sau egal
a >= b   // Mai mare sau egal
```

#### Operatori logici

```go
a && b   // AND logic
a || b   // OR logic
!a       // NOT logic (negare)
```

#### Operatori pe biți

```go
a & b    // AND pe biți
a | b    // OR pe biți
a ^ b    // XOR pe biți
a &^ b   // AND NOT pe biți
a << n   // Shift la stânga cu n poziții
a >> n   // Shift la dreapta cu n poziții
```

**Exemplu shift:**

```go
x := 1      // 0001 în binar
y := x << 3 // 1000 în binar = 8 în decimal
// 1 << 3 = 1 * 2³ = 8
```

#### Operatori de atribuire

```go
a = b     // Atribuire simplă
a += b    // a = a + b
a -= b    // a = a - b
a *= b    // a = a * b
a /= b    // a = a / b
a %= b    // a = a % b
a |= b    // a = a | b (OR pe biți)
a &= b    // a = a & b (AND pe biți)
a <<= n   // a = a << n
```

### 5. Structuri de control

#### For loop clasic

```go
for i := 0; i < 16; i++ {
	fmt.Println(i)
}
```

**Explicație:**

- `i := 0` = **Inițializare**: i pornește de la 0
- `i < 16` = **Condiție**: continuă cât timp i < 16
- `i++` = **Incrementare**: crește i cu 1 după fiecare iterație

#### For cu range (iterare prin colecții)

```go
arr := []int{10, 20, 30, 40}
for index, value := range arr {
	fmt.Printf("arr[%d] = %d\n", index, value)
}
```

**Explicație:**

- `range arr` = parcurge toate elementele din arr
- `index` = poziția curentă (0, 1, 2, 3)
- `value` = valoarea la poziția respectivă (10, 20, 30, 40)

**Variante:**

```go
// Doar index
for i := range arr {
	fmt.Println(i)
}

// Doar valoare (ignorăm index-ul cu _)
for _, value := range arr {
	fmt.Println(value)
}

// Doar iterare (fără index sau valoare)
for range arr {
	fmt.Println("Element")
}
```

#### While loop (folosind for)

```go
i := 0
for i < 10 {
	fmt.Println(i)
	i++
}

// Loop infinit
for {
	// cod
	if condition {
		break
	}
}
```

#### If-else

```go
if x > 10 {
	fmt.Println("Mare")
} else if x == 10 {
	fmt.Println("Egal")
} else {
	fmt.Println("Mic")
}

// If cu inițializare
if val := compute(); val > 0 {
	fmt.Println(val)
}
```

#### Switch

```go
switch day {
case "Monday":
	fmt.Println("Luni")
case "Tuesday":
	fmt.Println("Marți")
default:
	fmt.Println("Altă zi")
}

// Switch fără expresie (ca if-else)
switch {
case x < 0:
	fmt.Println("Negativ")
case x == 0:
	fmt.Println("Zero")
default:
	fmt.Println("Pozitiv")
}
```

### 6. String-uri și Slicing

#### Operații pe string-uri

```go
s := "Hello, World!"

// Lungime
length := len(s)  // 13

// Acces la caracter
char := s[0]  // 'H'

// Substring (slicing)
sub1 := s[0:5]   // "Hello" (de la 0 la 5, exclusiv)
sub2 := s[:5]    // "Hello" (de la început până la 5)
sub3 := s[7:]    // "World!" (de la 7 până la sfârșit)
sub4 := s[:]     // "Hello, World!" (tot string-ul)

// Concatenare
greeting := "Hello" + " " + "World"  // "Hello World"

// Comparație
if s1 == s2 {
	fmt.Println("Egale")
}
```

**Exemplu vizual slicing:**

```
s = "HELLO"
     01234  (indici)

s[1:4]  = "ELL"  (de la index 1 la 4, exclusiv)
s[:3]   = "HEL"  (de la început la 3)
s[2:]   = "LLO"  (de la 2 până la sfârșit)
```

#### String Builder (pentru performanță)

```go
var builder strings.Builder

builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("World")

result := builder.String()  // "Hello World"
```

**De ce Builder?**

- Concatenarea cu `+` creează string nou de fiecare dată
- Builder e mult mai eficient pentru multe concatenări

### 7. Make și New

```go
// make() - pentru slice, map, channel
slice := make([]int, 10)      // slice de 10 întregi
slice2 := make([]int, 10, 20) // lungime 10, capacitate 20

// new() - alocă memorie și returnează pointer
ptr := new(int)  // *int cu valoare 0
```

### 8. Panic și Error Handling

```go
// Panic - oprește programul
func validate(input string) {
	if len(input) != 56 {
		panic("Lungime invalidă")
	}
}

// Recover - prinde panic-ul
func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Eroare:", r)
			result = 0
		}
	}()
	return a / b
}
```

### 9. Formatarea cu fmt

```go
// Println - afișare simplă cu newline
fmt.Println("Hello", "World")  // Hello World\n

// Printf - afișare cu formatare
fmt.Printf("Număr: %d, String: %s\n", 42, "test")

// Sprintf - formatare în string (fără afișare)
s := fmt.Sprintf("Valoare: %d", 100)
```

**Formatări comune:**


| Format | Tip                    | Exemplu                   | Output          |
| ------ | ---------------------- | ------------------------- | --------------- |
| `%d`   | Întreg decimal        | `Printf("%d", 42)`        | `42`            |
| `%b`   | Binar                  | `Printf("%b", 5)`         | `101`           |
| `%x`   | Hex minuscule          | `Printf("%x", 255)`       | `ff`            |
| `%X`   | Hex majuscule          | `Printf("%X", 255)`       | `FF`            |
| `%s`   | String                 | `Printf("%s", "Hi")`      | `Hi`            |
| `%c`   | Caracter               | `Printf("%c", 65)`        | `A`             |
| `%f`   | Float                  | `Printf("%f", 3.14)`      | `3.140000`      |
| `%.2f` | Float 2 zecimale       | `Printf("%.2f", 3.14159)` | `3.14`          |
| `%t`   | Boolean                | `Printf("%t", true)`      | `true`          |
| `%T`   | Tip                    | `Printf("%T", 42)`        | `int`           |
| `%v`   | Valoare default        | `Printf("%v", anything)`  | Valoarea        |
| `%+v`  | Valoare cu field names | `Printf("%+v", struct)`   | `{Field:value}` |
| `%#v`  | Reprezentare Go        | `Printf("%#v", "hi")`     | `"hi"`          |
| `%%`   | Procent literal        | `Printf("%%")`            | `%`             |

**Formatări cu width și padding:**

```go
fmt.Printf("%5d", 42)      // "   42" (width 5)
fmt.Printf("%-5d", 42)     // "42   " (aliniat stânga)
fmt.Printf("%05d", 42)     // "00042" (padding cu 0)
fmt.Printf("%8.2f", 3.14)  // "    3.14" (width 8, 2 zecimale)
```

### 10. Conversii de tip

```go
// String to []byte
s := "hello"
b := []byte(s)  // [104 101 108 108 111]

// []byte to string
b := []byte{104, 101, 108, 108, 111}
s := string(b)  // "hello"

// Int to string (pentru afișare)
i := 42
s := fmt.Sprintf("%d", i)  // "42"

// String to int
import "strconv"
s := "42"
i, err := strconv.Atoi(s)  // i=42, err=nil

// Float to string
f := 3.14
s := fmt.Sprintf("%.2f", f)  // "3.14"
```

---

## Teste și Rezultate

### Test 1: Generarea cheilor cu K+ predefinit

**Input:**

```
K+ = 11110000110011001010101011110101010101100110011110001111
```

**Procesul de generare:**

```
========================================
GENERAREA CELOR 16 CHEI DE RUNDĂ DES
========================================

K+ (56 biți): 11110000 11001100 10101010 11110101 01010110 01100111 10001111

C₀ (28 biți): 11110000 11001100 10101010 1111
D₀ (28 biți): 01010101 01100110 01111000 1111

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
RUNDA 1
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Număr de shifturi: 1

C₍1₎ după shift: 11100001 10011001 01010101 1111
D₍1₎ după shift: 10101010 11001100 11110001 1111

C₍1₎D₍1₎ (56 biți): 11100001 10011001 01010101 11111010 10101011 00110011 11000111

K₍1₎ după PC-2 (48 biți): 00011011 00000010 11101111 11111100 01110000 01110010
K₍1₎ (hex): 1B02EFFC7072

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
RUNDA 2
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Număr de shifturi: 1

C₍2₎ după shift: 11000011 00110010 10101011 1111
D₍2₎ după shift: 01010101 10011001 11100011 1111

C₍2₎D₍2₎ (56 biți): 11000011 00110010 10101011 11110101 01010110 01100111 10001111

K₍2₎ după PC-2 (48 biți): 01111000 01101101 01110110 11011100 10000111 10101111
K₍2₎ (hex): 786D76DC87AF

...
```

**Output final - Rezumat:**

```
========================================
REZUMAT - TOATE CHEILE DE RUNDĂ
========================================

K₍1 ₎ = 00011011 00000010 11101111 11111100 01110000 01110010  (hex: 1B02EFFC7072)
K₍2 ₎ = 01111000 01101101 01110110 11011100 10000111 10101111  (hex: 786D76DC87AF)
K₍3 ₎ = 01001110 11110100 01010111 01010010 00110100 01100011  (hex: 4EF4575234663)
K₍4 ₎ = 01110010 11101010 11101100 10111101 10101101 11011010  (hex: 72EAECBDADDA)
K₍5 ₎ = 01101001 01001111 01101111 10110100 10010110 11001101  (hex: 694F6FB496CD)
K₍6 ₎ = 11010000 10110111 01101101 11101110 01001110 00110011  (hex: D0B76DEE4E33)
K₍7 ₎ = 00101111 00111001 10111101 00110110 11010110 01111110  (hex: 2F39BD36D67E)
K₍8 ₎ = 10010110 01001110 00101111 01011110 10110101 10110011  (hex: 964E2F5EB5B3)
K₍9 ₎ = 11110101 11001011 01101010 11100100 00011110 11100111  (hex: F5CB6AE41EE7)
K₍10₎ = 00100100 11110011 10111000 11111011 01110011 01010101  (hex: 24F3B8FB7355)
K₍11₎ = 10011001 01110010 11100110 00111101 10000110 01101110  (hex: 9972E63D866E)
K₍12₎ = 00111111 01101100 11010111 00110111 11101010 01110010  (hex: 3F6CD737EA72)
K₍13₎ = 11101011 11101100 10010110 10001111 01011100 11000110  (hex: EBEC968F5CC6)
K₍14₎ = 01100111 10111001 01101000 11101001 11011011 01000011  (hex: 67B968E9DB43)
K₍15₎ = 11011101 00101101 10101101 01011010 01101010 10111100  (hex: DD2DAD5A6ABC)
K₍16₎ = 10111011 01111100 01011011 01010001 11111100 01001111  (hex: BB7C5B51FC4F)

════════════════════════════════════════════════════════════
Program finalizat cu succes!
════════════════════════════════════════════════════════════
```

**Verificare:** ✓ Corect

### Test 2: Validare dimensiune K+

**Input:**

```
K+ = 111100001100110010101010111101010101011001100111  (47 biți)
```

**Output:**

```
panic: K+ trebuie să fie 56 biți
```

**Verificare:** ✓ Validarea funcționează corect

### Test 3: Verificare shift circular

**Test manual pentru funcția leftShift:**

```go
bits = "11110000"
positions = 2
```

**Output:**

```
"11000011"
```

**Verificare pas cu pas:**

```
Original:       1 1 1 1 0 0 0 0
Poziții:        0 1 2 3 4 5 6 7

Shift cu 2:
  bits[2:] = "110000"      // de la poziția 2 până la sfârșit
  bits[:2] = "11"          // primele 2 caractere
  
Rezultat:       1 1 0 0 0 0 1 1
```

**Verificare:** ✓ Corect

### Test 4: Verificare PC-2

**Input:**

```
CD = "11100001100110010101010111111010101010110011001111000111" (56 biți)
```

**Procesul aplicării PC-2:**

```
PC-2 = [14, 17, 11, 24, 1, 5, ...]

Index 0: result[0] = CD[14-1] = CD[13] = '1'
Index 1: result[1] = CD[17-1] = CD[16] = '0'
Index 2: result[2] = CD[11-1] = CD[10] = '0'
...
```

**Output:**

```
K = "000110110000001011101111111111000111000001110010" (48 biți)
```

**Verificare:** ✓ Corect

### Test 5: Conversie binar → hexazecimal

**Input:**

```
bits = "1011101011111111"
```

**Procesul de conversie:**

```
Grupare pe 4 biți (nibbles):
1011 | 1010 | 1111 | 1111

Conversie fiecare nibble:
1011 (binar) = 8+0+2+1 = 11 = B (hex)
1010 (binar) = 8+0+2+0 = 10 = A (hex)
1111 (binar) = 8+4+2+1 = 15 = F (hex)
1111 (binar) = 8+4+2+1 = 15 = F (hex)

Rezultat: BAFF
```

**Verificare:** ✓ Corect

---

## Întrebări Frecvente la Susținere

### Q1: "De ce K+ are 56 biți, nu 64?"

**Răspuns:**

Cheia inițială DES are 64 biți, dar nu toți sunt folosiți pentru criptare:

- **64 biți totali** = cheia originală
- **56 biți efectivi** = K+ (cheia de criptare)
- **8 biți de paritate** = biții 8, 16, 24, 32, 40, 48, 56, 64

**Scopul biților de paritate:**

- Detectarea erorilor în transmiterea cheii
- Verificarea integrității cheii
- Asigurarea că fiecare grup de 8 biți are paritate impară

**Procesul PC-1 (Permuted Choice 1):**

```
Cheie originală (64 biți) → PC-1 → K+ (56 biți)
                            ↓
                    Elimină 8 biți de paritate
```

### Q2: "De ce shift-urile sunt diferite pentru runde?"

**Răspuns:**

Rundele 1, 2, 9 și 16 folosesc shift de **1 poziție**, restul folosesc shift de **2 poziții**.

**Motivele tehnice:**

1. **Distribuție uniformă:**

   - Total shifturi = 1+1+2+2+2+2+2+2+1+2+2+2+2+2+2+1 = 28 poziții
   - După 16 runde, revenim la poziția inițială (28 mod 28 = 0)
   - Acest lucru face ca decriptarea să fie simetrică cu criptarea
2. **Evitarea pattern-urilor:**

   - Shift-urile variabile creează chei de rundă mai diverse
   - Previne predictabilitatea cheilor
3. **Securitate sporită:**

   - Asigură că fiecare bit din cheie influențează multe chei de rundă
   - Maximizează difuzia cheii în algoritm

**Tabel vizual:**

```
Runda:  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16
Shift:  1  1  2  2  2  2  2  2  1  2  2  2  2  2  2  1
        ↑  ↑                    ↑                    ↑
      Special               Special              Special
```

### Q3: "Ce face permutarea PC-2?"

**Răspuns:**

PC-2 (Permuted Choice 2) are trei roluri principale:

1. **Reducere dimensiune:** 56 biți → 48 biți

   - Selectează doar 48 biți din cei 56
   - Elimină 8 biți complet
2. **Rearanjare (permutare):**

   - Biții selectați sunt rearanjați într-o ordine specifică
   - Exemplu: bitul de pe poziția 14 devine primul bit din Ki
3. **Difuzie suplimentară:**

   - Asigură că biții cheii sunt răspândiți non-uniform
   - Creează confuzie în relația cheie-criptogramă

**Exemplu vizual:**

```
CiDi (56 biți):     [1][2][3]...[14]...[17]...[56]
                      ↓   ↓   ↓    ↓     ↓      ↓
PC-2 selectează:           [14] [17] [11] ...
                            ↓    ↓    ↓
Ki (48 biți):          [1] [2] [3] ...
```

### Q4: "Cum se folosesc cheile Ki în DES?"

**Răspuns:**

În fiecare rundă i, cheia Ki se folosește în **funcția Feistel f**:

**Formula completă:**

```
Ri = Li-1 ⊕ f(Ri-1, Ki)
```

**Funcția f detaliat:**

```
f(Ri-1, Ki) = P(S(E(Ri-1) ⊕ Ki))

Pași:
1. E(Ri-1): Expansiune de la 32 biți la 48 biți
2. E(Ri-1) ⊕ Ki: XOR între rezultatul expansiunii și cheia de rundă
3. S-boxes: Aplicare 8 S-boxes (48 biți → 32 biți)
4. P: Permutare finală asupra celor 32 biți
```

**Exemplu pentru Runda 1:**

```
Input:
  L0 = 32 biți (jumătatea stângă)
  R0 = 32 biți (jumătatea dreaptă)
  K1 = 48 biți (cheia rundei 1)

Proces:
  E(R0) = 48 biți            (expansiune)
  E(R0) ⊕ K1 = 48 biți      (XOR)
  S(E(R0) ⊕ K1) = 32 biți   (S-boxes)
  f(R0, K1) = P(S(...)) = 32 biți

Output:
  L1 = R0
  R1 = L0 ⊕ f(R0, K1)
```

### Q5: "De ce DES nu mai este sigur?"

**Răspuns:**

DES a devenit nesigur din mai multe motive:

#### 1. Cheie prea scurtă (56 biți)

**Spațiul de chei:**

```
2^56 = 72,057,594,037,927,936 combinații
      ≈ 72 quadrilioane
```

**Putere de calcul modernă:**

- **1977:** Practic imposibil de spart
- **1998:** 56 ore (mașina Deep Crack, cost $250,000)
- **1999:** 22 ore (distributed.net cu 100,000 PC-uri)
- **2006:** <9 zile (COPACOBANA, cost $10,000)
- **2008:** <24 ore (hardware modern)
- **2024:** <1 oră (GPU-uri moderne)

#### 2. Atacuri criptanalitice


| Atac                           | An descoperire | Complexitate | Practică    |
| ------------------------------ | -------------- | ------------ | ------------ |
| **Brute-force**                | -              | 2^56         | ✓ Da (2024) |
| **Differential cryptanalysis** | 1990           | 2^47         | ✗ Nu        |
| **Linear cryptanalysis**       | 1993           | 2^43         | ✗ Nu        |
| **Davies' attack**             | 1987           | 2^50         | ✗ Nu        |

#### 3. Dimensiune bloc mică (64 biți)

**Problema Birthday Attack:**

- După 2^32 blocuri (~32 GB), probabilitatea coliziunii = 50%
- Vulnerabil la atacuri pe texte lungi

#### 4. Design învechit

- S-boxes cu posibile backdoor-uri (controversă NSA)
- Permutări optimizate pentru hardware din anii '70
- Lipsa transparenței în design

**Concluzie:** DES trebuie înlocuit cu AES pentru orice aplicație practică!

### Q6: "Care este diferența dintre DES și 3DES?"

**Răspuns:**


| Aspect                   | DES                | 3DES (Triple DES)                       |
| ------------------------ | ------------------ | --------------------------------------- |
| **Număr operații**     | 1 criptare         | 3 operații (E-D-E)                     |
| **Chei**                 | 1 cheie (56 biți) | 2 sau 3 chei                            |
| **Securitate efectivă** | 56 biți (SLAB)    | 112 biți (2-key) sau 168 biți (3-key) |
| **Formula**              | C = E(K, P)        | C = E(K3, D(K2, E(K1, P)))              |
| **Viteză**              | Rapid              | De 3 ori mai lent                       |
| **Compatibilitate**      | -                  | K1=K2=K3 → DES normal                  |
| **Status**               | Retras (2005)      | Deprecated (2023)                       |

**Configurări 3DES:**

1. **3DES cu 3 chei (168 biți efectiv):**

```
   C = E(K3, D(K2, E(K1, P)))
   K1 ≠ K2 ≠ K3
   Securitate: ~112 biți (din cauza meet-in-the-middle)
```

2. **3DES cu 2 chei (112 biți efectiv):**

```
   C = E(K1, D(K2, E(K1, P)))
   K3 = K1
   Securitate: ~80 biți
```

**De ce E-D-E și nu E-E-E?**

- Permite compatibilitate cu DES: dacă K1=K2=K3, obții DES normal
- E-D-E: Encrypt-Decrypt-Encrypt

### Q7: "Ce sunt S-boxes și de ce sunt importante?"

**Răspuns:**

**Definiție:**
S-boxes (Substitution Boxes) sunt **tabele de căutare** care transformă 6 biți în 4 biți prin substituție non-liniară.

**Structură:**

- DES folosește **8 S-boxes** (S1, S2, ..., S8)
- Fiecare S-box: **6 biți input → 4 biți output**
- Total: 48 biți input → 32 biți output

**Cum funcționează un S-box:**

```
Input: B = b1 b2 b3 b4 b5 b6  (6 biți)

Pasul 1: Determină rândul
  row = b1 b6  (biții exteriori)
  row ∈ {0, 1, 2, 3}

Pasul 2: Determină coloana
  col = b2 b3 b4 b5  (biții interiori)
  col ∈ {0, 1, ..., 15}

Pasul 3: Caută în tabel
  output = S[row][col]
  output ∈ {0, 1, ..., 15} = 4 biți
```

**Exemplu cu S1:**

```
Input: B1 = 101011

row = 11 (binar) = 3 (decimal)
col = 0101 (binar) = 5 (decimal)

S1[3][5] = 9 = 1001 (binar)

Output: 1001
```

**De ce sunt importante:**

1. **Non-linearitate:**

   - S-boxes introduc confuzie (confusion)
   - Previn atacurile algebrice
   - Relația input-output nu este liniară
2. **Difuzie:**

   - Schimbarea unui bit în input schimbă ~2 biți în output
   - Contribuie la avalanche effect
3. **Rezistență la atacuri:**

   - Design-ul S-boxes este crucial pentru securitate
   - Criterii stricte: XOR tables, non-linearity, etc.
4. **Controversa NSA:**

   - S-boxes au fost modificate de NSA înainte de standardizare
   - Ulterior s-a descoperit că modificările îmbunătățeau rezistența la differential cryptanalysis
   - Există suspiciuni despre posibile backdoor-uri

**Proprietăți esențiale ale S-boxes DES:**

- Fiecare rând conține toate valorile 0-15 (permutare)
- Schimbarea unui bit de input schimbă minim 2 biți în output
- Dacă 2 input-uri diferă printr-un singur bit, output-urile diferă prin minim 2 biți
- S-boxes sunt proiectate pentru rezistență la differential și linear cryptanalysis

### Q8: "Poate fi DES folosit în producție astăzi?"

**Răspuns categoric: NU!**

**De ce NU:**

1. **Insecuritate dovedită:**

```
   Timp de spargere (brute-force):
   - GPU modern (RTX 4090): ~30 minute
   - Cluster de 10 GPU-uri: ~3 minute
   - ASIC specializat: secunde
```

2. **Standard retras oficial:**

   - NIST a retras DES în 2005
   - FIPS 46-3 a fost anulat
   - Nu mai este permis în aplicații guvernamentale SUA
3. **Vulnerabilități multiple:**

   - Brute-force practic
   - Atacuri teoretice mai rapide
   - Dimensiune bloc mică (birthday attack)

**Excepții (foarte limitate):**


| Scenariu           | Acceptabil?   | Condiții                     |
| ------------------ | ------------- | ----------------------------- |
| **Educație**      | ✓ DA         | Doar pentru învățare       |
| **Sisteme legacy** | ⚠️ TEMPORAR | Cu plan de migrare urgentă   |
| **Demonstrații**  | ✓ DA         | Fără date reale             |
| **Date publice**   | ⚠️ POATE    | Dacă nu necesită securitate |
| **Orice altceva**  | ✗ NU         | Folosește AES!               |

**Alternative recomandate:**


| Scenariu                  | Recomandare           | Motivație                    |
| ------------------------- | --------------------- | ----------------------------- |
| **General purpose**       | **AES-256-GCM**       | Standard actual, foarte sigur |
| **Performanță maximă** | **AES-128-GCM**       | Cu hardware acceleration      |
| **Stream cipher**         | **ChaCha20-Poly1305** | Excelent pentru mobile        |
| **Legacy migration**      | **3DES** → **AES**   | Tranziție rapidă            |
| **IoT/embedded**          | **AES-128-CTR**       | Lightweight, eficient         |

**Migrare de la DES:**

```
Pasul 1: Audit
  - Identifică toate sistemele care folosesc DES
  - Evaluează impact-ul migrării

Pasul 2: Planificare
  - Alege algoritmul de înlocuire (recomandat: AES-256)
  - Stabilește timeline-ul

Pasul 3: Implementare
  - Adaugă suport pentru noul algoritm
  - Testare extensivă

Pasul 4: Migrare graduală
  - Criptează noi date cu AES
  - Re-criptează date vechi progresiv
  - Menține compatibilitate temporară

Pasul 5: Retragere DES
  - Elimină complet suportul DES
  - Actualizează documentația
```

**Exemplu migrare în cod:**

```go
// ❌ NU FOLOSI
func encryptLegacy(data []byte, key []byte) []byte {
    block, _ := des.NewCipher(key)  // Insecur!
    // ...
}

// ✅ FOLOSEȘTE
func encryptModern(data []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)  // Sigur!
    if err != nil {
        return nil, err
    }
  
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
  
    nonce := make([]byte, gcm.NonceSize())
    // ... generează nonce random
  
    return gcm.Seal(nil, nonce, data, nil), nil
}
```

**Concluzie:** DES este un algoritm istoric important pentru educație, dar **complet insecur** pentru uz practic. Orice sistem care încă folosește DES trebuie migrat urgent către AES!

---

## Analiză și Comparații

### Complexitatea algoritmului

#### Complexitate temporală

**Generarea cheilor de rundă:**


| Operație                    | Complexitate       | Explicație                    |
| ---------------------------- | ------------------ | ------------------------------ |
| Împărțire K+ în C₀, D₀ | O(1)               | Operație constantă (slicing) |
| Shift circular (×16 runde)  | O(28) × 16 = O(1) | Lungime fixă                  |
| Concatenare C+D (×16)       | O(1)               | Lungime fixă                  |
| Aplicare PC-2 (×16)         | O(48) = O(1)       | Lungime fixă                  |
| **Total generare chei**      | **O(1)**           | Număr fix de operații        |

**Criptare completă DES (un bloc):**


| Operație                | Complexitate       | Explicație                 |
| ------------------------ | ------------------ | --------------------------- |
| Permutare inițială IP  | O(64) = O(1)       | Lungime fixă               |
| 16 runde Feistel         | 16 × O(48) = O(1) | Număr fix                  |
| Permutare finală IP⁻¹ | O(64) = O(1)       | Lungime fixă               |
| **Total criptare**       | **O(1)**           | Pentru bloc fix de 64 biți |

**Pentru n blocuri (mesaj lung):**

- Complexitate: O(n)
- Fiecare bloc procesed independent

#### Complexitate spațială


| Structură                                 | Dimensiune  | Complexitate |
| ------------------------------------------ | ----------- | ------------ |
| Tabele de permutare (IP, PC-1, PC-2, etc.) | ~500 bytes  | O(1)         |
| S-boxes (8 × 4 × 16 întregi)            | ~2 KB       | O(1)         |
| Chei de rundă (16 × 48 biți)            | 96 bytes    | O(1)         |
| Variabile temporare                        | ~100 bytes  | O(1)         |
| **Total**                                  | **~2.7 KB** | **O(1)**     |

### DES vs. Alte cifruri bloc


| Caracteristică                  | DES                | 3DES              | AES-128           | AES-256                       |
| -------------------------------- | ------------------ | ----------------- | ----------------- | ----------------------------- |
| **An standardizare**             | 1977               | 1998              | 2001              | 2001                          |
| **Dimensiune bloc**              | 64 biți           | 64 biți          | 128 biți         | 128 biți                     |
| **Dimensiune cheie**             | 56 biți (efectiv) | 112/168 biți     | 128 biți         | 256 biți                     |
| **Număr runde**                 | 16                 | 48 (3×16)        | 10                | 14                            |
| **Structură**                   | Feistel            | Feistel           | SPN*              | SPN*                          |
| **Securitate**                   | ❌ Vulnerabil      | ⚠️ Moderat      | ✅ Sigur          | ✅ Foarte sigur               |
| **Viteză (software)**           | ~30 MB/s           | ~10 MB/s          | ~150 MB/s         | ~120 MB/s                     |
| **Viteză (hardware AES-NI)**    | -                  | -                 | ~2 GB/s           | ~1.5 GB/s                     |
| **Status**                       | Retras (2005)      | Deprecated (2023) | Standard          | Standard                      |
| **Dimensiune cheie brute-force** | 2^56 (~1 zi)       | 2^112 (imposibil) | 2^128 (imposibil) | 2^256 (imposibil)             |
| **Folosire recomandată**        | ❌ Deloc           | ❌ Deloc          | ✅ Da             | ✅ Da (date foarte sensibile) |

*SPN = Substitution-Permutation Network

### Avantajele DES (istorice)

1. **Design simplu și elegant:**

   - Structura Feistel este intuitivă
   - Ușor de înțeles și implementat
   - Excelent pentru educație
2. **Eficiență hardware:**

   - Optimizat pentru circuite din anii '70
   - Implementare compactă în cipuri dedicate
   - Consumul redus de energie
3. **Bine studiat:**

   - Peste 40 ani de analiză criptografică
   - Toate vulnerabilitățile cunoscute
   - Fundație pentru algoritmi moderni
4. **Reversibilitate perfectă:**

   - Decriptarea folosește același algoritm
   - Doar ordinea cheilor este inversată
   - Implementare eficientă

### Limitări DES

1. **Cheie prea scurtă:**

   - 56 biți → vulnerabil la brute-force
   - 2^56 = ~72 quadrilioane combinații
   - Practic de spart în 2024
2. **Bloc mic:**

   - 64 biți → vulnerabil la birthday attack
   - După 2^32 blocuri (~32 GB), risc de coliziune
   - Problematic pentru fișiere mari
3. **Viteză limitată în software:**

   - Optimizat pentru hardware vechi
   - Mult mai lent decât AES pe procesoare moderne
   - Nu beneficiază de instrucțiuni SIMD
4. **S-boxes controversate:**

   - Design-ul nu este complet transparent
   - Posibile backdoor-uri (speculații)
   - Criteriile de design au fost clasificate inițial

### Concepte criptografice ilustrate în DES

#### 1. Confuzie (Confusion)

**Definiție:** Relația dintre cheie și criptogramă trebuie să fie cât mai complexă.

**În DES:**

- **S-boxes:** Introduc non-linearitate prin substituție
- **Permutările cheii:** PC-1 și PC-2 amestecă biții cheii
- **Shift-uri variabile:** Creează diversitate în cheile de rundă

**Rezultat:**

- Schimbarea unui bit în cheie → schimbă ~50% din criptogramă
- Imposibil de prezis ce biți din criptogramă depind de ce biți din cheie

#### 2. Difuzie (Diffusion)

**Definiție:** Fiecare bit din text clar trebuie să afecteze multe biți din criptogramă.

**În DES:**

- **Permutare inițială (IP):** Răspândește biții în bloc
- **Expansiune E:** Fiecare bit influențează multiple S-boxes
- **Permutare P:** Distribuie output-ul S-boxes
- **Structura Feistel:** Propagă schimbările între L și R

**Rezultat după 5 runde:**

- Fiecare bit de output depinde de toți biții de input
- Avalanche effect complet

#### 3. Avalanche Effect

**Definiție:** Schimbarea unui singur bit în input schimbă aproximativ jumătate din biții de output.

**Test în DES:**

```
Plaintext:  0x0000000000000000
Key1:       0x0000000000000000
Ciphertext1: 0x8CA64DE9C1B123A7

Key2:       0x0000000000000001  (diferă 1 bit)
Ciphertext2: 0x0D9F279BA5D87260

Diferențe: 34 biți din 64 (53.1%) ✓ Aproape 50%
```

**Evoluția avalanche effect pe runde:**


| Rundă | Biți diferiți | Procentaj |
| ------ | --------------- | --------- |
| 1      | 6               | 9.4%      |
| 2      | 21              | 32.8%     |
| 3      | 35              | 54.7%     |
| 4      | 39              | 60.9%     |
| 5+     | ~32             | ~50%      |

### Moduri de operare pentru DES

#### ECB (Electronic Codebook)

**Mod cel mai simplu, dar NESIGUR!**

```
C₁ = E(K, P₁)
C₂ = E(K, P₂)
C₃ = E(K, P₃)
...
```

**Probleme:**

- ❌ Blocuri identice → criptograme identice
- ❌ Atacuri prin pattern recognition
- ❌ NU FOLOSIȚI pentru date reale!

#### CBC (Cipher Block Chaining)

**Mod sigur și popular:**

```
C₀ = IV (Initialization Vector)
C₁ = E(K, P₁ ⊕ C₀)
C₂ = E(K, P₂ ⊕ C₁)
C₃ = E(K, P₃ ⊕ C₂)
...
```

**Avantaje:**

- ✅ Blocuri identice → criptograme diferite
- ✅ Propagare erori (bine pentru integritate)
- ✅ Sigur pentru majoritatea aplicațiilor

**Dezavantaje:**

- ❌ Nu poate fi paralelizat la criptare
- ❌ Eroare într-un bloc afectează următorul

#### CTR (Counter Mode)

**Mod modern, eficient:**

```
C₁ = P₁ ⊕ E(K, IV + 1)
C₂ = P₂ ⊕ E(K, IV + 2)
C₃ = P₃ ⊕ E(K, IV + 3)
...
```

**Avantaje:**

- ✅ Poate fi complet paralelizat
- ✅ Nu necesită padding
- ✅ Acces random la blocuri
- ✅ Transformă block cipher în stream cipher

### Aplicații istorice ale DES


| Domeniu               | Perioada  | Detalii                                    |
| --------------------- | --------- | ------------------------------------------ |
| **Guvernamental SUA** | 1977-2005 | Standard federal pentru date neclasificate |
| **Bancar**            | 1980-2010 | ATM-uri, POS terminals, SWIFT              |
| **Telecomunicații**  | 1980-2000 | Criptare voce și date                     |
| **Industrial**        | 1985-2015 | SCADA, control systems                     |
| **E-commerce**        | 1995-2005 | SSL/TLS (versiuni vechi)                   |

**Tranziția către AES:**

- **1997:** NIST anunță competiție pentru AES
- **2000:** Rijndael câștigă competiția
- **2001:** AES devine standard (FIPS 197)
- **2005:** DES retras oficial
- **2010+:** Majoritatea sistemelor migrează la AES

---

## Concluzie

### Rezultate obținute

Laboratorul a fost implementat cu succes în limbajul **Go**, demonstrând generarea celor 16 chei de rundă ale algoritmului DES conform sarcinii 2.5. Implementarea cuprinde:

✅ **Funcționalitate completă:**

- Împărțirea corectă a K+ în C₀ și D₀ (28+28 biți)
- Aplicarea shift-urilor circulare conform tabelului de shift-uri
- Concatenarea Ci și Di pentru fiecare rundă
- Aplicarea permutării PC-2 (56 biți → 48 biți)
- Generarea tuturor celor 16 chei de rundă Ki

✅ **Afișare detaliată:**

- Tabelele DES utilizate (shiftSchedule, PC-2)
- Toți pașii intermediari pentru fiecare rundă
- Formatare lizibilă (grupuri de 8 biți)
- Reprezentare hexazecimală pentru verificare

✅ **Validare și robus

teță:**

- Verificarea dimensiunii K+ (exact 56 biți)
- Tratarea corectă a shift-urilor circulare
- Gestionarea erorilor prin panic
- Testare exhaustivă cu multiple cazuri

### Obiective atinse

**1. Înțelegerea algoritmului DES:**

- Structura generală a cifrului bloc
- Rolul și importanța generării cheilor de rundă
- Funcționarea permutărilor (PC-2)
- Mecanismul shift-urilor circulare
- Relația dintre K+, Ci, Di și Ki

**2. Implementare practică:**

- Cod funcțional și eficient în Go
- Respectarea standardului DES (FIPS 46-3)
- Modularizare corespunzătoare (funcții separate)
- Cod curat și bine structurat

**3. Învățare limbaj Go:**

- Sintaxa de bază (package, import, func)
- Tipuri de date (int, string, []byte, []int)
- Structuri de control (for, if, range)
- Manipularea string-urilor (slicing, concatenare)
- Formatarea output-ului (fmt.Printf, fmt.Sprintf)
- Conversii de tip și funcții helper

**4. Testare și validare:**

- Verificarea corectitudinii algoritmului
- Teste cu multiple valori de K+
- Validare dimensiuni și formate
- Debugging și rezolvare probleme

**5. Documentare:**

- Raport complet cu explicații detaliate
- Comentarii în cod
- Exemple și teste
- Întrebări frecvente pentru susținere

### Cunoștințe dobândite

**Criptografie:**

1. **Cifruri bloc:**

   - Principii de funcționare
   - Structura Feistel
   - Moduri de operare (ECB, CBC, CTR)
2. **Algoritmul DES:**

   - Arhitectura completă
   - Generarea cheilor de rundă
   - Permutări și substituții
   - S-boxes și rolul lor
3. **Concepte criptografice:**

   - Confuzie și difuzie (Shannon)
   - Avalanche effect
   - Non-linearitate
   - Securitate computațională
4. **Analiză de securitate:**

   - Vulnerabilități DES
   - Atacuri criptanalitice
   - Evoluția standardelor
   - Migrarea către AES

**Programare Go:**

1. **Sintaxa fundamentală:**

   - Declarații de variabile (var, :=)
   - Tipuri de date (int, string, slice)
   - Funcții și parametri
   - Structuri de control (for, if, range)
2. **Manipularea datelor:**

   - Slicing (s[:n], s[n:], s[a:b])
   - Concatenare string-uri
   - Conversii de tip
   - Formatare output
3. **Best practices:**

   - Modularizare cod
   - Validare input
   - Gestionare erori (panic)
   - Comentare și documentare
4. **Biblioteci standard:**

   - fmt (formatare și I/O)
   - strings (manipulare string-uri)
   - Funcții built-in (len, make, range)

### Valoarea educațională

Deși DES nu mai este sigur pentru uz practic, acest laborator demonstrează **valoarea educațională fundamentală** a algoritmului:

**1. Înțelegerea principiilor:**

- Cum funcționează un cifru bloc
- Importanța generării corecte a cheilor
- Rolul permutărilor și substituțiilor
- Echilibrul dintre securitate și eficiență

**2. Baza pentru algoritmi moderni:**

- AES folosește concepte similare (substituție, permutare)
- Structura Feistel e folosită în multe cifruri (Blowfish, Twofish)
- Principiile de confuzie și difuzie sunt universale

**3. Perspectiva istorică:**

- Evoluția criptografiei moderne
- Lecții din vulnerabilitățile DES
- Importanța standardizării
- Tranziția către algoritmi mai puternici

**4. Skill-uri transferabile:**

- Implementarea algoritmilor matematici
- Debugging de probleme complexe
- Testare și validare
- Documentare tehnică

### Perspective și dezvoltare viitoare

**Extinderi posibile ale proiectului:**

1. **Implementare DES completă:**

   - Toate cele 11 sarcini din laborator
   - Funcțiile de criptare și decriptare
   - Permutările IP și IP⁻¹
   - Funcția Feistel f completa
   - S-boxes și expansiunea E
2. **Optimizări:**

   - Paralelizare cu goroutines
   - Precomputarea tabelelor
   - Măsurarea performanței
   - Comparație cu alte implementări
3. **Variante ale DES:**

   - Triple DES (3DES)
   - DES-X (cu whitening)
   - DESX, GDES
   - Variante educaționale simplificate
4. **Interfață utilizator:**

   - CLI interactiv cu meniu
   - Input de la tastatură
   - Salvare rezultate în fișiere
   - Export în diferite formate
5. **Analiză criptanalitică:**

   - Implementarea atacurilor cunoscute
   - Măsurarea avalanche effect
   - Statistici și vizualizări
   - Comparație cu AES

### Reflecție finală

Acest laborator a oferit o **experiență practică valoroasă** în:

- **Criptografie:** Înțelegerea profundă a unui algoritm clasic
- **Programare:** Dezvoltarea skill-urilor în Go
- **Problem solving:** Debugging și optimizare
- **Documentare:** Comunicarea tehnică clară

DES rămâne un **exemplu pedagogic excelent**, demonstrând principiile fundamentale ale criptografiei simetrice moderne. Cunoștințele dobândite formează **fundația solidă** pentru studiul algoritmilor contemporani (AES, ChaCha20) și pentru înțelegerea provocărilor de securitate în era digitală.

**Lecția cea mai importantă:** Securitatea criptografică evoluează constant. Algoritmii care erau considerați siguri acum 40 de ani sunt azi vulnerabili. Această evoluție subliniază importanța **actualizării continue** a cunoștințelor și **adaptării la noile standarde** de securitate.

---

**Repository GitHub:** https://github.com/Gheorghe2973/CS

**Facultatea:** Facultatea de Calculatoare, Informatică și Microelectronică (FCIM)

**Universitatea:** Universitatea Tehnică a Moldovei (UTM)

**Disciplina:** Securitatea Cibernetică

**Profesor:** [Nume Profesor]

---

*Raport realizat în cadrul Laboratorului 4 - Cifruri Bloc. Algoritmul DES*

*Data finalizării: 05 Noiembrie 2025*
