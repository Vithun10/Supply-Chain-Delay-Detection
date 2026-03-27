🚚 Supply Chain Delay Detection and Monitoring System

A concurrent logistics analytics platform built with Go and React that processes shipment datasets, detects delivery delays, evaluates shipment risks, and visualizes operational insights through an interactive dashboard.

This system demonstrates Go concurrency, REST API design, and modern frontend visualization to simulate a real-world logistics monitoring platform.

📌 Project Overview

Supply chains generate large volumes of shipment data that need to be analyzed to identify delivery delays and operational risks.

This system processes shipment datasets using Go’s concurrency features (worker pools and goroutines) to efficiently detect delays and classify high-risk shipments.
The processed data is exposed through REST APIs and visualized via a React-based analytics dashboard.

The system also allows dynamic dataset uploads, enabling administrators to reprocess new shipment data and refresh analytics instantly.

🧠 Key Features

🚛 Shipment delay detection

⚠️ Risk classification for high-risk deliveries

📊 Logistics analytics dashboard

📦 Shipment tracking by ID

📂 Dynamic dataset upload

🔎 Filtering and pagination

👥 Role-based dashboards (Owner / Admin / Customer)

⚡ Concurrent dataset processing using Go worker pools

🏗 System Architecture
React Dashboard (Frontend)
        │
        ▼
REST API Requests
        │
        ▼
Go Backend (Gin Framework)
        │
        ▼
Concurrent Worker Pool
        │
        ▼
Shipment Dataset (CSV)

The backend processes datasets and exposes analytics APIs, while the frontend visualizes results through dashboards and tables.

⚙️ Technology Stack
Backend

Go (Golang)

Gin Web Framework

Goroutines and Worker Pools

CSV Data Processing

REST APIs

Frontend

React

TailwindCSS

Axios

React Router

Tools

Git & GitHub

Vite

Node.js

📂 Project Structure
supply-chain-delay-monitor
│
├── backend
│   ├── cmd
│   │   └── main.go
│   ├── controllers
│   │   ├── shipment_controller.go
│   │   ├── analytics_controller.go
│   │   ├── tracking_controller.go
│   │   └── dataset_controller.go
│   │
│   ├── services
│   ├── repository
│   ├── routes
│   ├── utils
│   │   └── csv_loader.go
│   │
│   └── models
│
├── frontend
│   ├── src
│   │   ├── components
│   │   ├── pages
│   │   ├── App.jsx
│   │   └── main.jsx
│   │
│   └── package.json
│
├── data
│   └── sample_dataset.csv
│
├── screenshots
│
├── README.md
└── .gitignore
🚀 Running the Project
1️⃣ Clone the Repository
git clone https://github.com/yourusername/supply-chain-delay-monitor.git
cd supply-chain-delay-monitor
2️⃣ Run the Backend
cd backend
go mod tidy
go run cmd/main.go

Backend runs on:

http://localhost:8080
3️⃣ Run the Frontend
cd frontend
npm install
npm run dev

Frontend runs on:

http://localhost:5173
📊 Available API Endpoints
Shipment APIs
GET /shipments
GET /shipments/:id

Returns shipment information and shipment details.

Delay APIs
GET /delays
GET /delays/high-risk

Returns delayed shipments and high-risk shipments.

Analytics APIs
GET /analytics/delay-rate
GET /analytics/top-delayed-routes
GET /analytics/carrier-performance
GET /analytics/avg-delivery-time

Provides analytics insights for the dashboard.

Tracking API
GET /track/:shipment_id

Allows customers to track shipment details.

Dataset Upload API
POST /upload-dataset

Allows administrators to upload a new shipment dataset and reprocess analytics.

📈 Dashboard Features

The React dashboard includes:

Owner Dashboard

KPI cards

delay analytics

shipment trends

logistics insights

Admin Dashboard

shipment management

delayed shipment monitoring

dataset upload

Customer Interface

shipment tracking

⚡ Concurrency in the Project

The backend processes shipment datasets using Go worker pools.

Example:

pool := concurrency.NewWorkerPool(cfg.WorkerCount)
processed := pool.Process(shipments)

This distributes dataset processing across multiple goroutines for efficient execution.

🧪 Error Handling

Go uses explicit error returns rather than exceptions.

Example:

shipments, err := utils.LoadCSV(cfg.DataFilePath)
if err != nil {
    utils.Logger.Fatalf("Failed to load dataset: %v", err)
}

This ensures robust handling of dataset errors and file processing issues.

🔮 Future Improvements

Possible enhancements include:

Real-time shipment updates using WebSockets

Database integration (PostgreSQL / MongoDB)

Authentication and role-based access control

Dataset versioning

Machine learning-based delay prediction

👨‍💻 Author

Developed by Vithun TR and Raju P as part of Go Programming Course Project (2026)
CHRIST (Deemed to be University), Central Campus, Bangalore

⭐ Acknowledgements

This project demonstrates practical applications of:

Go concurrency

REST API architecture

Logistics data analytics

Full-stack web development