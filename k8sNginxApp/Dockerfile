# Specifies a parent image
FROM golang:bookworm

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod tidy

# Builds your app with optional configuration
RUN go build -o /k8sApp

# Tells Docker which network port your container listens on
EXPOSE 8080

# Specifies the executable command that runs when the container starts
CMD [ "/k8sApp" ]