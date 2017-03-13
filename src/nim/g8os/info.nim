import redis
import json
import os
import base

type
   InfoManager* = ref object of Rootobj
      client*: G8OS

proc cpu*(i: InfoManager): json.JsonNode =
   var response = i.client.getJson("info.cpu", json.newJNull())
   return response

proc nic*(i: InfoManager): json.JsonNode =
   var response = i.client.getJson("info.nic", json.newJNull())
   return response

proc mem*(i: InfoManager): json.JsonNode =
   var response = i.client.getJson("info.mem", json.newJNull())
   return response

proc disk*(i: InfoManager): json.JsonNode =
   var response = i.client.getJson("info.disk", json.newJNull())
   return response

proc os*(i: InfoManager): json.JsonNode =
   var response = i.client.getJson("info.os", json.newJNull())
   return response
