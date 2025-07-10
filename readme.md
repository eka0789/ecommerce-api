# ecommerce-api 🚀

[![Build Status](https://img.shields.io/github/actions/workflow/status/eka0789/ecommerce-api/ci.yml?branch=main)](https://github.com/eka0789/ecommerce-api/actions)
[![Go Version](https://img.shields.io/github/go-mod/go-version/eka0789/ecommerce-api)](https://golang.org/)
[![License](https://img.shields.io/github/license/eka0789/ecommerce-api)](LICENSE)

API backend modular untuk aplikasi e-commerce, dibangun dengan Go, Gin, MongoDB, Redis, RabbitMQ, dan Kafka. Proyek ini mengadopsi clean architecture, siap dikembangkan dan mudah diintegrasikan!

---

## ✨ Fitur Utama

- Clean architecture, modular & maintainable
- Gin sebagai web framework yang cepat
- MongoDB untuk penyimpanan data
- Redis sebagai cache layer
- RabbitMQ untuk async order processing
- Kafka untuk event streaming (notifikasi email/log order)
- Swagger API docs
- Docker Compose ready

---

## 🗂️ Struktur Folder

```
ecommerce-api/
├── cmd/
│   └── main.go
├── configs/
├── internal/
│   ├── domain/
│   ├── repository/
│   ├── usecase/
│   ├── delivery/
│   └── handler/
├── pkg/
├── docs/           # Swagger docs
├── scripts/
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## 🚀 Instalasi & Setup Lokal

### 1. Clone Repository

```bash
git clone https://github.com/eka0789/ecommerce-api.git
cd ecommerce-api
```

### 2. Jalankan dengan Docker Compose

```bash
docker-compose up --build
```

> Pastikan Docker & Docker Compose sudah terinstall.

### 3. Akses Swagger API Docs

Swagger UI tersedia di:  
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 🛠️ Konfigurasi Manual (Opsional)

1. Jalankan MongoDB, Redis, RabbitMQ, dan Kafka secara lokal.
2. Jalankan aplikasi:

```bash
go run cmd/main.go
```

---

## 📦 Endpoint Utama

| Method | Endpoint         | Deskripsi                |
|--------|------------------|-------------------------|
| POST   | `/orders`        | Membuat order baru      |
| GET    | `/orders`        | Ambil semua order       |
| GET    | `/orders/{id}`   | Ambil order by ID       |

---

## 📝 Contoh Payload `POST /orders`

```json
{
    "customer_id": "64a1f2c3e7b8a9d0e1f2a3b4",
    "items": [
        {
            "product_id": "64a1f2c3e7b8a9d0e1f2a3b5",
            "quantity": 2
        }
    ],
    "shipping_address": "Jl. Contoh No. 123, Jakarta"
}
```

---

## 🤝 Kontribusi

Kontribusi sangat terbuka!  
Silakan fork, buat branch, dan ajukan pull request.  
Lihat [CONTRIBUTING.md](CONTRIBUTING.md) untuk panduan lengkap.

---

## 📄 Lisensi

Proyek ini berlisensi [MIT](LICENSE).

---

> Made with ❤️ by [eka0789](https://github.com/eka0789)