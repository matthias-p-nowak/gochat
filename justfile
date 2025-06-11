run:
    #!/bin/bash
    go mod tidy
    go generate
    go run main.go