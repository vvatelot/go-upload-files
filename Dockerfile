FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main .
RUN touch .env
ADD public public
ADD templates templates
EXPOSE 8080
CMD ["./main"] 