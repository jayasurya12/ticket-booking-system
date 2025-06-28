# 🚆 IRCTC-Style Ticket Booking System (Golang + Kafka + Kubernetes)

A high-performance, microservice-based ticket booking system simulating IRCTC-like operations, using:

- ✅ Golang (Login, Booking, Payment services)
- ✅ Kafka for async communication
- ✅ Kubernetes for deployment (via Minikube)
- ✅ Optional: Load test with 10 million users

---

## 📁 Microservices

| Service   | Port  | Description                        |
|-----------|-------|------------------------------------|
| Login     | 8081  | Authenticates user, returns token |
| Booking   | 8082  | Accepts booking, pushes to Kafka  |
| Payment   | 8083  | Simulates payment transaction     |

---

## 🛠️ Tech Stack

- Golang 1.21
- Apache Kafka
- Kubernetes (Minikube)
- Docker
- Shell (for load script)

---

## 🚀 Local Setup Using Minikube

### 1️⃣ Install Minikube (Linux)
```bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube


2️⃣ Start Minikube (Low Resource)
minikube start --cpus=2 --memory=4096 --driver=docker


3️⃣ Use Minikube’s Docker
eval $(minikube docker-env)


🧱 Build Docker Images
# Login
cd login-service
docker build -t login-service:latest .

# Booking
cd ../booking-service
docker build -t booking-service:latest .

# Payment
cd ../payment-service
docker build -t payment-service:latest .

☸️ Deploy Kafka in Minikube
Simple Kafka Setup (recommended YAML)
Use the provided k8s/kafka-single.yaml:

kubectl apply -f k8s/kafka-single.yaml
☸️ Deploy Services
kubectl apply -f k8s/login-deployment.yaml
kubectl apply -f k8s/booking-deployment.yaml
kubectl apply -f k8s/payment-deployment.yaml
Check:

kubectl get pods
kubectl get svc
🌐 Access Services
To expose via browser:

minikube service login-service
minikube service booking-service
minikube service payment-service
🧪 Load Test Script
File: simulator.sh

chmod +x simulator.sh
./simulator.sh
It simulates 10 million users performing:

Login → Token

Booking → Kafka → Payment

📦 Folder Structure
project-root/
├── login-service/
├── booking-service/
├── payment-service/
├── k8s/
│   ├── login-deployment.yaml
│   ├── booking-deployment.yaml
│   ├── payment-deployment.yaml
│   └── kafka-single.yaml
├── simulator.sh
└── README.md
📌 Notes
Kafka topic: bookings

Booking service runs Kafka producer + consumer

Payment service is called via HTTP from consumer

All logs are printed to console (no DB used yet)

✨ To Do / Extensions
Add persistent database (ClickHouse / Mongo / Postgres)

JWT validation in Login & Booking

Circuit breakers (using go-resilience)

Rate limiting

Helm Charts

👨‍💻 Author
Made with ❤️ by Surya

yaml

---

Would you like me to generate this README as a downloadable file, or bundle it into a GitHub-ready [starter repo structure](f)?