FROM docker.io/library/golang:1.12-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -a -o /podinfo .

FROM scratch
ENV PORT=8080
COPY --from=build /podinfo /podinfo
CMD [ "/podinfo" ]
