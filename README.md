# **Backend Image Processing Service**

## **📌 Description**

This service processes images collected from stores. It:

1. Accepts jobs with multiple store visits, each containing multiple image URLs.
2. Downloads images and calculates their perimeter using `2 * (Height + Width)`.
3. Simulates GPU processing with a random delay of 0.1 to 0.4 seconds.
4. Stores the results and tracks job status (ongoing, completed, or failed).
5. Provides an API to submit jobs and retrieve job status.

---

## **🔍 Assumptions**

- Images are accessible via public URLs.
- `store_id` is validated based on an external **Store Master** dataset.
- Job processing is handled asynchronously.
- Jobs may take a few minutes to an hour to complete.
- A simple in-memory data store (`sync.Map`) is used for tracking jobs.
- API follows RESTful principles.

---

## **⚙️ Installation & Setup Instructions**

### **1️⃣ Prerequisites**

Ensure the following are installed on your system:

- [Go (1.18+)](https://go.dev/doc/install)
- [Docker](https://www.docker.com/get-started) (Optional)

### **2️⃣ Clone the Repository**

```sh
git clone https://github.com/your-repo/backend-image-service.git
cd backend-image-service
```

### **3️⃣ Setup Go Modules**

```sh
go mod tidy
```

### **4️⃣ Run the Server**

#### Without Docker:

```sh
go run main.go
```

#### With Docker:

```sh
docker build -t backend-image-service .
docker run -p 8080:8080 backend-image-service
```

---

## **🛠️ API Endpoints & Usage**

### **1️⃣ Submit Job**

- **URL:** `POST /api/submit/`
- **Request Body:**

```json
{
  "count": 2,
  "visits": [
    {
      "store_id": "S00339218",
      "image_url": [
        "https://www.gstatic.com/webp/gallery/2.jpg",
        "https://www.gstatic.com/webp/gallery/3.jpg"
      ],
      "visit_time": "2025-03-11"
    }
  ]
}
```

- **Response (201 Created):**

```json
{ "job_id": 1 }
```

### **2️⃣ Get Job Status**

- **URL:** `GET /api/status?jobid=1`
- **Response Examples:**
  - **Ongoing:**
    ```json
    { "status": "ongoing", "job_id": 1 }
    ```
  - **Completed:**
    ```json
    { "status": "completed", "job_id": 1 }
    ```
  - **Failed:**
    ```json
    {
      "status": "failed",
      "job_id": 1,
      "error": [{ "store_id": "S00339218", "error": "Image download failed" }]
    }
    ```

---

## **🧪 Running Tests**

Run unit tests using:

```sh
go test ./test/
```

---

## **💻 Development Environment**

- **OS:** Ubuntu 22.04 / macOS / Windows WSL
- **IDE:** VS Code / GoLand
- **Libraries:** `net/http`, `image`, `sync`, `math/rand`

---

## **🚀 Future Improvements**

If given more time, improvements would include:

- **Database Integration** (PostgreSQL) for persistent job tracking.
- **Queue System** (RabbitMQ/Kafka) for better job handling.
- **Logging & Monitoring** (Prometheus, Grafana).
- **Distributed Processing** using worker nodes for scalability.
- **Authentication & Rate Limiting** for security.

---

### 🎯 **Thank You! Happy Coding!** 🚀
