package job

type JobHashTable map[int]*CronJob

func (table *JobHashTable) SetJob(id int, job *CronJob) bool {
	tmp_job := (*table)[id]
	if tmp_job != nil {
		return false
	}
	(*table)[id] = job
	return true
}

func (list *JobHashTable) GetJob(id int) *CronJob {
	return (*list)[id]
}

func (list *JobHashTable) DelJob(id int) {
	delete(*list, id)
}
