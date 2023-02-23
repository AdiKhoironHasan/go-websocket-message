- this service to run server websocket
- running on localhost:8080
- send message using JSON data
- always send back the received mesage data, on broadcast or private message
- get all user in localhost:8080/get-user

- JSON:
  - broadcast message:
    {
    "From":"User 1",
    "ForUser":"",
    "Type":"Chat",
    "Message":"Hello All"
    }
  - broadcast message:
    {
    "From":"User 1",
    "ForUser":"User 2",
    "Type":"Chat",
    "Message":"Hello User 2"
    }
