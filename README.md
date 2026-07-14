# Go Agent Framework 🚀

A lightweight, extensible AI Agent Framework built in Go that demonstrates how to build AI agents from scratch using clean architecture, tool execution, session memory, and Model Context Protocol (MCP).

The project focuses on:

* Building AI agents in Go
* LLM integration using Ollama
* Tool-based agents
* Session memory and persistence
* MCP integration
* Clean Architecture and Dependency Injection

---

# ✨ Features

## AI & Agent Capabilities

* ✅ Chat API
* ✅ Session-based conversation memory
* ✅ Tool execution framework
* ✅ Dynamic tool registration
* ✅ Planner and Executor architecture
* ✅ MCP (Model Context Protocol) integration
* ✅ Dynamic MCP tool discovery
* ✅ Dynamic MCP tool execution

## Built-in Tools

* ✅ Time Tool
* ✅ Calculator Tool
* ✅ Weather Tool

## Persistence

* ✅ SQLite database
* ✅ Conversation history persistence
* ✅ GORM integration

## Architecture & Engineering

* ✅ Dependency Injection
* ✅ Repository Pattern
* ✅ Clean Architecture
* ✅ Interface-driven design
* ✅ Modular package structure

## Infrastructure

* ✅ Docker Support
* ✅ Docker Compose Support
* ✅ Ollama Integration
* ✅ Local LLM execution (Phi3:Mini)

---

# 🏗 Architecture

```text
                         ┌──────────────────┐
                         │      User        │
                         └────────┬─────────┘
                                  │
                                  ▼
                         ┌──────────────────┐
                         │     HTTP API     │
                         └────────┬─────────┘
                                  │
                                  ▼
                         ┌──────────────────┐
                         │   Agent Service  │
                         └────────┬─────────┘
                                  │
                  ┌───────────────┴───────────────┐
                  ▼                               ▼
         ┌────────────────┐              ┌────────────────┐
         │    Planner     │              │    Executor    │
         └────────────────┘              └────────────────┘
                  │                               │
                  └───────────────┬───────────────┘
                                  │
                                  ▼
                         ┌──────────────────┐
                         │   Tool Registry  │
                         └────────┬─────────┘
                                  │
         ┌──────────────┬─────────┼─────────┬──────────────┐
         ▼              ▼         ▼         ▼              ▼
     Time Tool     Calculator   Weather   MCP Tools    Future Tools

                                  │
                                  ▼
                         ┌──────────────────┐
                         │ Session Memory   │
                         │ SQLite + GORM    │
                         └──────────────────┘

                                  │
                                  ▼
                         ┌──────────────────┐
                         │ Ollama + Phi3    │
                         └──────────────────┘
```

---

# 🧩 MCP Architecture

```text
Go Agent
     │
     ▼
Go MCP SDK (Client)
     │
     ▼
Filesystem MCP Server (Node.js)
     │
     ▼
Local Filesystem
```

The framework can connect to any MCP server:

* Filesystem MCP Server
* GitHub MCP Server
* PostgreSQL MCP Server
* Slack MCP Server
* Custom MCP Servers

---

# 📂 Project Structure

```text
go-agent-framework/
├── cmd/
│   ├── server/
│   └── mcp-test/
│
├── internal/
│   ├── agent/
│   ├── api/
│   ├── clients/
│   ├── config/
│   ├── db/
│   ├── handler/
│   ├── llm/
│   ├── mcp/
│   ├── models/
│   ├── repository/
│   └── tools/
│
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

---

# 🛠 Tech Stack

| Category         | Technology                   |
| ---------------- | ---------------------------- |
| Language         | Go 1.25                      |
| HTTP Framework   | Chi Router                   |
| Database         | SQLite                       |
| ORM              | GORM                         |
| LLM Runtime      | Ollama                       |
| Model            | Phi3:Mini                    |
| Containerization | Docker                       |
| Orchestration    | Docker Compose               |
| Protocol         | MCP (Model Context Protocol) |
| MCP SDK          | Go MCP SDK                   |
| MCP Server       | Filesystem MCP Server        |

---

# 🚀 Getting Started

## Prerequisites

* Go 1.25+
* Docker
* Docker Compose

Optional:

* Node.js
* npm

---

## Clone Repository

```bash
git clone https://github.com/BalamuruganP25/go-agent-framework.git

cd go-agent-framework
```

---

## Start the Application

```bash
make setup
make run
```

---

# ⚠️ First Startup

The first startup may take several minutes because Ollama needs to download the LLM model (`phi3:mini`).

The model is downloaded only once and stored in a Docker volume.

---

## Download the Model

```bash
make setup
```

Equivalent command:

```bash
docker exec -it ollama ollama pull phi3:mini
```

---

## Verify the Model

```bash
docker exec -it ollama ollama list
```

Expected output:

```text
NAME       ID        SIZE
phi3:mini  xxxxxxxx  xxx MB
```

---

## Expected Startup Time

| Action               | Approximate Time |
| -------------------- | ---------------- |
| First model download | 1-5 minutes      |
| Subsequent startups  | A few seconds    |

---

## View Logs

```bash
make logs
```

---

## Stop Application

```bash
make down
```

---

## Rebuild Docker Images

```bash
make rebuild
```

---

# 🌐 API Endpoints

## Health Check

```bash
curl http://localhost:8084/health
```

---

## Chat API

```bash
curl -X POST http://localhost:8084/api/v1/chat \
-H "Content-Type: application/json" \
-d '{
  "session_id":"bala",
  "message":"What is the weather in Chennai?"
}'
```

---

# 💬 Example Requests

## Time

```text
What time is it?
```

---

## Calculator

```text
100 + 200
```

---

## Weather

```text
What is the weather in Chennai?
```

---

## MCP Examples

### List Files

```text
List files in the current directory
```

### Read README

```text
Read README.md
```

### Find Go Files

```text
Find all Go files
```

---

# 🧠 Current Agent Capabilities

### Single Tool Planning

```text
User
 ↓
Planner
 ↓
Executor
 ↓
Tool
```

### Session Memory

```text
Session ID
     ↓
SQLite
     ↓
Conversation History
```

### Dynamic MCP Tools

```text
MCP Server
      ↓
Tool Discovery
      ↓
Tool Registry
      ↓
Agent
```

---

# 🗺 Roadmap

## Completed

* [x] Chat API
* [x] SQLite Persistence
* [x] Repository Pattern
* [x] Dependency Injection
* [x] Session Memory
* [x] Tool Registry
* [x] Planner
* [x] Executor
* [x] MCP Integration
* [x] Docker Support

---

## Upcoming

### Agent Improvements

* [ ] Multi-step Planning
* [ ] ReAct Loop
* [ ] Streaming Responses
* [ ] Structured Outputs

### AI Features

* [ ] RAG Integration
* [ ] Vector Database Support
* [ ] Semantic Search
* [ ] Document Ingestion

### Multi-Agent System

* [ ] Supervisor Agent
* [ ] Research Agent
* [ ] Code Agent
* [ ] File Agent

### Production Features

* [ ] Authentication
* [ ] Authorization
* [ ] Metrics
* [ ] Tracing
* [ ] Kubernetes Deployment

---

# 📖 Learning Goals

This project demonstrates:

* Building an AI Agent Framework from scratch in Go
* Clean Architecture
* Dependency Injection
* Repository Pattern
* Tool-based Agents
* LLM Integration with Ollama
* Session Memory using SQLite
* MCP Integration
* Dockerized Development

---

# 📄 License

MIT License.

---

# ⭐ If you found this project useful, please consider giving it a star.
