# EisenhowerDo-Backend
Eine Go-API zur Verwaltung von Todos

<br/>

### So kann das Backend gestartet werden

Zuerst muss Docker installiert und aktiv sein, 
danach kann mit folgendem Befehl das Backend inkl. 
MySQL-Datenbank gestartet werden:

```shell
$ docker compose up
```

<br/>

### Das kann die API

- Anmeldung und Registrierung
- Benutzerauthentifizierung per JSON Web Token
- Hinzufügen, Ändern und Löschen von Todos
- Teilen von Todos mit anderen Benutzern
- Änderung der Position eines Todos
- Kategorisierung von Todos
- JSON als Austauschformat

<br/>

### Das habe ich benutzt

- Go für die API
- MySQL für die Datenbank
- Go Gin zur Verarbeitung von HTTP-Anfragen
- Go-MySQL-Driver zur Kommunikation mit der Datenbank
- go-jwt für JSON Web Tokens
- GoDotEnv für Umgebungsvariablen

<br/>

### So habe ich meine Datenbank entworfen

- Tabellen: user, todo, category, sharedTodo
- sharedTodo enthält die geteilten Todos und löst eine _n:m_ Beziehung auf, weil:
  - ein Benutzer mehrere seiner Todos mit einem anderen Benutzer teilen kann
  - ein Todo mit mehreren Benutzern geteilt werden kann

![erd](https://github.com/tim-w97/Todo24-API/assets/63613014/142b8630-38b6-496b-9c33-6a26a8e8b50e)

<br/>
