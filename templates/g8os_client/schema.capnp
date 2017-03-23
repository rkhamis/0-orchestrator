@0xc9bbc42833473edc;

struct Schema {
    redisAddr @0 :Text;
    redisPort @1 :UInt32 = 6379;
    redisPassword @2 :Text;
}
