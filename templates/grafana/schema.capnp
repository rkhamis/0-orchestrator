@0x9b086fb7291fcce5; 

struct Schema {
    node @0 :Text; # Pointer to the parent service
    influx @1 :Text; # influx db to connect to
    container @2 :Text; # Container spawned by this service
    status @3 :Status;

    enum Status{
        halted @0;
        running @1;
    }
}
