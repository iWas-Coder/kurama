FROM golang as builder
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
WORKDIR /kurama
COPY . .
RUN go build -v -a

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /kurama/kurama .
USER 65532:65532
ENTRYPOINT ["/kurama"]
