@0xdde10a317425c8d2;

struct Schema {
    id @0: Text; # mac address of the mngt network card
    status @1: NodeStatus;
    hostname @2: Text;

    client @3: Text; # name of the  g8os client service consumed
    gridConfig @4: Text; # name of the gridConfig service consumed

    enum NodeStatus {
        running @0;
        halted @1;
    }
}
