# 🌟 TaskBoard(в разработке) 🌟

![Go](https://img.shields.io/badge/Go-1.18%2B-blue)
![Gin](https://img.shields.io/badge/Gin-%E2%9C%94%EF%B8%8F-green)
![GORM](https://img.shields.io/badge/GORM-%E2%9C%94%EF%B8%8F-green)
![Docker](https://img.shields.io/badge/Docker-%E2%9C%94%EF%B8%8F-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-%E2%9C%94%EF%B8%8F-blue)

TaskBoard — это backend для проекта по управлению задачами на проекте, похожий на GitHub issues, написанный на Go. 

## 🚀 Возможности

- 🔒 Аутентификация и авторизация пользователей
- 📝 Операции CRUD для задач
- 👥 Назначение задач и отслеживание статуса
- 🌐 RESTful API
- 🐳 Поддержка Docker
- 🔄 Использование CompileDaemon для hot reload

## 🛠️ Технологии

- **Go**
- **Gin**
- **GORM**
- **Docker**
- **PostgreSQL**

## 🚀 Начало работы

### 📋 Предварительные требования

- Docker
- Docker Compose

### 🛠️ Установка

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/Permyakov-Dmitriy/TaskBoard.git
   cd TaskBoard
   ```

2. Создайте файл .env и настройте переменные окружения;

3. Соберите и запустите:
   ```bash
   docker compose up --build
   ```
