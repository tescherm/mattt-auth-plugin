FROM alpine:3.4

RUN mkdir -p /run/docker/plugins

COPY mattt-auth-plugin mattt-auth-plugin

CMD ["mattt-auth-plugin"]