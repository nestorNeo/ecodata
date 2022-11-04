import json
import base64

with open('android.json') as mfile:
    data = json.load(mfile)
    for record in data:
        with open(record["_id"]["$oid"] + record["name"], "wb") as fh:
            fh.write(base64.b64decode(record["content"]["$binary"]["base64"].encode("ascii")))
