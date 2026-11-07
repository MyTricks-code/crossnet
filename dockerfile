FROM golang:1.27
WORKDIR /app
COPY . .
EXPOSE 8080
WORKDIR /app/main
CMD ["go", "run", "."]

