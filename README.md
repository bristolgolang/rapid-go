# Rapid Go

This is a demo project doing simple things with Tilt in Go to achieve rapid development in containers.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Tilt](https://docs.tilt.dev/install.html)
- [Go](https://golang.org/doc/install)
- [Cue](https://cuelang.org)

## Getting Started

1. Clone the repository
2. Run `git checkout step1` to get to the first step
3. Run `tilt up` in the root of the repository
4. Open your browser to `http://localhost:10350` to see the TILT UI
5. Curl to `http://localhost:32400` to hit the api (available routes are `/` and `/ready`)

## How to Use

Try making changes to the codebase and see tilt rebuild the container and update the running container.
Some steps have already been outlined in the repository, but feel free to experiment with the codebase.
To see a diff between steps run `git diff step1 step2` or `git diff step2 step3` etc.

## Steps

### Step 1

Creates the basic structure for us to use.
A simple Go API, dockerfile, and tiltfile.

### Step 2

Adds caching to the dockerfile (particularly in the go build step)

### Step 3

Expands the project to use go 1.22's new routing to provide a better structure going forwards.

### Step 4

Moves the `/` endpoint to `/greet/{name}` and greets the user with a message.

### Step 5

Adds CI to ensure the Tiltfile workflow is working as expected.