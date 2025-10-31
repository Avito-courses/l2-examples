# Примеры по воркшопу подключение БД

Разверните БД по инструкции  
https://stepik.org/lesson/1882164/step/1?unit=1907580

## Работа с goose из терминала
Установка  
``go install github.com/pressly/goose/v3/cmd/goose@latest``

Установка переменных окружения
```
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="postgres://myuser:mypassword@localhost:5432/test_db"
```

Создание миграции  
``goose create init_schema sql``

Применение и откат миграций
```
goose up
goose down
```

