# 📦 Snippetbox

A web application built while studying the book **"Let's Go" by Alex Edwards**.

This project follows the book step-by-step and implements the complete Snippetbox application. It demonstrates how to build a secure, production-style web application using Go and MySQL, along with Docker containerization.

---

# Live demo on

[Live app](https://snippetbox-1948.onrender.com)

## 🚀 Features

- Create, view, and list text snippets
- User signup and authentication
- Secure password hashing (bcrypt)
- Session management
- CSRF protection
- Middleware chaining
- Structured logging
- TLS (HTTPS) support
- MySQL integration
- Docker & Docker Compose support

---

## 🛠 Requirements

### To Run Without Docker

- Go 1.25.x
- MySQL (local or cloud instance)

### To Run With Docker

- Docker
- Docker Compose

---

## ⚙️ Configuration

The application uses command-line flags for configuration.

### Supported Flags

| Flag  | Description | Example |
|-------|------------|----------|
| -addr | HTTP network address | :4000 |
| -dsn  | MySQL Data Source Name | user:pass@tcp(host:3306)/snippetbox?parseTime=true |

---

## 🧩 Database Setup

⚠️ This project does NOT automatically create the database or tables.  
You must create them manually before running the application.

### 1️⃣ Create Database

```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 2️⃣ Create Tables

#### Snippets Table

```sql
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);
```

#### Users Table

```sql
CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

---

## ▶️ Running Without Docker

### 1️⃣ Clone Repository

```bash
git clone https://github.com/YOUR_USERNAME/snippetbox.git
cd snippetbox
```

### 2️⃣ Run Application

```bash
go run ./cmd/web \
  -addr=:4000 \
  -dsn="user:password@tcp(host:3306)/snippetbox?parseTime=true"
```

we can also use .env file for passing secret
create .env file at Snippetbox/cmd/web with

```.env
app_port=4000
database_dsn= "user:password@tcp(host:3306)/snippetbox?parseTime=true"

```

Then open:

<https://localhost:4000>

---

## 🐳 Running With Docker

This project includes a multi-stage Docker build and Docker Compose configuration.

### 1️⃣ Create .env File

Create a file named:

.env

Add:

```
app_port=4000
database_dsn=user:password@tcp(host:3306)/snippetbox?parseTime=true
```

⚠️ Important:

- Do NOT commit `.env` to GitHub.
- Add `.env` to `.gitignore`.

### 2️⃣ Run With Docker Compose

```bash
docker compose up --build
```

Then visit:

<https://localhost:4000>

---

## 🐳 Docker Notes

- Multi-stage build is used for a smaller final image.
- The application runs as a non-root user inside the container.
- Environment variables are passed as flags using Docker Compose.
- Database is expected to be external (cloud or local).

---

## 📂 Project Structure

```
.
├── cmd/
│   └── web/            # Application entry point
├── internal/           # Application logic and models
├── ui/                 # Templates and static files
├── tls/                # TLS certificates
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## 🔐 Security Notes

- Database credentials are passed via flags (not hardcoded).
- `.env` file must never be committed.
- Rotate credentials immediately if exposed.
- In production, consider using:
  - Docker secrets
  - Cloud secret managers
  - Reverse proxy with Let's Encrypt

---

## 🧪 Useful Commands

### Build Binary

```bash
go build -o app ./cmd/web
```

### Run Tests

```bash
go test ./...
```

### View Docker Logs

```bash
docker compose logs -f
```

---

## 📚 Learning Purpose

This project is built for educational purposes while studying:

**"Let's Go" by Alex Edwards**

It is intended to:

- Learn Go web development
- Understand structured backend architecture
- Practice Docker containerization
- Apply secure configuration practices

---

## 👨‍💻 Author

Built while learning Go, backend engineering, and containerized deployments.

---

## 📄 License

This project is for educational use.
