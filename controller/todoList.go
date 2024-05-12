package controller

import (
	"fmt"
	"strconv"
	"todoList-golang/model"
	"todoList-golang/service"

	"github.com/gin-gonic/gin"
)

// @Summary	获取任务详情
// @Produce	json
// @Param		tag_id			body		string		false	"标签ID"
// @Param		title			body		string		false	"文章标题"
// @Param		desc			body		string		false	"文章简述"
// @Param		cover_image_url	body		string		false	"封面图片地址"
// @Param		content			body		string		false	"文章内容"
// @Param		modified_by		body		string		true	"修改者"
// @Success	200				{object}	model.Todo	"成功"
// @Failure	400				{object}	string		"请求错误"
// @Failure	500				{object}	string		"内部错误"
// @Router		/api/v1/Todo/{id} [put]
func GetTodoTask(c *gin.Context) {
	id := c.Param("id")
	todo, err := model.GetTodo(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 5000,
			"msg":  "failed",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
}

// @Summary	获取任务列表
// @Produce	json
// @Param		tag_id			body		string		false	"标签ID"
// @Param		title			body		string		false	"文章标题"
// @Param		desc			body		string		false	"文章简述"
// @Param		cover_image_url	body		string		false	"封面图片地址"
// @Param		content			body		string		false	"文章内容"
// @Param		modified_by		body		string		true	"修改者"
// @Success	200				{object}	model.Todo	"成功"
// @Failure	400				{object}	string		"请求错误"
// @Failure	500				{object}	string		"内部错误"
// @Router		/api/v1/Todo/{id} [put]
func GetTodoTaskList(c *gin.Context) {
	todo := model.Todo{}
	id, _ := c.GetQuery("id")
	todo_type, _ := c.GetQuery("type")
	// completed_time, _ := c.GetQuery("completed_time")
	if id != "" {
		id_, _ := strconv.Atoi(id)
		todo.ID = uint(id_)
	}
	if todo_type != "" {
		type_, _ := strconv.Atoi(todo_type)
		todo.TaskType = uint(type_)
	}
	fmt.Printf("ShouldBind::\nid:%d,\ntype:%d\n", todo.ID, todo.TaskType)
	todoList, err := model.GetAllTodo(&todo)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 5000,
			"msg":  "failed",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todoList,
		})
	}
}

// @Summary	创建任务
// @Produce	json
// @Param		tag_id			body		string		false	"标签ID"
// @Param		title			body		string		false	"文章标题"
// @Param		desc			body		string		false	"文章简述"
// @Param		cover_image_url	body		string		false	"封面图片地址"
// @Param		content			body		string		false	"文章内容"
// @Param		modified_by		body		string		true	"修改者"
// @Success	200				{object}	model.Todo	"成功"
// @Failure	400				{object}	string		"请求错误"
// @Failure	500				{object}	string		"内部错误"
// @Router		/api/v1/Todo/{id} [put]
func CreateTodoTask(c *gin.Context) {
	todo := model.Todo{}
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"code": 4001,
			"msg":  "请求参数错误",
			"data": nil,
		})
		return
	}
	err := service.CreateTodoTaskServer(&todo)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 5000,
			"msg":  err.Error(),
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": nil,
		})
	}
}

// @Summary	更新任务
// @Produce	json
// @Param		tag_id			body		string		false	"标签ID"
// @Param		title			body		string		false	"文章标题"
// @Param		desc			body		string		false	"文章简述"
// @Param		cover_image_url	body		string		false	"封面图片地址"
// @Param		content			body		string		false	"文章内容"
// @Param		modified_by		body		string		true	"修改者"
// @Success	200				{object}	model.Todo	"成功"
// @Failure	400				{object}	string		"请求错误"
// @Failure	500				{object}	string		"内部错误"
// @Router		/api/v1/Todo/{id} [put]
func UpdateTodoTask(c *gin.Context) {
	id := c.Param("id")
	todo := model.Todo{}
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"code": 4001,
			"msg":  "请求参数错误",
			"data": nil,
		})
		return
	}
	err := model.UpdateTodo(id, &todo)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 5000,
			"msg":  "update failed",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": nil,
		})
	}
}

// @Summary	删除任务
// @Produce	json
// @Param		tag_id			body		string		false	"标签ID"
// @Param		title			body		string		false	"文章标题"
// @Param		desc			body		string		false	"文章简述"
// @Param		cover_image_url	body		string		false	"封面图片地址"
// @Param		content			body		string		false	"文章内容"
// @Param		modified_by		body		string		true	"修改者"
// @Success	200				{object}	model.Todo	"成功"
// @Failure	400				{object}	string		"请求错误"
// @Failure	500				{object}	string		"内部错误"
// @Router		/api/v1/Todo/{id} [put]
func DeleteTodoTask(c *gin.Context) {
	id := c.Param("id")
	err := model.DeleteTodo(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 5000,
			"msg":  "delete failed",
			"data": nil,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": nil,
		})
	}
}
