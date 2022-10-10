# migrations

Database migrations are managed using [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) - install it locally before attempting to create or run migrations.

## Create new migration files

```bash
$ migrate create -ext sql -dir internal/database/migrations -seq create_users_table
/Users/meder/Code/shitlist/internal/database/migrations/000001_create_users_table.up.sql
/Users/meder/Code/shitlist/internal/database/migrations/000001_create_users_table.down.sql
```
