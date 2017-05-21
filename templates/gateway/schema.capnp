@0xdaf186973a8e7936;


struct Schema {
    node @0 :Text; # pointer to the parent service
    status @1 :Status;
    hostname @2 :Text;
    nics @3 :List(Nic); # Configuration of the attached nics to the container
    portforwards @4 :List(PortForward);
    portforwards6 @5 :List(PortForward);
    httpproxies @6 :List(HTTPProxy);
    dhcps @7 :List(DHCP);
    container @8 :Text; # Container spawned by this service

    struct Nic {
        type @0: NicType;
        id @1: Text;
        config @2: NicConfig;
    }

    struct CloudInit {
        userdata @0: Text;
        metadata @1: Text;
    }

    struct Host {
        macaddress @0: Text;
        hostname @1: Text;
        ipaddress @2: Text;
        ip6address @3: Text;
        cloudinit @4: CloudInit;
    }

    struct DHCP {
        nameservers @0: List(Text);
        hosts @1: List(Host);
        interface @2: Text;
        domain @3: Text;
    }

    struct NicConfig {
        dhcp @0: Bool;
        cidr @1: Text;
        gateway @2: Text;
        dns @3: List(Text);
    }

    enum HTTPType {
        http @0;
        https @1;
    }

    struct HTTPProxy {
        host @0: Text;
        distinations @1: List(Text);
        types @2: List(HTTPType);
    }

    enum Status{
        halted @0;
        running @1;
    }

    enum IPProtocol{
        tcp @0;
        udp @1;
    }

    struct PortForward{
        protocols @0: List(IPProtocol);
        srcport @1: Int32;
        srcip @2: Text;
        dstport @3: Int32;
        dstip @4: Text;
    }
    enum NicType {
        default @0;
        zerotier @1;
        vlan @2;
        vxlan @3;
        bridge @4;
    }
}