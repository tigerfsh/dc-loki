# syntax=docker/dockerfile:1.3-labs
FROM promtail:latest

# RUN curl -LO https://go.dev/dl/go1.18.4.linux-arm64.tar.gz

# RUN  tar -C /usr/local -xvzf go1.18.4.linux-arm64.tar.gz && rm -f go1.18.4.linux-arm64.tar.gz


# RUN <<EOF cat >> ~/.bashrc
# export PATH="/usr/local/go/bin:$PATH"
# EOF

COPY supervisor/app.conf /etc/supervisor/conf.d/app.conf
COPY app /app/app