# Image oficial golang
FROM golang:stretch
# create directory
RUN mkdir /app
# left source, right destination
ADD . /app
# Directory on the that working
WORKDIR /app
# Add this go mod download command to pull in any dependencies
RUN go mod download
# Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
CMD ["go", "run", "main.go"]