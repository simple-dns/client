#!/bin/sh

set -e
set -x

# linux
tar -cvzf release/dns_client_linux_amd64.tar.gz -C release/linux/amd64 dns-client
tar -cvzf release/dns_client_linux_arm64.tar.gz -C release/linux/arm64 dns-client
tar -cvzf release/dns_client_linux_arm.tar.gz   -C release/linux/arm   dns-client
tar -cvzf release/dns_client_linux_386.tar.gz   -C release/linux/386   dns-client

# windows
tar -cvzf release/dns_client_windows_amd64.tar.gz -C release/windows/amd64 dns-client.exe
tar -cvzf release/dns_client_windows_386.tar.gz   -C release/windows/386   dns-client.exe

# darwin
tar -cvzf release/dns_client_darwin_amd64.tar.gz -C release/darwin/amd64  dns-client

# freebase
tar -cvzf release/dns_client_freebsd_amd64.tar.gz -C release/freebsd/amd64 dns-client
tar -cvzf release/dns_client_freebsd_arm.tar.gz   -C release/freebsd/arm   dns-client
tar -cvzf release/dns_client_freebsd_386.tar.gz   -C release/freebsd/386   dns-client

# netbsd
tar -cvzf release/dns_client_netbsd_amd64.tar.gz -C release/netbsd/amd64 dns-client
tar -cvzf release/dns_client_netbsd_arm.tar.gz   -C release/netbsd/arm   dns-client

# openbsd
tar -cvzf release/dns_client_openbsd_amd64.tar.gz -C release/openbsd/amd64 dns-client
tar -cvzf release/dns_client_openbsd_arm.tar.gz   -C release/openbsd/arm   dns-client
tar -cvzf release/dns_client_openbsd_386.tar.gz   -C release/openbsd/386   dns-client

# dragonfly
tar -cvzf release/dns_client_dragonfly_amd64.tar.gz -C release/dragonfly/amd64  dns-client

# solaris
tar -cvzf release/dns_client_solaris_amd64.tar.gz -C release/solaris/amd64  dns-client

# generate shas for tar files
shasum release/*.tar.gz > release/dns_client_checksums.txt
