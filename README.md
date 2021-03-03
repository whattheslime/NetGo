# NetGo
```
    ___       ___       ___       ___       ___   
   /\__\     /\  \     /\  \     /\  \     /\  \  
  /:| _|_   /::\  \    \:\  \   /::\  \   /::\  \ 
 /::|/\__\ /::\:\__\   /::\__\ /:/\:\__\ /:/\:\__\
 \/|::/  / \:\:\/__/  /:/\/__/ \:\:\/__/ \:\/:/  /
   |:/  /   \:\/__/   \/__/     \::/  /   \::/  / 
   \/__/     \/__/               \/__/     \/__/  
```

## 1. Get the project
```bash
$ git clone https://github
```

## 2. Generate new TLS certificates
```bash
$ ./gencerts.sh
```
## 3. Compile the project
**For linux**
```bash
$ GOOS=linux GOARCH=amd64 go build ngo
```
**For windows**
```bash
$ GOOS=windows GOARCH=amd64 go build ngo
```
**For mac**
```bash
$ GOOS=darwin GOARCH=amd64 go build ngo
```

## 4. use it

### conn
### dump
### scan