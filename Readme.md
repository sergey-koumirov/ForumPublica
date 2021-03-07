# about

utilities for EVE online

- market monitor: check items stocks, prices, market volume, warnings about renewal of orders
- build calculator: exact calculation of materials needed for construction of specific amount of selected items

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

`GRANT PROCESS ON *.* TO 'publica'@'localhost';`

`./migrate -source=file://server/migrations -database=mysql://$(cat server/vars.json | jq '.DBC' -r) up`

`mysql -upublica -ppublica publica < publica.dump`

`cp file1 file1.gz`

`rm file1`

`gzip -d file1.gz`

# systemd

`sudo nano /etc/systemd/system/publica.service`

```
[Unit]
Description=Publica(EVE) server service
[Service]
Type=simple
WorkingDirectory=/home/user/eve/ForumPublica/
ExecStart=/home/user/eve/ForumPublica/publica
Restart=on-failure
RestartSec=5s
```

`systemctl start publica`

`systemctl status publica`

# thanks

to JetBrains for IDE
https://www.jetbrains.com/go/
