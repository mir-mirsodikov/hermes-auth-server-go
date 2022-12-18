<h1>Hermes - Auth Server</h1>

# Overview

This is the authentication server for the Hermes project. It is a simple REST API that allows users to register and login to the system. 
This server creates and validates JWT tokens for the user.

# Getting Started

## Prerequisites

- [Go](https://golang.org/dl/)

## Installation

### Checkout code from git

```bash
git clone https://github.com/Hermes-Chat-App/hermes-auth-server
```

## Environment Variables

There is a `.env.example` file in the root of the project. Copy this file to `.env` and fill in the values.

| Variable Name | Description                                 |
| ------------- | ------------------------------------------- |
| `PORT`        | Port to run the server on. Defaults to 8080 |

# Build and Run

## Commands

| Command      | Description                    |
| ------------ | ------------------------------ |
| `make build` | Build the binary               |
| `make run`   | Run the binary                 |
| `make tests` | Run the tests                  |
| `make clean` | Clean the build artifacts      |
| `make setup` | Setup local packages           |
| `air`        | Run the server with hot reload |

## Build

```bash
make build
```

## Run

```bash
make run
```

## Dev Run (with hot reload)

### Install Air

```bash
go install github.com/cosmtrek/air@latest
```

### Run

```bash
air
```