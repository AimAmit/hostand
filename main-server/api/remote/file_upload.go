package remote

import (
	"bufio"
	"context"
	"errors"
	"github.com/aimamit/hostand/main-server/proto"
	"io"
	"log"
)

func (c ClientCall) FileUpload(file io.Reader, domain, version string) error {
	cli := c.Docker

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	
	stream, err := cli.FileUpload(context.Background())
	if err != nil {
		log.Printf("Error calling file upload api: %v\n", err)
		return err
	}

	err = stream.Send(&proto.FileRequest{
		Data: &proto.FileRequest_AppVersion{
			AppVersion: &proto.AppVersion{
				Domain: domain,
				Version: version,
			},
		},
	})

	if err != nil {
		log.Println("cannot send file appInfo to server: ", err, stream.RecvMsg(nil))
		return err
	}

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		req := &proto.FileRequest{
			Data: &proto.FileRequest_Chunk{
				Chunk: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Println("cannot send chunk to server: ", err, stream.RecvMsg(nil))
			return err
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("cannot receive response: ", err)
		return err
	}
	if res.Error != "" {
		log.Println("cannot receive response: ", res.Error)
		return errors.New(res.Error)
	} else {
		log.Println("File sent successfully")
	}
	return nil
}