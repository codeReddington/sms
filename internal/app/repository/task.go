package repository

import (
	"errors"
	"go.uber.org/zap"
	"go_fiber/database"
	"go_fiber/internal/app/model"
	"go_fiber/logger"
	"gorm.io/gorm"
)

func GetTasks() ([]model.Task, error) {

	var tasks []model.Task
	err := database.Db.Find(&tasks).Error

	if err != nil {
		logger.Error("Error retrieving tasks :: ", zap.Error(err))
	}

	return tasks, nil
}

func GetTasksByID(taskID string) (model.Task, error) {
	var task model.Task

	err := database.Db.First(&task, taskID).Error
	if err != nil {
		logger.Error("Error retrieving tasks :: ", zap.Error(err))
	}

	return task, nil
}

func CreateTask(task model.Task) bool {

	if err := database.Db.Create(&task).Error; err != nil {
		if err != nil {
			logger.Error("Error creating task :: ", zap.Error(err))
			return false
		}
	}

	return true
}

func UpdateTask(task model.Task) bool {
	var exists bool

	err := database.Db.Model(&model.Task{}).
		Select("count(*) > 0").
		Where("id = ?", task.ID).
		Find(&exists).
		Error

	if err != nil {
		logger.Error("Error executing task ID query :: ", zap.Error(err))
		return false
	}

	if !exists {
		logger.Error("Error validating task ID :: ", zap.Error(err))
		return false
	}

	if err := database.Db.Model(&model.Task{}).Where("id = ?", task.ID).Updates(&task).Error; err != nil {
		logger.Error("Error executing update query :: ", zap.Error(err))
		return false
	}

	return true
}

func DeleteTask(taskID string) bool {
	var task model.Task

	if err := database.Db.First(&task, taskID).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error("Error checking if task ID exists :: ", zap.Error(err))
			return false
		}

		logger.Error("Error while validating task ID :: ", zap.Error(err))
		return false
	}

	if err := database.Db.Delete(&task, taskID).Error; err != nil {
		logger.Error("Error deleting task :: ", zap.Error(err))
		return false
	}

	return true
}
