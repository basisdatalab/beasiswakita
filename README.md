# BeasiswaKita
BeasiswaKita is simple scholarship tools for organizations and scholarship hunter.

## Development Guide
Make sure:
- `go` and `mysql` has been installed

1. Get this repository
  ```
  $ go get github.com/harkce/beasiswakita
  ```
2. Install dependencies
  ```
  $ make dep
  ```
3. Init database
  ```
  mysql> create database beasiswakita_development;
  $ make migrate
  ```

To create new migration file
  ```
  $ make migration <migration name>
  ```
The command above will create two files in directory `db/migrations`
  ```
  <current_timestamp>_<migration_name>_down.sql
  <current_timestamp>_<migration_name>_up.sql
  ```
