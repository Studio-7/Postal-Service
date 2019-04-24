## APIDOC
**Endpoint** - '/post/createpost'  
Required
- *token* string: JWT token
- *username* string
- *image* file
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

**Endpoint** - '/post/getpost'  
Required  
- *token* string: JWT token
- *username* string
- *ids* comma separted list of post ids Eg:- 917498-13313,18204-912-131,80357-28375
```json
Success
{
    "result": [
        {
            "comments": [
                {
                    "CommentBody": {
                        "Message": "First comment",
                        "Img": {
                            "Link": "",
                            "CreatedOn": "",
                            "UploadedOn": "0001-01-01T00:00:00Z",
                            "Manufacturer": "",
                            "Model": ""
                        }
                    },
                    "CreatedOn": "2019-04-23T21:48:03.904+05:30",
                    "CreatedBy": "user1",
                    "Parent": "15b64971-91aa-4bc1-9ad1-24745a97d68e",
                    "Likes": 0
                },
                {
                    "CommentBody": {
                        "Message": "Hello",
                        "Img": {
                            "Link": "",
                            "CreatedOn": "",
                            "UploadedOn": "0001-01-01T00:00:00Z",
                            "Manufacturer": "",
                            "Model": ""
                        }
                    },
                    "CreatedOn": "2019-04-23T21:35:52.537+05:30",
                    "CreatedBy": "user2",
                    "Parent": "15b64971-91aa-4bc1-9ad1-24745a97d68e",
                    "Likes": 0
                },
                {
                    "CommentBody": {
                        "Message": "Comment",
                        "Img": {
                            "Link": "",
                            "CreatedOn": "",
                            "UploadedOn": "0001-01-01T00:00:00Z",
                            "Manufacturer": "",
                            "Model": ""
                        }
                    },
                    "CreatedOn": "2019-04-23T21:34:22.611+05:30",
                    "CreatedBy": "user1",
                    "Parent": "15b64971-91aa-4bc1-9ad1-24745a97d68e",
                    "Likes": 0
                }
            ],
            "post": {
                "Id": "15b64971-91aa-4bc1-9ad1-24745a97d68e",
                "Title": "Post",
                "CreatedOn": "2019-04-11T11:52:54.118Z",
                "CreatedBy": "user4",
                "PostBody": {
                    "Message": "Post body",
                    "Img": {
                        "Link": "link to image",
                        "CreatedOn": "",
                        "UploadedOn": "2019-04-11T11:52:55.817Z",
                        "Manufacturer": "",
                        "Model": ""
                    }
                },
                "Hashtags": [
                    "#HashTag"
                ],
                "Likes": 0
            }
        }
    ],
    "token": "dGhvcg==.UkFqV3doVEhjdGN1QXhoeEtRRkRhRnBMU2pGYmNYb0U=.UWk0Q3NvU0xod3RCVlNrdE81aFRJY2cvTG9nb3VpOGRCZmdaWldrWTZGRT0="
}
Error
{ "error" : <ERROR>,
  "token": <NEWTOKEN>
}
The error could be "Not Authorized" in which case a jwt is not returned
```

**Endpoint** - '/post/addcomment'  
Required  
- *token* string: JWT token
- *username* string
- *postid* string
- *message* string
```json
Success
{
    "result": "comment added successfully",
    "token": "am9obndpY2s=.Y2N0TnN3WU5zR1J1c3NWbWFvekZaQnNiT0ppRlFHWnM=.c3ZsaGw3cDF3bS9DYTVXdWZwNGdmdjQvVTIyZGt1MksrZXBBekV4N002OD0="
}
Error
{ "error" : <ERROR>,
  "token": <NEWTOKEN>
}
The error could be "Not Authorized" in which case a jwt is not returned
```