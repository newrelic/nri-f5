FROM golang:1.10 as builder
RUN go get -d github.com/newrelic/nri-f5/... && \
    cd /go/src/github.com/newrelic/nri-f5 && \
    make && \
    strip ./bin/nr-f5

FROM newrelic/infrastructure:latest
ENV NRIA_IS_FORWARD_ONLY true
ENV NRIA_K8S_INTEGRATION true
COPY --from=builder /go/src/github.com/newrelic/nri-f5/bin/nr-f5 /var/db/newrelic-infra/newrelic-integrations/bin/nr-f5
COPY --from=builder /go/src/github.com/newrelic/nri-f5/f5-definition.yml /var/db/newrelic-infra/newrelic-integrations/definition.yml
