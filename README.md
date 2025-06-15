🏦 Banking Application
A console-based banking application developed in Go that simulates real-world banking operations like account creation, transactions. 
This project follows a layered architecture for clean separation of concerns.

📌 Features
✅ Create and manage customer accounts (Savings, Current)
💵 Deposit, Withdraw, and Transfer funds
📜 View transaction history
📊 Generate account and transaction reports
🧪 Unit testing support
🐳 Dockerized for easy deployment

🛠️ Technologies Used
Go (Golang) — Main programming language
File I/O — For data persistence
Environment Variables — Configuration management
Docker — Containerization
Unit Testing — Built-in Go test framework

🚀 Getting Started
Prerequisites
Go 1.16+
Git

Clone the Repository: git clone https://github.com/rakeshdatti/banking_application.git
cd banking_application
error and logger libary need to import: go get https://github.com/rakeshdatti/banking_lib
Install Dependencies: go mod tidy
Run the Application: go run main.go
🧪 Running Tests
go test ./...
🐳 Docker Setup
Build the Docker image: docker build -t banking_application .
Run the container: docker run -it banking_application


📄 Documentation
Flow.txt: Step-by-step user interaction flow
doc.txt: Functional overview
docTest.txt: Test case summaries
docker.txt: Instructions for Docker setup

👨‍💻 Author
Rakesh Datti
🔗 GitHub

📜 License
This project is open-source and available under the MIT License.

