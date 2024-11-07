package views

import "commitinder/models"

type DiffRequest struct {
	Content string `json:"content"`
}

func (d *DiffRequest) ToModel() (diff models.Diff) {
	diff.Content = d.Content
	return diff
}

type DiffResponse struct {
	Id             int    `json:"id"`
	Content        string `json:"content"`
	CommitMessages []int  `json:"commit_messages"`
}

func (d *DiffResponse) FormModelToView(diff *models.Diff) {
	d.Id = diff.Id
	d.Content = diff.Content

	for _, commitMessage := range diff.CommitMessages {
		d.CommitMessages = append(d.CommitMessages, commitMessage.Id)
	}
}
