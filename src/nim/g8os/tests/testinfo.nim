import ../g8os
import os
import json

when isMainModule:
   var server = os.commandLineParams()[0]
   var client = newG8OS(server)
   echo client.getJson("core.ping", json.newJNull()).str
   echo $client.system(@["ip", "a"]).get().streams[0]
   echo $client.bash("hostname").get().streams[0]

   var infoc = InfoManager(client: client)
   echo "CPU", json.pretty(infoc.cpu())
   echo "NIC", json.pretty(infoc.nic())
   echo "MEM", json.pretty(infoc.mem())
   echo "DISK", json.pretty(infoc.disk())
   echo "OS", json.pretty(infoc.os())
