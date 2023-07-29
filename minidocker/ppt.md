Unsharing is caring

# UTS namespace demo

```
T1: unshare -u bash
T1: hostname erge
T1: hostname

T2: hostname
```

# user namespace
```
T1: unshare -U bash
T1: id

T2: id
```

# mount namespace
```
T1: unshare -m bash
T1: cat /proc/$$/mounts
# child mount ns automatically filled with parent's list of mounts
```

# pid namespace

# network namespace
```
P: ip netns add erge
P: ip netns list
P: ls /var/run/netns
P: sudo ip netns exec erge bash

C: ip link list
C: ping 127.0.0.1
C: ip link set lo up
C: ping 127.0.0.1

# add virtual ethenet pair for communicating between two network namespaces
P: ip link add veth0 type veth peer name veth1
P: ip link set veth1 netns erge
P: ip link list
P: ip addr add 10.0.0.1/24 dev veth0
P: ip link set veth0 up
P: ping 10.0.0.2

C: ip link list
C: ip addr add 10.0.0.2/24 dev veth1
C: ip link set veth1 up
C: ping 10.0.0.1
```
