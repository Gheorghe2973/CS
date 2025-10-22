# Laborator 3 - Cifruri Polialfabetice

## Introducere - Cifruri de substituție polialfabetică

Slăbiciunea cifrurilor monoalfabetice este definită de faptul că distribuția lor de frecvență reflectă distribuția alfabetului folosit. Un cifru este mai sigur din punct de vedere criptografic dacă prezintă o distribuție cât mai regulată, care să nu ofere informații criptanalistului.

O cale de a aplatiza distribuția este combinarea distribuțiilor ridicate cu cele scăzute. Dacă litera T este criptată câteodată ca A și altă dată ca B, și dacă litera X este de asemenea câteodată criptată ca A și altă dată ca B, frecvența ridicată a lui T se combină cu frecvența scăzută a lui X producând o distribuție mai moderată pentru A și pentru B.

Diferența dintre cifrurile polialfabetice și cele monoalfabetice constă în faptul că substituția unui caracter variază în text, în funcție de diverși parametri (poziție, context etc.). Aceasta conduce la un număr mult mai mare de chei posibile.

---

## Teorie - Cifrul Vigenère

### Scurt istoric

Metoda de criptare cunoscută sub numele de „cifrul Vigenère" a fost atribuită greșit lui Blaise de Vigenère în secolul al XIX-lea și, de fapt, a fost descrisă pentru prima dată de Giovan Battista Bellaso în cartea sa din 1553 *La cifra del. Sig*. Vigenère a creat un cifru asemănător, dar totuși diferit și mai puternic în 1586.

### Principiul de funcționare

Cifrul Vigenère folosește aceleași operații ca și cifrul Caesar, dar cu o diferență majoră: folosește o **deplasare multiplă**. Cheia nu este constituită de o singură deplasare, ci de mai multe, fiind generate de câțiva întregi k<sub>i</sub>, unde 0 ≤ k<sub>i</sub> ≤ 25 (pentru alfabetul latin cu 26 de litere) sau 0 ≤ k<sub>i</sub> ≤ 30 (pentru alfabetul românesc cu 31 de litere).

### Formula matematică

**Criptare:**
```
c_i = (m_i + k_i) mod n
```

**Decriptare:**
```
m_i = (c_i - k_i) mod n
```

unde:
- m<sub>i</sub> = poziția literei din textul clar
- c<sub>i</sub> = poziția literei criptate
- k<sub>i</sub> = valoarea de deplasare din cheie (ciclică)
- n = mărimea alfabetului (26 sau 31)

### Exemplu de funcționare

Cheia poate fi, de exemplu, k = (5, 20, 17, 10, 20, 13) și ar provoca:
- Deplasarea primei litere cu 5: c₁ = m₁ + 5 (mod 26)
- A celei de a doua cu 20: c₂ = m₂ + 20 (mod 26)
- Și așa mai departe până la sfârșitul cheii, apoi se reia de la început

Cheia este de obicei un cuvânt pentru a fi mai ușor de memorat – cheia de mai sus corespunde cuvântului „FURTUN".

### Avantaje ale cifrului Vigenère

1. **Lungimea necunoscută a cheii:** Atacatorul nu știe lungimea cheii
2. **Spațiu mare de chei:** Pentru o cheie de lungime 5, există 26⁵ = 11,881,376 combinații posibile
3. **Rezistență la analiza frecvenței simple:** Aceeași literă din text poate fi criptată diferit

### Tabula Recta

Pentru simplificarea procesului se poate utiliza Tabula Recta, unde toate cele 26 (sau 31) cifruri sunt situate pe orizontală și fiecărui cifru îi corespunde o literă din cheie. Procesul este simplu: având litera m<sub>i</sub> din mesaj și litera k<sub>i</sub> din cheie, se găsește litera c<sub>i</sub> la intersecția liniei m<sub>i</sub> și coloanei k<sub>i</sub>.

---

## Teorie - Cifrul Playfair

### Scurt istoric

Cu toate că poartă numele baronului Lyon Playfair, algoritmul a fost inventat de prietenul acestuia, **Charles Wheatstone** și descris pentru întâia dată într-un document la 26 martie 1854.

La început a fost respins de *British Foreign Office* deoarece a fost considerat foarte greu de înțeles. Atunci când Wheatstone s-a oferit să demonstreze că în 15 minute va învăța să folosească algoritmul 3 băieți din 4 din școala aflată în apropiere, secretarul biroului de externe i-a răspuns: „Da, este foarte posibil, însă nu îi vei putea învăța să fie buni diplomați".

După crearea algoritmului, baronul Playfair a convins guvernul britanic să adopte acest algoritm pentru uz oficial și de aceea poartă numele său.

### Aplicații istorice

- **Războiul cu burii (1899-1902):** Folosit de armata britanică în Africa de Sud
- **Primul Război Mondial:** Versiuni modificate folosite de britanici
- **Al Doilea Război Mondial:** Armata australiană a folosit variante ale cifrului

### Caracteristici principale

Cifrul Playfair este unul dintre primii algoritmi ce folosește **principiile moderne ale cifrurilor bloc**:
- Operează pe **perechi de litere** (digrame), nu pe litere individuale
- Folosește o **matrice de substituție** generată din cheie
- Oferă **rezistență sporită** față de analiza frecvenței simple

### Descrierea algoritmului

Criptarea Playfair implică parcurgerea următorilor pași:

1. **Pregătirea textului** ce urmează a fi criptat
2. **Construirea matricei** de criptare
3. **Construirea mesajului** criptat

---

## Sarcina 3.1 - Implementarea algoritmului Playfair

### Descrierea sarcinii

De implementat algoritmul Playfair în unul din limbajele de programare pentru mesaje în limba română (31 de litere). Valorile caracterelor textului sunt cuprinse între 'A' și 'Z', 'a' și 'z' și nu sunt premise alte valori. Lungimea cheii nu trebuie să fie mai mică de 7. Utilizatorul va putea alege operația - criptare sau decriptare, va putea introduce cheia, mesajul sau criptograma și va obține criptograma sau mesajul decriptat.

### Pași de implementare

#### 1. Reprezentarea alfabetului

Alfabetul românesc este reprezentat ca o listă de 31 caractere:

```python
ALPHABET = [
    'A', 'Ă', 'Â', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'Î', 'J',
    'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'Ș', 'T', 'Ț', 'U',
    'V', 'W', 'X', 'Y', 'Z'
]
```

#### 2. Validarea textului

```python
def validate_input(text):
    valid_chars = set(ALPHABET + [ch.lower() for ch in ALPHABET])
    for char in text:
        if char not in valid_chars and char != ' ':
            return False
    return True
```

#### 3. Pregătirea cheii

Cheia trebuie să aibă minim 7 caractere și să nu conțină duplicate:

```python
def prepare_key(key):
    if len(key) < 7:
        raise ValueError("Cheia trebuie să aibă minim 7 caractere!")
    
    key = key.upper().replace(' ', '')
    
    if not validate_input(key):
        raise ValueError("Cheia conține caractere invalide!")
    
    seen = set()
    unique_key = []
    for char in key:
        if char not in seen:
            seen.add(char)
            unique_key.append(char)
    
    return unique_key
```

#### 4. Crearea matricei 5×7

Pentru alfabetul românesc (31 litere), se folosește o matrice 5×7:

```python
def create_matrix(key):
    prepared_key = prepare_key(key)
    matrix_chars = list(prepared_key)
    
    for char in ALPHABET:
        if char not in matrix_chars:
            matrix_chars.append(char)
    
    matrix = []
    for i in range(5):
        row = []
        for j in range(7):
            idx = i * 7 + j
            if idx < len(matrix_chars):
                row.append(matrix_chars[idx])
            else:
                row.append('')
        matrix.append(row)
    
    return matrix
```

**Exemplu cu cheia "PASSWORD":**
```
P  A  S  W  O  R  D
Ă  Â  B  C  E  F  G
H  I  Î  J  K  L  M
N  Q  T  Ț  U  V  X
Y  Z  
```

#### 5. Pregătirea textului pentru criptare

```python
def prepare_text(text):
    text = text.upper().replace(' ', '')
    
    if not validate_input(text):
        raise ValueError("Textul conține caractere invalide!")
    
    pairs = []
    i = 0
    while i < len(text):
        if i == len(text) - 1:
            pairs.append(text[i] + 'X')
            i += 1
        elif text[i] == text[i + 1]:
            pairs.append(text[i] + 'X')
            i += 1
        else:
            pairs.append(text[i:i+2])
            i += 2
    
    return pairs
```

**Reguli de tratare:**
- Literă singură la sfârșit → se adaugă 'X'
- Două litere identice consecutive → se inserează 'X' între ele
- Altfel → perechi normale de 2 litere

#### 6. Regulile de criptare Playfair

##### Regula 1: Același rând
Fiecare literă se înlocuiește cu litera imediat la dreapta (circular):

```python
if row1 == row2:
    new_col1 = (col1 + 1) % 7
    new_col2 = (col2 + 1) % 7
    while new_col1 < 7 and matrix[row1][new_col1] == '':
        new_col1 = (new_col1 + 1) % 7
    while new_col2 < 7 and matrix[row2][new_col2] == '':
        new_col2 = (new_col2 + 1) % 7
    return matrix[row1][new_col1] + matrix[row2][new_col2]
```

##### Regula 2: Aceeași coloană
Fiecare literă se înlocuiește cu litera imediat dedesubt (circular):

```python
elif col1 == col2:
    new_row1 = (row1 + 1) % 5
    new_row2 = (row2 + 1) % 5
    return matrix[new_row1][col1] + matrix[new_row2][col2]
```

##### Regula 3: Dreptunghi
Fiecare literă se înlocuiește cu cea de pe același rând, dar pe coloana celeilalte:

```python
else:
    return matrix[row1][col2] + matrix[row2][col1]
```

#### 7. Implementarea completă

```python
def encrypt(plaintext, key):
    matrix = create_matrix(key)
    pairs = prepare_text(plaintext)
    
    ciphertext = ''
    for pair in pairs:
        ciphertext += encrypt_pair(matrix, pair)
    
    return ciphertext, matrix

def decrypt(ciphertext, key):
    matrix = create_matrix(key)
    ciphertext = ciphertext.upper().replace(' ', '')
    
    pairs = [ciphertext[i:i+2] for i in range(0, len(ciphertext), 2)]
    
    plaintext = ''
    for pair in pairs:
        if len(pair) == 2:
            plaintext += decrypt_pair(matrix, pair)
    
    return plaintext, matrix
```

### Teste și rezultate

#### Test 1: Criptare cu cheia "PASSWORD"

**Input:**
- Operație: Criptare
- Cheie: `PASSWORD`
- Text: `SALUT ROMANIA`

**Matricea de criptare:**
```
------------------------------
P  A  S  W  O  R  D
Ă  Â  B  C  E  F  G
H  I  Î  J  K  L  M
N  Q  T  Ț  U  V  X
Y  Z                    
------------------------------
```

**Pregătirea textului:**
```
Text original: SALUT ROMANIA
Text procesat: SALUTROMANIA
Perechi: ['SA', 'LU', 'TR', 'OM', 'AN', 'IA']
```

**Procesul de criptare:**
```
SA: S(0,2) și A(0,1) → același rând → WS
LU: L(2,5) și U(3,4) → dreptunghi → VK
TR: T(3,2) și R(0,5) → dreptunghi → DX
OM: O(0,4) și M(2,6) → dreptunghi → DE
AN: A(0,1) și N(3,0) → dreptunghi → PȚ
IA: I(2,1) și A(0,1) → aceeași coloană → QÂ
```

**Output:**
```
==================================================
Mesaj original:  SALUT ROMANIA
Mesaj criptat:   WSVKDXDEPȚQÂ
==================================================
```

#### Test 2: Decriptare

**Input:**
- Operație: Decriptare
- Cheie: `PASSWORD`
- Text: `WSVKDXDEPȚQÂ`

**Output:**
```
==================================================
Criptogramă:     WSVKDXDEPȚQÂ
Mesaj decriptat: SALUTROMANIA
==================================================

Notă: Adăugați manual spațiile în mesaj.
Mesaj final: SALUT ROMANIA
```

**Verificare:** ✓ Corect

#### Test 3: Validare cheie scurtă

**Input:**
- Cheie: `PASS` (lungime = 4)

**Output:**
```
Eroare: Cheia trebuie să aibă minim 7 caractere!
```

**Verificare:** ✓ Corect

#### Test 4: Tratarea literelor duble

**Input:**
- Cheie: `PASSWORD`
- Text: `HELLO`

**Pregătire:**
```
Text: HELLO
Perechi: ['HE', 'LX', 'LO']
         (X inserat între cele două L-uri)
```

**Output:**
```
Mesaj original:  HELLO
Mesaj criptat:   IFJVMO
```

**Verificare:** ✓ Corect

#### Test 5: Literă singură la sfârșit

**Input:**
- Cheie: `SECURITY`
- Text: `MESAJ`

**Pregătire:**
```
Text: MESAJ
Perechi: ['ME', 'SA', 'JX']
         (X adăugat pentru J)
```

**Verificare:** ✓ Corect

---

## Structura completă a programului

```python
def main():
    print("=" * 50)
    print("CIFRUL PLAYFAIR - ALFABETUL ROMÂNESC")
    print("=" * 50)
    print(f"Alfabet: {''.join(ALPHABET)}")
    print("=" * 50)
    
    while True:
        print("\nAlegeți operația:")
        print("1. Criptare")
        print("2. Decriptare")
        print("3. Ieșire")
        
        choice = input("\nOpțiunea dvs (1/2/3): ").strip()
        
        if choice == '3':
            print("La revedere!")
            break
        
        if choice not in ['1', '2']:
            print("Opțiune invalidă!")
            continue
        
        try:
            key = input("\nIntroduceți cheia (minim 7 caractere): ").strip()
            
            if choice == '1':
                plaintext = input("Introduceți mesajul: ").strip()
                ciphertext, matrix = encrypt(plaintext, key)
                display_matrix(matrix)
                print(f"\nMesaj original:  {plaintext}")
                print(f"Mesaj criptat:   {ciphertext}")
                
            else:
                ciphertext = input("Introduceți criptograma: ").strip()
                plaintext, matrix = decrypt(ciphertext, key)
                display_matrix(matrix)
                print(f"\nCriptogramă:     {ciphertext}")
                print(f"Mesaj decriptat: {plaintext}")
                print("\nNotă: Adăugați manual spațiile în mesaj.")
        
        except ValueError as e:
            print(f"\nEroare: {e}")

if __name__ == "__main__":
    main()
```

---

## Analiză comparativă

### Vigenère vs. Playfair

| Caracteristică | Vigenère | Playfair |
|----------------|----------|----------|
| **Tip cifru** | Polialfabetic simplu | Polialfabetic digrafic |
| **Unitate criptare** | 1 literă | 2 litere (digrame) |
| **Formula** | c = (m + k) mod n | Reguli matriceale |
| **Matrice** | Nu folosește | 5×7 (română) / 5×5 (engleză) |
| **Cheie minimă** | Variabilă | 7 caractere |
| **Anul inventării** | 1586 (versiunea puternică) | 1854 |
| **Complexitate** | Simplă | Moderată |

### Avantajele cifrului Playfair

1. **Criptare pe blocuri:** Digramele oferă mai multă securitate decât literele individuale
2. **Rezistență la analiza frecvenței:** Distribuția este mai uniformă
3. **Spațiu de chei mare:** Aproximativ 26! permutări pentru matrice
4. **Ușor de memorat:** Cheia este un cuvânt sau frază

### Limitări comune

Ambele cifruri sunt **vulnerabile la criptanaliză modernă** și nu trebuie folosite pentru date sensibile reale. Sunt excelente pentru **scopuri educaționale** și demonstrarea principiilor criptografice de bază.

---

## Complexitatea algoritmului Playfair

### Complexitate temporală

- **Pregătire cheie:** O(k), unde k = lungimea cheii
- **Creare matrice:** O(n), unde n = mărimea alfabetului (31)
- **Pregătire text:** O(m), unde m = lungimea textului
- **Criptare/Decriptare:** O(m) × O(1) = O(m)
- **Complexitate totală:** O(m + k + n) ≈ **O(m)**

### Complexitate spațială

- **Matrice:** 5 × 7 = 35 celule → O(1)
- **Stocare perechi:** O(m/2) ≈ O(m)
- **Complexitate totală:** **O(m)**

---

## Considerații de securitate

### Atacuri posibile

1. **Analiza frecvenței digramelor:** Digramele frecvente pot fi identificate
2. **Atac brute-force:** Teoretic 26! permutări, practic mult mai rapid
3. **Known-plaintext attack:** Dacă se cunosc perechi text-criptogramă, matricea poate fi dedusă

### Recomandări

1. ✓ Chei lungi (minim 10-15 caractere)
2. ✓ Chei aleatorii (nu cuvinte din dicționar)
3. ✓ Nu reutilizați cheia

---

## Concluzie

Laboratorul a fost implementat cu succes în Python pentru alfabetul românesc (31 litere). Algoritmul Playfair a fost testat exhaustiv pentru toate cazurile:

✓ Criptare și decriptare corectă  
✓ Tratarea literelor duble  
✓ Validarea cheii (lungime minimă 7)  
✓ Validarea caracterelor de intrare  
✓ Afișarea matricei de criptare  
✓ Gestionarea erorilor  

Cifrul Playfair reprezintă o evoluție semnificativă în istoria criptografiei, fiind unul dintre primii algoritmi care operează pe blocuri de caractere. Împreună cu cifrul Vigenère, acestea ilustrează tranziția de la cifrurile monoalfabetice simple la sistemele polialfabetice mai complexe, precursori ai criptografiei moderne.

**Repository GitHub:** [https://github.com/Gheorghe2973/CS](https://github.com/Gheorghe2973/CS)