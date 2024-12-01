# repoview

## Description

A tool to keep track of your repository statuses across multiple users and hosts.

## Usage

```shell
# First, enroll a repository on all relevant users and systems
repoview enroll .

# Add and enable systemd unit to run repoview push on shutdown
<TBA> # not yet implemented. well, the push command is, just not the unit file

# Whenever you're in doubt, check the statuses of a repository
repoview status
```

## Install

```shell
make build && make install # Installs the binary into ${HOME}/.local/bin by default. make sure its in your PATH
```

## Roadmap

- [ ] Unit file for auto pushes
- [ ] ?? something about branch info, maybe
