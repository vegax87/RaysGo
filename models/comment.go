package models

func GetNodeComments(nid int64) ([]Comment, error) {
	comments := make([]Comment, 0)
	err := Engine.Where("Nid = ? AND Status = ?", nid, ACTIVE).OrderBy("Id desc").Cascade(true).Find(&comments)
	return comments, err
}

func (this *Comment) ActiveComment() error {
	this.Status = ACTIVE
	_, err := Engine.Id(this.Id).Update(this)
	return err
}

func (this *Comment) BlockComment(id int64) error {
	this.Status = BLOCKED
	_, err := Engine.Id(this.Id).Update(this)
	return err
}
