# Go Todo Backend API

This is a simple RESTful Todo Backend API built using Go and the Fiber framework.

The purpose of this project is to demonstrate:

- REST API design
- Proper HTTP method usage
- JSON file persistence
- Business rule enforcement at application level
- Clean backend logic without frontend dependency

---

## 📌 Core Features

• Add a new todo  
• Get all todos  
• Delete a todo  
• Mark a todo as completed  
• Live count endpoint (total / completed / pending)  
• Data persistence using JSON file  
• Constraint: Only one todo can be completed at a time  

---

## 🧠 Design Decisions

- Todos are stored in-memory using a slice.
- Data is persisted to a `todos.json` file after every modification.
- On server restart, data is loaded from the JSON file.
- Business constraint (single completed todo) is enforced inside the PATCH logic.
- Counts are calculated dynamically to ensure accuracy.

---

## 🛠 Tech Stack

- Go
- Fiber (github.com/gofiber/fiber/v2)
- Standard library (encoding/json, os, time)

---


---

## 🚀 How to Run

1. Clone the repository

git clone <https://github.com/RajjakAhmed/Go_Todo_backend_api>

2. Navigate to the project folder

cd go-todo-backend-api

3. Install dependencies

go mod tidy

4. Run the server

go run main.go

Server runs on:
http://localhost:3000

---

## 🧪 API Testing

Test using Postman or any API client.

---

## 👤 Author

Rajjak Ahmed  
Computer Science Engineering Student  
