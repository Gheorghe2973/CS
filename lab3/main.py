ALPHABET = [
    'A', 'Ă', 'Â', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'Î', 'J',
    'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'Ș', 'T', 'Ț', 'U',
    'V', 'W', 'X', 'Y', 'Z'
]


def validate_input(text):
    valid_chars = set(ALPHABET + [ch.lower() for ch in ALPHABET])
    for char in text:
        if char not in valid_chars and char != ' ':
            return False
    return True


def prepare_key(key):
    if len(key) < 7:
        raise ValueError("Cheia trebuie să aibă minim 7 caractere!")
    
    key = key.upper().replace(' ', '')
    
    if not validate_input(key):
        raise ValueError("Cheia conține caractere invalide! Folosiți doar A-Z, Ă, Â, Î, Ș, Ț")
    
    seen = set()
    unique_key = []
    for char in key:
        if char not in seen:
            seen.add(char)
            unique_key.append(char)
    
    return unique_key


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


def find_position(matrix, char):
    for i, row in enumerate(matrix):
        for j, cell in enumerate(row):
            if cell == char:
                return (i, j)
    return None


def prepare_text(text):
    text = text.upper().replace(' ', '')
    
    if not validate_input(text):
        raise ValueError("Textul conține caractere invalide! Folosiți doar A-Z, Ă, Â, Î, Ș, Ț")
    
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


def encrypt_pair(matrix, pair):
    char1, char2 = pair[0], pair[1]
    pos1 = find_position(matrix, char1)
    pos2 = find_position(matrix, char2)
    
    if pos1 is None or pos2 is None:
        return pair
    
    row1, col1 = pos1
    row2, col2 = pos2
    
    if row1 == row2:
        new_col1 = (col1 + 1) % 7
        new_col2 = (col2 + 1) % 7
        while new_col1 < 7 and matrix[row1][new_col1] == '':
            new_col1 = (new_col1 + 1) % 7
        while new_col2 < 7 and matrix[row2][new_col2] == '':
            new_col2 = (new_col2 + 1) % 7
        return matrix[row1][new_col1] + matrix[row2][new_col2]
    
    elif col1 == col2:
        new_row1 = (row1 + 1) % 5
        new_row2 = (row2 + 1) % 5
        return matrix[new_row1][col1] + matrix[new_row2][col2]
    
    else:
        return matrix[row1][col2] + matrix[row2][col1]


def decrypt_pair(matrix, pair):
    char1, char2 = pair[0], pair[1]
    pos1 = find_position(matrix, char1)
    pos2 = find_position(matrix, char2)
    
    if pos1 is None or pos2 is None:
        return pair
    
    row1, col1 = pos1
    row2, col2 = pos2
    
    if row1 == row2:
        new_col1 = (col1 - 1) % 7
        new_col2 = (col2 - 1) % 7
        while new_col1 >= 0 and matrix[row1][new_col1] == '':
            new_col1 = (new_col1 - 1) % 7
        while new_col2 >= 0 and matrix[row2][new_col2] == '':
            new_col2 = (new_col2 - 1) % 7
        return matrix[row1][new_col1] + matrix[row2][new_col2]
    
    elif col1 == col2:
        new_row1 = (row1 - 1) % 5
        new_row2 = (row2 - 1) % 5
        return matrix[new_row1][col1] + matrix[new_row2][col2]
    
    else:
        return matrix[row1][col2] + matrix[row2][col1]


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
    
    if not validate_input(ciphertext):
        raise ValueError("Textul conține caractere invalide! Folosiți doar A-Z, Ă, Â, Î, Ș, Ț")
    
    pairs = [ciphertext[i:i+2] for i in range(0, len(ciphertext), 2)]
    
    plaintext = ''
    for pair in pairs:
        if len(pair) == 2:
            plaintext += decrypt_pair(matrix, pair)
    
    return plaintext, matrix


def display_matrix(matrix):
    print("\nMatricea de criptare:")
    print("-" * 30)
    for row in matrix:
        print(' '.join(f'{cell:2}' for cell in row))
    print("-" * 30)


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
            print("Opțiune invalidă! Alegeți 1, 2 sau 3.")
            continue
        
        try:
            key = input("\nIntroduceți cheia (minim 7 caractere): ").strip()
            
            if choice == '1':
                plaintext = input("Introduceți mesajul de criptat: ").strip()
                
                ciphertext, matrix = encrypt(plaintext, key)
                display_matrix(matrix)
                
                print(f"\n{'='*50}")
                print(f"Mesaj original:  {plaintext}")
                print(f"Mesaj criptat:   {ciphertext}")
                print(f"{'='*50}")
                
            else:
                ciphertext = input("Introduceți criptograma de decriptat: ").strip()
                
                plaintext, matrix = decrypt(ciphertext, key)
                display_matrix(matrix)
                
                print(f"\n{'='*50}")
                print(f"Criptogramă:     {ciphertext}")
                print(f"Mesaj decriptat: {plaintext}")
                print(f"{'='*50}")
                print("\nNotă: Adăugați manual spațiile în mesaj.")
        
        except ValueError as e:
            print(f"\nEroare: {e}")
        except Exception as e:
            print(f"\nEroare neașteptată: {e}")


if __name__ == "__main__":
    main()