ARG FEDORA_VERSION=36
ARG FEDORA_ARCH=x86_64
FROM registry.fedoraproject.org/fedora:$FEDORA_VERSION-$FEDORA_ARCH

#VSCode Dependencies
RUN dnf install -y --refresh glibc libgcc libstdc++ python ca-certificates tar wget tree vim make \ 
    && dnf clean all

WORKDIR /src
ARG GO_VERSION=1.19.4
RUN wget -q "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" -O /src/go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm /src/go.tar.gz \
    && echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bashrc \
    && echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc \
    && echo 'export PATH=$PATH:$GOPATH/bin' >> $HOME/.bashrc \
    && /usr/local/go/bin/go install golang.org/x/tools/cmd/goimports@latest \
    && /usr/local/go/bin/go install golang.org/x/tools/gopls@latest \
    && /usr/local/go/bin/go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

