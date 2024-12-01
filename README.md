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

Install the binary into `${HOME}/.local/bin` by default. Make sure it's in your PATH. Alternatively choose location with
`PREFIX`

```shell
make build && make install 
```

## Roadmap

- [ ] Unit file for auto pushes
- [ ] ?? something about branch info, maybe
