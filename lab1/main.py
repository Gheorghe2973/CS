def encode_letter(letter):
    alphabet = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 
                'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
    
    for i in range(26):
        if alphabet[i] == letter:
            return i
    return -1

def decode_letter(code):
    alphabet = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 
                'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
    
    if code >= 0 and code <= 25:
        return alphabet[code]
    return ''

def to_uppercase(letter):
    lowercase = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
                 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
    uppercase = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
                 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
    
    for i in range(26):
        if lowercase[i] == letter:
            return uppercase[i]
    return letter

def is_letter(char):
    uppercase = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
                 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
    lowercase = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
                 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
    
    for letter in uppercase:
        if char == letter:
            return True
    for letter in lowercase:
        if char == letter:
            return True
    return False

def prepare_text(text):
    cleaned_text = ""
    for character in text:
        if character != ' ':
            if is_letter(character):
                cleaned_text += to_uppercase(character)
    return cleaned_text

def encrypt(text, key):
    result = ""
    text = prepare_text(text)
    
    for i in range(len(text)):
        letter = text[i]
        code = encode_letter(letter)
        
        if code != -1:
            new_code = (code + key) % 26
            result += decode_letter(new_code)
    
    return result

def decrypt(text, key):
    result = ""
    text = prepare_text(text)
    
    for i in range(len(text)):
        letter = text[i]
        code = encode_letter(letter)
        
        if code != -1:
            new_code = (code - key) % 26
            result += decode_letter(new_code)
    
    return result

def validate_key(key_str):
    digits = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
    
    if len(key_str) == 0:
        return False, 0
    
    for char in key_str:
        is_digit = False
        for digit in digits:
            if char == digit:
                is_digit = True
                break
        if not is_digit:
            return False, 0
    
    key = 0
    for char in key_str:
        for i in range(10):
            if char == digits[i]:
                key = key * 10 + i
                break
    
    if key >= 1 and key <= 25:
        return True, key
    else:
        return False, 0

def main_menu():
    print("=" * 50)
    print("CAESAR CIPHER - ENCRYPTION/DECRYPTION PROGRAM")
    print("=" * 50)
    
    while True:
        print("\nChoose operation:")
        print("1. Encryption")
        print("2. Decryption")
        print("3. Exit")
        
        option = input("\nEnter option (1/2/3): ")
        
        if option == '3':
            print("Goodbye!")
            break
        
        if option != '1' and option != '2':
            print("Invalid option! Choose 1, 2 or 3.")
            continue
        
        while True:
            key_input = input("\nEnter key (1-25): ")
            valid, key = validate_key(key_input)
            
            if valid:
                break
            else:
                print("Invalid key! Enter a number between 1 and 25.")
        
        if option == '1':
            text = input("\nEnter message to encrypt: ")
            processed_text = prepare_text(text)
            result = encrypt(text, key)
            
            print("\n" + "-" * 50)
            print("ENCRYPTION RESULTS:")
            print("-" * 50)
            print("Original text: " + text)
            print("Processed text: " + processed_text)
            print("Key: " + str(key))
            print("Ciphertext: " + result)
            print("-" * 50)
        
        else:
            text = input("\nEnter ciphertext to decrypt: ")
            processed_text = prepare_text(text)
            result = decrypt(text, key)
            
            print("\n" + "-" * 50)
            print("DECRYPTION RESULTS:")
            print("-" * 50)
            print("Ciphertext: " + processed_text)
            print("Key: " + str(key))
            print("Decrypted message: " + result)
            print("-" * 50)

def test_example():
    print("TEST EXAMPLE:")
    print("-" * 50)
    text = "ATTACK AT ONCE"
    key = 4
    print("Text: " + text)
    print("Key: " + str(key))
    
    ciphertext = encrypt(text, key)
    print("Ciphertext: " + ciphertext)
    
    decrypted = decrypt(ciphertext, key)
    print("Decrypted text: " + decrypted)
    print("-" * 50)
    print()

if __name__ == "__main__":
    test_example()
    main_menu()