## Example app for the Microservices wp

## Installation

``` bash
go get -v 

go build
./cmd/recipe-server/
```

## Route handles & endpoints

* `GET` /recipes
* `GET` /recipes/{id}
* `GET` /recipes/random/{mealType}
* `POST` /recipes
* `DELETE` /recipes/{id}

### Examples 

create a new recipe

``` bash
POST /recipes

Request sample
{
    "mealtype": "Breakfast",
    "name": "Pancakes",
    "Ingredients": [ "150g all purpose flour",
    				 "150ml of milk"],
    "preparation": "Add all ingredients and mix. Put in Pan."
}
```

get a specific  recipe

``` bash
GET /recipes/{id}

Request sample
{
  "id":"1"
}
```

get a random recipe of specific mealType

``` bash
GET /recipes/random/{mealType}

Request sample
{
  "mealType":"Breakfast"
}
```

delete an existing recipe

``` bash
DELETE /recipes/{id}

Request sample
{
  "id":"1"
}
```

for example using curl

```bash
curl -d '{"id":1}' -H 'Content-Type: application/json' -X DELETE http://localhost:8080/recipes/1

```
