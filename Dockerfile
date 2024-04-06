FROM golang:1.22.1 as development

ARG UID=1000
ARG GID=1001
ENV CODE_DIR=/src

USER root

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.1

RUN groupadd -g $GID gophers
RUN useradd -rm -d /home/gopher -g $GID -u $UID -s /bin/bash gopher

ENV CODE_DIR=/src
WORKDIR $CODE_DIR

COPY .$CODE_DIR $CODE_DIR
RUN chown -R gopher:gophers $CODE_DIR
RUN chmod -R 775 ${CODE_DIR}

USER gopher

### Step 1: Build stage
# FROM golang:1.22.1-alpine as builder

# WORKDIR /app

# # Copy the Go module files and download dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the application source code and build the binary
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

# ### 
# ## Step 2: Runtime stage
# FROM scratch

# # Copy only the binary from the build stage to the final image
# COPY --from=builder /app/myapp /

# # Set the entry point for the container
# ENTRYPOINT ["/myapp"]
