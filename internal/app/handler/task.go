package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"go_fiber/internal/app/model"
	"go_fiber/internal/app/repository"
	"go_fiber/logger"
	"go_fiber/structs"
)

func GetTasks(c *fiber.Ctx) error {
	tasks, err := repository.GetTasks()

	if err != nil {
		HostHeaderInfo := structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "E001",
			ResponseMessage: "Oops.. something went wrong",
		}
		log.Panic(err)
		return c.JSON(HostHeaderInfo)
	}

	if len(tasks) == 0 {
		HostHeaderInfo := structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "000",
			ResponseMessage: "No tasks found",
		}
		return c.JSON(HostHeaderInfo)
	}

	TaskResponse := structs.TasksResponse{
		HostHeaderInfo: structs.HostHeaderInfo{
			Channel:         "000",
			ResponseCode:    "mobile",
			ResponseMessage: "success",
		},
		Tasks: nil,
	}

	var responseTask []structs.Tasks
	for _, task := range tasks {
		taskResponse := structs.Tasks{
			ID:          int(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		}
		responseTask = append(responseTask, taskResponse)
	}

	TaskResponse.Tasks = responseTask

	return c.JSON(TaskResponse)
}

func GetTasksByID(c *fiber.Ctx) error {
	taskID := c.Params("id")
	task, err := repository.GetTasksByID(taskID)

	if err != nil {
		HostHeaderInfo := structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "E001",
			ResponseMessage: "Oops.. something went wrong",
		}
		return c.JSON(HostHeaderInfo)
	}

	if task.ID == 0 {
		HostHeaderInfo := structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "000",
			ResponseMessage: "task not found",
		}
		return c.JSON(HostHeaderInfo)
	}

	TaskResponse := structs.TaskResponse{
		HostHeaderInfo: structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "000",
			ResponseMessage: "success",
		},
		Tasks: structs.Tasks{
			ID:          int(task.ID),
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		},
	}
	return c.JSON(TaskResponse)

}

func CreateTask(c *fiber.Ctx) error {

	var task model.Task
	var status bool

	if err := c.BodyParser(&task); err != nil {

		hostHeaderInfo := structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "002",
			ResponseMessage: "Oops.. something went wrong.",
		}
		logger.Error("Error parsing body :: ", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(hostHeaderInfo)
	}

	if err := task.Validate(); err != nil {
		logger.Error("Error validating required fields in body :: ", zap.Error(err))
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	status = repository.CreateTask(task)

	if !status {
		hostHeaderInfo := structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "002",
			ResponseMessage: "Oops.. something went wrong.",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(hostHeaderInfo)
	}

	hostHeaderInfo := structs.HostHeaderInfo{
		Channel:         "mobile",
		ResponseCode:    "000",
		ResponseMessage: "Task successfully created.",
	}
	return c.JSON(hostHeaderInfo)
}

func UpdateTask(c *fiber.Ctx) error {
	var task model.Task
	var status bool

	if err := c.BodyParser(&task); err != nil {
		hostHeaderInfo := structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "002",
			ResponseMessage: "Oops.. something went wrong.",
		}
		logger.Error("Error parsing body :: ", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(hostHeaderInfo)
	}

	if err := task.Validate(); err != nil {
		logger.Error("Error validating required fields in body :: ", zap.Error(err))
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	status = repository.UpdateTask(task)

	if !status {
		hostHeaderInfo := structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "002",
			ResponseMessage: "Oops.. something went wrong.",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(hostHeaderInfo)
	}

	hostHeaderInfo := structs.HostHeaderInfo{
		Channel:         "mobile",
		ResponseCode:    "000",
		ResponseMessage: "Task successfully updated.",
	}
	return c.JSON(hostHeaderInfo)
}

func DeleteTask(c *fiber.Ctx) error {
	var status bool
	taskID := c.Params("id")
	status = repository.DeleteTask(taskID)

	if !status {
		hostHeaderInfo := structs.HostHeaderInfo{
			Channel:         "mobile",
			ResponseCode:    "002",
			ResponseMessage: "Oops.. something went wrong.",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(hostHeaderInfo)
	}

	hostHeaderInfo := structs.HostHeaderInfo{
		Channel:         "mobile",
		ResponseCode:    "000",
		ResponseMessage: "Task successfully deleted.",
	}
	return c.JSON(hostHeaderInfo)
}
