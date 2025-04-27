# 📎 URL Shortener Service

Простое приложение на **Go (Golang)** для сокращения ссылок с возможностью хранения статистики переходов.

## 🚀 Возможности

- Генерация коротких ссылок (автоматическая или с вашим кастомным алиасом)
    
- Перенаправление по короткой ссылке на оригинальный URL
    
- Сбор статистики переходов (время, IP-адрес, User-Agent, страна, устройство)
    
- Срок жизни короткой ссылки по умолчанию — **24 часа**
    
- API на базе **Gin**
    

## 🛠 Стек технологий

- **Golang** (Gin Web Framework)
    
- **PostgreSQL** (для хранения ссылок и статистики)
    
- **Docker** / **Docker Compose** (опционально)
    
- **Railway / GitHub Actions** для деплоя
    

---

## 📦 Установка и локальный запуск

### 1. Клонирование репозитория

```
git clone https://github.com/yourusername/urlshortener.git
cd urlshortener
```

### 2. Настройка окружения

Создайте файл `config/config.yaml`, если его нет. Пример настроек:

```
POSTGRES:
  POSTGRES_HOST: postgres
  POSTGRES_PORT: 5432
  POSTGRES_USER: root
  POSTGRES_PASSWORD: 1234
  POSTGRES_DB: postgres

REST_PORT: "8080"
```

Или используйте переменные окружения.

### 3. Запуск через Docker Compose

```
docker-compose up --build
```

Это поднимет:

- Базу данных PostgreSQL
    
- Сам сервис на Go
    

### 4. Билд вручную

Если хотите собрать локально:

```
go mod download
go build -o service ./cmd/main.go
./service
```

---

## 📚 API Документация

### 1. Создать короткую ссылку

**POST** `/api/short`

**Тело запроса (JSON):**

```
{
  "url": "https://example.com",
  "custom_alias": "myalias" // необязательный параметр
}
```

**Ответ:**

```
{
  "url": "https://example.com",
  "short_url": "myalias",
  "custom_alias": "myalias",
  "expireData": "2025-04-27T22:00:00Z"
}
```

---

### 2. Перейти по короткой ссылке

**GET** `/{short_link}`

- Происходит редирект на оригинальный URL.
    

---

### 3. Получить статистику по ссылке

**GET** `/api/stat/{short_link}`

**Ответ:**

```
{
  "short_id": "myalias",
  "clicks": 3,
  "details": [
    {
      "clicked": "2025-04-27T21:00:00Z",
      "ip_address": "192.168.1.1",
      "user_agent": "Mozilla/5.0",
      "country": "US",
      "device_type": "mobile"
    },
    ...
  ]
}
```

---

## 🐳 Быстрый деплой на Railway

1. Перейдите на [railway.app](https://railway.app/).
    
2. Нажмите **New Project** → **Deploy from GitHub Repo**.
    
3. Выберите этот репозиторий.
    
4. Railway автоматически определит Dockerfile и запустит приложение.
    

> ⚡️ Если нужны переменные окружения, настройте их в разделе **Variables**.

---

## 📄 Структура проекта

```
.
├── cmd/
│   └── main.go       # Точка входа в приложение
├── pkg/
│   ├── postgres/     # Работа с базой данных PostgreSQL
│   └── short/        # Генерация коротких ссылок
├── db/
│   └── migrations/   # SQL миграции для базы данных
├── config/
│   └── config.yaml   # Конфиг приложения
├── docker/
│   └── Dockerfile    # Инструкция по сборке образа
├── docker-compose.yml
├── README.md
└── go.mod
```

---

## ✨ Планы на будущее

- Аутентификация для создания приватных ссылок
    
- Веб-интерфейс
    
- Расширенная аналитика
    

---

## 🧑‍💻 Автор

- Telegram: @timurghub
    
- GitHub: [vizurth](https://github.com/vizurth)
