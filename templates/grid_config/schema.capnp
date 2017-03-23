@0x9e47335012092925;

struct Schema {
    network @0: NetworkConfig;


    NetworkConfig struct{
        mgmtNetwork @0: Text;
        backplaneNetwork @1:Text;
    }
}
