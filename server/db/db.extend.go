package db

import "fmt"

func (m *DatasetVersionData) HasParent() bool{
	return len(m.ParentUid)>0
}


func (t *DatasetData) Caption() string {
	if len(t.Title) > 0 {
		return t.Title
	}
	return t.Name
}

func (t *ProjectData) Caption() string {
	if len(t.Title) > 0 {
		return t.Title
	}
	return t.Name
}

func (t *Job) Caption() string {
	if len(t.Title) > 0 {
		return t.Title
	}
	return t.Name
}

func (t *JobRunData) Caption() string {
	if len(t.Title) > 0 {
		return t.Title
	}
	return fmt.Sprintf("#%d", t.RunNum)
}

func (t *SystemData) Caption() string {


	if len(t.Title) > 0 {
		return t.Title
	}
	return t.Name
}