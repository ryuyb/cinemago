FROM golang:1.24.2 AS builder

COPY . /src
WORKDIR /src

RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin/ app
COPY --from=builder /src/configs/config.yml data/conf/config.yml

WORKDIR /app

EXPOSE 8000
VOLUME [ "/data/conf" ]

CMD [ "./app", "--env", "prod", "-c", "/data/conf/config.yml" ]
