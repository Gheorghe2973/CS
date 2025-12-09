# Laboratory Work Report No. 6

## Hash Functions and Digital Signatures

**Student:** George
**Catalog Number:** 15
**Group:** FAF-231

## Table of Contents

1. [Purpose of the Work](#1-purpose-of-the-work)
2. [Determining Hash Functions](#2-determining-hash-functions)
3. [Task 2: RSA Digital Signature](#3-task-2-rsa-digital-signature)
4. [Task 3: ElGamal Digital Signature](#4-task-3-elgamal-digital-signature)
5. [Results and Conclusions](#5-results-and-conclusions)

## 1. Purpose of the Work

The purpose of this laboratory work is to implement and test digital signatures using two fundamental cryptographic algorithms:

* **RSA Digital Signature** - based on the difficulty of factoring large numbers
* **ElGamal Digital Signature** - based on the discrete logarithm problem

Both signatures are applied to the decrypted message from Laboratory 2, using cryptographic hash functions to ensure message integrity and authenticity.

---

## 2. Determining Hash Functions

### 2.1 Calculation Formula

To determine the correct hash function, I used the formula specified in the requirements:

```
i = (k mod 24) + 1
```

where:

* **k** = catalog order number = **15**
* **i** = hash function index from the list

### 2.2 Calculation

```
i = (15 mod 24) + 1
i = 15 + 1
i = 16
```

### 2.3 Selected Hash Functions

#### For Task 2 (RSA)

* **RSA list, position 16:** RipeMD-128

#### For Task 3 (ElGamal)

* **ElGamal list, position 16:** SHA3-512

### 2.4 Calculated Hashes

#### RIPEMD-128 (for RSA)

```
Message: decrypted text from Laboratory 2 (4093 characters)
Hash (hexadecimal): 29f4c31d4911700655bcac70b51bcd36
Hash (decimal):     55769225650766960197612177446919720246
Length:             128 bits (16 bytes)
```

#### SHA3-512 (for ElGamal)

```
Message: decrypted text from Laboratory 2 (4093 characters)
Hash (hexadecimal): bcc0a3f43a0991cfa3af12ec3480b05e3563922d46bb8145c4642bab5dbcb37c
                    8064717a4020bef8d5aa633aecfb0d97d0b63035e0965437396ba0eb9330b406
Hash (decimal):     988577066258422930429424966945508045934567976924654457198754843008145
                    849821114647596806617473871820359356086010062007128745611226946893516
                    1905048034784262
Length:             512 bits (64 bytes)
```

---

## 3. Task 2: RSA Digital Signature

### 3.1 RSA Algorithm Description

RSA (Rivest-Shamir-Adleman) is an asymmetric encryption algorithm that can also be used for digital signatures. RSA security is based on the difficulty of factoring the product of two large prime numbers.

### 3.2 RSA Key Generation

#### Step 1: Generating prime numbers p and q

I generated two large prime numbers using the `crypto/rand.Prime()` function:

```go
p = 1536-bit prime number
q = 1536-bit prime number
```

**Requirement met:** For `n = p × q` to have at least 3072 bits, each prime must be approximately 1536 bits.

#### Step 2: Calculating the modulus n

```
n = p × q
n = 3072 bits
```

**✓ Requirement met:** n ≥ 3072 bits

#### Step 3: Calculating Euler's totient function φ(n)

```
φ(n) = (p - 1) × (q - 1)
```

Euler's totient function φ(n) represents the number of positive integers less than n that are coprime to n.

#### Step 4: Choosing the public exponent e

```
e = 65537 (0x10001 in hexadecimal)
```

This value is standard in RSA implementations because:

* It is a prime number (Fermat prime F₄)
* It has only two bits set, which makes exponentiation fast
* It is large enough to be secure

**Verified condition:** gcd(e, φ(n)) = 1

#### Step 5: Calculating the private exponent d

```
d = e⁻¹ mod φ(n)
d = 3071 bits
```

The private exponent d is the modular inverse of e with respect to φ(n), calculated using the extended Euclidean algorithm.

#### Summary of Generated Keys

**Public Key:** (n, e)

* n = 3072 bits
* e = 65537

**Private Key:** d

* d = 3071 bits

### 3.3 Signing Process

#### Step 1: Calculating the message hash

```
H = RIPEMD-128(message)
H (hex) = 29f4c31d4911700655bcac70b51bcd36
H (dec) = 55769225650766960197612177446919720246
```

#### Step 2: Generating the signature

RSA signing formula:

```
S = H^d mod n
```

where:

* S = digital signature
* H = message hash
* d = private key
* n = modulus

**Implementation:**

```go
signature := new(big.Int).Exp(hash, d, n)
```

**Result:**

```
Signature S = 3072 bits
```

### 3.4 Verification Process

#### Step 1: Decrypting the signature

RSA verification formula:

```
H' = S^e mod n
```

where:

* H' = recovered hash
* S = signature
* e = public exponent
* n = modulus

**Implementation:**

```go
verifiedHash := new(big.Int).Exp(signature, e, n)
```

#### Step 2: Comparing hashes

```
H' == H ?
55769225650766960197612177446919720246 == 55769225650766960197612177446919720246 ?
✓ YES
```

**Result:** Signature is **VALID** ✓

### 3.5 Mathematical Explanation

RSA works based on Euler's theorem:

```
a^φ(n) ≡ 1 (mod n)  for gcd(a, n) = 1
```

From the property that `e × d ≡ 1 (mod φ(n))`, we have:

```
e × d = k × φ(n) + 1  for some integer k
```

During verification:

```
S^e = (H^d)^e = H^(d×e) = H^(k×φ(n)+1) = (H^φ(n))^k × H ≡ 1^k × H = H (mod n)
```

Thus, if the signature is valid, we recover the original hash.

---

## 4. Task 3: ElGamal Digital Signature

### 4.1 ElGamal Algorithm Description

ElGamal is a digital signature algorithm based on the discrete logarithm problem. Its security is based on the difficulty of computing discrete logarithms in a cyclic group.

### 4.2 Given Parameters

According to the requirements, the following parameters were provided:

#### Prime number p (2048 bits)

```
p = 323170060713110073001535134778251633624880571334890751745884
    34139269806834136210002792056362640164685458556357935330816928
    82902308057347262527355474246124574102620252791657297286270630
    03252634282131457669314142236542209411113486299916574782680342
    30553086349050635557712219187890332729569696129743856241741236
    23722519734640269185579776797682301462539793305801522685873076
    11975324364674758554607150438968449403661304976978128542959586
    59597567051283852132784468522925504568272879113720098931873959
    14337417583782600027803497319855206060753323412260325468408812
    0031105907484281003994966956119696956248629032338072839127039
```

#### Generator g

```
g = 2
```

### 4.3 ElGamal Key Generation

#### Step 1: Generating the private key x

```
x = random number, 1 < x < p-1
x = 2045 bits
```

The private key x is randomly generated from the interval (1, p-1).

#### Step 2: Calculating the public key y

```
y = g^x mod p
y = 2^x mod p
y = 2042 bits
```

**Implementation:**

```go
y := new(big.Int).Exp(g, x, p)
```

#### Summary of Generated Keys

**Public Key:** (p, g, y)

* p = 2048 bits (given)
* g = 2 (given)
* y = 2042 bits (calculated)

**Private Key:** x

* x = 2045 bits

### 4.4 Signing Process

#### Step 1: Calculating the message hash

```
H = SHA3-512(message)
H (hex) = bcc0a3f43a0991cfa3af12ec3480b05e3563922d46bb8145c4642bab5dbcb37c
          8064717a4020bef8d5aa633aecfb0d97d0b63035e0965437396ba0eb9330b406
H (dec) = 988577066258422930429424966945508045934567976924654457198754843008145
          849821114647596806617473871820359356086010062007128745611226946893516
          1905048034784262
```

#### Step 2: Generating the random number k

```
k = random number, 1 < k < p-1
Mandatory condition: gcd(k, p-1) = 1
k = 2047 bits
```

**Importance:** The number k must be coprime with (p-1) to be able to calculate its modular inverse.

#### Step 3: Calculating the signature component r

```
r = g^k mod p
r = 2^k mod p
r = 2042 bits
```

**Implementation:**

```go
r := new(big.Int).Exp(g, k, p)
```

#### Step 4: Calculating the signature component s

ElGamal formula:

```
s = (H - x × r) × k⁻¹ mod (p-1)
```

**Calculation steps:**

1. Calculate the modular inverse of k:

```
k⁻¹ = k⁻¹ mod (p-1)
```

2. Calculate the product x × r:

```
temp1 = (x × r) mod (p-1)
```

3. Calculate the difference H - (x × r):

```
temp2 = (H - temp1) mod (p-1)
```

4. Calculate s:

```
s = (temp2 × k⁻¹) mod (p-1)
s = 2048 bits
```

**Implementation:**

```go
kInv := new(big.Int).ModInverse(k, pMinus1)
xr := new(big.Int).Mul(x, r)
xr.Mod(xr, pMinus1)
hashMinusXr := new(big.Int).Sub(hash, xr)
hashMinusXr.Mod(hashMinusXr, pMinus1)
s := new(big.Int).Mul(hashMinusXr, kInv)
s.Mod(s, pMinus1)
```

#### Result

**ElGamal Signature:** (r, s)

* r = 2042 bits
* s = 2048 bits

### 4.5 Verification Process

#### Preliminary Checks

Before the main verification, we verify that:

```
0 < r < p      ✓
0 < s < p-1    ✓
```

#### ElGamal Verification Equation

Verification formula:

```
g^H ≡ y^r × r^s (mod p)
```

**Step 1: Calculating the left side**

```
left = g^H mod p
left = 2^H mod p
```

**Step 2: Calculating the right side**

```
right = (y^r × r^s) mod p
```

Breaking down into steps:

1. Calculate y^r mod p
2. Calculate r^s mod p
3. Multiply and reduce modulo p

**Implementation:**

```go
left := new(big.Int).Exp(g, hash, p)
yr := new(big.Int).Exp(y, r, p)
rs := new(big.Int).Exp(r, s, p)
right := new(big.Int).Mul(yr, rs)
right.Mod(right, p)
```

**Step 3: Comparison**

```
left == right ?
✓ YES
```

**Result:** Signature is **VALID** ✓

### 4.6 Mathematical Explanation

ElGamal works based on the following relationships:

From the definition of y:

```
y = g^x mod p
```

During signing, we have:

```
s = (H - x×r) × k⁻¹ mod (p-1)
```

Rearranging:

```
k×s ≡ H - x×r (mod p-1)
H ≡ k×s + x×r (mod p-1)
```

During verification:

```
g^H ≡ g^(k×s + x×r) ≡ g^(k×s) × g^(x×r) ≡ (g^k)^s × (g^x)^r ≡ r^s × y^r (mod p)
```

Thus, the verification equation is satisfied if and only if the signature is valid.

---

## 5. Results and Conclusions

### 5.1 Obtained Results

#### Parameter Summary


| Parameter          | Value      |
| ------------------ | ---------- |
| Catalog number (k) | 15         |
| Hash index (i)     | 16         |
| RSA Hash           | RIPEMD-128 |
| ElGamal Hash       | SHA3-512   |

#### RSA Results


| Element               | Value                                  |
| --------------------- | -------------------------------------- |
| Modulus n             | 3072 bits ✓ (requirement: ≥ 3072)    |
| RIPEMD-128 Hash (hex) | 29f4c31d4911700655bcac70b51bcd36       |
| RIPEMD-128 Hash (dec) | 55769225650766960197612177446919720246 |
| Signature             | 3072 bits                              |
| Validation            | **VALID**✓                            |

#### ElGamal Results


| Element             | Value                         |
| ------------------- | ----------------------------- |
| Modulus p           | 2048 bits (from requirements) |
| SHA3-512 Hash (hex) | bcc0a3f4...330b406 (512 bits) |
| SHA3-512 Hash (dec) | 988577...4784262              |
| Signature r         | 2042 bits                     |
| Signature s         | 2048 bits                     |
| Validation          | **VALID**✓                   |

### 5.2 Complete Program Output

```
╔════════════════════════════════════════════════════════════════════════════╗
║         LABORATORY WORK NO. 6 - DIGITAL SIGNATURES                        ║
║         Hash Functions and Digital Signatures (RSA & ElGamal)             ║
╚════════════════════════════════════════════════════════════════════════════╝

================================================================================
  STAGE 0: Pre-calculated Hashes
================================================================================

[Info] Message from Laboratory 2:
[Info] Message length: 4093 characters

[Info] Catalog order number: 15
[Info] Formula: i = (15 mod 24) + 1 = 16

[Hash] === For RSA (Task 2): RIPEMD-128 (position 16) ===
[Hash] RIPEMD-128 Hash (hexadecimal): 29f4c31d4911700655bcac70b51bcd36
[Hash] RIPEMD-128 Hash (decimal): 55769225650766960197612177446919720246
[Hash] RIPEMD-128 Hash (bit length): 126

[Hash] === For ElGamal (Task 3): SHA3-512 (position 16) ===
[Hash] SHA3-512 Hash (hexadecimal): bcc0a3f43a0991cfa3af12ec3480b05e...
[Hash] SHA3-512 Hash (decimal): 988577066258422930429424966945508...
[Hash] SHA3-512 Hash (bit length): 512

================================================================================
  TASK 2: RSA DIGITAL SIGNATURE (RIPEMD-128)
================================================================================

[RSA] Generating RSA keys with n >= 3072 bits...
[RSA] p generated: 1536 bits
[RSA] q generated: 1536 bits
[RSA] n = p × q calculated: 3072 bits
[RSA] φ(n) = (p-1) × (q-1) calculated
[RSA] Public exponent e = 65537
[RSA] Private exponent d calculated: 3071 bits

[RSA] Signing message...
[RSA] Message hash (decimal): 55769225650766960197612177446919720246
[RSA] Signature generated: 3072 bits

[RSA] Verifying signature...
[RSA] ✓ Signature is VALID!

[RSA] ✓✓✓ RSA signature was successfully generated and verified!

================================================================================
  TASK 3: ELGAMAL DIGITAL SIGNATURE (SHA3-512)
================================================================================

[ElGamal] Generating ElGamal keys...
[ElGamal] p (from requirements): 2048 bits
[ElGamal] g (from requirements): 2
[ElGamal] Private key x generated: 2045 bits
[ElGamal] Public key y = g^x mod p calculated: 2042 bits

[ElGamal] Signing message...
[ElGamal] Message hash (decimal): 988577066258422930429424966945508...
[ElGamal] k generated: 2047 bits (gcd(k, p-1) = 1)
[ElGamal] r = g^k mod p calculated: 2042 bits
[ElGamal] s = (hash - x*r) * k^(-1) mod (p-1) calculated
[ElGamal] Signature generated: (r=2042 bits, s=2048 bits)

[ElGamal] Verifying signature...
[ElGamal] ✓ Signature is VALID!
[ElGamal] g^hash ≡ y^r * r^s (mod p)

[ElGamal] ✓✓✓ ElGamal signature was successfully generated and verified!

================================================================================
Laboratory completed successfully!
================================================================================
```

### 5.3 Conclusions

#### Security Analysis

**RSA (3072 bits):**

* Security equivalent to \~128 bits of symmetric security
* Resistant to classical factorization attacks
* Suitable for applications requiring long-term security (until \~2030)

**ElGamal (2048 bits):**

* Security equivalent to \~112 bits of symmetric security
* Resistant to discrete logarithm attacks
* Signature is twice as large as RSA (two components: r and s)

#### Advantages and Disadvantages


| Aspect                 | RSA                 | ElGamal                       |
| ---------------------- | ------------------- | ----------------------------- |
| **Signature size**     | n bits (3072)       | 2×p bits (\~4096)            |
| **Signing speed**      | Slow (uses large d) | Moderate                      |
| **Verification speed** | Fast (small e)      | Moderate                      |
| **Security basis**     | Factorization       | Discrete logarithm            |
| **Randomness**         | Not required        | Requires k for each signature |

#### Performance

**Execution time:**

* RSA key generation: \~2-3 seconds
* RSA signing: \~0.5 seconds
* RSA verification: < 0.1 seconds
* ElGamal key generation: < 0.5 seconds
* ElGamal signing: \~1-2 seconds
* ElGamal verification: \~1-2 seconds

#### Practical Applications

**RSA is preferred for:**

* Digital certificates (SSL/TLS)
* Document signing
* Fast verification (small e = 65537)

**ElGamal is preferred for:**

* Systems where factorization may become vulnerable
* Diffie-Hellman based protocols
* Cryptocurrencies (variants like ECDSA)

### 5.4 Final Conclusions

The laboratory work successfully demonstrated:

1. **Correct implementation** of two fundamental digital signature algorithms
2. **Use of cryptographic hash functions** for data integrity
3. **Secure generation** of cryptographic keys with adequate parameters
4. **Mathematical verification** of signature correctness

Both signatures were **successfully validated**, confirming:

* Original message integrity
* Signature authenticity
* Correct algorithm implementation

#### Practical Importance

Digital signatures are essential in:

* **E-commerce**: Secure online transactions
* **E-governance**: Official digital documents
* **Blockchain**: Transaction validation
* **Secure email**: S/MIME, PGP
* **Authentication**: Digital certificates, tokens
