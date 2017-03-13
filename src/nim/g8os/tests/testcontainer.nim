import ../g8os
import os
import json

when isMainModule:
   var server = os.commandLineParams()[0]
   var client = newG8OS(server)

   var cont = ContainerManager(client: client)
   echo json.pretty(cont.list())
   var bridge = ContainerBridge(name: "core-0", settings: "")
   #echo $cont.create(root_url = "http://home.maxux.net/gig/flist-ubuntu1604.db.tar.gz", bridges = @[bridge])
