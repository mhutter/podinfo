FROM registry.mhnet.dev/library/golang:1.17-alpine AS build
WORKDIR /src/podinfo
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s" -o /podinfo .

FROM registry.mhnet.dev/library/busybox:latest AS runtime
ENV PORT=8080
EXPOSE $PORT
COPY --from=build /podinfo /usr/local/bin/podinfo
CMD [ "podinfo" ]
