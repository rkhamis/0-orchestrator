### Prepare GateWay to be able to connect remote VXLANs with each other

#### Funcionality 

How this will work :
  1. create a zerotier network, but without!! allocated route or ip addresses
  1. have a running gateway with a pub interface and a leg on the vxlan (PRIV)
  1. create (maybe foresee it by default) a (linux) bridge inthe GW, and attach the PRIV interface to it
  1. migrate the ip of the PRIV interface to the bridge, and bring the bridge up
  1. start zerotier-one, join network id
  1. configure the zerotier interface to have config/bridge=true
  1. attach zt0 to the bridge


