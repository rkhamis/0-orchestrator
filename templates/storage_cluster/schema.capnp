@0x80eab0d1fa3f13a0;

struct Schema {
    label @0 :Text;
    status @2 :Status = empty;
    instanceNbr @6 :UInt32 = 256;
    diskClass @7:DiskClass = ssd;
    hasSlave @8 :Bool = false;

    storagePool @3 :List(Text); # storagepools availables to deploy ardb 
    dataArdb  @4 :List(Text); # ardbs for data
    metaArdb  @5 :Text; # ardb for metadata

    enum Status{
        empty @0;
        deploying @1;
        ready @2;
        error @3;
    }

    enum DiskClass {
        nvme @0;
        ssd @1;
        hdd @2;
        archive @3;
    }
}
