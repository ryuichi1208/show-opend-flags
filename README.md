# show-opend-flags
show open(2) flags tool

## Usage

```
$ show-open-flags ${PID}

## Example
$ show-open-flags 1
/dev/null: [O_RDWR]
/dev/null: [O_RDWR]
/proc/1/mountinfo: [O_CLOEXEC]
/proc/swaps: [O_CLOEXEC]
/run/cloud-init/hook-hotplug-cmd: [O_RDWR O_CLOEXEC O_NDELAY]
/dev/rfkill: [O_RDWR O_CLOEXEC O_NDELAY]
/dev/null: [O_RDWR]
/dev/autofs: [O_CLOEXEC]
/dev/kmsg: [O_WRONLY O_CLOEXEC]
/sys/fs/cgroup: [O_CLOEXEC O_DIRECTORY O_NDELAY]
/run/dmeventd-server: [O_RDWR O_CLOEXEC O_NDELAY]
/run/dmeventd-client: [O_RDWR O_CLOEXEC O_NDELAY]
/run/initctl: [O_RDWR O_CLOEXEC O_NDELAY]
```
