import os
import base
import tables
import json

type
   ContainerManager* = ref object of Rootobj
      client*: G8OS
   ContainerMount* = ref object of Rootobj
      source*: string
      destination*: string
   ContainerBridge* = ref object of Rootobj
      name*: string
      settings*: string
   ContainerPort* = ref object of Rootobj
      source*: int
      destination*: int
   ContainerClient* = ref object of G8OS
      container*: int

proc list*(c: ContainerManager): json.JsonNode =
   var response = c.client.getJson("corex.list", json.newJObject())
   return response


proc create*(c: ContainerManager, root_url: string, mounts: seq[ContainerMount] = @[], zerotier: string = nil,
            bridges: seq[ContainerBridge] = @[], ports: seq[ContainerPort] = @[], hostname: string = ""): json.JsonNode =
   var args = %* {
      "root": root_url,
      "mount": json.newJObject(),
      "network": {
         "zerotier": zerotier,
         "bridge": json.newJArray(),
      },
      "port": json.newJObject(),
      "hostname": hostname
   }
   for mount in mounts:
      args["mount"][mount.source] = json.newJString(mount.destination)
   for bridge in bridges:
      var element = json.newJArray()
      element.add(json.newJString(bridge.name))
      element.add(json.newJString(bridge.settings))
      args["network"]["bridge"].add(element)
   for port in ports:
      args["port"][$port.source] = json.newJInt(port.destination)
   echo json.pretty(args)
   return c.client.getJson("corex.create", args)


proc terminate*(c: ContainerManager, container: int): json.JsonNode =
   var args = %*{
      "container": container
   }
   discard c.client.sync("corex.terminate", args)


proc raw*(c: ContainerClient, command: string, arguments: json.JsonNode): Response =
   var args = %* {
      "container": c.container,
      "command": {
         "command": command,
         "arguments": arguments
      }
   }
   var data = raw(c.G8OS, "corex.dispatch", args).getJson()
   return Response(client: c.client, id: $data.str)
