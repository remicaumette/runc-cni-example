Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/trusty64"
    config.vm.box_check_update = true

    config.vm.network "private_network", type: "dhcp"

    config.vm.synced_folder ".", "/home/vagrant/go/src/gitlab.com/expected.sh/agent"
    config.vm.synced_folder "./scripts/net.d", "/etc/cni/net.d"

    config.vm.provider "virtualbox" do |vb|
        vb.memory = "1024"
    end

    config.vm.provision "shell", path: "./scripts/provisioning.sh"
end
