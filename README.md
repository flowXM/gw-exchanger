Запуск сервиса
```bash
go run cmd/main.go -c config.env
```

Миграция базы данных
```bash
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE TYPE currency AS ENUM ('RUB', 'USD', 'EUR');

    CREATE TABLE IF NOT EXISTS exchange_rates (
        currency currency PRIMARY KEY,
        rate numeric NOT NULL
    );


    INSERT INTO
        exchange_rates
    VALUES
        ('RUB', 1),
        ('USD', 84.0981),
        ('EUR', 90.648);    
EOSQL
```