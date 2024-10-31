# Stage 1
FROM golang:1.23.2-alpine AS builder 

WORKDIR /

COPY  go.mod go.sum ./
 

RUN go mod download

COPY . .

RUN go build -o bin/job_exc /cmd/main.go


# Stage 2

FROM  alpine:latest 

WORKDIR /root/

COPY --from=builder /bin/job_exc .

# exposing port on which the app will run 

EXPOSE 8080

CMD [ "./bin/job_exc" ]








