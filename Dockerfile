FROM golang:latest as app-builder
WORKDIR /go/src/app
COPY . .
#COPY ./static/ /go/src/app/static
COPY ./views/index.html .
#COPY ./templates /go/src/app/templates

RUN echo "Cache break counter: 7"
# Static build required so that we can safely copy the binary over.
# `-tags timetzdata` embeds zone info from the "time/tzdata" package.
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata ./...\

FROM scratch
# the test program:
COPY --from=app-builder /go/bin/cyeam /cyeam
# the tls certificates:
# NB: this pulls directly from the upstream image, which already has ca-certificates:
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /go/src/app/views/ /go/src/app/

CMD ["/cyeam"]
