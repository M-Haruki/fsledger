FROM debian:13-slim

RUN apt-get update && apt-get install -y git curl ca-certificates tar xz-utils make libatomic1
# libatomic1 for pnpm

# Go
ENV GO_VERSION=1.26.5
RUN curl -LO https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
  rm -rf /usr/local/go && \
  tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
  rm go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:/usr/local/node/bin:${PATH}"

# Node
ENV NODE_VERSION=24.19.0
RUN curl -LO https://nodejs.org/dist/v${NODE_VERSION}/node-v${NODE_VERSION}-linux-x64.tar.xz && \
  tar -xJf node-v${NODE_VERSION}-linux-x64.tar.xz && \
  rm -rf /usr/local/node && \
  mv node-v${NODE_VERSION}-linux-x64 /usr/local/node && \
  rm node-v${NODE_VERSION}-linux-x64.tar.xz

# pnpm
ENV PNPM_VERSION=11.20.0
ENV PNPM_HOME="/root/.local/share/pnpm"
ENV PATH="${PNPM_HOME}/bin:${PATH}"
RUN curl -fsSL https://get.pnpm.io/install.sh | env PNPM_VERSION=${PNPM_VERSION} SHELL=/bin/bash sh -
RUN pnpm config set store-dir /root/.local/share/pnpm/store --global

WORKDIR /workspace
