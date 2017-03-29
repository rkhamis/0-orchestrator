@0x935023b5e21bf041;

struct Schema {
    size @0 :UInt64;
    blocksize @1 :UInt32;
    deduped @2 :Bool;
    templateVolume @3 :Text; # in case it's a copy of another volume
    readOnly @4 :Bool;
    driver @5 :Text;
    status @6 :Status;

    storageCluster @6 :Text; # parent

    enum Status {
        running @0;
        halter @1;
        rollingback @3;
    }
}
