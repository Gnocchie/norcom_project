# Go File Scanner with RabbitMQ

This program recursively scans a local directory and sends a JSON message to a RabbitMQ queue for each file it finds.

## Features

- Recursively scans directories
- Captires file metadata such as file path, file size, last modified and hostname
- Publishes file events to RabbitMQ queue
- Supports configs cia enviroment variables

## Requirements

- Go 1.22 or better
- RabbitMQ
- Docker (for local testing)

## Setup
1. Clone the repository
2. Install required dependencies
3. Start RabbitMQ using Docker
4. Configure RabbitMQ environment variables

## Running the program
- Run with ```go run ./cmd/scanner```

Expected output:
```
2025/12/05 23:55:00 Config loaded as: host=localhost:5672 queue=file_events root=/tmp/test_files
2025/12/05 23:55:00 Starting file scan in /tmp/test_files
2025/12/05 23:55:00 Sent event for /tmp/test_files/file1.txt
...
2025/12/05 23:55:05 Scan completed successfully.
```

- Verify messages in RabbitMQ

Expected output:
```
{
  "path": "/tmp/test_files/file1.txt",
  "size": 1024,
  "modified": "2025-12-05T23:55:00Z",
  "hostname": "myComputer"
}
```

