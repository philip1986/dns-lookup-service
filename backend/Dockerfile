FROM golang:1.17

# Set destination for COPY
WORKDIR /app

ENV GOPROXY direct

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# add app
COPY . ./

# Build
RUN make clean
RUN make test_and_build

# Run
CMD [ "./dns-lookup-service" ]