
FROM soulteary/milvus-openblas:0.3.20-intel-x86-ubuntu-22.04 AS OpenBLAS

FROM ubuntu:22.04 AS Base
LABEL maintainer=soulteary@gmail.com
# https://soulteary.com/2022/06/21/building-a-cost-effective-linux-learning-environment-on-a-laptop-the-basics.html
RUN sed -i -e "s/archive.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/" /etc/apt/sources.list && \
    sed -i -e "s/security.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/" /etc/apt/sources.list
# https://soulteary.com/2022/07/09/building-a-vector-database-from-scratch-source-code-compilation-and-installation-of-milvus-1.html
RUN apt-get update && apt-get install -y \
    cmake clang-format clang-tidy \
    g++ gcc make lcov libtool m4 autoconf automake ccache libssl-dev zlib1g-dev libboost-regex-dev libboost-program-options-dev libboost-system-dev libboost-filesystem-dev libboost-serialization-dev python3-dev libboost-python-dev libcurl4-openssl-dev gfortran libtbb-dev pkg-config && \
    apt-get remove --purge -y
# https://soulteary.com/2022/07/31/into-vector-computing-making-openblas-docker-prebuilt-product-images.html
COPY --from=OpenBLAS /usr/lib/libopenblas-r0.3.20.so /usr/lib/
RUN ln -s /usr/lib/libopenblas-r0.3.20.so /usr/lib/libopenblas.so.0 && \
    ln -s /usr/lib/libopenblas.so.0 /usr/lib/libopenblas.so
# https://soulteary.com/2022/07/04/build-a-maintainable-golang-development-environment.html
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN apt-get install -y binutils bison gcc make git curl && apt-get remove --purge -y
RUN git clone --branch master --depth 1 https://gitcode.net/soulteary/gvm.git && \
    SRC_REPO=https://gitcode.net/soulteary/gvm.git bash gvm/binscripts/gvm-installer
ENV GO_BINARY_BASE_URL=https://mirrors.aliyun.com/golang/
ENV GVM_ROOT=/root/.gvm
ENV GOROOT_BOOTSTRAP=$GOROOT
RUN source "$HOME/.gvm/scripts/gvm" && \
    gvm install go1.19.1 -B && gvm use go1.19.1 --default
ENV GO111MODULE=on
ENV GOPATH="$HOME/go"
ENV PATH="$GOPATH/bin:$PATH"
ENV GOPROXY="https://goproxy.cn,direct"