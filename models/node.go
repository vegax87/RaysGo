package models

import (
	"RaysGo/helpers"
	"time"
)

func GetNodeType(name string) int64 {
	nodeType := &NodeType{Name: name}
	if has, e := Engine.Get(nodeType); has && e == nil {
		return nodeType.Id
	}

	return 0
}

func AddNodeType(typeName string, description string) *NodeType {
	nodeType := &NodeType{Name: typeName, Description: description}
	if _, e := Engine.Insert(nodeType); e == nil {
		return nodeType
	} else {
		return nil
	}
}

func GetNode(id int64) *Node {
	node := &Node{Id: id}
	if has, e := Engine.Cascade(true).Get(node); has && e == nil {
		return node
	}
	return nil
}

func (this *Node) ParseContent() string {
	if this.ContentType == CONTENT_TYPE_MARKDOWN {
		return string(helpers.Markdown([]byte(this.Content)))
	}
	return this.Content
}

func (this *Node) UpdateCounter() {
	if this.ICounter.Id > 0 {
		this.ICounter.DayCount += 1
		this.ICounter.WeekCount += 1
		this.ICounter.TotalCount += 1
		if helpers.ThisDate().After(this.ICounter.Timestamp) {
			this.ICounter.DayCount = 1
		}
		if helpers.ThisWeek().After(this.ICounter.Timestamp) {
			this.ICounter.WeekCount = 1
		}
		this.ICounter.Timestamp = time.Now()
		// TODO: 这里在xorm设置了缓存的情况下，更新数据出错，并没有每次都更新
		Engine.NoCache().Id(this.ICounter.Id).Update(&this.ICounter)
	} else {
		c := Counter{DayCount : 1,WeekCount : 1, TotalCount : 1, Timestamp : time.Now()}
		Engine.Insert(&c)
		this.ICounter = c
		Engine.NoCache().Id(this.Id).Update(this)
	}
}

func GetStatusName(status int) string {
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

func GetStatus(name string) int {
	switch name {
	case "published":
		return PUBLISHED
	case "private":
		return PRIVATE
	case "draft":
		return DRAFT
	}
	return PUBLISHED
}
