package controllers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"gitlab.com/jaderabbit/go-rabbit/tee/task"
)

// TeeTaskController for tee task creates, executes, and various queries
type TeeTaskController struct {
	beego.Controller
}

// URLMapping if you register this function, you will use the interface to execute the function. The performance will
// be improved a lot.
func (o *TeeTaskController) URLMapping() {
	o.Mapping("Create", o.Create)
	o.Mapping("Get", o.Get)
	o.Mapping("GetAll", o.GetAll)
	o.Mapping("Execute", o.Execute)
}

// Create a tee task to share data
// @Title Create
// @Description Create a tee task to calculate the data, the task creator should provide all the data ID and algorithm ID, and the location where the result is stored
// @Param	data_id[]		formData 	[]string	true		"All shared data IDs, this shared data stores the raw data address."
// @Param	algorithm_id	formData 	string		true		"Shared algorithm data ID."
// @Param	result_address	formData 	string		true		"The result of the trusted calculation is the address."
// @Success 200 {string} tee.SharedData.ID
// @router / [post]
func (o *TeeTaskController) Create() {
	var req task.CreateRequest
	if err := o.ParseForm(&req); err != nil {
		responseJSON(o.Controller, false, err.Error())
		return
	}

	if len(req.DataIDs) == 1 && strings.Contains(req.DataIDs[0], ",") {
		req.DataIDs = strings.Split(req.DataIDs[0], ",")
	}

	logger.Infof("Create a tee task by req: %v", req)
	taskID, err := task.Create(&req)
	if err != nil {
		responseJSON(o.Controller, false, err.Error())
		return
	}

	responseJSON(o.Controller, true, taskID)
}

// GetAll get all the tee task
// @Title GetAll
// @Description get all the tee task
// @Success 200 {object} []tee.Task
// @router / [get]
func (o *TeeTaskController) GetAll() {
	tasks, err := task.GetAll()
	if err != nil {
		responseJSON(o.Controller, false, err.Error())
		return
	}

	responseJSON(o.Controller, true, tasks)
}

// Get a tee task
// @Title Get
// @Description Get a tee task by task id
// @Param	id		path 	string	true		"Task id for tee task. It is generated by the server after the task is created."
// @Success 200 {object} tee.Task
// @router /:id [get]
func (o *TeeTaskController) Get() {
	id := o.GetString(":id")

	task, err := task.GetByID(id)
	if err != nil {
		responseJSON(o.Controller, false, err.Error())
		return
	}

	responseJSON(o.Controller, true, task)
}

// Execute the tee task with algorithm
// @Title Execute
// @Description Execute the tee task with algorithm
// @Param	task_id				formData 	string		true	"Task id for tee task. It is generated by the server after the task is created."
// @Param	executor			formData 	string		true	"The executor is the public key that performs this trusted execution task."
// @Param	container_type		formData 	string		true	"The choice of trusted execution environment, 0-docker 1-subdocker, default is 0."
// @Param	hash				formData 	string		false	"All parameters except the signature are sequentially connected to obtain a hash."
// @Param	signatures			formData 	[]string	false	"Signature of the data summary by the user's private key."
// @Success 200 string Successful execution
// @router / [put]
func (o *TeeTaskController) Execute() {
	var req task.ExecuteRequest
	if err := o.ParseForm(&req); err != nil {
		responseJSON(o.Controller, false, err.Error())
		return
	}

	signatures := o.GetStrings("signatures")
	if len(signatures) == 0 {
		signatures = make([]string, 1)
	}

	signature, err := json.Marshal(signatures)
	if err != nil {
		responseJSON(o.Controller, false, err.Error())
		return
	}
	req.Signature = string(signature)

	result, err := task.Execute(&req)
	if err != nil {
		responseJSON(o.Controller, false, err.Error())
		return
	}

	responseJSON(o.Controller, true, result)
}