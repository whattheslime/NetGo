# NetGo

NetGo is a basic implementation of netcat in golang.

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
git clone git@github.com:WhatTheSlime/NetGo.git
cd NetGo
```

### 2. Compile the project

**For Linux and Windows**
```bash
make
```

**For Linux**
```bash
make linux
```

**For windows**
```bash
make windows
```

**For mac**
```bash
make darwin
````

### 3. Use it

```bash
cd build
./netgo -h
```

### 2. (Optional) Generate TLS certificates

```bash
make certifs
```

## Features

| Implemented | Flags                        | Description
|:-----------:|-----------------------------:|:-
| No          | -d, --delay <time>           | Wait between read/writes        
| Yes         | -e, --exec <command>         | Executes the given command
| Yes         | -h, --help                   | Display help screen
| Yes         | -k, --keep-open              | Accept multiple connections in listen mode
| Yes         | -l, --listen                 | Bind and listen for incoming connections
| Yes         | -m, --max-conns <number>     | Maximum simultaneous connections (default: 50)
| No          | -n, --nodns                  | Do not resolve hostnames via DNS
| No          | -u, --udp                    | Use UDP instead of default TCP
| No          | -v, --verbose                | Set verbosity level
| No          | -w, --wait <time>            | Connect timeout
| No          | -b, --broker                 | Enable connection brokering mode
| http        | -x, --proxy <proxy>          | Specify address of host to proxy through (<http|socks5|socks5h://[login:password@]host:port>)
| Yes         | --send                       | Only send data, ignoring received; quit on EOF (print md5)        
| Yes         | --recv                       | Only receive data, never send anything (print md5)
| Yes         | --tls                        | Connect or listen with TLS
| Yes         | --tls-cert                   | Specify TLS certificate file (PEM) for listening 
| Yes         | --tls-key                    | Specify TLS private key (PEM) for listening
| Yes         | --version                    | Display version information and exit


## References

- [Ncat](https://nmap.org/ncat/)
- [TLS Certificates](https://golang.org/src/crypto/tls/generate_cert.go)
- [Pflag](https://github.com/spf13/pflag)
