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

## Install
### 1. Get the project
```bash
$ git clone git@github.com:WhatTheSlime/netgo.git
$ cd netgo
```

### 2. Generate new TLS certificates
```bash
$ ./gencerts.sh
```

### 3. Compile the project
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

### 4. Drop it and use it wherever you want
```bash
$ ./netgo -h
```

## Features
- [X] Executes the given command (-e, --exec <command>)
- [ ] Maximum <n> simultaneous connections (-m, --max-conns <n>)
- [X] Display help screen (-h, --help)
- [ ] Wait between read/writes (-d, --delay <time>)        
- [X] Bind and listen for incoming connections (-l, --listen)
- [ ] Accept multiple connections in listen mode (-k, --keep-open)
- [ ] Do not resolve hostnames via DNS (-n, --nodns)
- [ ] Use UDP instead of default TCP (-u, --udp)
- [ ] Set verbosity level (-v, --verbose)
- [ ] Connect timeout (-w, --wait <time>)
- [ ] Specify address of host to proxy through (--proxy <scheme://addr[:port]>)
- [X] Connect or listen with TLS (--tls)
- [ ] Display version information and exit (--version)


