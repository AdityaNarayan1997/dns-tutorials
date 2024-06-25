# Latest golang image
FROM golang:latest

# Work directory
WORKDIR /docker-go

# Copying all the files
COPY . .

# Starting our application
CMD ["go", "run", "main.go"]

# Exposing server port
EXPOSE 8080