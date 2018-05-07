### fetch
  Fetch an image from any other docker host. (docker save => sftp => docker load)
- #### usage: 

```bash
# Prepare
mkdir -p $GOPATH/src/golang.org/x/crypto
git clone https://github.com/golang/crypto.git $GOPATH/src/golang.org/x/crypto
mkdir -p $GOPATH/src/golang.org/x/sys
git clone https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys
go get github.com/pkg/sftp
go get github.com/urfave/cli
go get gopkg.in/cheggaaa/pb.v1

# Prepare (Add following line to the remote node)
echo "source /etc/profile" >> /root/.bashrc

#Build
go build -o cli
```
- #### example:
```bash
➜  ./cli fetch root@192.168.1.1:golang:1.5.1  
password: 
Download: 693.11 MiB / 693.11 MiB [====================] 100.00% 14.08 MiB/s 49s
Loaded image: golang:1.5.1

```
