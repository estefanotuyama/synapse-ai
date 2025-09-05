# GoRAG Agent Platform

![Status](https://img.shields.io/badge/status-work%20in%20progress-yellow)

> A web-based system for creating, training, and interacting with custom RAG-powered AI agents.

## About The Project

This project aims to build a complete system where users can define custom AI "agents" with specific purposes. Users can upload documents (like PDFs, TXT, etc.) to provide a knowledge base for each agent. The system will then use a Retrieval-Augmented Generation (RAG) pipeline to allow users to chat with their agents, receiving answers grounded in the provided documents.

## Planned Technology Stack

* **Backend:** Go
* **Frontend:** React
* **Vector Database:** Weaviate
* **Persistence Database:** MongoDB

## Project Status

This project is in the **initial setup phase**. The foundational structure for the Go backend is currently being built.

### High-Level Roadmap

- Set up Go backend server skeleton & API structure.
- Implement document ingestion pipeline (parsing, chunking, embedding).
- Develop the core RAG chat endpoint.
- Integrate MongoDB for persistence and Weaviate for vector storage.
- Build the React frontend for agent and chat management.

---
