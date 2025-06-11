run:
    #!/bin/bash
    git add -f ./db/ent/schema/*
    go generate ./db/ent
    go mod tidy
    go generate
    go run main.go config.ini