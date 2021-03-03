# NetGo
NetGo is a basic implementation of ncat in go language.
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
$ git clone git@github.com:WhatTheSlime/netgo.git
$ cd netgo
```

## 2. Generate new TLS certificates
```bash
$ ./gencerts.sh
```

## 3. Compile the project
**For linux**
```bash
$ GOOS=linux GOARCH=amd64 go build .
```
**For windows**
```bash
$ GOOS=windows GOARCH=amd64 go build .
```
**For mac**
```bash
$ GOOS=darwin GOARCH=amd64 go build .
```

## 4. Drop it and use it wherever you want
```bash
$ ./netgo -h
```
