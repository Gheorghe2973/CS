1. Caesar 1 key (Romanian)
   k = 29
   11 4 17 23 29 12 15 14 4
   I B N S X I L K B

* Î(11) - 29 = -18 (mod 31) ≡ 13 → **K**
* C(4) - 29 = -25 (mod 31) ≡ 6 → **E**
* N(17) - 29 = -12 (mod 31) ≡ 19 → **Q**
* T(23) - 29 = -6 (mod 31) ≡ 25 → **V**
* X(29) - 29 (mod 31) = 0 → **A**
* J(12) - 29 = -17 (mod 31) ≡ 14 → **L**
* M(15) - 29 = -14 (mod 31) ≡ 17 → **N**
* L(14) - 29 = -15 (mod 31) ≡ 16 → **M**
* C(4) - 29 = -25 (mod 31) ≡ 6 → **F**

Decrypted message: KFQVALNMF


|   | 1 | 2  | 3  | 4  | 5    | 6  |
| - | - | -- | -- | -- | ---- | -- |
| 1 | A | Ă | Â | B  | C    | D  |
| 2 | E | F  | G  | H  | I/Î | J  |
| 3 | K | L  | M  | N  | O    | P  |
| 4 | Q | R  | S  | Ș | T    | Ț |
| 5 | U | V  | W  | X  | Y    | Z  |

k2 = 43 21 15 42 21 45 53 35 42 16
SECRETWORD

# Affine Cipher Decoding: Step-by-Step Solution

**Given:**

- Ciphertext: `qdrbkddkdwxwyxtdy`
- a = 5
- b = 9

## Step 1: Find the Modular Multiplicative Inverse of a

We need to find a⁻¹ such that (5 × a⁻¹) ≡ 1 (mod 26)

Testing values:

- 5 × 1 = 5 ≡ 5 (mod 26) ❌
- 5 × 2 = 10 ≡ 10 (mod 26) ❌
- 5 × 3 = 15 ≡ 15 (mod 26) ❌
- ...
- 5 × 21 = 105 = 4×26 + 1 ≡ 1 (mod 26) ✅

**Result: a⁻¹ = 21**

## Step 2: Decryption Formula

x = 21(y - 9) mod 26

Where:

- y = numerical value of encrypted letter (a=0, b=1, ..., z=25)
- x = numerical value of decrypted letter

## Step 3: Decrypt Each Letter


| Cipher | y  | Calculation                              | x  | Plain |
| ------ | -- | ---------------------------------------- | -- | ----- |
| q      | 16 | 21(16-9) = 21×7 = 147 ≡ 17 (mod 26)    | 17 | r     |
| d      | 3  | 21(3-9) = 21×(-6) = -126 ≡ 4 (mod 26)  | 4  | e     |
| r      | 17 | 21(17-9) = 21×8 = 168 ≡ 12 (mod 26)    | 12 | m     |
| b      | 1  | 21(1-9) = 21×(-8) = -168 ≡ 14 (mod 26) | 14 | o     |
| k      | 10 | 21(10-9) = 21×1 = 21 ≡ 21 (mod 26)     | 21 | v     |
| d      | 3  | 21(3-9) = 21×(-6) = -126 ≡ 4 (mod 26)  | 4  | e     |
| d      | 3  | 21(3-9) = 21×(-6) = -126 ≡ 4 (mod 26)  | 4  | e     |
| k      | 10 | 21(10-9) = 21×1 = 21 ≡ 21 (mod 26)     | 21 | v     |
| d      | 3  | 21(3-9) = 21×(-6) = -126 ≡ 4 (mod 26)  | 4  | e     |
| w      | 22 | 21(22-9) = 21×13 = 273 ≡ 13 (mod 26)   | 13 | n     |
| x      | 23 | 21(23-9) = 21×14 = 294 ≡ 8 (mod 26)    | 8  | i     |
| w      | 22 | 21(22-9) = 21×13 = 273 ≡ 13 (mod 26)   | 13 | n     |
| y      | 24 | 21(24-9) = 21×15 = 315 ≡ 3 (mod 26)    | 3  | d     |
| x      | 23 | 21(23-9) = 21×14 = 294 ≡ 8 (mod 26)    | 8  | i     |
| t      | 19 | 21(19-9) = 21×10 = 210 ≡ 2 (mod 26)    | 2  | c     |
| d      | 3  | 21(3-9) = 21×(-6) = -126 ≡ 4 (mod 26)  | 4  | e     |
| y      | 24 | 21(24-9) = 21×15 = 315 ≡ 3 (mod 26)    | 3  | d     |

## Step 4: Final Result

**Decrypted text: `removeevenindiced
