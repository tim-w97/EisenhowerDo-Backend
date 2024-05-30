# EisenhowerDo-Backend
Eine Go-API zur Verwaltung von Todos

<br/>

### So kann das Backend gestartet werden

Sie benötigen Docker, um das Backend zu starten.

Sobald Docker installiert und aktiv ist, 
kann mit folgendem Befehl das Backend inkl. 
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
- JSON als Austauschformat

<br/>

### Das habe ich benutzt

- Go für die API
- MySQL für die Datenbank
- Go Gin zur Verarbeitung von HTTP-Anfragen
- Go-MySQL-Driver zur Kommunikation mit der Datenbank
- go-jwt für JSON Web Tokens
- GoDotEnv für Umgebungsvariablen