# Project Golang API

## Requirements

Before cloning the project please make sure you already satisfied the Requirements below:

- [Git](https://git-scm.com/)
- [Golang](https://go.dev/doc/install)
- [Postgres](https://www.postgresql.org/)
- [VSCode](https://code.visualstudio.com/) // or you can use your favorite text editor

### SQL Client

Don't forget to download SQL Client to make your life easier
Here is my recommendations :

- [TablePlus](https://tableplus.com/) this one is paid but the free version is enough to do most basic task
- [AntaresSQL](https://antares-sql.app/) free and open source
- [DBGate](https://dbgate.org/) free and open source
- [PgAdmin](https://www.pgadmin.org/download/) free and open source

## Installation

After installing all the Requirements follow these steps to run the program

1. open terminal and run the code below\
   ```console
   git clone https://git.enigmacamp.com/enigma-20/benedictus-jullian-pradana/challenge-goapi.git
   cd challenge-go-api
   code .
   ```
   or you can just open the project directly through the text editor after clone
2. before continue to the next step, open your favorite sql client,
   create a new connection and then create a new database 'enigma_laundry'
3. go back to the project folder, copy and paste the content of the file DDL.sql into the sql client query editor
   and then do the same thing with the DML.sql and then run the sql command
4. go back to the project folder navigate to config/ folder and open database-connection.go
5. on line 10 of the code

   ```go
   const (
   	 host     = "localhost"
   	 port     = 5432
   	 user     = "enigmacamp" // change this into your own psql user
   	 password = "1234" // change this into your own psql user password
   	 dbname   = "enigma_laundry"
   )
   ```

6. Open terminal and run
   ```console
   go mod tidy
   ```
   to install all the dependencies
7. and then run
   ```console
   go run main.go
   ```
   to run the project
