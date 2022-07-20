docker run command:
```bash
sudo docker run -d \
  -p5432:5432 \
	--name mvp \
	-e POSTGRES_PASSWORD=postgres \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v /home/rjg/capstone/project/pgdata:/var/lib/postgresql/data \
	postgres
```
The above downloads the image, and runs it locally (in detached mode).

It also creates a _mount_ with the `-v` flag. _Choose this path to match a location that makes sense on your machine_.
This is the location on your local machine where the container stores it's data.

Get into the container to load the seed data
(This is just one way to load the data)
interactive shell command:
```bash
sudo docker exec -it mvp /bin/bash
```

From inside the container, you can of course access psql this way:
```bash
psql -U postgres 
```
All the data references the specified location on the container, which is _mounted_ to the specified storage (volume) on your local machine we specified with the `-v` flag when running the image

You can access the DB remotely from your _local_ instance of psql like this:
```bash
psql -h localhost -p5432 -U postgres
```
