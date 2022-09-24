# syntax=docker/dockerfile:1.3-labs
FROM promtail:latest

ENV JOB=app-1
ENV APP=app-test
ENV INSTANCE=app-test-1

COPY supervisor/app.conf /etc/supervisor/conf.d/app.conf
COPY app /app/app