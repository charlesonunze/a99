# Cars API

API to let an automotive industry register their cars, list and view them.

## Running the application

Make sure you have [Git](https://git-scm.com/downloads) and [Docker](https://docs.docker.com/get-docker/) installed locally.

Also make sure nothing is running on port **8090**

#### Clone the project locally

```bash
git clone https://github.com/charlesonunze/a99.git && cd a99
```

#### Running the API server

```bash
docker run -p 8090:8090 --env-file .env charlesonunze/a99
```

#### Swagger documentation

```bash
docker run -p 8888:8080 -e SWAGGER_JSON=/pb/api.swagger.json -v $PWD/pb/:/pb swaggerapi/swagger-ui
```

#### Swagger UI

Visit the Swagger UI at [localhost:8888](http://localhost:8888/)

## Possible Improvements

#### Pagination

Add cusor based pagination for endpoints fetching multiple records.

#### API Responses

Currently when you don't find a record for example, it returns 200 and not 404. I could not find an easy way to do this, at least in time.

#### Database Modelling

Indexes to the models.

There is a One to many relationship of cars to features. This is not the best in a real world scenario.

For starters you might want to have an endpoint that you can add features independent of cars, in a hypotetical client application there could be an endpoint to fetch all features to populate a dropdown list or generate checkboxes. This also makes filtering by features easier a lot easier.

#### Tests

I did not add tests for the `car_service.go`. Mostly because the service doesn't do much at the moment. They only call the car repo internally. I real world service would do much more and therefore require additional tests.

#### Secrets

Only an example env file should be checked into source control.
