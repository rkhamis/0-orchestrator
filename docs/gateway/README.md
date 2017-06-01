# Table of contents
- [What is the gateway?](#%20What%20is%20the%20gateway?)
- [DHCP service](#DHCP service)
- [Port forwarding](#Port%20forwarding)
- [Reverse proxying](#Reverse%20proxying)
- [Advanced modus](#Advanced%20modus)
- [Cloud init](#Cloud%20init)
- [V(x)Lan to V(x)Lan bridge](#V(x)Lan%20to%20V(x)Lan%20bridge)

# What is the gateway?
![gateway](gateway.png)
The gateway is the networking Swiss army knife of the Zero-OS stack. It provides the following functions towards private V(x)Lan's:
- DHCP service for handing out networking configuration to containers and virtual machines
- A firewalled public ipaddress
- Internet connectivity towards the containers and virtual machines in the VxLans
- Port forwarding public ip traffic to hosted resources in the VxLan
- Reverse proxying http & https to hosted containers & virtual machines in the connected VxLans
- Cloud init server to initiate new virtual machines with passwords, ssh keys, configure swap, ...
- An OSI Layer 2 bridge between remote VxLans spread over different capacity pools

A gateway supports a mix of up to 100 network interfaces on VxLans, Vlans, zerotier networks and bridges.
- V(x)Lan networks need to have distinct subnets, and will be routed automatically by the Gateway.
- Routing configuration of connected zerotier networks needs to be handled in zerotier, as well as allowing the Gateway into the zerotier network (In case of a private zerotier network, when the token is not provided along with the zerotier network id).

# Creating a gateway
Creating a Zero-OS gateway is actually fairly easy. By submitting a POST request to the 0-orchestration server as specified in the api docs, gateways are created instantaneously:
https://rawgit.com/zero-os/0-orchestrator/master/raml/api.html#nodes__nodeid__gws_post

# DHCP service
After the gateway has been created, additional hosts can be added using the RESTful api of the 0-orchestrator. See https://rawgit.com/zero-os/0-orchestrator/master/raml/api.html#nodes__nodeid__gws__gwname__dhcp__interface__hosts_post to find out what needs to be posted to the 0-orchestrator server for adding / removing hosts.

# Port forwarding
Exposing tcp/udp based services hosted in the connected V(x)Lan networks is achieved via the portforwarding service of the Gateway. See https://rawgit.com/g8os/grid/master/raml/api.html#nodes__nodeid__gws__gwname__firewall_forwards_post to find out what needs to be posted to the 0-orchestrator server for adding / removing port forwards.

# Reverse proxying
The reverse proxy service in the Gateway can be used to expose http(s) services hosted in the connected V(x)Lans. It can do SSL-offloading and act as a load balancer towards multiple http servers.

# Advanced modus

## nftables
Advanced firewalling rules can be configured by just posting the [nftables](https://en.wikipedia.org/wiki/Nftables) configuration file that will be used in the Gateway. **Important** to note is that when the advanced firewalling configuration is set the portforwarding api as discussed above will no longer function. See https://rawgit.com/zero-os/0-orchestrator/master/raml/api.html#nodes__nodeid__gws__gwname__advanced_firewall_post to find how to use the advanced firewalling modus.

## Caddy
Advanced reverse proxy configuration can be configured by uploading the [Caddy](https://caddyserver.com/) configuration file onto the 0-orchestrator api. **Important** to note is that when the advanced reverse proxy configuration is set the reverse proxy api as discussed above will no longer function. See https://rawgit.com/zero-os/0-orchestrator/master/raml/api.html#nodes__nodeid__gws__gwname__advanced_http_post to find out how to use the advanced reverse proxying modus.

# Cloud init
TODO: Complete this with Jo.

# V(x)Lan to V(x)Lan bridge
Probably the coolest feature of the Gateway is this function. It allows to connect V(x)Lans in remote sites into one logical L2 network using a specially configured zerotier network. See https://github.com/zero-os/0-orchestrator/blob/master/docs/Network/spec_inter-vxlan-zerotier.md for detailed information on how to configure the zerotier network.
The bridge can be configured by setting the **zerotierbridge** property of the V(x)Lan interface of the Gateway. For more information how to create the bridge, see file:///C:/Users/geert/gig/code/github/g8os/resourcepool/raml/api.html#nodes__nodeid__gws_post
