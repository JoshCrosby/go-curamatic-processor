# FHIR Go-Curamatic-Processor Application

The FHIR Go-Curamatic-Processor Application is designed to periodically fetch, validate, and ingest FHIR files from a specified source into a relational database. This application is built with Go and utilizes Docker for easy setup and deployment.

## Requirements

To run this application, you will need:

- Docker: For creating and managing the application and database containers.
- Docker Compose: For orchestrating multi-container Docker applications.
- Make (optional): For simplifying Docker and Docker Compose commands with a Makefile.

## Setup Instructions

### 1. Install Docker and Docker Compose

- **Docker:** Follow the installation guide for Docker at [Docker's official site](https://docs.docker.com/get-docker/).
- **Docker Compose:** Docker Desktop for Windows and Mac includes Docker Compose. For Linux, follow the [Docker Compose installation guide](https://docs.docker.com/compose/install/).

### 2. Install Make (Optional)

- **Linux:** Usually pre-installed. If not, you can install it using your distribution's package manager, for example, `sudo apt-get install make` on Debian/Ubuntu.
- **Mac:** Comes pre-installed with Command Line Tools for Xcode. To install Command Line Tools, run `xcode-select --install` in the terminal.
- **Windows:** You can install Make through various packages like [Chocolatey](https://chocolatey.org/) (`choco install make`) or [Cygwin](https://www.cygwin.com/).

### 3. Clone the Project

Clone the project repository to your local machine:

```
git clone https://github.com/JoshCrosby/go-curamatic-processor
cd go-curamatic-processor
```

### 4. Build and Run the Application

Using Docker Compose directly:

- **Build the application:** `docker-compose build`
- **Start the application:** `docker-compose up -d`
- **Stop the application:** `docker-compose down`

Using the Makefile commands:

- **Build the application:** `make build`
- **Start the application:** `make up`
- **Stop the application:** `make down`
- **View logs:** `make logs`

## Additional Commands

- **Rebuild and restart the application:** `make rebuild`
- **Access the application container (bash):** `make exec-bash`
- **Clean up Docker volumes (removes all data):** `make clean`

**Note:** If you have not installed Make, you can run the Docker Compose commands listed in the Makefile directly.

## Support

For support, please contact Josh Crosby joshdcrosby@gmail.com or open an issue in the project repository. 