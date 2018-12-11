package server

import "gitlab.com/expected.sh/agent/pkg/protocol"

type Container struct {
	Name    string
	Root    string
	Running bool
	Pid     uint32
	Network struct {
		Ip      string
		Gateway string
		Mask    string
	}
}

func (container *Container) ToFunction() *protocol.Function {
	return &protocol.Function{
		Name:    container.Name,
		Root:    container.Root,
		Running: container.Running,
		Pid:     container.Pid,
		Network: &protocol.Function_Network{
			Ip:      container.Network.Ip,
			Gateway: container.Network.Gateway,
			Mask:    container.Network.Mask,
		},
	}
}

func (server *Server) CreateContainer(name string) (container *Container, err error) {
	return &Container{}, nil
}
