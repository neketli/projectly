package entity

type StatisticData struct {
	ID                  int     `json:"id" example:"1"`
	Code                string  `json:"code" example:"proj1"`
	TotalTasksCount     int     `json:"total_tasks_count" example:"10"`
	CompletedTasksCount int     `json:"completed_tasks_count" example:"5"`
	AvgTaskDuration     float64 `json:"avg_task_duration" example:"2.5"`
	Cluster             int     `json:"cluster" example:"1"`
}
