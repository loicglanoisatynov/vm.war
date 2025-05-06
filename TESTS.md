Lancer le projet : 
```bash
go run main.go
```

Créer une session : 
```bash
curl -X POST -H "Content-Type: application/json" localhost:8080/v-network -d '{"username":"test"}' -v
```