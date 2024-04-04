# Todo24 API
### Eine API zur Verwaltung von Todos, mein erstes Go-Projekt 😊

![todo24_gopher](https://github.com/tim-w97/My-awesome-Todo-API/assets/63613014/ab4aced2-1833-40ec-be87-a4bb0cc2f0e4)

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

### Link zur Dokumentation der API-Endpunkte

<a href="https://app.swaggerhub.com/apis-docs/TimWagner/Todo24/1.0.0">
  <img width="200px" alt="swagger button" src="https://github.com/tim-w97/Todo24-API/assets/63613014/9ad378fc-aa0a-4de6-b1be-d50057cf7ba6"> 
</a>
