package models

import(
	"RaysGo/helpers"
)


func GetNodeType(name string) int64 {
	nodeType := &NodeType{Name : name}
	if has, e := Engine.Get(nodeType); has && e == nil{
		return nodeType.Id
	}

	return 0
}

func AddNodeType(typeName string, description string) *NodeType{
	nodeType := &NodeType{Name: typeName, Description : description}
	if _, e := Engine.Insert(nodeType); e == nil{
		return nodeType
	} else {
		return nil
	}
}

func GetNode(id int64) *Node{
	node := &Node{Id: id}
	if has, e := Engine.Cascade(true).Get(node); has && e == nil{
		return node
	}
	return nil
}

func (this *Node) ParseContent() string{
	if this.ContentType == CONTENT_TYPE_MARKDOWN {
		return string(helpers.Markdown([]byte(this.Content)))
	}
	return this.Content
}

func GetStatusName(status int) string{
	switch status {
	case DRAFT:
		return "draft"
	case PUBLISHED:
		return "published"
	case PRIVATE:
		return "private"
	}
	return "published"
}

func GetStatus(name string) int{
	switch name{
	case "published":
		return PUBLISHED 
	case "private":
		return PRIVATE
	case "draft":
		return DRAFT
	}
	return PUBLISHED
}