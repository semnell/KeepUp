# KeepUp
[![Go Report Card](https://goreportcard.com/badge/github.com/semnell/KeepUp)](https://goreportcard.com/report/github.com/semnell/KeepUp)
[![Tests](https://github.com/semnell/KeepUp/actions/workflows/go-test.yml/badge.svg)](https://github.com/semnell/KeepUp/actions/workflows/go-test.yml)
[![Docker Builds](https://github.com/semnell/KeepUp/actions/workflows/docker-image.yml/badge.svg)](https://github.com/semnell/KeepUp/actions/workflows/docker-image.yml)
[![codebeat badge](https://codebeat.co/badges/2d1784dc-7b42-4c49-9082-7d96a034b2e2)](https://codebeat.co/projects/github-com-semnell-keepup-main)
[![CodeFactor](https://www.codefactor.io/repository/github/semnell/keepup/badge/main)](https://www.codefactor.io/repository/github/semnell/keepup/overview/main)

KeepUp is a robust and simple-to-use uptime monitoring tool designed to be both lightweight and scalable, with the capability to distribute tasks across multiple workers for efficient handling of requests. Whether you're monitoring a few endpoints or hundreds, KeepUp aims to provide a straightforward and performant solution.

If you find KeepUp useful, consider giving it a star to show your support.

## Roadmap
- **Setup Helm:** Streamline deployment with Helm charts.
- **Metrics:** Provide more metrics.

## Features
- **Deadly Simple Config Language:** Configure monitoring tasks effortlessly using an intuitive config language.
- **Prometheus Metrics Endpoint:** Integration-friendly metrics endpoint for seamless integration into various monitoring systems.
- **Minimal External Dependencies:** Relies on [Faktory](https://github.com/contribsys/faktory), a single binary job queue, as the only external dependency.
- **Rich Logging:** Detailed logging to keep you informed about the monitoring process.

## License
This tool is released under the MIT license, providing you with the freedom to use, modify, and distribute it as you see fit. Go ahead, monitor the world, or choose not to—it's up to you.

## Is it any good?
We certainly hope so. Your feedback and contributions are always welcome to make KeepUp even better.


## Usage

1. **Download the Latest Release:**
   - Choose the appropriate release for your OS and architecture (e.g., Linux arm64).

2. **Configuration:**
   - Create a `config.yaml` file (refer to the example in the project's root).
   - Set up the `.env` file or configure environment variables in your shell.

3. **Options:**
   - Run the binary in either standalone or distributed mode.

### Standalone Mode
- Execute the binary without any arguments to run in standalone mode.

### Distributed Mode
- The binary supports two modes in a distributed setup: server and worker.

#### Server
- Run the binary with the correct `.env` values and the `server` argument.

#### Worker
- Run the binary with the correct `.env` values and the `worker` argument.

In most setups, it's recommended to run the server on a single machine and distribute workers across multiple machines. Note that running the server multiple times is not allowed due to unique Prometheus metrics and the server's exclusive ability to add new jobs to the queue.