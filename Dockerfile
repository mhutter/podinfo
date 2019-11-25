FROM docker.io/library/golang:1.13-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -a -o /podinfo .

FROM scratch
ENV PORT=8080
EXPOSE $PORT
COPY --from=build /podinfo /podinfo
CMD [ "/podinfo" ]
