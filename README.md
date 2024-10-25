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

## Setup Instructions

1. **Clone the Repository:**
   ```sh
   git clone https://github.com/semnell/KeepUp.git
   cd KeepUp
   ```

2. **Install Dependencies:**
   ```sh
   go mod download
   ```

3. **Build the Application:**
   ```sh
   make build
   ```

4. **Run the Application:**
   ```sh
   make run
   ```

## Configuration Options and Environment Variables

The application can be configured using a `config.yaml` file and environment variables. Below are the available options:

### `config.yaml`

```yaml
version: v1 # Version of the config
jobs:
  - name: self # Name of the job
    scheme: http # Type of the job
    url: 127.0.0.1:8080/status # URL of the job
    interval: 1 # Interval of the job
    timeout: 5 # Timeout of the job
    headers: # Headers of the job
      - key: Content-Type
        value: application/json
    method: GET # Method of the job
    expect: # Expect of the job
      status: 200
      contains:
        - "ok"
```

### Environment Variables

| Variable               | Description                          | Default Value               |
|------------------------|--------------------------------------|-----------------------------|
| `FAKTORY_URL`          | URL of the Faktory server            | `http://yourFaktoryAdress:7419` |
| `CONFIG_FILE_PATH`     | Path to the configuration file       | `./config.yaml`             |
| `SERVER_PORT`          | Port for the server to listen on     | `8080`                      |
| `WORKER_CONCURRENCY`   | Number of concurrent workers         | `20`                        |
| `JOB_QUEUE_NAME`       | Name of the job queue                | `keepup`                    |
| `SERVER_CALLBACK_URL`  | Callback URL for the server          | `http://127.0.0.1:8080/callback` |
| `GIN_MODE`             | Mode for the Gin framework           | `release`                   |

## Testing and Running the Application

### Running Tests

To run the tests, use the following command:

```sh
make test
```

### Running the Application

To run the application, use the following command:

```sh
make run
```

You can also run the server and worker separately:

#### Running the Server

```sh
make run-server-standalone
```

#### Running the Worker

```sh
make run-worker-standalone
```

## Contribution Guidelines

We welcome contributions to improve KeepUp. To contribute, please follow these guidelines:

1. **Fork the Repository:**
   - Click the "Fork" button at the top right of the repository page.

2. **Clone Your Fork:**
   ```sh
   git clone https://github.com/your-username/KeepUp.git
   cd KeepUp
   ```

3. **Create a New Branch:**
   ```sh
   git checkout -b feature/your-feature-name
   ```

4. **Make Your Changes:**
   - Ensure your code follows the project's coding style and conventions.
   - Add tests for your changes if applicable.

5. **Commit Your Changes:**
   ```sh
   git commit -m "Add your commit message here"
   ```

6. **Push Your Changes:**
   ```sh
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request:**
   - Go to the repository page on GitHub and click the "New Pull Request" button.
   - Provide a clear description of your changes and any related issues.

Thank you for contributing to KeepUp!
