# repoview

## Description

A tool to keep track of your repository statuses across multiple users and hosts.

## Usage

### Initial

Create a repository for data. Point to that data in your `${HOME}/.config/repoview/config.yaml`. An example
configuration file can be found [here](./docs/config.yaml).

### Everyday usage

```shell
# First, enroll a repository on all relevant users and systems
repoview enroll .

# Add and enable systemd unit to run repoview push on shutdown
<TBA> # not yet implemented. well, the push command is, just not the unit file
# in other words; find a neat place to put `repoview push`

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
