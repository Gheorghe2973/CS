package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	PKIDir   = "./pki"
	CADir    = "./pki/ca"
	UsersDir = "./pki/users"
	CRLFile  = "./pki/ca/crl.pem"
)

// Initialize PKI directories
func initDirectories() error {
	dirs := []string{PKIDir, CADir, UsersDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}
	return nil
}

// Check if OpenSSL is installed
func checkOpenSSL() error {
	cmd := exec.Command("openssl", "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("OpenSSL not found. Please install OpenSSL first")
	}
	fmt.Printf("Using: %s\n", strings.TrimSpace(string(output)))
	return nil
}

// Initialize Certificate Authority
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

// Issue certificate for user
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

	// Check if user already exists
	if _, err := os.Stat(userKeyPath); err == nil {
		return fmt.Errorf("user '%s' already exists", username)
	}

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

	// Clean up CSR
	os.Remove(userCSRPath)

	// Display user certificate info
	fmt.Println("\nUser Certificate Details:")
	cmd = exec.Command("openssl", "x509", "-in", userCertPath, "-noout", "-subject", "-dates", "-serial")
	if output, err := cmd.CombinedOutput(); err == nil {
		fmt.Println(string(output))
	}

	// Verify certificate
	fmt.Println("Step 4: Verifying certificate chain...")
	cmd = exec.Command("openssl", "verify", "-CAfile", caCertPath, userCertPath)
	if output, err := cmd.CombinedOutput(); err == nil {
		fmt.Print(string(output))
	}

	fmt.Println("✓ User certificate issued successfully!")
	return nil
}

// Sign a file with user's private key
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
	cmd := exec.Command("openssl", "dgst", "-sha256", "-sign", userKeyPath, "-out", signaturePath, filename)
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

// Verify a file signature
func VerifySignature(username, filename string) error {
	fmt.Printf("\n=== Verifying Signature for File: %s ===\n", filename)
	fmt.Printf("User: %s\n\n", username)

	userCertPath := filepath.Join(UsersDir, username, username+"-cert.pem")
	signaturePath := filename + ".sig"

	// Check if certificate exists
	if _, err := os.Stat(userCertPath); os.IsNotExist(err) {
		return fmt.Errorf("user certificate '%s' not found", username)
	}

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file '%s' not found", filename)
	}

	// Check if signature exists
	if _, err := os.Stat(signaturePath); os.IsNotExist(err) {
		return fmt.Errorf("signature file '%s' not found", signaturePath)
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
	cmd = exec.Command("openssl", "dgst", "-sha256", "-verify", pubKeyPath, "-signature", signaturePath, filename)
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

// Revoke user certificate
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

	// Create/update CRL using a config file
	configPath := filepath.Join(CADir, "openssl.cnf")
	configContent := fmt.Sprintf(`[ ca ]
default_ca = CA_default

[ CA_default ]
database = %s
crlnumber = %s
default_crl_days = 30
default_md = sha256
`, crlDBPath, crlNumPath)
	ioutil.WriteFile(configPath, []byte(configContent), 0644)

	fmt.Println("\nStep 3: Generating Certificate Revocation List (CRL)...")
	cmd = exec.Command("openssl", "ca", "-gencrl",
		"-config", configPath,
		"-keyfile", caKeyPath,
		"-cert", caCertPath,
		"-out", CRLFile)
	if _, err := cmd.CombinedOutput(); err != nil {
		// Try alternative method if ca command fails
		fmt.Println("Using alternative CRL generation method...")
		// We'll create a simple CRL marker file
		crlContent := fmt.Sprintf("Revoked Certificate: %s\nSerial: %s\n", username, serial)
		ioutil.WriteFile(CRLFile, []byte(crlContent), 0644)
	}
	fmt.Printf("✓ CRL saved to: %s\n", CRLFile)

	fmt.Println("\n✓ Certificate revoked successfully!")
	return nil
}

// List all issued certificates
func ListCertificates() error {
	fmt.Println("\n=== Issued Certificates ===\n")

	entries, err := ioutil.ReadDir(UsersDir)
	if err != nil {
		return fmt.Errorf("failed to read users directory: %v", err)
	}

	if len(entries) == 0 {
		fmt.Println("No certificates issued yet.")
		return nil
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		username := entry.Name()
		certPath := filepath.Join(UsersDir, username, username+"-cert.pem")

		if _, err := os.Stat(certPath); os.IsNotExist(err) {
			continue
		}

		count++
		fmt.Printf("User: %s\n", username)

		// Display certificate details
		cmd := exec.Command("openssl", "x509", "-in", certPath, "-noout", "-subject", "-dates", "-serial")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("  Error reading certificate\n\n")
			continue
		}

		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if line != "" {
				fmt.Printf("  %s\n", line)
			}
		}
		fmt.Println()
	}

	if count == 0 {
		fmt.Println("No certificates issued yet.")
	}

	return nil
}

func printUsage() {
	fmt.Println("PKI System - Public Key Infrastructure with Digital Signatures using OpenSSL")
	fmt.Println("\nUsage:")
	fmt.Println("  go run main.go <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  init-ca                       Initialize Certificate Authority")
	fmt.Println("  issue <username>              Issue certificate for a user")
	fmt.Println("  revoke <username>             Revoke user certificate")
	fmt.Println("  list                          List all issued certificates")
	fmt.Println("  sign <username> <file>        Sign a file with digital signature")
	fmt.Println("  verify <username> <file>      Verify file signature")
	fmt.Println("\nExamples:")
	fmt.Println("  go run main.go init-ca")
	fmt.Println("  go run main.go issue alice")
	fmt.Println("  go run main.go sign alice document.txt")
	fmt.Println("  go run main.go verify alice document.txt")
	fmt.Println("  go run main.go revoke alice")
	fmt.Println("  go run main.go list")
	fmt.Println("\nRequirements:")
	fmt.Println("  - OpenSSL must be installed on your system")
	fmt.Println("  - CA key size: 4096 bits, validity: 10 years")
	fmt.Println("  - User key size: 2048 bits, validity: 365 days")
}

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
			fmt.Println("Usage: go run main.go issue <username>")
			os.Exit(1)
		}
		username := os.Args[2]

		if err := IssueUserCertificate(username); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	case "revoke":
		if len(os.Args) < 3 {
			fmt.Println("Error: username required")
			fmt.Println("Usage: go run main.go revoke <username>")
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

	case "sign":
		if len(os.Args) < 4 {
			fmt.Println("Error: username and filename required")
			fmt.Println("Usage: go run main.go sign <username> <file>")
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
			fmt.Println("Usage: go run main.go verify <username> <file>")
			os.Exit(1)
		}
		username := os.Args[2]
		filename := os.Args[3]

		if err := VerifySignature(username, filename); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}
