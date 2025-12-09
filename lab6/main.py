#!/usr/bin/env python3
import random
from math import gcd

message = """the first one was waiting in the wings..."""  # mesajul complet

# Hash SHA3-512 CORECT
hash_hex = "bcc0a3f43a0991cfa3af12ec3480b05e3563922d46bb8145c4642bab5dbcb37c8064717a4020bef8d5aa633aecfb0d97d0b63035e0965437396ba0eb9330b406"
hash_int = int(hash_hex, 16)

p = int("32317006071311007300153513477825163362488057133489075174588434139269806834136210002792056362640164685458556357935330816928882902308057347262527355474246124574102620252791657297286270630032526342821314576693141422365422094111134862999165747826803423055308634905063555771221918789033272956969612974385624174162372251973464026918557977679768230146253979330580152268587307611975324364674758554607150438968449403661304976978128542959586595675670512838521327844685229255045682728791137200989318739591433741758378260002780349731985520606075332341226032546840881200311059074842810039949669561196969562486290323380728391270390")
g = 2
p_minus_1 = p - 1

print("="*70)
print("    LABORATOR 6 - SARCINA 3: ELGAMAL cu SHA3-512")
print("="*70)
print()

print("--- HASH SHA3-512 ---")
print(f"Hash (hex): {hash_hex}")
print(f"Hash (decimal): {hash_int}")
print(f"Hash (mărime): {hash_int.bit_length()} biți")
print()

# Reducem hash-ul
hash_mod = hash_int % p_minus_1
print(f"Hash redus mod (p-1): {hash_mod}")
print()

# Generare chei
print("--- GENERARE CHEI ---")
x = random.randint(2, p_minus_1 - 1)
y = pow(g, x, p)
print(f"x (primii 60): {str(x)[:60]}...")
print(f"y (primii 60): {str(y)[:60]}...")
print()

# Semnare
print("--- SEMNARE ---")
while True:
    k = random.randint(2, p_minus_1 - 1)
    if gcd(k, p_minus_1) == 1:
        break

r = pow(g, k, p)
print(f"k (primii 60): {str(k)[:60]}...")
print(f"r (primii 60): {str(r)[:60]}...")

k_inv = pow(k, -1, p_minus_1)
x_r = (x * r) % p_minus_1
h_minus_xr = (hash_mod - x_r) % p_minus_1
s = (k_inv * h_minus_xr) % p_minus_1

print(f"s (primii 60): {str(s)[:60]}...")
print()

# Verificare algebrică
print("--- VERIFICARE ALGEBRICĂ ---")
k_s = (k * s) % p_minus_1
xr_plus_ks = (x_r + k_s) % p_minus_1
print(f"h ≡ x*r + k*s (mod p-1)? {hash_mod == xr_plus_ks}")
print()

# Verificare ElGamal
print("--- VERIFICARE ELGAMAL ---")
left_side = pow(g, hash_mod, p)
y_r = pow(y, r, p)
r_s = pow(r, s, p)
right_side = (y_r * r_s) % p

print(f"g^h mod p (primii 60):       {str(left_side)[:60]}...")
print(f"y^r * r^s mod p (primii 60): {str(right_side)[:60]}...")
print()

if left_side == right_side:
    print("✓✓✓ SEMNĂTURA ESTE VALIDĂ! ✓✓✓")
else:
    print("✗✗✗ SEMNĂTURA NU ESTE VALIDĂ! ✗✗✗")

print()
print("="*70)