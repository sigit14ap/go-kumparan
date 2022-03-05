
## Golang Kumparan

How to run

- Compose docker container with command :
```bash
  docker-compose up
```
- Open new terminal without kill first step and run command:
```bash
  docker exec -it go_kumparan go run cmd/app/main.go
```

- Import Postman collection from  : <strong>Go-Kumparan.postman_collection.json</strong>


- You can access Mongo Express from url :
```bash
  http://localhost:8081
```

- Run Unit Test
```bash
  cd tests && go test
```