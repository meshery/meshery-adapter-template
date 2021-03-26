FROM golang:1.15 as builder
# Replace <adapter-name> to what you want to create
# Replace <service-mesh> to the service-mesh application name
ARG ENVIRONMENT="development"
ARG CONFIG_PROVIDER="viper"
WORKDIR /build
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download
# Copy the go source
COPY main.go main.go
COPY internal/ internal/
COPY <service-mesh>/ <service-mesh>/
# Build
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags="-w -s -X main.environment=$ENVIRONMENT -X main.provider=$CONFIG_PROVIDER" -a -o <adapter-name> main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/base:nonroot-amd64
ENV DISTRO="debian"
ENV GOARCH="amd64"
WORKDIR /templates
COPY templates/* .
WORKDIR /
COPY --from=builder /build/<adapter-name> .
ENTRYPOINT ["/<adapter-name>"]