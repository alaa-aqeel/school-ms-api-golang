## School Management System API 

###  [go-chi](https://go-chi.io) A lightweight, idiomatic and composable router for building Go HTTP services.

### Server
```shell
# run docker compose
$ sh scripts/start.sh 
```

### Migrate
```shell
# make new 
$ sh scripts/migrate.sh make "name_migrate_file"

# up --all
$ sh scripts/migrate.sh up `$args`

# drops down
$ sh scripts/migrate.sh down --all
```