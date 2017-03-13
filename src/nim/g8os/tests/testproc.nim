import ../g8os
import os
import json

when isMainModule:
   var server = os.commandLineParams()[0]
   var client = newG8OS(server)

   var process = ProcessManager(client: client)
   var procs = process.list()
   echo "PROC", json.pretty(process.get(id = procs[0]["cmd"]["id"].str))
