# ssh

`ssh` is a SSH client wrapper with reconnection functionality.

[![Latest Release](https://img.shields.io/github/v/release/shibijm/ssh?label=Latest%20Release)](https://github.com/shibijm/ssh/releases/latest)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shibijm/ssh/release.yml?label=Build&logo=github)](https://github.com/shibijm/ssh/actions/workflows/release.yml)

## Download

Downloadable builds are available on the [releases page](https://github.com/shibijm/ssh/releases).

## Usage

```text
$ /path/to/this/ssh root@192.168.0.2
Enter passphrase for key '~/.ssh/id_rsa':
root:~ ❯ echo Hi
Hi
root:~ ❯ Read from remote host 192.168.0.2: Connection reset by peer
Connection to 192.168.0.2 closed.
client_loop: send disconnect: Connection reset by peer
Reconnecting... (exit status 255)
ssh: connect to host 192.168.0.2 port 22: Connection timed out
Reconnecting... (exit status 255)
Enter passphrase for key '~/.ssh/id_rsa':
root:~ ❯
```
