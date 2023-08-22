package remote

import (
	"net/http"
	"os/exec"

	"github.com/docker/cli/cli/connhelper"
	"github.com/docker/docker/client"
)

func Exec(host string, cmd string) ([]byte, error) {
	return exec.Command("ssh", host, cmd).CombinedOutput()
}

func (dockerSSHClient DockerSSHClient) Client() *client.Client {
	helper, err := connhelper.GetConnectionHelper(dockerSSHClient.Host)
	if err != nil {
		panic(err)
	}
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: helper.Dialer,
		},
	}
	docker, err := client.NewClientWithOpts(
		client.WithHTTPClient(httpClient),
		client.WithHost(helper.Host),
		client.WithDialContext(helper.Dialer),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	return docker
}

type DockerSSHClient struct {
	Host string
}
