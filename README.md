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
baykara/sample-app:main
```

Package helm charts

```bash
export CR_PAT=
echo $CR_PAT |docker login ghcr.io -u mbaykara --password-stdin #echo $CR_PAT |helm registry login ghcr.io -u mbaykara --password-stdin
helm push sample-app-1.1.5.tgz oci://ghcr.io/mbaykara/charts/sample-app
```
