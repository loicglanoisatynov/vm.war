# Projet VM.WAR

## Description

Ce projet est le rendu du projet fil rouge d'infra du cours de 2ème année de Bachelor Informatique chez Lyon Ynov Campus. Il s'agit d'un serveur d'hébergement de machines virtuelles vulnérables en réseaux fermés, sur lesquels les étudiants peuvent s'entraîner à la cybersécurité en s'attaquant mutuellement.

## Installation

### Prérequis
- Git
- Go 1.23 ou supérieur
  

### Clonage du dépôt
```bash
git clone https://github.com/loicglanoisatynov/vm.war.git
cd vm.war
```

### Installation des dépendances
```bash
go mod tidy
```

### Lancer le projet
```bash
go run main.go
```

## Créer une session (remplacer "test" par le nom d'utilisateur souhaité)
```bash
curl -X POST -H "Content-Type: application/json" localhost:8080/v-network -d '{"username":"test"}' -v
```

## Récupérer la session (remplacer "test" par le nom d'utilisateur souhaité)
```bash
curl -X GET -H "Content-Type: application/json" localhost:8080/v-network -v -d '{"username":"test"}'
```

## Supprimer la session (remplacer "test" par le nom d'utilisateur souhaité)
```bash
curl -X DELETE -H "Content-Type: application/json" localhost:8080/v-network -d '{"username":"test"}' -v
```