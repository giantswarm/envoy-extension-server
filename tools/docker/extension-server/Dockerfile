FROM gcr.io/distroless/static:nonroot@sha256:f512d819b8f109f2375e8b51d8cfd8aafe81034bc3e319740128b7d7f70d5036
COPY ./envoy-extension-server-linux-amd64 /usr/local/bin/envoy-extension-server

USER 65532:65532

ENTRYPOINT ["/usr/local/bin/envoy-extension-server"]
