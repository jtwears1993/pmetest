# PME TAKE HOME TEST 
## Design and Architecture 

This API is designed with a clean, idiomatic Go architecture that emphasizes separation of concerns, testability, and maintainability. The structure is framework-free — aside from using chi as a minimal, idiomatic router — and it leverages Go’s powerful standard library to its full potential.

I don’t believe full-featured frameworks like Gin or Echo offer meaningful performance or ergonomic advantages for this use case. Instead, they often introduce unnecessary abstractions and package bloat, which go against Go’s simplicity-first philosophy.

📁 Directory Overview
python
Copy
Edit
.
├── cmd/
│   └── api/                # Entry point: server, routes, middleware, handlers, config 
├── internal/
│   ├── services/
│   │   └── trip/           # Business logic for itinerary generation
│   ├── request/            # Request parsing & validation structs
│   └── response/           # Response formatting & helper logic
└── go.mod / go.sum         # Go module files

🔧 cmd/api/: Transport Layer

This layer is responsible for:

 - Starting the HTTP server
 - Registering routes
 - Defining middleware (e.g., logging, panic recovery)
 - Handling incoming HTTP requests and returning JSON responses
 - Defining API-specific error types
 - Each route has a corresponding handler that:
 - Parses and validates the request
 - Calls the appropriate service method
 - Constructs and returns a JSON response


💼 internal/services/trip: Business Logic

Contains the core domain logic, including:

 - TripService: Responsible for reconstructing the full itinerary from a list of flights
 - Stateless and framework-agnostic — no knowledge of HTTP
 - Easy to test with pure Go unit tests - run `make test` to run unit tests against our services.

🧩 internal/request & internal/response

 - request: Defines typed request payloads with optional validation helpers
 - response: Centralized helpers for consistent JSON responses and error formatting


✅ Benefits of This Design

1. Separation of Concerns: Transport logic is kept out of the business layer
2. Testability: Core logic can be tested without running an HTTP server
3. Extensibility: New services can be added easily (e.g., internal/services/booking)
4. Maintainability: Small, focused packages with clearly defined responsibilities
5. Idiomatic Go: No frameworks, no magic — just clean Go 

🔧 Potential Improvements
This project is designed to be extendable and production-ready with minimal additions. Here are some enhancements that can be layered in as needed:

🔍 Observability

 - Add tracing support (e.g., OpenTelemetry, Jaeger) to trace requests across services
 - Expose Prometheus metrics for runtime stats, request durations, error rates, etc.

🔐 Security

 - Add OAuth2 / JWT middleware for route-level authentication and authorization
 - Integrate rate limiting and request validation at the middleware level
 - Enable CORS configuration if exposing a public-facing API

🚀 Deployment Options

Kubernetes:

 - Package the service as a container and deploy via a Helm chart. 
 - Add a Kubernetes Deployment, Service, and Ingress for traffic routing

Bare Metal / VM:

 - Use a Docker container for portability
 - Use Caddy, Nginx, or Traefik as a reverse proxy with TLS termination and automatic HTTPS
 - Optionally use systemd to run the service as a managed unit

📦 Configuration & Secrets

 - Use a .env file or environment variables for config (via github.com/caarlos0/env)
 - For production, integrate with secret managers like Vault or AWS Secrets Manager

✅ CI/CD

 - Add a CI pipeline for build, test, lint, and Docker image publishing (e.g., GitHub Actions)
 - Push tagged releases to a container registry (Docker Hub, GHCR, ECR, etc.)


## Setup 

Two options to run server:

  1. Build docker image and then run in network mode host (on linux only - not the same for Mac or Windows due to how docker works)
  2. Run `make run` - easiest option, will build and execute the binary 

To query the endpoint:

```bash 
  curl -X POST http://localhost:4444/itinerary   -H "Content-Type: application/json"   -d '[["LAX","DXB"],["JFK","LAX"],["SFO","SJC"],["DXB","SFO"]]'

*Note! this was built and tested on Ubuntu 24.04 running on AMD ryzen 7 PRO processors. Performance and build experiences may differ on your machine.
