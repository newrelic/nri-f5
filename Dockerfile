FROM golang:1.10 as builder
COPY . /go/src/github.com/newrelic/nri-f5/
RUN cd /go/src/github.com/newrelic/nri-f5 && \
    make && \
    strip ./bin/nri-f5

FROM newrelic/infrastructure:latest
ENV NRIA_IS_FORWARD_ONLY true
ENV NRIA_K8S_INTEGRATION true
COPY --from=builder /go/src/github.com/newrelic/nri-f5/bin/nri-f5 /nri-sidecar/newrelic-infra/newrelic-integrations/bin/nri-f5
USER 1000
