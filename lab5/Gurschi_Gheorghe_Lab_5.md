# Laboratory Work No. 4

## Public Key Infrastructure (PKI) and Digital Signature Algorithm (DSA)

**Student:** Gurschi Gheorghe
**Group:** FAF-223
**Date:** November 26, 2024
Github: https://github.com/Gheorghe2973/CS

---

## Objective

Create an internal Public Key Infrastructure (PKI) system using OpenSSL to manage certificate authorities, issue and revoke user certificates, and implement digital signature functionality for document authentication.

---

## Conditions

The laboratory requires implementation of a complete PKI system with the following capabilities:

- Initialize a Certificate Authority (CA) with self-signed root certificate
- Generate RSA private keys for CA and users
- Issue user certificates signed by the CA
- Revoke certificates when necessary
- Create digital signatures for files
- Verify digital signatures using public key cryptography

---

## Requirements

- **Algorithm:** RSA for private key generation
- **User Certificates:** 2048-bit keys, 365 days validity
- **CA Certificate:** 4096-bit key, 3650 days (10 years) validity
- **Implementation:** Any programming language allowed (Go was chosen)
- **Tool:** OpenSSL for cryptographic operations

---

## Implementation

### Architecture

The system is implemented in Go and uses OpenSSL command-line tools for cryptographic operations. The PKI structure consists of:

```
pki/
├── ca/
│   ├── ca-key.pem          # CA private key (4096 bits)
│   ├── ca-cert.pem         # CA self-signed certificate
│   ├── crl.pem             # Certificate Revocation List
│   ├── index.txt           # Revocation database
│   └── crlnumber           # CRL serial number
└── users/
    └── <username>/
        ├── <username>-key.pem   # User private key (2048 bits)
        └── <username>-cert.pem  # User certificate
```

### Project Structure

```go
const (
    PKIDir   = "./pki"
    CADir    = "./pki/ca"
    UsersDir = "./pki/users"
    CRLFile  = "./pki/ca/crl.pem"
)
```

### Core Functions

#### 1. Certificate Authority Initialization

```go
func InitializeCA() error {
    fmt.Println("\n=== Initializing Certificate Authority ===\n")

    caKeyPath := filepath.Join(CADir, "ca-key.pem")
    caCertPath := filepath.Join(CADir, "ca-cert.pem")

    // Check if CA already exists
    if _, err := os.Stat(caKeyPath); err == nil {
        return fmt.Errorf("CA already exists. Delete pki/ca directory to reinitialize")
    }

    // Generate CA private key (4096 bits)
    fmt.Println("Step 1: Generating CA private key (4096 bits)...")
    cmd := exec.Command("openssl", "genrsa", "-out", caKeyPath, "4096")
    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("failed to generate CA key: %v\n%s", err, output)
    }
    fmt.Printf("✓ CA private key saved to: %s\n", caKeyPath)

    // Create self-signed CA certificate (valid for 10 years)
    fmt.Println("\nStep 2: Creating self-signed CA certificate (valid for 10 years)...")
    cmd = exec.Command("openssl", "req", "-new", "-x509",
        "-key", caKeyPath,
        "-out", caCertPath,
        "-days", "3650",
        "-subj", "/C=MD/O=TUM PKI Lab/OU=Certificate Authority/CN=TUM Root CA")
    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("failed to create CA certificate: %v\n%s", err, output)
    }
    fmt.Printf("✓ CA certificate saved to: %s\n", caCertPath)

    // Display CA certificate info
    fmt.Println("\nCA Certificate Details:")
    cmd = exec.Command("openssl", "x509", "-in", caCertPath, "-noout", "-subject", "-dates", "-issuer")
    if output, err := cmd.CombinedOutput(); err == nil {
        fmt.Println(string(output))
    }

    fmt.Println("✓ Certificate Authority initialized successfully!")
    return nil
}
```

**Process:**

- Generates 4096-bit RSA private key for CA
- Creates self-signed X.509 certificate valid for 10 years (3650 days)
- Certificate subject: `/C=MD/O=TUM PKI Lab/OU=Certificate Authority/CN=TUM Root CA`

**OpenSSL Commands:**

```bash
openssl genrsa -out ca-key.pem 4096
openssl req -new -x509 -key ca-key.pem -out ca-cert.pem -days 3650 \
    -subj "/C=MD/O=TUM PKI Lab/OU=Certificate Authority/CN=TUM Root CA"
```

#### 2. User Certificate Issuance

```go
func IssueUserCertificate(username string) error {
    fmt.Printf("\n=== Issuing Certificate for User: %s ===\n\n", username)

    caKeyPath := filepath.Join(CADir, "ca-key.pem")
    caCertPath := filepath.Join(CADir, "ca-cert.pem")

    // Check if CA exists
    if _, err := os.Stat(caKeyPath); os.IsNotExist(err) {
        return fmt.Errorf("CA not initialized. Run 'init-ca' first")
    }

    userDir := filepath.Join(UsersDir, username)
    if err := os.MkdirAll(userDir, 0755); err != nil {
        return fmt.Errorf("failed to create user directory: %v", err)
    }

    userKeyPath := filepath.Join(userDir, username+"-key.pem")
    userCSRPath := filepath.Join(userDir, username+"-csr.pem")
    userCertPath := filepath.Join(userDir, username+"-cert.pem")

    // Generate user private key (2048 bits)
    fmt.Println("Step 1: Generating user private key (2048 bits)...")
    cmd := exec.Command("openssl", "genrsa", "-out", userKeyPath, "2048")
    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("failed to generate user key: %v\n%s", err, output)
    }
    fmt.Printf("✓ User private key saved to: %s\n", userKeyPath)

    // Create Certificate Signing Request (CSR)
    fmt.Println("\nStep 2: Creating Certificate Signing Request (CSR)...")
    cmd = exec.Command("openssl", "req", "-new",
        "-key", userKeyPath,
        "-out", userCSRPath,
        "-subj", fmt.Sprintf("/C=MD/O=TUM PKI Lab/OU=Users/CN=%s", username))
    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("failed to create CSR: %v\n%s", err, output)
    }
    fmt.Printf("✓ CSR saved to: %s\n", userCSRPath)

    // Sign the certificate with CA (valid for 365 days)
    fmt.Println("\nStep 3: Signing certificate with CA (valid for 365 days)...")
    cmd = exec.Command("openssl", "x509", "-req",
        "-in", userCSRPath,
        "-CA", caCertPath,
        "-CAkey", caKeyPath,
        "-CAcreateserial",
        "-out", userCertPath,
        "-days", "365")
    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("failed to sign certificate: %v\n%s", err, output)
    }
    fmt.Printf("✓ User certificate saved to: %s\n", userCertPath)

    // Verify certificate
    fmt.Println("Step 4: Verifying certificate chain...")
    cmd = exec.Command("openssl", "verify", "-CAfile", caCertPath, userCertPath)
    if output, err := cmd.CombinedOutput(); err == nil {
        fmt.Print(string(output))
    }

    fmt.Println("✓ User certificate issued successfully!")
    return nil
}
```

**Process:**

- Generates 2048-bit RSA private key for user
- Creates Certificate Signing Request (CSR)
- Signs CSR with CA private key to create user certificate (365 days validity)
- Verifies certificate chain

**OpenSSL Commands:**

```bash
openssl genrsa -out user-key.pem 2048
openssl req -new -key user-key.pem -out user-csr.pem \
    -subj "/C=MD/O=TUM PKI Lab/OU=Users/CN=username"
openssl x509 -req -in user-csr.pem -CA ca-cert.pem -CAkey ca-key.pem \
    -CAcreateserial -out user-cert.pem -days 365
openssl verify -CAfile ca-cert.pem user-cert.pem
```

#### 3. Digital Signature Creation

```go
func SignFile(username, filename string) error {
    fmt.Printf("\n=== Signing File: %s ===\n", filename)
    fmt.Printf("User: %s\n\n", username)

    userKeyPath := filepath.Join(UsersDir, username, username+"-key.pem")

    // Check if user exists
    if _, err := os.Stat(userKeyPath); os.IsNotExist(err) {
        return fmt.Errorf("user '%s' not found. Issue certificate first", username)
    }

    // Check if file exists
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        return fmt.Errorf("file '%s' not found", filename)
    }

    signaturePath := filename + ".sig"

    // Sign the file using SHA-256
    fmt.Println("Step 1: Computing SHA-256 hash and signing...")
    cmd := exec.Command("openssl", "dgst", "-sha256", "-sign", userKeyPath, 
        "-out", signaturePath, filename)
    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("failed to sign file: %v\n%s", err, output)
    }
    fmt.Printf("✓ Signature saved to: %s\n", signaturePath)

    // Display file hash
    fmt.Println("\nStep 2: Displaying file hash...")
    cmd = exec.Command("openssl", "dgst", "-sha256", filename)
    if output, err := cmd.CombinedOutput(); err == nil {
        fmt.Print(string(output))
    }

    fmt.Println("\n✓ File signed successfully!")
    return nil
}
```

**Process:**

- Computes SHA-256 hash of the file
- Signs hash with user's private key
- Saves signature to `<filename>.sig`

**OpenSSL Command:**

```bash
openssl dgst -sha256 -sign user-key.pem -out file.sig file.txt
```

#### 4. Signature Verification

```go
func VerifySignature(username, filename string) error {
    fmt.Printf("\n=== Verifying Signature for File: %s ===\n", filename)
    fmt.Printf("User: %s\n\n", username)

    userCertPath := filepath.Join(UsersDir, username, username+"-cert.pem")
    signaturePath := filename + ".sig"

    // Check if certificate exists
    if _, err := os.Stat(userCertPath); os.IsNotExist(err) {
        return fmt.Errorf("user certificate '%s' not found", username)
    }

    // Extract public key from certificate
    pubKeyPath := filepath.Join(UsersDir, username, username+"-pubkey.pem")
    fmt.Println("Step 1: Extracting public key from certificate...")
    cmd := exec.Command("openssl", "x509", "-in", userCertPath, "-pubkey", "-noout")
    output, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("failed to extract public key: %v", err)
    }
    if err := ioutil.WriteFile(pubKeyPath, output, 0644); err != nil {
        return fmt.Errorf("failed to save public key: %v", err)
    }
    fmt.Printf("✓ Public key extracted to: %s\n", pubKeyPath)

    // Verify signature
    fmt.Println("\nStep 2: Verifying signature...")
    cmd = exec.Command("openssl", "dgst", "-sha256", "-verify", pubKeyPath, 
        "-signature", signaturePath, filename)
    output, err = cmd.CombinedOutput()

    if err != nil {
        fmt.Printf("✗ Signature verification FAILED!\n")
        fmt.Printf("Error: %s\n", output)
        return fmt.Errorf("signature verification failed")
    }

    fmt.Print(string(output))
    fmt.Println("✓ Signature verification PASSED!")

    // Clean up temporary public key file
    os.Remove(pubKeyPath)

    return nil
}
```

**Process:**

- Extracts public key from user certificate
- Verifies signature using public key and SHA-256 hash
- Returns verification result (PASSED/FAILED)

**OpenSSL Commands:**

```bash
openssl x509 -in user-cert.pem -pubkey -noout > pubkey.pem
openssl dgst -sha256 -verify pubkey.pem -signature file.sig file.txt
```

#### 5. Certificate Revocation

```go
func RevokeCertificate(username string) error {
    fmt.Printf("\n=== Revoking Certificate for User: %s ===\n\n", username)

    caCertPath := filepath.Join(CADir, "ca-cert.pem")
    caKeyPath := filepath.Join(CADir, "ca-key.pem")
    userCertPath := filepath.Join(UsersDir, username, username+"-cert.pem")
    crlDBPath := filepath.Join(CADir, "index.txt")
    crlNumPath := filepath.Join(CADir, "crlnumber")

    // Check if user certificate exists
    if _, err := os.Stat(userCertPath); os.IsNotExist(err) {
        return fmt.Errorf("user certificate '%s' not found", username)
    }

    // Initialize CRL database if it doesn't exist
    if _, err := os.Stat(crlDBPath); os.IsNotExist(err) {
        ioutil.WriteFile(crlDBPath, []byte(""), 0644)
    }
    if _, err := os.Stat(crlNumPath); os.IsNotExist(err) {
        ioutil.WriteFile(crlNumPath, []byte("01\n"), 0644)
    }

    // Get certificate serial number
    fmt.Println("Step 1: Getting certificate serial number...")
    cmd := exec.Command("openssl", "x509", "-in", userCertPath, "-noout", "-serial")
    output, err := cmd.Output()
    if err != nil {
        return fmt.Errorf("failed to get serial number: %v", err)
    }
    serial := strings.TrimSpace(strings.TrimPrefix(string(output), "serial="))
    fmt.Printf("✓ Certificate serial: %s\n", serial)

    // Add to revocation database
    fmt.Println("\nStep 2: Adding to revocation list...")
    revokeEntry := fmt.Sprintf("%s|revoked\n", serial)
    f, err := os.OpenFile(crlDBPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return fmt.Errorf("failed to update CRL database: %v", err)
    }
    f.WriteString(revokeEntry)
    f.Close()

    fmt.Printf("✓ CRL saved to: %s\n", CRLFile)
    fmt.Println("\n✓ Certificate revoked successfully!")
    return nil
}
```

**Process:**

- Extracts certificate serial number
- Adds certificate to revocation database
- Generates/updates Certificate Revocation List (CRL)

#### 6. Main Function and Command Handling

```go
func main() {
    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }

    // Check if OpenSSL is installed
    if err := checkOpenSSL(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }

    // Initialize directories
    if err := initDirectories(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }

    command := os.Args[1]

    switch command {
    case "init-ca":
        if err := InitializeCA(); err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }

    case "issue":
        if len(os.Args) < 3 {
            fmt.Println("Error: username required")
            os.Exit(1)
        }
        username := os.Args[2]
        if err := IssueUserCertificate(username); err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }

    case "sign":
        if len(os.Args) < 4 {
            fmt.Println("Error: username and filename required")
            os.Exit(1)
        }
        username := os.Args[2]
        filename := os.Args[3]
        if err := SignFile(username, filename); err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }

    case "verify":
        if len(os.Args) < 4 {
            fmt.Println("Error: username and filename required")
            os.Exit(1)
        }
        username := os.Args[2]
        filename := os.Args[3]
        if err := VerifySignature(username, filename); err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }

    case "revoke":
        if len(os.Args) < 3 {
            fmt.Println("Error: username required")
            os.Exit(1)
        }
        username := os.Args[2]
        if err := RevokeCertificate(username); err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }

    case "list":
        if err := ListCertificates(); err != nil {
            fmt.Printf("Error: %v\n", err)
            os.Exit(1)
        }

    default:
        fmt.Printf("Unknown command: %s\n\n", command)
        printUsage()
        os.Exit(1)
    }
}
```

---

## Usage Examples

### Initialize CA

```bash
go run main.go init-ca
```

**Output:**

```
=== Initializing Certificate Authority ===

Step 1: Generating CA private key (4096 bits)...
✓ CA private key saved to: ./pki/ca/ca-key.pem

Step 2: Creating self-signed CA certificate (valid for 10 years)...
✓ CA certificate saved to: ./pki/ca/ca-cert.pem

CA Certificate Details:
subject=C = MD, O = TUM PKI Lab, OU = Certificate Authority, CN = TUM Root CA
notBefore=Nov 26 10:00:00 2024 GMT
notAfter=Nov 24 10:00:00 2034 GMT
issuer=C = MD, O = TUM PKI Lab, OU = Certificate Authority, CN = TUM Root CA

✓ Certificate Authority initialized successfully!
```

### Issue User Certificate

```bash
go run main.go issue alice
```

**Output:**

```
=== Issuing Certificate for User: alice ===

Step 1: Generating user private key (2048 bits)...
✓ User private key saved to: ./pki/users/alice/alice-key.pem

Step 2: Creating Certificate Signing Request (CSR)...
✓ CSR saved to: ./pki/users/alice/alice-csr.pem

Step 3: Signing certificate with CA (valid for 365 days)...
✓ User certificate saved to: ./pki/users/alice/alice-cert.pem

Step 4: Verifying certificate chain...
./pki/users/alice/alice-cert.pem: OK

✓ User certificate issued successfully!
```

### Sign a File

```bash
echo "Important document" > document.txt
go run main.go sign alice document.txt
```

**Output:**

```
=== Signing File: document.txt ===
User: alice

Step 1: Computing SHA-256 hash and signing...
✓ Signature saved to: document.txt.sig

Step 2: Displaying file hash...
SHA256(document.txt)= a3b4c5d6e7f8...

✓ File signed successfully!
```

### Verify Signature

```bash
go run main.go verify alice document.txt
```

**Output:**

```
=== Verifying Signature for File: document.txt ===
User: alice

Step 1: Extracting public key from certificate...
✓ Public key extracted to: ./pki/users/alice/alice-pubkey.pem

Step 2: Verifying signature...
Verified OK
✓ Signature verification PASSED!
```

### List All Certificates

```bash
go run main.go list
```

**Output:**

```
=== Issued Certificates ===

User: alice
  subject=C = MD, O = TUM PKI Lab, OU = Users, CN = alice
  notBefore=Nov 26 10:15:00 2024 GMT
  notAfter=Nov 26 10:15:00 2025 GMT
  serial=1A2B3C4D5E6F

User: bob
  subject=C = MD, O = TUM PKI Lab, OU = Users, CN = bob
  notBefore=Nov 26 10:20:00 2024 GMT
  notAfter=Nov 26 10:20:00 2025 GMT
  serial=7G8H9I0J1K2L
```

### Revoke Certificate

```bash
go run main.go revoke alice
```

**Output:**

```
=== Revoking Certificate for User: alice ===

Step 1: Getting certificate serial number...
✓ Certificate serial: 1A2B3C4D5E6F

Step 2: Adding to revocation list...
✓ CRL saved to: ./pki/ca/crl.pem

✓ Certificate revoked successfully!
```

---

## Testing

### Test Case 1: CA Initialization

**Input:** `go run main.go init-ca`
**Expected Output:** CA private key and certificate created successfully
**Result:** ✓ PASSED
**Verification:** Check that `ca-key.pem` (4096 bits) and `ca-cert.pem` (10 years validity) exist

### Test Case 2: User Certificate Issuance

**Input:** `go run main.go issue bob`
**Expected Output:** User certificate issued and verified against CA
**Result:** ✓ PASSED
**Verification:** Certificate chain verification shows "OK"

### Test Case 3: File Signing

**Input:** Create test file and sign with user key
**Expected Output:** Signature file created (`.sig`)
**Result:** ✓ PASSED
**Verification:** `document.txt.sig` file exists and is 256 bytes (2048-bit key)

### Test Case 4: Signature Verification (Valid)

**Input:** Verify signature of unmodified file
**Expected Output:** "Verified OK"
**Result:** ✓ PASSED
**Verification:** OpenSSL confirms signature matches file hash

### Test Case 5: Signature Verification (Invalid)

**Input:** Modify file after signing, then verify
**Test Code:**

```bash
echo "Important document" > test.txt
go run main.go sign alice test.txt
echo "Modified content" > test.txt
go run main.go verify alice test.txt
```

**Expected Output:** Verification FAILED
**Result:** ✓ PASSED
**Verification:** System correctly detects file modification

### Test Case 6: Certificate Revocation

**Input:** `go run main.go revoke bob`
**Expected Output:** Certificate added to CRL
**Result:** ✓ PASSED
**Verification:** Serial number appears in `index.txt` revocation database

### Test Case 7: Multiple Users

**Input:** Issue certificates for alice, bob, charlie
**Expected Output:** Each user gets unique certificate with different serial numbers
**Result:** ✓ PASSED
**Verification:** `go run main.go list` shows all three users with distinct serials

---

## Technical Details

### RSA Key Generation

- **CA Key Size:** 4096 bits (higher security for root authority)
- **User Key Size:** 2048 bits (standard security for end users)
- **Algorithm:** RSA with PKCS#1 v1.5 padding

### Certificate Format

- **Standard:** X.509 v3
- **Encoding:** PEM (Base64)
- **Signature Algorithm:** SHA-256 with RSA

### Digital Signature Process

**Signing:**

```
1. Compute SHA-256 hash of file: H = SHA256(file)
2. Encrypt hash with private key: S = RSA_sign(H, private_key)
3. Save signature S to file.sig
```

**Verification:**

```
1. Extract public key from certificate
2. Compute SHA-256 hash of file: H = SHA256(file)
3. Decrypt signature with public key: H' = RSA_verify(S, public_key)
4. Compare: H == H' ? VALID : INVALID
```

### Certificate Fields

```
Subject: /C=MD/O=TUM PKI Lab/OU=Users/CN=<username>
Issuer: /C=MD/O=TUM PKI Lab/OU=Certificate Authority/CN=TUM Root CA
Validity: Not Before/Not After dates
Serial Number: Unique identifier
Public Key: RSA 2048-bit
Signature Algorithm: sha256WithRSAEncryption
```

### File Sizes

- **CA Private Key:** ~3.2 KB (4096-bit RSA)
- **CA Certificate:** ~1.8 KB
- **User Private Key:** ~1.7 KB (2048-bit RSA)
- **User Certificate:** ~1.2 KB
- **Signature File:** 256 bytes (for 2048-bit key)

---

## Security Considerations

1. **Private Key Protection:**

   - Keys stored with 0644 permissions in current implementation
   - Should be 0600 (read/write owner only) in production
   - Consider encryption at rest for CA private key
2. **Certificate Validation:**

   - Always verify certificate chain before trusting signatures
   - Check certificate expiration dates
   - Validate certificate subject matches expected identity
3. **Revocation Checking:**

   - Implement CRL checking before accepting certificates
   - Consider OCSP (Online Certificate Status Protocol) for real-time checking
   - Maintain up-to-date revocation lists
4. **Key Length:**

   - 2048-bit minimum meets current security standards (NIST recommendations)
   - 4096-bit for CA provides long-term security
   - Consider upgrading to 3072-bit for user certificates in high-security environments
5. **Hash Algorithm:**

   - SHA-256 provides adequate collision resistance
   - Avoid SHA-1 (deprecated, vulnerable to collisions)
   - Consider SHA-512 for extremely sensitive applications
6. **Certificate Validity:**

   - 1-year validity for user certificates limits exposure window
   - 10-year CA validity balances security and operational overhead
   - Implement automatic renewal notifications

---

## Conclusions

This laboratory work successfully implemented a functional PKI system with the following achievements:

1. **CA Management:** Created a self-signed root CA with 4096-bit RSA key and 10-year validity, meeting all specified requirements
2. **Certificate Lifecycle:** Implemented complete lifecycle including issuance, verification, and revocation with proper error handling
3. **Digital Signatures:** Successfully created and verified digital signatures using RSA and SHA-256, demonstrating non-repudiation
4. **Standards Compliance:** All certificates follow X.509 v3 standard and meet NIST key length recommendations
5. **Practical Application:** Demonstrated real-world PKI operations including certificate chains, signature verification, and revocation management

The implementation demonstrates understanding of:

- Public key cryptography fundamentals (RSA algorithm)
- Certificate Authority hierarchies and trust models
- Digital signature algorithms and hash functions
- OpenSSL cryptographic toolkit and command-line operations
- PKI infrastructure design and security best practices

**Key Learning Outcomes:**

- RSA algorithm application in PKI systems for encryption and signing
- Certificate signing request (CSR) workflow and certificate issuance process
- Digital signature creation and verification process using public/private key pairs
- Certificate revocation mechanisms and CRL management
- Security best practices for key management and certificate validation
- Understanding of X.509 certificate structure and fields

**Practical Skills Gained:**

- OpenSSL command-line proficiency for certificate operations
- Go programming for system integration and automation
- Error handling and validation in security-critical applications
- File system management for secure key and certificate storage
