package runc

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Client struct {
	Command	string
}

func New(command string) Client {
	return Client{
		Command: command,
	}
}

func (client *Client) runCommand(args ...string) (string, error) {
	outReader, outWriter, err := os.Pipe()
	if err != nil {
		return "", err
	}
	errReader, errWriter, err := os.Pipe()
	if err != nil {
		return "", err
	}
	inWriter, err := os.OpenFile(os.DevNull, os.O_RDWR, 0)

	files := []*os.File{inWriter, outWriter, errWriter}
	process, err := os.StartProcess(client.Command, append([]string{client.Command}, args...), &os.ProcAttr{
		Env: os.Environ(),
		Files: files,
		Sys: nil,
	})

	state, err := process.Wait()

	fmt.Println("%v %v %v\n", state.Exited(), state.UserTime(),  err)

	outWriter.Close()
	errWriter.Close()
	loly, err := ioutil.ReadAll(ioutil.NopCloser(outReader))
	fmt.Println(err)
	lolz, err := ioutil.ReadAll(ioutil.NopCloser(errReader))
	fmt.Println(err)
	fmt.Println(string(loly), string(lolz))
	/*if statusCode == 0 {
		return string(loly), nil
	}*/
	return "", nil
}
