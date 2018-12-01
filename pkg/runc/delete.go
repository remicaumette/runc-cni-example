package runc

type DeleteOps struct {
	Force	bool
}

func (client *Client) Delete(id string, opts *DeleteOps) error {
	args := []string{}
}
