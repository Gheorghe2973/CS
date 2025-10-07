# Laborator 1

## Sarcina 1.1 - Implementarea algoritmului Caesar simplu

### Descrierea sarcinii

De implementat algoritmul Cezar pentru alfabetul limbii engleze în unul din limbajele de programare. Utilizați doar codificarea literelor cum este arătat în tabelul 1 (nu se permite de folosit codificările specificate în limbajul de programare, de ex. ASCII sau Unicode). Valorile cheii vor fi cuprinse între 1 și 25 inclusiv și nu se permit alte valori. Valorile caracterelor textului sunt cuprinse între 'A' și 'Z', 'a' și 'z' și nu sunt premise alte valori. În cazul în care utilizatorul introduce alte valori - i se va sugera diapazonul corect. Înainte de criptare textul va fi transformat în majuscule și vor fi eliminate spațiile. Utilizatorul va putea alege operația - criptare sau decriptare, va putea introduce cheia, mesajul sau criptograma și va obține respectiv criptograma sau mesajul decriptat.

### Pași de rezolvare

#### 1. Reprezentarea alfabetului

Alfabetul este reprezentat ca o listă de caractere, fiecare literă având un index de la 0 la 25:

```python
ALPHABET = [
    'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
    'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'
]
```

#### 2. Conversie literă ↔ index

Funcțiile care permit conversii între litere și poziții numerice fără a folosi ASCII:

```python
def letter_to_index(letter, alphabet=None):
    if alphabet is None:
        alphabet = ALPHABET
    for i in range(len(alphabet)):
        if alphabet[i] == letter:
            return i
    return -1

def index_to_letter(index, alphabet=None):
    if alphabet is None:
        alphabet = ALPHABET
    return alphabet[index % 26]
```

#### 3. Validarea textului

Verifică că textul conține doar litere valide:

```python
def validate_text(text):
    for c in text:
        if c not in ALPHABET:
            return False
    return True
```

#### 4. Formulele de criptare și decriptare

**Formula de criptare:**

```
c = (m + k) mod 26
```

unde:

- `m` = poziția literei în alfabet (0-25)
- `k` = cheia (1-25)
- `c` = poziția literei criptate

**Formula de decriptare:**

```
m = (c - k) mod 26
```

#### 5. Implementarea criptării

```python
def caesar_encrypt(text, key):
    encrypted = ''
    for c in text:
        pos = letter_to_index(c)
        new_pos = (pos + key) % 26
        encrypted += index_to_letter(new_pos)
    return encrypted
```

**Exemplu:**

- Text: `HELLO`, Cheie: `3`
- `H` (7) → (7+3) mod 26 = 10 → `K`
- `E` (4) → (4+3) mod 26 = 7 → `H`
- `L` (11) → (11+3) mod 26 = 14 → `O`
- `L` (11) → (11+3) mod 26 = 14 → `O`
- `O` (14) → (14+3) mod 26 = 17 → `R`
- Rezultat: `KHOOR`

#### 6. Implementarea decriptării

```python
def caesar_decrypt(text, key):
    decrypted = ''
    for c in text:
        pos = letter_to_index(c)
        new_pos = (pos - key) % 26
        decrypted += index_to_letter(new_pos)
    return decrypted
```

---

## Sarcina 1.2 - Caesar cu permutare (2 chei)

### Descrierea sarcinii

De implementat algoritmul Cezar cu 2 chei, cu păstrarea condițiilor exprimate în Sarcina 1.1. În plus, cheia 2 trebuie să conțină doar litere ale alfabetului latin, și să aibă o lungime nu mai mică de 7.

### Pași de rezolvare

#### 1. Generarea alfabetului permutat

Alfabetul permutat se creează plasând literele din `key2` la început (fără duplicate), urmate de restul literelor în ordine naturală:

```python
def generate_permuted_alphabet(key2):
    key2_cleaned = key2.replace(" ", "").upper()
    permuted_alphabet = []
  
    for c in key2_cleaned:
        if c in ALPHABET and c not in permuted_alphabet:
            permuted_alphabet.append(c)
  
    for c in ALPHABET:
        if c not in permuted_alphabet:
            permuted_alphabet.append(c)
  
    return permuted_alphabet
```

**Exemplu:**

- `key2 = "cryptography"`
- Alfabet standard: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
- Alfabet permutat: `CRYPTOGAHBDEFIJKLMNQSUVWXZ`

#### 2. Formula de criptare cu permutare

**Procesul:**

1. Găsește poziția literei în alfabetul **standard**: `pos_standard`
2. Aplică deplasarea: `new_pos = (pos_standard + k1) mod 26`
3. Ia litera de la poziția `new_pos` din alfabetul **permutat**

```python
def caesar_encrypt_permuted(text, key1, key2):
    permuted_alphabet = generate_permuted_alphabet(key2)
    encrypted = ''
  
    for c in text:
        standard_pos = letter_to_index(c, ALPHABET)
        new_pos = (standard_pos + key1) % 26
        encrypted += permuted_alphabet[new_pos]
  
    return encrypted
```

**Exemplu:**

- Text: `ABC`, key1: `3`, key2: `cryptography`
- Alfabet permutat: `CRYPTOGAHBDEFIJKLMNQSUVWXZ`

```
'A': poziție standard = 0 → (0+3) mod 26 = 3 → permutat[3] = 'P'
'B': poziție standard = 1 → (1+3) mod 26 = 4 → permutat[4] = 'T'
'C': poziție standard = 2 → (2+3) mod 26 = 5 → permutat[5] = 'O'
Rezultat: "PTO"
```

#### 3. Formula de decriptare cu permutare

**Procesul:**

1. Găsește poziția literei în alfabetul **permutat**: `pos_permuted`
2. Aplică deplasarea inversă: `standard_pos = (pos_permuted - k1) mod 26`
3. Ia litera de la poziția `standard_pos` din alfabetul **standard**

```python
def caesar_decrypt_permuted(text, key1, key2):
    permuted_alphabet = generate_permuted_alphabet(key2)
    decrypted = ''
  
    for c in text:
        permuted_pos = letter_to_index(c, permuted_alphabet)
        standard_pos = (permuted_pos - key1) % 26
        decrypted += ALPHABET[standard_pos]
  
    return decrypted
```

#### 4. Validarea cheii secundare

```python
key2 = input("Enter key 2 (minimum 7 letters): ").strip()
key2_cleaned = key2.replace(" ", "").upper()

if len(key2_cleaned) < 7:
    print("Key 2 must have at least 7 characters.")
    return

if not validate_text(key2_cleaned):
    print("Key 2 can only contain letters A-Z or a-z.")
    return
```

---

## Structura completă a programului

```python
def main_menu():
    while True:
        print("\n1. Task 1.1 - Simple Caesar")
        print("2. Task 1.2 - Caesar with permutation")
        print("0. Exit")
      
        option = input("\nChoose option (0-2): ").strip()
      
        if option == '0':
            print("Goodbye!")
            break
        elif option == '1':
            task_1_1()
        elif option == '2':
            task_1_2()
        else:
            print("Invalid option!")

if __name__ == "__main__":
    main_menu()
```

---

## Concluzie

Laboratul a fost implementat cu succes în Python, respectând toate cerințele specificate în sarcini. Algoritmul Caesar a fost implementat pentru ambele variante - simplu și cu permutare.
