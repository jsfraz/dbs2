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
