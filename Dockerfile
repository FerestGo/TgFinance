FROM alpine

RUN apk --no-cache add ca-certificates
RUN mkdir -p /opt/tsm && mkdir -p /etc/tsm
COPY TgFinanceh /opt/tsm/
