ALPHABET = [
    'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
    'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'
]

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

def validate_text(text):
    for c in text:
        if c not in ALPHABET:
            return False
    return True

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

def caesar_encrypt(text, key):
    encrypted = ''
    for c in text:
        pos = letter_to_index(c)
        new_pos = (pos + key) % 26
        encrypted += index_to_letter(new_pos)
    return encrypted

def caesar_decrypt(text, key):
    decrypted = ''
    for c in text:
        pos = letter_to_index(c)
        new_pos = (pos - key) % 26
        decrypted += index_to_letter(new_pos)
    return decrypted

def caesar_encrypt_permuted(text, key1, key2):
    permuted_alphabet = generate_permuted_alphabet(key2)
    encrypted = ''
    
    for c in text:
        standard_pos = letter_to_index(c, ALPHABET)
        new_pos = (standard_pos + key1) % 26
        encrypted += permuted_alphabet[new_pos]
    
    return encrypted

def caesar_decrypt_permuted(text, key1, key2):
    permuted_alphabet = generate_permuted_alphabet(key2)
    decrypted = ''
    
    for c in text:
        permuted_pos = letter_to_index(c, permuted_alphabet)
        standard_pos = (permuted_pos - key1) % 26
        decrypted += ALPHABET[standard_pos]
    
    return decrypted

def task_1_1():
    print("\n" + "="*60)
    print("TASK 1.1 - Simple Caesar Cipher")
    print("="*60)
    
    print("\nChoose operation:")
    print("1. Encryption")
    print("2. Decryption")
    
    option = input("\nEnter option (1/2): ").strip()
    if option not in ['1', '2']:
        print("Invalid option. Choose '1' or '2'.")
        return

    try:
        key = int(input("Enter key (1-25): "))
        if key < 1 or key > 25:
            print("Key must be between 1 and 25.")
            return
    except ValueError:
        print("Key must be an integer.")
        return

    text = input("Enter text: ").replace(" ", "").upper()
    if not validate_text(text):
        print("Text can only contain letters A-Z or a-z.")
        return

    print("\n" + "-"*60)
    if option == '1':
        result = caesar_encrypt(text, key)
        print(f"Original text:  {text}")
        print(f"Key:            {key}")
        print(f"Encrypted text: {result}")
    else:
        result = caesar_decrypt(text, key)
        print(f"Ciphertext:     {text}")
        print(f"Key:            {key}")
        print(f"Decrypted text: {result}")
    print("-"*60)

def task_1_2():
    print("\n" + "="*60)
    print("TASK 1.2 - Caesar Cipher with Permutation")
    print("="*60)
    
    print("\nChoose operation:")
    print("1. Encryption")
    print("2. Decryption")
    
    option = input("\nEnter option (1/2): ").strip()
    if option not in ['1', '2']:
        print("Invalid option. Choose '1' or '2'.")
        return

    try:
        key1 = int(input("Enter key 1 (1-25): "))
        if key1 < 1 or key1 > 25:
            print("Key 1 must be between 1 and 25.")
            return
    except ValueError:
        print("Key 1 must be an integer.")
        return

    key2 = input("Enter key 2 (minimum 7 letters): ").strip()
    key2_cleaned = key2.replace(" ", "").upper()
    
    if len(key2_cleaned) < 7:
        print("Key 2 must have at least 7 characters.")
        return
    
    if not validate_text(key2_cleaned):
        print("Key 2 can only contain letters A-Z or a-z.")
        return

    text = input("Enter text: ").replace(" ", "").upper()
    if not validate_text(text):
        print("Text can only contain letters A-Z or a-z.")
        return

    permuted_alphabet = generate_permuted_alphabet(key2)
    alphabet_str = ''.join(permuted_alphabet)

    print("\n" + "-"*60)
    print(f"Standard alphabet: {''.join(ALPHABET)}")
    print(f"Permuted alphabet: {alphabet_str}")
    print("-"*60)
    
    if option == '1':
        result = caesar_encrypt_permuted(text, key1, key2)
        print(f"Original text:    {text}")
        print(f"Key 1:            {key1}")
        print(f"Key 2:            {key2}")
        print(f"Encrypted text:   {result}")
    else:
        result = caesar_decrypt_permuted(text, key1, key2)
        print(f"Ciphertext:       {text}")
        print(f"Key 1:            {key1}")
        print(f"Key 2:            {key2}")
        print(f"Decrypted text:   {result}")
    print("-"*60)

def main_menu():
    while True:
        print("\n" + "="*60)
        print("CAESAR CIPHER - LABORATORY 1")
        print("="*60)
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