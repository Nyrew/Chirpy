
# Chirpy 🐦

**Chirpy** je jednoduchá mikroblogovací platforma implementovaná v Go (Golang). Chirpy umožňuje uživatelům vytvářet, upravovat a mazat příspěvky (tzv. "chirps") a spravovat uživatelské účty.

## 🚀 Funkce
- **Autentizace uživatelů**: Přihlášení pomocí tokenů.
- **CRUD operace pro příspěvky**:
  - Vytváření, čtení, aktualizace a mazání chirpů.
- **Chirpy Red Membership**:
  - Speciální členství umožňující editaci příspěvků.
- **Webhooks pro integraci**:
  - API klíče pro bezpečnou komunikaci s Polka (platební platforma).
- **Filtrování a řazení**:
  - Možnost řadit a filtrovat příspěvky podle autora a data.

## 📁 Struktura projektu
```
.
├── cmd/
│   └── main.go         # Hlavní vstupní bod aplikace
├── database/
│   ├── queries.sql     # SQL dotazy pro práci s databází
│   └── models.go       # Modely pro databázové entity
├── auth/
│   ├── jwt.go          # Logika autentizace
│   └── api_key.go      # Práce s API klíči
├── handlers/
│   ├── chirps.go       # Handlery pro správu chirpů
│   └── users.go        # Handlery pro správu uživatelů
├── migrations/
│   └── *.sql           # Migrace pro databázi
├── .env.example        # Ukázkový .env soubor
└── README.md           # Dokumentace
```

## 🛠️ Instalace
### Požadavky
- Go 1.20 nebo novější
- PostgreSQL databáze
- Nainstalovaný [sqlc](https://docs.sqlc.dev/en/stable/)

### Kroky
1. **Klonování repozitáře**:
   ```bash
   git clone https://github.com/username/chirpy.git
   cd chirpy
   ```

2. **Nastavení prostředí**:
   Zkopírujte `.env.example` do `.env` a upravte podle potřeby:
   ```bash
   cp .env.example .env
   ```

3. **Spuštění migrací**:
   Spusťte migrace pro inicializaci databáze:
   ```bash
   make migrate
   ```

4. **Instalace závislostí**:
   ```bash
   go mod tidy
   ```

5. **Spuštění serveru**:
   ```bash
   go run cmd/main.go
   ```

Server bude spuštěn na `http://localhost:8080`.

## 🔑 API
Chirpy API používá JWT pro autentizaci a API klíče pro správu webhooků.

### Přehled endpointů
| Metoda   | Endpoint                 | Popis                          |
|----------|--------------------------|--------------------------------|
| `POST`   | `/api/users`             | Vytvoření nového uživatele     |
| `PUT`    | `/api/users`             | Aktualizace emailu a hesla     |
| `GET`    | `/api/chirps`            | Získání seznamu chirpů         |
| `POST`   | `/api/chirps`            | Vytvoření nového chirpu        |
| `PUT`    | `/api/chirps/{chirpID}`  | Aktualizace chirpu (Chirpy Red)|
| `DELETE` | `/api/chirps/{chirpID}`  | Smazání chirpu                 |
| `POST`   | `/api/polka/webhooks`    | Správa členství Chirpy Red     |

### Příklad API požadavku
```bash
curl -X POST http://localhost:8080/api/chirps -H "Authorization: Bearer YOUR_JWT_TOKEN" -H "Content-Type: application/json" -d '{
  "body": "Hello Chirpy!"
}'
```

## 🧪 Testování
Pro testování CLI příkazů a funkcí:
```bash
make test
```


Happy Chirping! 🐦
