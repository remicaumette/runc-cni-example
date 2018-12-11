#!/usr/bin/env bash
apt-get update
apt-get upgrade -y
apt-get install -y unzip git build-essential openjdk-8-jdk

wget -q https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
tar xvf go1.11.2.linux-amd64.tar.gz -C /usr/local

wget -q https://github.com/opencontainers/runc/releases/download/v1.0.0-rc6/runc.amd64 -O /usr/local/bin/runc
chmod a+x /usr/local/bin/runc

wget -q https://storage.googleapis.com/gvisor/releases/nightly/latest/runsc -O /usr/local/bin/runsc
chmod a+x /usr/local/bin/runsc

wget -q https://github.com/containernetworking/cni/releases/download/v0.6.0/cni-amd64-v0.6.0.tgz
tar xvf cni-amd64-v0.6.0.tgz -C /usr/local/bin

wget -q https://github.com/containernetworking/plugins/releases/download/v0.7.4/cni-plugins-amd64-v0.7.4.tgz
mkdir -p /opt/cni/plugins
tar xvf cni-plugins-amd64-v0.7.4.tgz -C /opt/cni/plugins

wget -q https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
unzip -o protoc-3.3.0-linux-x86_64.zip -d /usr/local bin/protoc

mkdir -p /etc/cni/net.d

echo "export GOPATH=~/go" >> /home/vagrant/.bashrc
echo "export GOROOT=/usr/local/go" >> /home/vagrant/.bashrc
echo "export PATH=\"\$GOPATH/bin:\$PATH\"" >> /home/vagrant/.bashrc
echo "export PATH=\"\$GOROOT/bin:\$PATH\"" >> /home/vagrant/.bashrc
echo "export GO111MODULE=on" >> /home/vagrant/.bashrc
