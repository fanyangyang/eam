package helper

import (
	"encoding/json"
	"github.com/fanyangyang/eam/models"
)

const (
	DEPARTMENT_NOT_SET      = "部门未分配"
	POSITION_NOT_SET        = "岗位未分配"
)

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// TODO 改进任务的结构，将经常会出现的任务（部门、岗位为分配）如何考虑
func AddTODO(todo models.TODO, userId int, name string) {
	created, id, err := models.Ormer.ReadOrCreate(&todo, "title") // title是作为条件，不是查询结果
	if err != nil {
		return
	}

	if created == false {
		todo.Id = int(id)
		task := Task{userId,name,DEPARTMENT_NOT_SET}
		bs,_ := json.Marshal(todo.Body)
		var tasks []Task
		json.Unmarshal(bs,&tasks)
		tasks = append(tasks, task)
		newTasks, _:= json.Marshal(tasks)
		todo.Body = string(newTasks)
		models.Ormer.Update(&todo)
	}
	return
}

func MakeTODOBody(userId int, name , desc string)[]Task{
	task := Task{userId,name,DEPARTMENT_NOT_SET}

	return []Task{task}
}
