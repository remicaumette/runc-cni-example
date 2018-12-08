NS=$(cat /var/run/runc/hello/state.json | jq '.namespace_paths.NEWNET' -r)
mkdir -p /var/run/netns
ln -sf $NS /var/run/netns/hello

CNI_PATH=/opt/cni/
CNI_CONTAINERID=hello
CNI_NETNS=/var/run/netns/hello
CNI_IFNAME=eth0
CNI_COMMAND=ADD

cat /home/vagrant/go/src/gitlab.com/expected.sh/agent/net/expected.conf | $CNI_PATH/bridge
