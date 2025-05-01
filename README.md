# Itinerary Reconstruction Web App

This is a Go web application that reconstructs a travel itinerary from a list of flight tickets provided as source-destination pairs. The application uses the **Fiber** framework, follows the **Model-View-Controller (MVC)** pattern, and incorporates **generics** for type flexibility.

It exposes a single endpoint:

```
POST /itinerary
```

## ✨ Features

- **Endpoint**: `POST /itinerary`
- **Input**: JSON array of ticket pairs  
  Example:
  ```json
  [["LAX", "DXB"], ["JFK", "LAX"], ["SFO", "SJC"], ["DXB", "SFO"]]
  ```
- **Output**: JSON array of airport codes  
  Example:
  ```json
  ["JFK", "LAX", "DXB", "SFO", "SJC"]
  ```
- **Error Handling**: Detects and returns detailed messages for:
    - Empty inputs
    - Cycles
    - Disconnected paths
    - Invalid/malformed itineraries
- **Performance**: Linear time complexity `O(n)`
- **Architecture**: MVC + Generics
- **Generics Support**: Works with any comparable type, not just strings

## 🗂 Project Structure

```
itinerary-app/
├── controllers/           # HTTP request handlers
│   └── itinerary_controller.go
├── models/                # Data structures with generics
│   └── ticket.go
├── services/              # Business logic with generics
│   └── itinerary_service.go
├── main.go                # Application entry point
├── go.mod                 # Go module definition
└── README.md              # Project documentation
```

## ✅ Prerequisites

- [Go](https://golang.org/dl/) (version **1.21+** for generics support)
- Git

## 🚀 Setup Instructions

### 1. Clone the Repository

```bash
git clone <repository-url>
cd proffesional
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Application

```bash
go run main.go
```

The server will start on:  
📍 `http://localhost:8080`

## 📬 Testing the Endpoint

You can use `curl`, Postman, or any HTTP client.

### Example Request

```bash
curl -X POST http://localhost:8080/itinerary \
  -H "Content-Type: application/json" \
  -d '[["LAX","DXB"], ["JFK","LAX"], ["SFO","SJC"], ["DXB","SFO"]]'
```

### Example Response

```json
["JFK", "LAX", "DXB", "SFO", "SJC"]
```

### Error Response (e.g., empty input)

```json
{
  "error": "tickets array cannot be empty"
}
```

## 🔎 Test Case Coverage

The application handles:

- ✅ Valid itineraries  
  e.g., `[["JFK","LAX"], ["LAX","DXB"]]`
- ❌ Empty input  
  e.g., `[]`
- ❌ Invalid JSON  
  e.g., `{invalid}`
- ❌ Cycles  
  e.g., `[["JFK","LAX"], ["LAX","JFK"]]`
- ❌ Disconnected itineraries  
  e.g., `[["JFK","LAX"], ["SFO","SJC"]]`
- ❌ Missing start/end nodes  
  e.g., incomplete chains

## ⚙️ Development Design

### 🌐 Framework: Fiber
Chosen for its speed (fasthttp) and simplicity (Express-like syntax).

### 🧱 MVC Breakdown

- **Model**:  
  `Ticket[T]` struct in `models/ticket.go`, supports generics

- **Controller**:  
  `ItineraryController` in `controllers/itinerary_controller.go`  
  Handles HTTP, JSON parsing, and response formatting

- **Service**:  
  `ItineraryService[T]` in `services/itinerary_service.go`  
  Core logic for reconstructing itineraries using generics

- **View**:  
  Handled via JSON responses, no frontend needed

### 🧬 Generics

Using Go 1.18+ generics:
- `Ticket[T]` and `ItineraryService[T]` are generic over any comparable type
- Reusable and type-safe
- Zero performance cost over concrete types

### 🧠 Algorithm

- Constructs a **directed graph** (adjacency list)
- Tracks in-degrees to find starting node
- Traverses and builds a full path
- Detects:
    - Cycles
    - Disconnected nodes
    - Invalid paths

⏱ **Time complexity**: `O(n)`

### 🧪 Input Handling

- Converts raw input `[["source", "dest"]]` to `[]Ticket[string]`
- Validates each ticket

### 🔒 Error Handling

Detects and handles:
- Empty or malformed JSON
- Tickets with missing/empty source or destination
- Cycles or unreachable nodes
- Incomplete paths

## 📦 Dependencies

- [`github.com/gofiber/fiber/v2`](https://github.com/gofiber/fiber) – HTTP framework
- Minimal additional dependencies

---

## 📄 License

MIT

## 👨‍💻 Author

Built with Go 1.21+, Fiber, and love ❤️
