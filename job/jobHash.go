package job

type JobHashTable map[int]*CronJob

func (table *JobHashTable) SetJob(id int, job *CronJob) {
	(*table)[id] = job
}

func (list *JobHashTable) GetJob(id int) *CronJob {
	return (*list)[id]
}

func (list *JobHashTable) DelJob(id int) {
	delete(*list, id)
}
