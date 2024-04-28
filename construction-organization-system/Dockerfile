FROM golang:1.22.1-alpine  as builder
LABEL authors="ykhdr"

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd cmd
COPY internal internal
COPY pkg pkg
RUN go build -o construction-ogranization-system cmd/main.go

FROM alpine:3.19.1
COPY --from=builder /build/construction-ogranization-system /bin/construction-ogranization-system
ENTRYPOINT ["/bin/construction-ogranization-system"]