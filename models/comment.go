package models

import(
	"time"
)

// TODO

func AddComment(title string, content string, contentType string, uid int64, postId int64, pid int64, status int) error {
	comment := Comment{
		IUser: User{Id: uid}, 
		Pid: pid, 
		INode: Node{Id : postId}, 
		Title : title, 
		Content : content, 
		ContentType : contentType, 
		CreateTime : time.Now(), 
		Status : status,
	}
	var err error = nil
	_, err = Engine.Insert(&comment)
	return err
}

func GetNodeComments(nid int64) ([]Comment, error){
	comments := make([]Comment, 0)
	err := Engine.Where("Nid = ? AND Status = ?", nid, ACTIVE).Cascade(true).Find(&comments)
	return comments, err
}

func ActiveComment(id int64) error {
	return nil
}

func BlockComment(id int64) error {
	return nil
}

func DelComment(id int64) error {
	return nil
}
