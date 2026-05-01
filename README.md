![image alt](https://github.com/gulmmdoff/GOfensive/blob/2d6c7f408b1773e9c0c37be5a49f1c9092e5febc/GOfensive.png).
# ⚔️ GOfensive: Lightweight & Malleable C2 Framework



**GOfensive** is a modern Command & Control (C2) framework designed to simulate real-world adversary infrastructure. Built with the power of Golang, it provides a lightweight, extensible architecture for managing agents, executing remote commands, and monitoring compromised systems with minimal footprint.

---

## 🚀 Features

- **Agent-Based Architecture:** Efficient and stable management of remote sessions.
- **Remote Shell Execution:** Execute system commands directly on target machines.
- **Real-time Monitoring:** Instant tracking of active agents and their status reports.
- **Minimalist Web UI:** Fast, responsive, and effective operator control panel.
- **Lab & CTF Ready:** Optimized specifically for controlled lab environments and CTF challenges.

---

## 🏗️ Project Structure

- `cmd/` — Entry points: Main source code for both Server and Agent.
- `internal/` — Core logic: Handlers, database models, and security algorithms.
- `ui/` — Frontend: The operator's web-based management console.

---

## 🌐 Architecture

GOfensive operates on a three-way synchronized communication model:
`Agent <----HTTP----> C2 Server <----> Web UI`

1. **Agents:** Connect back to the server at defined intervals to "check-in" and fetch pending tasks.
2. **Server:** Handles tasking, processes agent responses, and maintains the persistent database.
3. **UI:** Allows the operator to interact with the system and visualize agent activity.

---

## ⚡ Quick Start

### 1. Clone the Repository
```bash
git clone [https://github.com/gulmmdoff/GOfensive.git](https://github.com/gulmmdoff/GOfensive.git)
cd GOfensive


---
##

2. Clone the Repository
```bash
git clone [https://github.com/gulmmdoff/GOfensive.git](https://github.com/gulmmdoff/GOfensive.git)
cd GOfensive
