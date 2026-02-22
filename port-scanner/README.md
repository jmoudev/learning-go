# Port Scanner

CLI port scanner

## Usage

Command line help:
```
port-scanner --help

NAME:
   port-scanner - Scan a number of ports for a given url

USAGE:
   port-scanner [global options] <url>

GLOBAL OPTIONS:
   --startPort int  The lowest port in the port range (default: 1)
   --endPort int    The highest port in the port range (default: 1000)
   --help, -h       show help
```

Example output:
```
port-scanner --startPort 1 --endPort 500 golang.org

Scanning ports 1-500 of golang.org
Port 80 OK
Port 443 OK
```

