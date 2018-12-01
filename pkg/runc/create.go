package runc

type CreateOpts struct {
	Bundle	string
	ConsoleSocket string
	PidFile string
	NoPivot bool
	NoNewKeyring bool
}

func (client *Client) Create(id string, opts *CreateOpts) error {
	args := []string{"create"}
	//if opts != nil {
	//	if opts.Bundle != "" {
	//		path, err := filepath.Abs(opts.Bundle)
	//		if err != nil {
	//			return err
	//		}
	//		args = append(args, "--bundle", path)
	//	}
	//	if opts.ConsoleSocket != "" {
	//		path, err := filepath.Abs(opts.ConsoleSocket)
	//		if err != nil {
	//			return err
	//		}
	//		args = append(args, "--console-socket", path)
	//	}
	//	if opts.PidFile != "" {
	//		path, err := filepath.Abs(opts.PidFile)
	//		if err != nil {
	//			return err
	//		}
	//		args = append(args, "--pid-file", path)
	//	}
	//	if opts.NoPivot {
	//		args = append(args, "--no-pivot")
	//	}
	//	if opts.NoNewKeyring {
	//		args = append(args, "--no-new-keyring")
	//	}
	//}
	args = append(args, id)
	_, err := client.runCommand(args...)
	return err
}
