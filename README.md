# tools
### search 
  Search docker images and tags on hub.docker.com
#### usage: 

```bash
go get -u github.com/AEGQ/tools/search
```
#### example:
```go
package main

import (
  "fmt"
  "log"

  "github.com/AEGQ/tools/search"
)

func main() {
  images, err := search.Images("ubuntu", 25) 
  if err != nil {
    log.Fatal(err)
  }
  search.PrintImages(images)
}
```
```bash
➜  go run main.go

STAR      OFFICIAL   NAME                                                   URL
6768      [OK]       ubuntu                                                 https://hub.docker.com/r/library/ubuntu
141                  dorowu/ubuntu-desktop-lxde-vnc                         https://hub.docker.com/r/dorowu/ubuntu-desktop-lxde-vnc
115                  rastasheep/ubuntu-sshd                                 https://hub.docker.com/r/rastasheep/ubuntu-sshd
88                   ansible/ubuntu14.04-ansible                            https://hub.docker.com/r/ansible/ubuntu14.04-ansible
80        [OK]       ubuntu-upstart                                         https://hub.docker.com/r/library/ubuntu-upstart
40        [OK]       neurodebian                                            https://hub.docker.com/r/library/neurodebian
32        [OK]       ubuntu-debootstrap                                     https://hub.docker.com/r/library/ubuntu-debootstrap
22                   nuagebec/ubuntu                                        https://hub.docker.com/r/nuagebec/ubuntu
19                   tutum/ubuntu                                           https://hub.docker.com/r/tutum/ubuntu
17                   1and1internet/ubuntu-16-nginx-php-phpmyadmin-mysql-5   https://hub.docker.com/r/1and1internet/ubuntu-16-nginx-php-phpmyadmin-mysql-5
11                   ppc64le/ubuntu                                         https://hub.docker.com/r/ppc64le/ubuntu
9                    aarch64/ubuntu                                         https://hub.docker.com/r/aarch64/ubuntu
8                    i386/ubuntu                                            https://hub.docker.com/r/i386/ubuntu
3                    darksheer/ubuntu                                       https://hub.docker.com/r/darksheer/ubuntu
3                    codenvy/ubuntu_jdk8                                    https://hub.docker.com/r/codenvy/ubuntu_jdk8
2                    1and1internet/ubuntu-16-apache-php-7.0                 https://hub.docker.com/r/1and1internet/ubuntu-16-apache-php-7.0
2                    1and1internet/ubuntu-16-nginx-php-5.6-wordpress-4      https://hub.docker.com/r/1and1internet/ubuntu-16-nginx-php-5.6-wordpress-4
0                    pivotaldata/ubuntu-gpdb-dev                            https://hub.docker.com/r/pivotaldata/ubuntu-gpdb-dev
0                    1and1internet/ubuntu-16-healthcheck                    https://hub.docker.com/r/1and1internet/ubuntu-16-healthcheck
0                    pivotaldata/ubuntu                                     https://hub.docker.com/r/pivotaldata/ubuntu
0                    thatsamguy/ubuntu-build-image                          https://hub.docker.com/r/thatsamguy/ubuntu-build-image
0                    ossobv/ubuntu                                          https://hub.docker.com/r/ossobv/ubuntu
0                    1and1internet/ubuntu-16-sshd                           https://hub.docker.com/r/1and1internet/ubuntu-16-sshd
0                    defensative/socat-ubuntu                               https://hub.docker.com/r/defensative/socat-ubuntu
0                    smartentry/ubuntu                                      https://hub.docker.com/r/smartentry/ubuntu

```
