# SIST Admission Management System Backend
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)


Backend for SIST Admission Management System built using the Go programming language.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Gin](https://gin-gonic.com/docs/introduction/installation/)
- [GORM](https://gorm.io/docs/installation.html)
- [Docker (optional)](https://www.docker.com/products/docker-desktop)

### Installing

- Clone the repository

```bash
git clone https://github.com/SIST-Admission/adm-backend.git
```

- Install dependencies

```bash
go mod tidy
```

### Build and Run the server natively

- Build the project
```bash
go build -o adm-backend app/sevice.go
```

- Run the project Using Unix or Linux
```bash
./adm-backend
```

- Using Windows
```bash
adm-backend.exe
```

### Run Using Docker
 If you don't have docker installed, you can download it from [here](https://www.docker.com/products/docker-desktop)

 You don't need to install Go, PostgreSQL or any other dependencies. Docker will take care of that.

- Build the image

```bash
docker build -t adm-backend .
```

- Run the container

```bash
docker run -p 8080:8080 adm-backend
```
## Authors

- [**Bijay Sharma**](https://github.com/BijaySharma)
- [**Sefali Basnet**](https://github.com/sefali20)


