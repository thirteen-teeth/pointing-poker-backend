# pointing-poker-backend

```
new_session=`curl -sX POST 'http://localhost:8000/newsession' -d 'userID=exampleUserID' | jq .sessionID`
join_session=`curl -sX POST 'http://localhost:8000/join' -d 'sessionID='$new_session'&userID=exampleUserID'`
echo $join_session
```