# Studiengang Restful API
Example app for the Microservices wp

## Installation

``` bash
go get -v 

go build
./cmd/recipe-server/
```

## Route handles & endpoints

* `GET` /recipes
* `GET` /recipes/{id}
* `POST` /recipes
* `DELETE` /recipes/{id}

### Examples 

create a new recipe

``` bash
POST /studiengaenge

Request sample
{
  "id":"4545454",
  "name":"CloudWP",
  "beschreibung": "Foo bar",
  "kontakt":{"vorname":"Christian",  "nachname":"Bargmann"}
}
```

update an existing studiengang
``` bash
PUT /studiengaenge/{id}

Request sample
{
  "id":"4545454",
  "name":"Aktualisiertes CloudWP",
  "beschreibung": "Foo bar",
  "kontakt":{"vorname":"Christian",  "nachname":"Bargmann"}
}
```

delete an existing studiengang

``` bash
DELETE /studiengaenge/{id}

Request sample
{
  "id":"4545454"
}
```

for example using curl

```bash
curl -d '{"id":1}' -H 'Content-Type: application/json' -X DELETE http://localhost:8080/studiengaenge/1

```
