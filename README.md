<div align="center">
  <br>
  <h1>Feed Me More</h1>
</div>
## Running

- **Backend and database**

First, You should install [Docker](https://www.docker.com/) in your machine

Second, Open the terminal in project directory, and type the following commands:

```bash
docker-compose build
```

```bash
docker-compose up
```

now DB container on port 5432 and backend container are running on , you can test the endpoints using [this postman collection](postman%20collection/Feed-Me-More.postman_collection.json) by [Postman application](https://www.postman.com/), or run the frontend container then test.

- **Frontend**

got o this [directory](./frontend/feed-me/) and run the fo;;owing command:

```
docker-compose up
```

