# TruckStreamer Pro: High-Performance Telemetry System

TruckStreamer Pro is a distributed microservices project designed to handle real-time telemetry data from virtual trucking fleets. This project demonstrates a hybrid architecture utilizing .NET 8 for robust API ingestion and Go (Golang) for high-concurrency data analytics.

## 🏗️ Architecture Overview
The system is split into two specialized services to maximize performance and scalability:

- **Ingestion API (.NET 8):** A high-availability REST API that acts as the entry point for telemetry packets (Speed, Fuel, Location) sent by trucks. It focuses on data validation and persistence.
- **Analytics Processor (Go):** A high-performance background worker that processes large batches of data in parallel. It utilizes Goroutines and WaitGroups to calculate real-time statistics with minimal CPU overhead.
- **Shared Storage:** A decoupled data layer that allows both services to communicate asynchronously via JSON streaming.

## 🛠️ Tech Stack
- **Backend:** C# / .NET 8 (Minimal APIs & Controllers)
- **Performance Layer:** Go 1.21 (Concurrency-first design)
- **Orchestration:** Docker & Docker Compose
- **Documentation:** Swagger UI (OpenAPI)

## 📁 Project Structure
```
├── IngestionAPI/           # .NET 8 Service - Data Ingestion
├── AnalyticsProcessor/     # Go Service - Parallel Processing Engine
├── Data/                   # Shared volume for inter-service communication
└── docker-compose.yml      # Infrastructure orchestration
```

## 🚀 Getting Started

### Prerequisites
- Docker Desktop installed.
- .NET 8 SDK (for local development).
- Go 1.21+ (for local development).

### Installation & Execution
Clone the repository:
```bash
git clone https://github.com/your-username/TruckStreamer-Pro.git
cd TruckStreamer-Pro
```
Run using Docker Compose:
```bash
docker-compose up --build
```

### Access the API Documentation
Once the containers are running, open your browser at:
```
http://localhost:5000/swagger
```

## 🧪 Testing the Flow
You can send a sample telemetry packet using curl or any API client (like Postman):
```bash
curl -X 'POST' \
  'http://localhost:5000/api/Telemetry' \
  -H 'Content-Type: application/json' \
  -d '{
  "truckId": "Peterbilt_389",
  "speed": 79.5,
  "fuelLevel": 0.92,
  "timestamp": "2026-05-12T10:00:00Z"
}'
```
The Analytics Processor will automatically pick up the data in the next cycle and log the processed statistics to the console.

## 📈 Key Technical Implementation
- **Concurrency in Go:** Uses sync.WaitGroup to ensure all parallel calculations are completed before finishing a processing cycle.
- **Scalability:** The decoupled nature allows the Analytics Processor to be scaled independently if the volume of trucks increases.
- **Clean Code:** Implements modern C# features like record types and async/await patterns for non-blocking I/O operations.

---
Developed by Keo Coelho de Almeida – Software Engineer specializing in high-performance backend systems.
