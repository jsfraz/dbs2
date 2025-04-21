# dbs2

Backend pro projekt pro KIKM/DBS2 a KIT/TNPW2.

- [frontend projektu](https://github.com/DomDomiX/DBS2_Frontend)

## Použité technologie

- jazyk [Go](https://go.dev/)
- [PostgreSQL databáze](https://www.postgresql.org/)
- [OpenAPI 3 specifikace](https://swagger.io/specification/) ([openapi.json](openapi.json))

## Použité knihovny

| Název                                  | Popis                                             | URL projektu                                   |
|----------------------------------------|---------------------------------------------------|------------------------------------------------|
| gorm.io/gorm                           | ORM knihovna pro práci s relační databází         | <https://gorm.io/>                             |
| github.com/gin-gonic/gin               | Webový framework                                  | <https://gin-gonic.com/>                       |
| github.com/wI2L/fizz                   | gin nadstavba pro Open API 3                      | <https://github.com/wI2L/fizz>                 |
| github.com/go-playground/validator/v10 | Knihovna pro validaci HTTP požadavků a jiných dat | <https://github.com/go-playground/validator>   |
| github.com/golang-jwt/jwt              | Knihovna pro JWT                                  | <https://golang-jwt.github.io/jwt/>            |
| github.com/kelseyhightower/envconfig   | Knihovna pro .env konfiguraci                     | <https://github.com/kelseyhightower/envconfig> |

## Požadavky

- Docker
- Nginx
- Certbot

## Proměnné prostředí

- definováno v [.env](.env) souboru

| Název                 | Popis                                                   | Příklad                                 | Povinné |
|-----------------------|---------------------------------------------------------|-----------------------------------------|---------|
| GIN_MODE              | Určuje zda je program v debug módu nebo v produkci      | `debug` nebo `release`                  | Ano     |
| APP_URL               | Adresa aplikace (kvůli OpenAPI specifikaci)             | `https://example.com`                   | Ano     |
| SWAGGER               | Určuje zda je na adrese APP_URL/swagger.json Swagger UI | `true` nebo `false`                     | Ano     |
| ADMIN_MAIL            | Adminův e-mail                                          | `user@example.com`                      | Ano     |
| ADMIN_PW              | Adminovo heslo                                          | tajné heslo...                          | Ano     |
| PG_USER               | PostgreSQL uživatel                                     | `dbs2`                                  | Ano     |
| PG_PW                 | Heslo pro PostgreSQL uživatele                          | tajné heslo...                          | Ano     |
| PG_HOST               | Adresa PostgreSQL serveru                               | `localhost`                             | Ano     |
| PG_PORT               | Port PostgreSQL serveru                                 | `5432`                                  | Ano     |
| PG_DB                 | Název PostgreSQL databáze                               | `dbs2`                                  | Ano     |
| ACCESS_TOKEN_LIFESPAN | Životnost access tokenu ve vteřinách                    | `111600`                                | Ano     |
| ACCESS_TOKEN_SECRET   | Tajný klíč pro podepisování access tokenů               | dlouhý text... (třeba sloka z písničky) | Ano     |

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
- ❌ alespoň jeden číselník – viz. například <https://cs.wikipedia.org/wiki/%C4%8C%C3%ADseln%C3%ADk>
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
