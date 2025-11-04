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
		result[i] = input[pos-1] // PC-2 folosește indexare de la 1
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

		// Numărul de shifturi pentru această rundă
		shifts := shiftSchedule[i]
		fmt.Printf("Număr de shifturi: %d\n\n", shifts)

		// Aplică left shift
		C = leftShift(C, shifts)
		D = leftShift(D, shifts)

		fmt.Printf("C₍%d₎ după shift: %s\n", i+1, formatBits(C, 28))
		fmt.Printf("D₍%d₎ după shift: %s\n\n", i+1, formatBits(D, 28))

		// Concatenează C și D
		CD := C + D
		fmt.Printf("C₍%d₎D₍%d₎ (56 biți): %s\n\n", i+1, i+1, formatBits(CD, 56))

		// Aplică PC-2 pentru a obține cheia de rundă
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
