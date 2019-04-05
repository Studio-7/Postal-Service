## APIDOC
**Endpoint** - '/post/createpost'  
Required
- *token* string: JWT token
- *username* string
- *title* string
- *message* string
- *hashtags* string array: List of hastags separated by comma without any spaces
- *travelcapsule* string: ID of the travelcapsule  
```json
Success
{
    "result": "successfully uploaded",
    "token": "am9obndpY2s=.Ump4QXdud2VrckJFbWZkemRjRWtYQkFralFaTEN0VE0=.YzlxU3pvbnpJYXR6RC9Od1RLTFRKSTJhTW53MS9QWVkwSTNIR2d4ZjI1UT0=",
    "travelcapsule": "999f73c2-4df7-4057-9522-fa0d80151d4d"
}
Error
{ "error" : <ERROR>,
  "token": <NEWTOKEN>
}
The error could be "Not Authorized" in which case a jwt is not returned
```
  
**Endpoint** - 'post/createtravelcapsule'  
Required  
- *token* string: JWT token
- *username* string
- *title* string
```json
Success
{
    "result": "successfully created",
    "token": "am9obndpY2s=.Z2JhaUNNUkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2o=.elUwd3d4SjQwb3lsM0R2OUpWeExCNVRLS3hrUG5QRVFRV25sNHVGdFRvcz0=",
    "travelcapsule": "343b1f5a-4a0d-4167-a334-4e94115fb794"
}
Error
{ "error" : <ERROR>,
  "token": <NEWTOKEN>
}
The error could be "Not Authorized" in which case a jwt is not returned
```
  
**Endpoint** - '/post/like'  
Required  
- *token* string: JWT token
- *username* string
- *postid* string
```json
Success
{
    "result": "liked",
    "token": "am9obndpY2s=.Y2N0TnN3WU5zR1J1c3NWbWFvekZaQnNiT0ppRlFHWnM=.c3ZsaGw3cDF3bS9DYTVXdWZwNGdmdjQvVTIyZGt1MksrZXBBekV4N002OD0="
}
Error
{ "error" : <ERROR>,
  "token": <NEWTOKEN>
}
The error could be "Not Authorized" in which case a jwt is not returned
```
  
**Endpoint** - '/post/unlike'  
Required  
- *token* string: JWT token
- *username* string
- *postid* string
```json
Success
{
    "result": "unliked",
    "token": "am9obndpY2s=.Y2N0TnN3WU5zR1J1c3NWbWFvekZaQnNiT0ppRlFHWnM=.c3ZsaGw3cDF3bS9DYTVXdWZwNGdmdjQvVTIyZGt1MksrZXBBekV4N002OD0="
}
Error
{ "error" : <ERROR>,
  "token": <NEWTOKEN>
}
The error could be "Not Authorized" in which case a jwt is not returned
```
