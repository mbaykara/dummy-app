```
go run main.go
```
then get user
```
curl localhost:8090/users
```
or create user
```
curl -XPOST localhost:8090/createUser -d '{"ID":112, "Username":"Susan", "Email":"susan@sus.com"}'
```
Docker image:

```bash
baykara/sample-app:0.1.3
```