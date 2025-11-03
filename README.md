# Go Crypto Service 🔐

A **production-ready cryptographic microservice** built in **Go**, designed to handle encryption, decryption, and key management securely via **REST** and **gRPC** APIs.

---

## 🚀 Features
- **AES Encryption/Decryption** for symmetric cryptography.
- **RSA Public/Private Key** support for asymmetric encryption.
- **REST API** (port `8080`) and **gRPC API** (port `50051`).
- **Dependency Injection** using [Google Wire](https://github.com/google/wire).
- **Environment-based Configuration** (`dev`, `prod`).
- **Key Management** from `./keys` folder.
- **Error handling**, **logging**, and **graceful shutdown**.

---

## 🏗️ Project Structure
```
go-crypto-service/
├── README.md
├── go.mod
├── go.sum
├── Dockerfile
├── Makefile
├── scripts/
│ └── genkeys.sh
├── proto/
│ └── crypto.proto
├── cmd/
│ └── main.go
├── internal/
│ ├── config/
│ │ └── config.go
│ ├── crypto/
│ │ ├── aes.go
│ │ └── rsa.go
│ ├── handler/
│ │ ├── rest_handler.go
│ │ └── grpc_handler.go
│ ├── keys/
│ │ └── key_manager.go
│ ├── service/
│ │ └── crypto_service.go
│ ├── wire/
│ │ └── wire.go
│ └── util/
│ └── logger.go
├──keys/
  ├── aes.key        # Demo key, DO NOT USE IN PRODUCTION
  ├── private.pem    # Dummy RSA key for testing
  └── public.pem

```

---

## ⚙️ Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/<your-username>/go-crypto-service.git
cd go-crypto-service
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Generate Keys
#### AES Key
```bash
mkdir keys
echo "your-random-32-byte-key" > keys/aes.key
```
#### RSA Keys
```bash
openssl genpkey -algorithm RSA -out keys/private.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -in keys/private.pem -pubout -out keys/public.pem
```

### 4. Run the Application
```bash
go run cmd/main.go
```
Expected output:
```
2025/11/03 11:39:34 gRPC server listening on :50051
2025/11/03 11:39:34 REST server listening on :8080
```

---

## 🧠 Example API Usage
### REST Endpoint: `/encrypt`
```bash
curl -X POST http://localhost:8080/encrypt \
  -H "Content-Type: application/json" \
  -d '{"data": "Hello Crypto"}'
```

### gRPC Client (example snippet)
```go
conn, _ := grpc.Dial(":50051", grpc.WithInsecure())
client := pb.NewCryptoServiceClient(conn)
res, _ := client.Encrypt(context.Background(), &pb.EncryptRequest{Data: "Hello"})
fmt.Println(res.Encrypted)
```

---

## 🧰 Tech Stack
- **Go 1.22+**
- **Wire** (DI)
- **gRPC**
- **net/http** for REST
- **crypto/aes**, **crypto/rsa**, **x509** for cryptography

---

## ⚠️ Security Note
This project is for **learning purposes only**. Never commit or push **real keys** (private.pem, aes.key) to public repositories in production environments.

To ignore keys:
```bash
echo "keys/" >> .gitignore
```

---

## 🧩 Future Enhancements
- JWT signing & verification
- Hashing utilities (SHA256, bcrypt)
- Key rotation system
- Dockerfile and CI/CD setup

---

## 🧑‍💻 Author
**Devendra Pratap**  
Software Developer — Go 

[![GitHub](https://img.shields.io/badge/GitHub-DevendraPratap-black?logo=github)](https://github.com/devendrapratap307/)  
[![LinkedIn](https://img.shields.io/badge/LinkedIn-DevendraPratap-blue?logo=linkedin)](https://linkedin.com/in/devendrapratap307/) 
---
