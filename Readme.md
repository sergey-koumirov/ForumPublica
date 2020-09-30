# todo

- ore calc inside construction
- load expenses from wallet api (job cost, sell taxes)

# run dev
server

`go run ./server/app.go`

`go build -o publica ./server/app.go`

run import & reset cache

`curl https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/sde.zip --output sde.zip`

`go run ./sde/app.go --file sde.zip`

set variables in var.json by example

# db & migrations (MySQL)


`sudo mysql -u root`

`create database publica;`

`create user 'publica'@'localhost' identified by 'publica';`

`grant all privileges on publica.* to 'publica'@'localhost';`

`./migrate -source=file://server/migrations -database=mysql://$(cat server/vars.json | jq '.DBC' -r) up`

`mysql -upublica -ppublica publica < publica.dump`

# thanks

to JetBrains for IDE
https://www.jetbrains.com/go/
