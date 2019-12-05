# Certs container
FROM alpine:latest as certs

# Run container
FROM scratch

COPY kage kage
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

ENV KAGE_SERVER "true"
ENV PORT "80"

EXPOSE 80
CMD ["./kage", "agent"]