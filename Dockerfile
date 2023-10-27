FROM alpine:latest

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app

COPY ./bin/main /app/main

RUN chmod -R 755 /app && chown -R appuser:appuser /app

USER appuser:appuser

CMD [ "/app/main" ]