FROM gcr.io/distroless/base:latest
LABEL maintainer="Jonathan Wright <jon@than.io>" \
  org.opencontainers.image.url="https://github.com/n3tuk/action-pull-requester" \
  org.opencontainers.image.source="https://github.com/n3tuk/action-pull-requester/blob/Dockerfile" \
  org.opencontainers.image.title="pull-requester" \
  org.opencontainers.image.description="A GitHub Action for checking pull requests" \
  org.opencontainers.image.authors="Jonathan Wright <jon@than.io>" \
  org.opencontainers.image.vendor="n3t.uk"

COPY pull-requester /go/bin/pull

ENTRYPOINT ["/go/bin/pull-requester"]
CMD ["run"]
