# Sample Application

## Running the Application

To run the application, use the following command:

```bash
go run main.go
```

## Interacting with the API

### Get Users

To retrieve the list of users, execute:

```bash
curl localhost:8090/users
```

### Create a User

To create a new user, use the following command:

```bash
curl -XPOST localhost:8090/createUser -d '{"ID":112, "Username":"Susan", "Email":"susan@sus.com"}'
```

## Docker Image

The Docker image for this application is available at:

```bash
baykara/sample-app:main
```

## Package Helm Charts

To package and push Helm charts, use the following commands:

```bash
export CR_PAT=
echo $CR_PAT | docker login ghcr.io -u mbaykara --password-stdin
# echo $CR_PAT | helm registry login ghcr.io -u mbaykara --password-stdin
helm push sample-app-1.1.5.tgz oci://ghcr.io/mbaykara/charts/sample-app
```
