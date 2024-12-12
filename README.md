# repoview

⚠⚠ Beware ⚠⚠
This piece of software cannot be considered stable

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

# Whenever you're in doubt, check the statuses of a repository
repoview status
```

## Install

Install the binary into `${HOME}/.local/bin` by default. Make sure it's in your PATH. Alternatively choose location with
`PREFIX`. If PREFIX is used, make sure to use the same prefix when and if you are running `make install-service`

```shell
# Build and install Repoview
make build && make install 

# (Optional) Install systemd service if you want automatic push on shutdown
make install-service
systemd --user enable repoview-push.service
```



## Roadmap

- [x] Unit file for auto pushes
- [ ] ?? something about branch info, maybe
