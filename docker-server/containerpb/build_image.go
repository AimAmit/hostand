package containerpb

import (
	"bufio"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
	"log"
)


func (c *DockerClient)BuildImage(domain, version string,dockerBuildContext io.Reader) {

	options := types.ImageBuildOptions{
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		Tags:           []string{fmt.Sprintf("%s:%s",domain, version)},
		Memory: 		1.28e+8,
		MemorySwap: 	2.56e+8,
		Dockerfile:     "Dockerfile",
	}
	buildResponse, err := c.Client.ImageBuild(context.Background(), dockerBuildContext, options)
	if err != nil {
		log.Println(err, " :unable to build remote image")
	}
	defer func() {
		if err = buildResponse.Body.Close(); err != nil {
			log.Println("Unable to close build response: ", err)
		}
	}()

	rd := bufio.NewReader(buildResponse.Body)
	for {
		n, _, err := rd.ReadLine()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(n))
	}
}