![image alt](https://github.com/gulmmdoff/GOfensive/blob/2d6c7f408b1773e9c0c37be5a49f1c9092e5febc/GOfensive.png).
Overview

GOFensive is designed to simulate real-world adversary infrastructure. It provides a simple yet extensible architecture for managing agents, executing commands, and monitoring compromised systems in controlled environments.

Features
Agent-based architecture
Command execution (remote shell)
Lightweight Go implementation
Real-time agent monitoring
Simple web UI control panel
Modular structure (easy to extend)
Designed for lab & CTF environments
Project Structure
cmd/        → entry points (server, agent)
internal/   → core logic (handlers, models, logic)
ui/         → frontend panel
Architecture
Agent  <----HTTP---->  C2 Server  <---->  Web UI
Agents connect back to the server
Server handles tasking & responses
UI allows operator interaction
Quick Start
1. Clone Repository
git clone https://github.com/gulmmdoff/GOfensive.git
cd GOfensive
2. Run Server
go run cmd/server/main.go

Server will start on:

http://0.0.0.0:8080
3. Build Agent
go build -o agent.exe ./cmd/agent/main.go

Run the agent on target machine.

Usage
Open web panel
Wait for agent connection
Select agent
Execute commands

Example commands:

whoami
hostname
ipconfig
dir
