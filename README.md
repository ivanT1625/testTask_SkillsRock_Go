testTask_SkillsRock_Go - это REST API для управления задач. Приложение позволяет создавать, обновлять и удалять задачи
через HTTP - запросы.

Технологии:
- Язык программирования: Go (Golang)
- Фреймворк: Fiber
- База данных : PostgreSQL
- Драйвер PostgreSQL: pgx

Основные функции:
1. Создание задачи
      Метод: POST /api/tasks

2. Получение списка задач
      Метод: GET /api/tasks
   
3. Обновление задачи
      Метод: PUT /api/tasks/:id
   
4. Удаление задачи
      Метод: DELETE /api/tasks/:id


Установка и запуск
1. Требования
  - Установленный Go (версия 1.18 или выше)
  - Установленный PostgreSQL
  - Установленный Fiber (go get github.com/gofiber/fiber/v2)
  - Установленный pgx (go get github.com/jackc/pgx/v5)
    
2. Настройка базы данных
  2.1  Создайте базу данных "todo_db":
         psql -U postgres -c "CREATE DATABASE todo_db;"
  2.2  Создайте таблицу tasks, выполнив миграцию:
         psql -U postgres -d todo_db -f migrations/create_tasks_table.sql
    
3. Запуск приложения
  3.1  Перейдите в корень проекта
        cd путь_к_проекту
  3.2  Запустите приложение
        go run main.go
  3.3  Приложение будет доступно по адресу: http://localhost:3000.


Структура базы данных
  Таблица "tasks" содержит следующие поля:

     ПОЛЕ                       ТИП                         ОПИСАНИЕ
      id                  SERIAL PRIMARY KEY              Уникальный идентификатор задачи
    title                   TEXT NOT NULL                  Название задачи
    description                 TEXT                        Описание задачи
    status                    TEXT CHECK                   Статус задачи (new, in_progress, done)
    created_at                TIMESTAMP                    Время создания задачи
    updated_at                TIMESTAMP                    Время последнего обновления задачи
 
Автор 
Иван Т.
