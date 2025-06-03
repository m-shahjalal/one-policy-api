# OnePolicy API

OnePolicy API is a Go-based backend service that provides policy management functionality. Built with modern Go practices and a clean architecture approach.

## 🚀 Features

- RESTful API endpoints
- JWT Authentication
- PostgreSQL Database Integration
- Environment-based Configuration
- CORS Support
- Input Validation
- Clean Architecture

## 🛠️ Tech Stack

- Go 1.24.3
- Gin Web Framework
- GORM (ORM)
- PostgreSQL
- JWT for Authentication
- Viper for Configuration
- Air for Live Reload

## 📋 Prerequisites

- Go 1.24.3 or higher
- PostgreSQL
- Make (for using Makefile commands)

## 🔧 Installation

1. Clone the repository:

```bash
git clone https://github.com/m-shahjalal/onepolicy-api.git
cd one-policy-api
```

2. Install dependencies:

```bash
go mod download
```

3. Set up environment variables:
   Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
JWT_SECRET=your_jwt_secret
```

## 🚀 Running the Application

### Development Mode

```bash
make dev
```

This will start the application with hot-reload enabled using Air.

### Production Mode

```bash
make build
./bin/one-policy-api
```

## 📁 Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── config/              # Configuration files
├── internal/
│   ├── controller/      # HTTP request handlers
│   ├── middleware/      # HTTP middleware
│   ├── model/          # Data models
│   ├── router/         # Route definitions
│   ├── service/        # Business logic
│   └── validator/      # Input validation
├── utils/              # Utility functions
├── go.mod             # Go module file
├── go.sum             # Go module checksum
└── Makefile           # Build automation
```

## 🔐 API Documentation

API documentation is available at `/swagger` when running the application.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👥 Authors

- Shahjalal - Initial work

## 🙏 Acknowledgments

- Gin Web Framework
- GORM
- All other open-source contributors
