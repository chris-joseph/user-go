FROM golang:1.18-stretch as builder
RUN mkdir /build
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN make build

# generate clean, final image
FROM scratch
WORKDIR /build
COPY --from=builder /build/bin/sc-integrations-orders-api ./app
CMD [ "/build/app" ]
