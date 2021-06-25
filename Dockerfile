FROM golang:1.13.5

WORKDIR /app/forum

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN make

CMD ["./forum"]