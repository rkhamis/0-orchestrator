import redis
import uuid
import json
import os

type
   G8osClient* = ref object of RootObj
     host*: string
     client: Redis

type
   Response* = ref object of RootObj
     id*: string
     client*: redis.Redis

type
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


proc get*(r: Response, timeout: int = 10, verify: bool = false): ResponseData =
   var queue = "result:" & r.id
   var rawresp = r.client.brpoplpush(queue, queue, timeout)
   var jsonresp = json.parseJson(rawresp)
   var resp = ResponseData(id: r.id,
                       command: jsonresp["command"].str,
                       state: jsonresp["state"].str,
                       tags: jsonresp["tags"].str,
                       starttime: int(jsonresp["starttime"].num),
                       time: int(jsonresp["time"].num),
                       level: int(jsonresp["level"].num),
                       streams: [jsonresp["streams"][0].str, jsonresp["streams"][1].str],
                       data: jsonresp["data"].str)
   if verify:
      if resp.state != "SUCCESS":
         raise newException(OSError, "Invalid response " & resp.state)
   return resp

proc getJson*(r: Response, timeout: int = 10): json.JsonNode =
   var data = r.get(timeout, true)
   if data.level != 20:
      raise newException(OSError, "Invalid result level, expecting json(20) got " & $data.level)
   return json.parseJson(data.data)


proc newG8osClient*(host: string): G8osClient =
   var osclient = G8osClient(host: host)
   osclient.client = redis.open(host)
   return osclient

proc raw*(c: G8osClient, command: string, arguments: json.JsonNode): Response =
   var id_uuid: uuid.Tuuid
   id_uuid.uuid_generate_random()
   var payload = %*{
      "id": id_uuid.to_hex(),
      "command": command,
      "arguments": arguments,
   }
   echo $payload
   var resp = c.client.rpush("core:default", $payload)
   return Response(id: id_uuid.to_hex(), client: c.client)


proc system*(c: G8osClient, command: seq[string], pwd: string = "", stdin: string = ""): Response =
   var args = %*{
      "name": command[0],
      "args": command[1..^1],
      "dir": pwd,
      "stdin": stdin
   }
   var response = c.raw("core.system", args)
   return response

proc bash*(c: G8osClient, command: string): Response =
   var args = %*{
      "stdin": command
   }
   var response = c.raw("bash", args)
   return response

proc sync*(c: G8osClient, command: string, arguments: json.JsonNode): ResponseData =
   var response = c.raw(command, arguments)
   var data = response.get(verify = true)
   return data

proc getJson*(c: G8osClient, command: string, arguments: json.JsonNode): json.JsonNode =
   var response = c.raw(command, arguments)
   return response.getJson()


when isMainModule:
   var server = os.commandLineParams()[0]
   var client = newG8osClient(server)
   echo client.getJson("core.ping", json.newJNull()).str
   echo $client.system(@["ip", "a"]).get().streams[0]
   echo $client.bash("hostname").get().streams[0]
