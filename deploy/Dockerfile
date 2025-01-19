FROM alpine:latest

COPY go-sister-design-site /go-sister-design-site
COPY ./configs /configs

RUN apk --no-cache add openssl curl tzdata && \
    chmod a+x /go-sister-design-site

CMD [ "/go-sister-design-site" ]
