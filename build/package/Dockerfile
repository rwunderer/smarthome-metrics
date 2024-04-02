#--------
# builder
#--------
ARG BUILDPLATFORM
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.22.1-alpine3.19@sha256:0466223b8544fb7d4ff04748acc4d75a608234bf4e79563bff208d2060c0dd79 as builder

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
FROM --platform=${BUILDPLATFORM:-linux/amd64} gcr.io/distroless/static:nonroot@sha256:f41b84cda410b05cc690c2e33d1973a31c6165a2721e2b5343aab50fecb63441

LABEL version=0.5.1

ENV CONFIG_FILE=/conf/config.yaml

COPY --from=builder /conf /conf

WORKDIR /app

# USER 1002:1002

COPY --from=builder /app/smarthome-metrics /app/smarthome-metrics

ENTRYPOINT ["./smarthome-metrics"]
