#--------
# builder
#--------
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -installsuffix cgo -ldflags '-extldflags "-static"' ./cmd/smarthome-metrics

RUN mkdir /conf

COPY examples/config.yaml /conf

#--------
# container
#--------
FROM --platform=${BUILDPLATFORM:-linux/amd64} gcr.io/distroless/static:nonroot

ENV CONFIG_FILE=/conf/config.yaml

COPY --from=builder /conf /conf

WORKDIR /app

USER 1002:1002

COPY --from=builder /app/smarthome-metrics /app/smarthome-metrics

ENTRYPOINT ["./smarthome-metrics"]
