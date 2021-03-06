# GOFLEXIFY

### introduction

_Jedna se o projekt urceny pro komercni strany, kdy mohou v real-time fungovar, sledovat zakazniky, provadet transakce atd..._

### Spusteni projektu

Projekt je psan v Golang, proto je potreba udelat nekolik prikazu

1. go mod init <nazev-projektu> _pro importovani projektovych pkg_
2. go mod download _pro stazeni dependencies_

### DOCKER

Dale cela aplikace je zabalena do containers v Docker
Pro spusteni virtualizace je potreba byt v root projektu

1. docker-compose build && docker-compose up

Nastaveni dockeru
Environment vars jsou pouze pro verzovani, v prod/devu se pouziva dotfile/ini
