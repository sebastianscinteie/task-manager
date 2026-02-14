package updater

import (
	"task-manager/api/internal/database"
)

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UpdateTask(task CreateTaskRequest) {

	database.WriteToDatabase(task.Title)
}
