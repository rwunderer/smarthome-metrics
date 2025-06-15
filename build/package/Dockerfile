#--------
# builder
#--------
ARG BUILDPLATFORM
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.23.4-alpine3.19@sha256:5f3336882ad15d10ac1b59fbaba7cb84c35d4623774198b36ae60edeba45fd84 AS builder

ARG TARGETPLATFORM
ARG TARGETOS="linux"
ARG TARGETARCH
ARG TARGETVARIANT

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT#"v"} go build -a -installsuffix cgo -ldflags '-extldflags "-static"' ./cmd/smarthome-metrics

RUN mkdir /conf

COPY examples/config.yaml /conf

#--------
# container
#--------
FROM --platform=${BUILDPLATFORM:-linux/amd64} gcr.io/distroless/static:nonroot@sha256:627d6c5a23ad24e6bdff827f16c7b60e0289029b0c79e9f7ccd54ae3279fb45f AS smarthome-metrics

LABEL version=0.6.6

ENV CONFIG_FILE=/conf/config.yaml

COPY --from=builder /conf /conf

WORKDIR /app

# USER 1002:1002

COPY --from=builder /app/smarthome-metrics /app/smarthome-metrics

ENTRYPOINT ["./smarthome-metrics"]
