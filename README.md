# 🧠 CEILO - Backend

Mental Health Web Platform using **Golang + Gin + GORM**.

---

## 🚀 How to Run this Project

### **PREPARATION**

#### 1. Install Dependencies
- **Golang** (version 1.21 or later)
- **Node.js** (version 18 or later)
- **PostgreSQL** (version 14 atau later)
- **Git**
- **VSCode** (recommended)

#### 2. Make PostgreSQL Database
```sql
CREATE DATABASE ceilo_db;
```

---

## 🔧 BACKEND SETUP (Golang)

### **Step 1: Open Backend Folder**
```bash
cd ceilo-backend
```

### **Step 2: Install Dependencies**
Dependencies are pre-installed, but if necessary update them:
```bash
go mod tidy
go mod download
```

### **Step 3: Environment Configuration**
File `.env` already at `internal/config/.env`. Customize with your PostgreSQL configuration:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ceilo_db

SERVER_PORT=8080
JWT_SECRET=ceilo-mental-health-secret-key-2025
```

### **Step 4: Add Missing Dependency**
```bash
go get golang.org/x/crypto/bcrypt
```

### **Step 5: Run Backend Server**
```bash
go run cmd/main.go
```

The server will be running on: **http://localhost:8080**

### **Step 6: Test API (Optional)**
Open a browser or Postman and access:
```
http://localhost:8080/api/health
```

Response:
```json
{
  "status": "ok",
  "message": "Ceilo API is running"
}
```

---

## 🛠️ TROUBLESHOOTING

### **Backend Issues**

#### Error: "Failed to connect to database"
```bash
# Make sure PostgreSQL is running
sudo systemctl status postgresql

# Or check it manual
psql -U postgres -d ceilo_db
```

#### Error: "Address already in use"
```bash
# Kill the process that is using port 8080
lsof -ti:8080 | xargs kill -9

# Or change the port in .env
SERVER_PORT=8081
```

#### Error: "Missing dependency"
```bash
# Reinstall dependencies
go mod tidy
go mod download
```

---

## 📂 File Structure

### **Backend (ceilo-backend/)**
```
ceilo-backend/
├── cmd/
│   └── main.go                           # Entry point
├── internal/
│   ├── config/
│   │   ├── config.go                    # Configuration loader
│   │   └── .env                         # Environment variables
│   ├── database/
│   │   ├── connection.go                # Database connection
│   │   ├── migration.go                 # Database migrations
│   │   └── seed.go                      # Database seeder
│   ├── models/
│   │   ├── user.go                      # User model
│   │   ├── forum.go                     # Forum model
│   │   ├── stress_test.go               # Stress test model
│   │   ├── consultation.go              # Consultation model
│   │   ├── article.go                   # Article model
│   │   └── event.go                     # Event model
│   ├── repository/
│   │   ├── user_repository.go           # User data access
│   │   ├── forum_repository.go          # Forum data access
│   │   └── consultation_repository.go   # Consultation data access
│   ├── service/
│   │   ├── user_service.go              # User business logic
│   │   ├── forum_service.go             # Forum business logic
│   │   └── consultation_service.go      # Consultation business logic
│   ├── handler/
│   │   ├── auth_handler.go              # Auth HTTP handlers
│   │   ├── forum_handler.go             # Forum HTTP handlers
│   │   └── consultation_handler.go      # Consultation HTTP handlers
│   ├── routes/
│   │   └── routes.go                    # API routes
│   ├── middleware/
│   │   ├── jwt_middleware.go            # JWT authentication
│   │   └── cors_middleware.go           # CORS configuration
│   └── utils/
│       ├── response.go                  # Response formatter
│       ├── hash.go                      # Password hashing
│       └── jwt.go                       # JWT utilities
├── go.mod                                # Go modules
└── go.sum                                # Dependencies checksum
```

---

## 🎯 FEATURES THAT HAVE BEEN IMPLEMENTED

✅ User Authentication (Register & Login)  
✅ JWT Token Management  
✅ Protected Routes  
✅ Dashboard with Health Check  
✅ Clean Architecture (Repository-Service-Handler)  
✅ CORS Configuration  
✅ Password Hashing  
✅ Form Validation  
✅ Error Handling  
✅ Responsive Design with Bootstrap  
✅ Database Migration & Seeding  

---

## 👨‍💻 DEVELOPMENT NOTES

### **Backend**
- Using Clean Architecture pattern
- Repository layer for data access
- Service layer for business logic
- Handler layer untuk HTTP requests
- Middleware for authentication & CORS

---

## 📞 SUPPORT

If you have any issues or questions:
1. Check this documentation first.
2. See the Troubleshooting section.
3. Check the console log for error details.
4. Make sure all dependencies are installed correctly.

---

## 📄 LICENSE

This project is for educational purposes.

---

**Happy Coding! 🚀**
