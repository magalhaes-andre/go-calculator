# go-calculator

REST calculator developed in Go, using help of [mux library](https://github.com/gorilla/mux "MUX Source") for HTTP routing.

#### Provisioning Scripts Included

The provisioning schema in this repo includes:
- Packer (builds the docker image)
- Ansible (provision mux into the image Packer will build)
