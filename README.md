# Portfolio Website

Welcome to my portfolio website! This repository contains the source code for my personal portfolio site, which is deployed on Vercel. The backend is built using Go (Golang), and all templates are custom-made for the site.

### Features
- **URL**: Available at [rojin.dev](https://rojin.dev)
- **Backend**: Written in Go (Golang)
- **Frontend**: Custom HTML templates
- **Deployment**: Deployed on Vercel for easy and fast hosting
- **Docker**: The application is containerized using Docker for simplified deployment and scalability
- **Responsive Design**: The website is designed to be responsive, making it accessible on various devices
- **Pages**: Includes pages like "Home", "About", "Skills", etc.

### Getting Started

#### Prerequisites

- **Go**: Ensure Go is installed on your system (version 1.22 or higher).
- **Docker**: Make sure Docker is installed for containerization.
- **Vercel Account**: If you want to deploy on Vercel, sign up for a free account at [Vercel](https://vercel.com/).

#### Running Locally

To run the website locally, follow these steps:

1. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/rojinebrahimi/portfolio.git
   cd portfolio
   ```

2. Build the Go backend (if not using Docker):
   ```bash
   go build -o server
   ```

3. Run the Go server locally:
   ```bash
   go run main.go
   ```

4. Open your browser and visit `http://localhost:8080` to view the website.

#### Running with Docker

1. Build the Docker image:
   ```bash
   docker build -t portfolio .
   ```

2. Run the Docker container:
   ```bash
   docker run -p 8080:8080 portfolio
   ```

3. Visit `http://localhost:8080` to view the website.
