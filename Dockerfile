FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /enterprise-licenses-report

# Run
CMD ["/enterprise-licenses-report", "generate-report"]