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
The above downloads the image, and runs it locally (in detached mode)

It also creates a _mount_ with the `-v` flag.
This is the location on your local machine where the container stores it's data.

Get into the container to load the seed data
(This is just one way to load the data)
interactive shell command:
```bash
sudo docker exec -it mvp /bin/bash
```

From inside the container, build the schema:
```bash
psql -U postgres -d postgres < /var/lib/postgresql/data/buildschema.sql
```
This all references the specified location on the container, which is _mounted_ to the specified storage (volume) on your local machine we specified with the `-v` flag when running the image

```bash
psql -h localhost -p 5432 -U postgres -d postgres -a -f /home/rjg/capstone/project/pgdata/buildschema.sql
```


https://app.diagrams.net/#G1GawNf-pdLUuXlt-w1ViFff-rzzkpnlUd