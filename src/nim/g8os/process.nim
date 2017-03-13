import redis
import json
import os
import base
import strutils

type
   ProcessManager* = ref object of Rootobj
      client*: G8OS

proc list*(i: ProcessManager): json.JsonNode =
   var response = i.client.getJson("process.list", json.newJObject())
   return response

proc get*(i: ProcessManager, id: string): json.JsonNode =
   var response = i.client.getJson("process.list", %*{"id": id})
   if len(response) != 1:
      raise newException(OSError, "Could not find process with id $1" % id)
   return response[0]

proc kill*(i: ProcessManager, id: string): json.JsonNode =
   var response = i.client.getJson("process.kill", %*{"id": id})
   return response
