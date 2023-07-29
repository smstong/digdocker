Dig Docker

# Learn docker without docker

- what a process can see? 
    - hostname
    - process tree
    - IPC objects
    - network stack
    - users and groups
    - files and folder (mounts)
    - ...

- Limit what a process can see (unsharing is caring)
    - command line demo
        - uts
        - user
        - network
    - code demo
        - minidocker

# docker ecosystem 

- union file system (e.g. overlay fs)
- image registry
- standard vs implementation, other implemenations
- survival cheat-sheet
    - pull
    - run
    - stop
    - exec

# Build your own images 
- base rootfs

- minimal image
  - C demo
  - Go demo

# compose 
- docker image only contains "mnt namespace", all other namespaces need to be configured outside
- docker compose for the rest namespaces

# K8s
- pod
- complex overlay network arch
- Deployment/Storage/Loadbanlacer, ....
- Do you need it?
