package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/AEGQ/tools/fetch/utils"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/urfave/cli"
)

func getFetchFlags() []cli.Flag {
	flag := []cli.Flag{
		cli.BoolFlag{
			Name:  "tar,t",
			Usage: "[Optional] fetch an image tar package from other host.",
		},
	}

	return flag
}

func fetchCmd(c *cli.Context) {
	tar := c.Bool("tar")
	username, password, ip, image := parseFlag(c, c.Args().First())

	//docker_image_tag_***_.tar
	tarFile := "docker_" +
		strings.Replace(strings.Replace(image, ":", "_", -1), "/", "_", -1) + "_" +
		strconv.FormatInt(time.Now().UTC().UnixNano(), 10)[9:] + ".tar"

	//r_docker_image_tag_***_.tar
	remoteFilePath := filepath.Join(os.TempDir(), "r_"+tarFile)
	//l_docker_image_tag_***_.tar
	localFilePath := filepath.Join(os.TempDir(), "l_"+tarFile)

	//docker save image:tag > /tmp/r_docker_image_tag_***_.tar
	err := utils.SshExec(username, password, ip, "docker save "+image+" > "+remoteFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//sftp user@IP:/tmp/r_docker_image_tag_***.tar /tmp/l_docker_image_tag_***.tar
	utils.SftpGet(username, password, ip, remoteFilePath, localFilePath)

	//rm /tmp/r_docker_image_tag_***_.tar
	err = utils.SshExec(username, password, ip, "rm "+remoteFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if !tar {
		//docker load < /tmp/l_docker_image_tag_***.tar
		cmd := exec.Command("docker", "load", "-i", localFilePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		//rm /tmp/l_docker_image_tag_***_.tar
		_, err = exec.Command("rm", localFilePath).Output()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		//mv /tmp/l_docker_image_tag_***_.tar .
		_, err := exec.Command("mv", localFilePath, ".").Output()
		if err != nil {
			log.Fatal(err)
		}
	}
}

//["user[:password]","IP:image[:tag]"]
func parseFlag(c *cli.Context, arg string) (username, password, ip, image string) {
	arr := strings.Split(arg, "@")
	if len(arr) != 2 {
		cli.ShowSubcommandHelp(c)
		os.Exit(0)
	}
	//arrPrev = "user[:password]"
	arrPrev := arr[0]
	//arrPrevArr = [user, password]
	arrPrevArr := strings.Split(arrPrev, ":")
	if 1 == len(arrPrevArr) {
		username = arrPrevArr[0]
		password = ""
	} else {
		username = arrPrevArr[0]
		password = arrPrevArr[1]
	}
	//arrAfter = "IP:image[:tag]"
	arrAfter := arr[1]
	//arrAfterArr = [IP,image,tag]
	arrAfterArr := strings.Split(arrAfter, ":")
	lenArrAfterArr := len(arrAfterArr)
	if lenArrAfterArr == 2 {
		ip = arrAfterArr[0]
		image = arrAfterArr[1] + ":latest"
	} else if lenArrAfterArr >= 3 {
		ip = arrAfterArr[0]
		image =  strings.Join(arr_after_arr[1:len_arr_after_arr-1], ":") + ":" + arr_after_arr[len_arr_after_arr-1]
	} else {
		cli.ShowSubcommandHelp(c)
		os.Exit(0)
	}

	//read password
	if password == "" {
		fmt.Printf("password: ")
		pass, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		password = string(pass)
		fmt.Println("")
		if password == "" {
			fmt.Println("Error: Password Required")
			os.Exit(0)
		}
	}

	return username, password, ip, image
}
