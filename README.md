# BeasiswaKita
BeasiswaKita is simple scholarship tools for organizations and scholarship hunter.

## Quick Start
### Environment Setup
Make sure:
- `go` and `mysql` has been installed
- `$GOPATH` exist in your path

1. Get this repository
    ```
    $ go get github.com/harkce/beasiswakita
    $ cd $GOPATH/src/github.com/harkce/beasiswakita/
    ```
2. Set up development environtment
    ```
    $ cp env.sample .env
    ```
    Edit `.env` file according to your environment setup.
    ```env
    # Fill this with random string
    APP_KEY=50m3r4nd0m4w3s0m35tr1ng

    # Change DB_USERNAME and DB_PASSWORD
    DEVELOPMENT_DATABASE_URL=DB_USERNAME:DB_PASSWORD@(127.0.0.1:3306)/beasiswakita_development

    # Host to serve image PATH
    BEASISWAKITA_HOST=http://localhost:8061
    ```
3. Install dependencies
    ```
    $ make init
    ```
3. Init database
    ```
    mysql> create database beasiswakita_development;
    $ make migrate
    ```
### Development Guide
This project use MySQL as default DBMS, and use [kamimai](https://github.com/Fs02/kamimai) for the migration.

To create new migration file, go to project root directory and run migration command.
  ```
  $ make migration <migration name>
  ```
The command above will create two files in directory `db/migrations`
  ```
  <current_timestamp>_<migration_name>_down.sql
  <current_timestamp>_<migration_name>_up.sql
  ```
To do a migration, run migrate on project root
  ```
  $ make migrate
  ```
Run rollback to revert a migration
  ```
  $ make rollback
  ```
If you add more dependencies to the project, don't forget to `dep ensure`
  ```
  $ dep ensure
  ```
### Build Project
Build project using build commnd.
  ```
  $ make build
  ```
Run project by executing:
  ```
  $ ./bin/beasiswakita
  ```
Beasiswakita will run at port `:8061`

## Beasiswakita endpoint

### beasiswakita.surge.sh/
-> Halaman ini untuk mengakses halaman utama
 
### beasiswakita.surge.sh/board
-> Halaman ini untuk mengakses halaman dashboard
-> Requirement :
   - Harus sudah dalam kondisi login sebagai user penerima beasiswa
   - 

### beasiswakita.surge.sh/search
-> Halaman ini untuk mengakses halaman pencarian 

### beasiswakita.surge.sh/detilBeasiswa
-> Halaman ini untuk mengakses halaman lihat detil beasiswa

### beasiswakita.surge.sh/partnership
-> Halaman ini untuk halaman pendaftar bagi partnership 

### beasiswakita.surge.sh/partnership/board
-> Halaman ini digunakan sebagai halaman dashboard perusahaan atau instansi untuk membuat lowongan beasiswa

### beasiswakita.surge.sh/login
-> Halaman ini untuk mengakses halaman login
-> Requirement :
   - 
   - 
-> Akun admin dummy :
   - username : basdat
   - password : password

### beasiswakita.surge.sh/daftar
-> Halaman ini untuk mengakses halaman signup

#endpoint

