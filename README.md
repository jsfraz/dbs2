# dbs2

Backend pro projekt pro KIKM/DBS2 a KIT/TNPW2.

- [frontend projektu](https://github.com/DomDomiX/DBS2_Frontend)

## Požadavky

- Docker
- Nginx
- Certbot

## Proměnné prostředí

- definováno v [.env](.env) souboru

- TODO proměnné prostředí

## Instalace

```bash
#Zkopírování Nginx konfigurace
sudo cp dbs2.conf /etc/nginx/conf.d/
# Vygenerování certifikátu
sudo certbot --nginx -d dbs2-backend.josefraz.cz
# Restart Nginx
sudo systemctl restart nginx
# Compose (po konfiguraci proměnných prostředí)
sudo docker compose up -d --build
```

## Swagger UI (pokud povoleno)

<https://dbs2-backend.josefraz.cz/swagger>

## Databázové schéma

![Databázové schéma](dbs2%20-%20public.png)

## Náležitosti

- ✔ cca. 10 tabulek navrženého datového modelu (individuální domluva)
- ❌ alespoň jeden číselník – viz. například https://cs.wikipedia.org/wiki/%C4%8C%C3%ADseln%C3%ADk
- ✔ alespoň tři pohledy. které budou volány z aplikace
- ✔ alespoň tři funkce různého typu s odpovídající složitostí
- ✔ alespoň tři uložené procedury
- ✔ alespoň dva triggery
- ✔ alespoň jedna transakce s ošetřeným chováním při rollbacku
- ❌ použití indexů na neklíčové sloupce
- ❌ použití kompozitních primárních klíčů
- ❌ vyzkoušet si použití datového typu JSON v moderních relačních databázích (rozumné použití včetně filtrace nad těmito sloupci může ovlivnit počet požadovaných databázových tabulek, případně odpuštění jednoho ze zde uvedených požadavků – záleží na domluvě se cvičícím)
- ❌ v databázovém serveru bude vytvořen uživatel s potřebnými právy pouze k databázovým objektům, které pro správný běh aplikace potřebuje – tzn. root (admin) účet nebude aplikací používán, vč. omezení přístupu pouze z potřebné IP adresy
- ❌ doporučené rozjetí projektu v Dockeru pomocí docker-compose – bude zajištěna inicializace struktury databáze a nahrání dat při startu
- ✔ verzování vývoje pomocí Gitu
- ✔ vhodným způsobem zajistit ukládání obrázků, které budou v aplikaci načteny a zobrazeny
- ❌ aplikace bude využívat minimálně 2 plnohodnotné formuláře (např. ošetření vstupních polí, apod.) pro vytváření nebo modifikaci dat v tabulkách