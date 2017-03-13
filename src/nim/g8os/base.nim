import redis
import uuid
import json
import os

type
   G8OS* = ref object of RootObj
     host*: string
     client: Redis

   Response* = ref object of RootObj
     id*: string
     client*: redis.Redis

   ResponseData* = ref object of Rootobj
      id*: string
      command*: string
      data*: string
      streams*: array[2, string]
      level*: int
      state*: string
      starttime*: int
      time*: int
      tags*: string

proc getResponse(resp: json.JsonNOde, verify: bool = false): ResponseData =
   var response = ResponseData(id: resp["id"].str,
                       command: resp["command"].str,
                       state: resp["state"].str,
                       tags: resp["tags"].str,
                       starttime: int(resp["starttime"].num),
                       time: int(resp["time"].num),
                       level: int(resp["level"].num),
                       streams: [resp["streams"][0].str, resp["streams"][1].str],
                       data: resp["data"].str)
   if verify:
      if response.state != "SUCCESS":
         raise newException(OSError, "Invalid response " & response.state)
   return response

proc get*(r: Response, timeout: int = 10, verify: bool = false): ResponseData =
   var queue = "result:" & r.id
   var rawresp = r.client.brpoplpush(queue, queue, timeout)
   var jsonresp = json.parseJson(rawresp)
   return getResponse(jsonresp)

proc getJson*(r: Response, timeout: int = 10): json.JsonNode =
   var data = r.get(timeout, true)
   if data.level != 20:
      raise newException(OSError, "Invalid result level, expecting json(20) got " & $data.level)
   return json.parseJson(data.data)

proc newG8OS*(host: string): G8OS =
   var osclient = G8OS(host: host)
   osclient.client = redis.open(host)
   return osclient

proc raw*(c: G8OS, command: string, arguments: json.JsonNode): Response =
   var id_uuid: uuid.Tuuid
   id_uuid.uuid_generate_random()
   var payload = %*{
      "id": id_uuid.to_hex(),
      "command": command,
      "arguments": arguments,
   }
   discard c.client.rpush("core:default", $payload)
   return Response(id: id_uuid.to_hex(), client: c.client)

proc sync*(c: G8OS, command: string, arguments: json.JsonNode): ResponseData =
   var response = c.raw(command, arguments)
   var data = response.get(verify = true)
   return data

proc getJson*(c: G8OS, command: string, arguments: json.JsonNode): json.JsonNode =
   var response = c.raw(command, arguments)
   return response.getJson()

proc system*(c: G8OS, command: seq[string], pwd: string = "", stdin: string = ""): Response =
   var args = %*{
      "name": command[0],
      "args": command[1..^1],
      "dir": pwd,
      "stdin": stdin
   }
   var response = c.raw("core.system", args)
   return response

proc bash*(c: G8OS, command: string): Response =
   var args = %*{
      "stdin": command
   }
   var response = c.raw("bash", args)
   return response
