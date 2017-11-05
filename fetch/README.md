### fetch
  Fetch an image from any other docker host.
- #### usage: 

```bash
go build -o cli
```
- #### example:
```bash
➜  ./cli fetch root@192.168.1.1:golang:1.5.1  
password: 
Download: 693.11 MiB / 693.11 MiB [====================] 100.00% 14.08 MiB/s 49s
Loaded image: golang:1.5.1

```
