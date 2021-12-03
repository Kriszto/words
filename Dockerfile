ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS builder

ARG APP_VERSION
ARG APP_NAME

WORKDIR /app

ADD go.mod go.sum /app/
RUN go mod download

ADD . /app
RUN export ROOT_PACKAGE="$(go list -m)" \
    && go build \
    -mod=readonly \
    -ldflags "-X $ROOT_PACKAGE/internal/buildinfo.Version=$APP_VERSION" \
    -ldflags "-X $ROOT_PACKAGE/internal/buildinfo.BuildTime=$(date +%FT%T%z)" \
    -ldflags "-X $ROOT_PACKAGE/internal/buildinfo.AppName=$APP_NAME" \
    -o "/tmp/$APP_NAME"



FROM alpine:latest

ARG APP_NAME
ENV APP_NAME ${APP_NAME}

ARG RUNTIME_PACKAGES=""

WORKDIR /usr/local/bin/

COPY --from=builder /tmp/$APP_NAME ./$APP_NAME
RUN chmod +x ./$APP_NAME

CMD ["sh", "-c", "exec $APP_NAME"]
