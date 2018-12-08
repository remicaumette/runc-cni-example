Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/xenial64"
    config.vm.box_check_update = false
    config.vm.network "private_network", type: "dhcp"

    config.vm.synced_folder ".", "/home/vagrant/go/src/gitlab.com/expected.sh/agent"

    config.vm.provider "virtualbox" do |vb|
        vb.memory = "1024"
    end

    config.vm.provision "shell", inline: <<-SHELL
        apt-get update
        apt-get upgrade -y
        apt-get install -y unzip git build-essential
        wget https://github.com/containerd/containerd/releases/download/v1.2.0/containerd-1.2.0.linux-amd64.tar.gz
        tar xvf containerd-1.2.0.linux-amd64.tar.gz -C /usr/local
        wget https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
        tar xvf go1.11.2.linux-amd64.tar.gz -C /usr/local
        wget https://github.com/opencontainers/runc/releases/download/v1.0.0-rc6/runc.amd64 -O /usr/local/bin/runc
        chmod a+x /usr/local/bin/runc
        wget https://storage.googleapis.com/gvisor/releases/nightly/latest/runsc -O /usr/local/bin/runsc
        chmod a+x /usr/local/bin/runsc
        wget https://github.com/containernetworking/cni/releases/download/v0.6.0/cni-amd64-v0.6.0.tgz
        tar xvf cni-amd64-v0.6.0.tgz -C /usr/local/bin
        mkdir /etc/containerd/
        containerd config default > /etc/containerd/config.toml
        mkdir /opt/cni
        wget https://github.com/containernetworking/plugins/releases/download/v0.7.4/cni-plugins-amd64-v0.7.4.tgz
        tar xvf cni-plugins-amd64-v0.7.4.tgz -C /opt/cni
        mkdir -p /etc/cni/net.d
        bash -c 'cat >/etc/cni/net.d/10-containerd-net.conflist <<EOF
        {
          "cniVersion": "0.3.1",
          "name": "containerd-net",
          "plugins": [
            {
              "type": "bridge",
              "bridge": "cni0",
              "isGateway": true,
              "ipMasq": true,
              "promiscMode": true,
              "ipam": {
                "type": "host-local",
                "subnet": "10.88.0.0/16",
                "routes": [
                  { "dst": "0.0.0.0/0" }
                ]
              }
            },
            {
              "type": "portmap",
              "capabilities": {"portMappings": true}
            }
          ]
        }
        EOF'
        wget https://raw.githubusercontent.com/containerd/containerd/master/containerd.service -O /etc/systemd/system/containerd.service
        systemctl enable containerd
        systemctl start containerd
        echo "export GOPATH=~/go" >> /home/vagrant/.bashrc
        echo "export GOROOT=/usr/local/go" >> /home/vagrant/.bashrc
        echo "export PATH=\"\$GOPATH/bin:\$PATH\"" >> /home/vagrant/.bashrc
        echo "export PATH=\"\$GOROOT/bin:\$PATH\"" >> /home/vagrant/.bashrc
        echo "export GO111MODULE=on" >> /home/vagrant/.bashrc
    SHELL
end
