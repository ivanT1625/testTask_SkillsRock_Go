package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"todo-api/db"
	"todo-api/models"
)

// Создание новой задачи
func CreateTask(c *fiber.Ctx) error {
	var task models.Task

	var err = c.BodyParser(&task)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	query :=
		`INSERT INTO tasks (title, description, status)
         VALUES ($1, $2, $3)
         RETURNING id, title, description, status, created_at, updated_at`

	var row = db.DB.QueryRow(context.Background(), query, task.Title, task.Description, task.Status)

	var newTask models.Task

	err = row.Scan(&newTask.ID, &newTask.Title, &newTask.Description, &newTask.Status, &newTask.CreatedAt, &newTask.UpdatedAt)
	if err != nil {
		log.Printf("Ошибка создания задачи: %v\n", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось создать задачу"})
	}
	
	return c.JSON(newTask)
}

// Получение всех задач из бд
func GetTasks(c *fiber.Ctx) error {
	var sql = `SELECT id, title, description, status, created_at, updated_at FROM tasks`
	rows, err := db.DB.Query(context.Background(), sql)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось получить список задач"})
	}

	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось считать задачи"})
		}

		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

// Обновление существующей задачи по ID
func UpdateTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Неверный ID задачи"})
	}

	var task models.Task

	err = c.BodyParser(&task)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Неверный формат данных"})
	}

	var query = `
				UPDATE tasks
				SET title=$1, description=$2, status=$3, updated_at=now()
				WHERE id=$4
				RETURNING id, title, description, status, created_at, updated_at
			`

	var row = db.DB.QueryRow(context.Background(), query, task.Title, task.Description, task.Status, id)
	var updatedTask models.Task
	err = row.Scan(&updatedTask.ID, &updatedTask.Title, &updatedTask.Description, &updatedTask.Status, &updatedTask.CreatedAt, &updatedTask.UpdatedAt)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Задача не найдена"})
	}

	return c.JSON(updatedTask)
}

// DeleteTask удаляет задачу по ID
func DeleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Неверный ID задачи"})
	}

	var query = "DELETE FROM tasks WHERE id=$1"
	result, err := db.DB.Exec(context.Background(), query, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Не удалось удалить задачу"})
	}

	var rowAffected = result.RowsAffected()
	if rowAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Задача не найдена"})
	}

	return c.SendStatus(http.StatusNoContent)
}
