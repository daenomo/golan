FROM ubuntu:24.04

# avoid interactive mode when installing packages
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y software-properties-common ca-certificates && \
    apt-get install -y sudo build-essential libgmp-dev git wget unzip which tar curl golang && \
    apt-get clean

ARG USERNAME

RUN groupadd $USERNAME && \
    useradd -m -s /bin/bash -g $USERNAME -G sudo $USERNAME && \
    echo "$USERNAME   ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
USER $USERNAME
WORKDIR /home/$USERNAME/app