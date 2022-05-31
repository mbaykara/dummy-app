```
go run main
```
then get user
```
curl localhost:8090/users
```
or create user
```
curl -XPOST localhost:8090/createUser -d '{"ID":"112", "Name":"Noone", "Age": 19, "Job":"Pilot"}'
```