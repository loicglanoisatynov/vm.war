Lancer le projet : 
```bash
go run main.go
```

Créer une session : 
```bash
curl -X POST -H "Content-Type: application/json" localhost:8080/v-network -d '{"username":"test"}' -v
```

Récupérer la session : 
```bash
curl -X GET -H "Content-Type: application/json" localhost:8080/v-network -v -d '{"username":"test"}'
```

Supprimer la session : 
```bash
curl -X DELETE -H "Content-Type: application/json" localhost:8080/v-network -d '{"username":"test"}' -v
```