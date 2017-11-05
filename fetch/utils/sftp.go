package utils

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	pb "gopkg.in/cheggaaa/pb.v1"

	"github.com/pkg/sftp"

	"golang.org/x/crypto/ssh"
)

func sftpConnect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}

func SftpGet(username, password, ip, remoteFilePath, localFilePath string) {
	var (
		err        error
		sftpClient *sftp.Client
	)

	sftpClient, err = sftpConnect(username, password, ip, 22)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	srcFile, err := sftpClient.Open(remoteFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(localFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()
	srcFileInfo, _ := srcFile.Stat()
	bar := pb.New(int(srcFileInfo.Size())).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = true
	bar.ShowTimeLeft = true
	bar.ShowFinalTime = true
	bar.SetMaxWidth(80)
	bar.Format("[=> ]")
	bar.Prefix("Download:")
	bar.Start()
	writer := io.MultiWriter(dstFile, bar)
	if _, err = srcFile.WriteTo(writer); err != nil {
		log.Fatal(err)
	}
	bar.Finish()
}
