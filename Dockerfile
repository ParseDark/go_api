FROM golang:1.13 as build

ENV GOPROXY="https://goproxy.io"
# https://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host
# build for static link
ENV CGO_ENABLED=0
WORKDIR /api
COPY . /api
RUN make build

# production stage
FROM alpine as production

WORKDIR /api
COPY ./conf/ /api/conf
COPY --from=build /api/web /api
EXPOSE 8081
ENTRYPOINT ["/api/web"]
CMD [ "-c", "./conf/config_docker.yaml" ]
