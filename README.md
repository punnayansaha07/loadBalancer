# Load Balancer Implementation

## Overview
This project is a **custom Load Balancer** implemented in **Go**, which efficiently distributes incoming traffic among multiple **Flask-based** backend servers. The system ensures high availability and optimized request distribution using a **Round Robin** load balancing algorithm.

## Features
- **Round Robin Load Balancing**: Requests are distributed evenly across available backend servers.
- **Health Check Mechanism**: Periodic pings ensure only healthy servers receive traffic.
- **Fault Tolerance**: Automatically skips failed servers and continues routing traffic.
- **Concurrency Support**: Handles multiple requests efficiently.
- **Flask-based Backend Servers**: Simulated microservices handling requests.

## Tech Stack
- **Go**: Load Balancer
- **Flask (Python)**: Backend Servers
- **PowerShell / Shell**: Process Management

## Installation and Usage

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/punnayansaha07/loadBalancer.git
cd LoadBalancer
```

### 2️⃣ Start Virtual Environment (For Flask Servers)
#### Windows (PowerShell)
```sh
.venv\Scripts\Activate.ps1
```
#### Mac/Linux
```sh
source .venv/bin/activate
```

### 3️⃣ Install Dependencies
```sh
pip install -r requirements.txt
```

### 4️⃣ Start Backend Servers
Run the following commands in separate terminals:
```sh
python server.py server-1 5001 &
python server.py server-2 5002 &
python server.py server-3 5003 &
python server.py server-4 5004 &
python server.py server-5 5005 &
```

### 5️⃣ Start Load Balancer
```sh
go run main.go
```

### 6️⃣ Test the Load Balancer
Open a browser or use `curl`:
```sh
curl http://localhost:8080/
```
Each request will be routed to a different backend server in a **Round Robin** manner.

## Algorithm Used
The Load Balancer uses a **Round Robin Algorithm**, where requests are forwarded to backend servers in sequential order. If a server is unresponsive, it is skipped.

## Contributing
Feel free to open issues or pull requests to enhance the project!

## License
This project is open-source and available under the [MIT License](LICENSE).
