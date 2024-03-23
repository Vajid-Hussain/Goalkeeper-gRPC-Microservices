# Goalkeeper - gRPC Microservices Task Management & Reminder System

Goalkeeper is a distributed system built using gRPC and microservices that helps you organize tasks, set reminders, and receive timely email notifications.

## Features

* **User Authentication:** Secure user registration and login managed by a dedicated service.
* **Task Management:** 
    * Create custom categories to organize tasks effectively.
    * Add tasks with labels, descriptions, and due dates.
* **Automated Reminders:** Receive daily email reminders for upcoming tasks.
* **Efficient Cleanup:** Expired reminders are deleted to maintain data hygiene.

## Architecture

* **API Gateway:** Handles client requests and routing.
* **Authentication Service (PostgreSQL):** Manages user data and authentication.
* **Vault Service (MongoDB):** Core task and reminder management.
* **Mail Service:** Integrated email notifications with cron-based scheduling.

## Technologies

* **gRPC:** For communication between microservices.
* **Languages:** (Specify for each service)
* **Databases:** PostgreSQL, MongoDB
* **Mail:** (Specify your provider/library)
* **Cron:** (Specify your tool)
* **Docker:** For containerization of individual microservices.
* **Docker Compose:** To orchestrate multi-container deployment.
* **Kubernetes:** For production-grade deployment and scaling.

## Getting Started

**Prerequisites:**

* (List all prerequisites, including Docker, Kubernetes if applicable)

**Installation**

1. **Clone:** `git clone https://github.com/Vajid-Hussain/Goalkeeper-gRPC-Microservices.git`
2. **Build Docker Images:**
   
**Deployment**

* **Development (Docker Compose):**
    `docker-compose up -d` 

* **Production (Kubernetes):**
    `kubectl apply -f deploy.sh` 

