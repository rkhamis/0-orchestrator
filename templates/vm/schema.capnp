@0x935023b5e21ba041;

struct Schema {
    node @0 :Text; # pointer to the parent service
    status @1 :Status;
    id @2 :Text;
    memroy @3 :UInt16; # Amount of memory in MiB
    cpu @3 :UInt16; # Number of virtual CUPs
    nic @3 :List(Text);
    # List of nic specifications.
    # Possible formats:
    # - "VxLAN:<<VxLAN id>>"
    # eg "VxLAN:200" Attaches the nic to VxLan 200
    # - "Zerotier:<<Zerotier network id>>"
    # eg "Zerotier:fsjyhgu76fsd87ydzf86t7dfygis" Attaches the nic to Zerotier network fsjyhgu76fsd87ydzf86t7dfygis
    disk @4 :List(Text);
    userCloudInit @5 :Text;
    systemCloudInit @6:Text;

    enum Status{
        running @0;
        halted @1;
        paused @2;
        halting @3;
        migrating @4;
    }

}
