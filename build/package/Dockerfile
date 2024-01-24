#--------
# builder
#--------
ARG BUILDPLATFORM
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.21.6-alpine3.19@sha256:51a7800206bc7b276a9d62a7229cdede7b1e0f45ec28259ed44c1603c6cda1e7 as builder

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
FROM --platform=${BUILDPLATFORM:-linux/amd64} gcr.io/distroless/static:nonroot@sha256:112a87f19e83c83711cc81ce8ed0b4d79acd65789682a6a272df57c4a0858534

ENV CONFIG_FILE=/conf/config.yaml

COPY --from=builder /conf /conf

WORKDIR /app

# USER 1002:1002

COPY --from=builder /app/smarthome-metrics /app/smarthome-metrics

ENTRYPOINT ["./smarthome-metrics"]
