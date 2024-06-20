# GoLang Exoplanet Service
Welcome to the GoLang Exoplanet Service! This service provides information about exoplanets discovered beyond our solar system.

# Table of Contents
Overview
Features
Prerequisites
Installation
Usage
API Documentation

# Overview
The GoLang Exoplanet Service is a RESTful API service that allows users to query and retrieve data about exoplanets. It provides endpoints to fetch details about exoplanets based on various criteria such as name, discovery method, and more.

# Features
Add new exoplanet data
Retrieve a list of all exoplanets
Fetch details of a specific exoplanet by ID
Update exoplanet details 
Delete exoplanet data
Estimate fuel requirements

# Prerequisites
Before running the GoLang Exoplanet Service, ensure you have the following installed:

Go Programming Language: Install Go
Docker (optional for containerized deployment): Install Docker

Installation
Follow these steps to set up and run the GoLang Exoplanet Service:

# Clone the repository:


git clone https://github.com/Tamil-k/exoplanet.git
cd exoplanet-service
# Build and run the service locally:

go build
./exoplanet-service

# Docker Installation (Optional)
Alternatively, you can use Docker for deployment:

# Build the Docker image:
docker build -t exoplanet-service .
# Run the Docker container:
docker run -d -p 9000:9000 exoplanet-service
The service will be accessible at http://localhost:9000.

# Usage
Once the service is running, you can interact with it via HTTP requests. Here are some example API endpoints:
POST /exoplanets: Adds a new exoplanet.
GET /exoplanets: Retrieves a list of all exoplanets.
GET /exoplanets/{id}: Retrieves details of a specific exoplanet by ID.
PUT /exoplanets/{id}: Updates details of a specific exoplanet.
DELETE /exoplanets/{id}: Deletes a specific exoplanet.
GET /exoplanets/{id}/{crewcapacity}/fuel: Estimates fuel requirements for travel to a specific exoplanet.

# Example usage with curl:

# Retrieve all exoplanets
curl http://localhost:9000/exoplanets

# Retrieve exoplanet details by ID
curl http://localhost:9000/exoplanets/1

# Estimate fuel for travel to an exoplanet
curl http://127.0.0.1:9000/exoplanets/bdaaac26-c0d5-49ea-8089-e15e34b81830/2/fuel
