# Notification Service

The Notification Service is responsible for sending notifications to users based on upcoming task deadlines in the TODO List API. This service checks for tasks with approaching deadlines and sends email reminders using Mailgun.

## Concepts Covered

- **Clean Architecture**: The service follows the principles of Clean Architecture, with a clear separation of concerns between different layers (handlers, services, repositories). This ensures maintainability, testability, and extensibility.
- **Email Notifications**: The service integrates with Mailgun to send email reminders to users.
- **Database Integration**: Uses a PostgreSQL database to store notification data and manage user preferences.
- **Docker Compose Integration**: Configured to run in a multi-container Docker environment alongside the TODO List API.

## Requirements

- Golang
- Docker
- Docker Compose
- Mailgun API credentials
- PostgreSQL

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repository/notification-service.git
   cd notification-service

2. Set up the environment variables by creating a `.env` file in the root of the project, containing the necessary configuration like database credentials, and Mailgun API keys.

3. Build and start the application with Docker Compose:
   ```bash
   docker-compose up --build
   ```