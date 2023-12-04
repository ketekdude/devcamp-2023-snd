# Backend Shipper Service for Tokopedia Devcamp 2023 Server and Database

You can run the service using docker detached option

```shell
docker-compose up -d
```

If you want to rebuild the service, run this command

```shell
make build
```

If you want to run the service, run this command

```shell
make run
```

## Code Structure

Code structure for shipper service.

```
service
 ├── database                # Database Initialization Configuration
 ├── model                   # Entity Definition
 ├── server                  # Server Initialization Configuration
        └─ handlers          # gRPC Handlers
              └─ shipper     # gRPC Handler for shipper service
                   └─ proto  # gRPC Protobuf for shipper service
 ├── shippermodule           # Business Logic
 ├── main.go                 # Service Initalization
 └── Dockerfile              # Dockerfile to build the database image
```
