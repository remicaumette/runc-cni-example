package runc

type DeleteOps struct {
	Force	bool
}

func (client *Client) Delete(id string, opts *DeleteOps) error {
	args := []string{"delete"}
	if opts != nil {
		if opts.Force {
			args = append(args, "--force")
		}
	}
	args = append(args, id)
	if _, err := client.runCommand(args...); err != nil {
		return err
	}
	return nil
}
