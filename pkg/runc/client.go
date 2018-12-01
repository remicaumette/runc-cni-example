package runc

type Client struct {
	Command	string
}

func New(command string) Client {
	return Client{
		Command: command,
	}
}
