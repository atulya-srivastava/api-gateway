# Go API Gateway

Hey there! I'm building this API Gateway from scratch to learn Go and get hands-on experience with concurrency, networking, and backend concepts.

## What's Inside

- Dynamic Service Routing: Mapping paths to backend service clusters
- Round-Robin Load Balancing: Distributing requests evenly and thread-safely across healthy backends
- Active Health Checks: Concurrently checking backend health using goroutines, channels, and a background ticker
- Custom Middleware: Simple authentication, rate limiting, and request logging middleware
- Reverse Proxying: Forwarding HTTP requests efficiently while reusing connection pools

## Project Structure

```
.
├── backend/        # Sample backend target servers for testing
├── healthcheck/    # Health check logic & background monitor
├── loadbalancer/   # Round-robin load balancer & proxy caching
├── middleware/     # Auth, rate limiting, and logging middleware
├── proxy/          # Reverse proxy setup
├── services/       # Service route definitions
└── main.go         # Application entry point
```

## Running it Locally

Make sure you have Go installed on your machine.

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/api-gateway.git
   cd api-gateway
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

The gateway will start listening on port 8080.

## Contributing

Since I'm making this project to learn Go, any suggestions, fixes, or improvements are super welcome! If you'd like to contribute:

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/cool-feature`)
3. Commit your changes (`git commit -m 'Add cool feature'`)
4. Push to your branch (`git push origin feature/cool-feature`)
5. Open a Pull Request :)

Happy coding!