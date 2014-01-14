package models

// TODO

const (
	TAGS = "tags"
)

func GetNodeTags(uid int64, nid int64) (*[]CategoryTerm, error){
	tags := make([]CategoryTerm, 0)
	var (
		cid int64
		err error
	)
	if cid, err = createTagCategory(uid); err == nil {
		if nid != 0 {
			err = Engine.Where("Uid = ? and Nid = ? and Cid = ?", uid, nid, cid).Find(&tags)
		} else {
			err = Engine.Where("Uid = ? and Cid = ?", uid, cid).Find(&tags);
		}
		
	}
	return &tags, err
}

func GetUserTags(uid int64) (*[]CategoryTerm, error){
	return GetNodeTags(uid, 0)
}

// Add tags to post
func AddTags(uid int64, pid int64, tags []string) (*[]CategoryTerm, error){
	_tags := make([]CategoryTerm, 0)
	if len(tags) == 0 {
		return &_tags, nil
	}
	var err error
	var cid int64
	if cid, err = createTagCategory(uid); err == nil {
		Engine.Where("Uid = ? and Nid = ? and Cid = ?", uid, pid, cid).Delete(&CategoryTerm{})
		var weight int64 = 0
		for _, t := range tags{
			term := CategoryTerm{
				Name : t, Uid : uid, Cid : cid, Pid : 0, Nid : pid, Weight : weight,
			}
			_tags = append(_tags, term)
			weight = weight + 1
		}
		if _, err = Engine.Insert(&_tags); err == nil {
			return &_tags, nil
		}
	}
	return &_tags, err
}

func ExistCategoryTerm(uid int64, nid int64, parentId int64, name string) (bool, error){
	term := CategoryTerm{Uid : uid, Name : name, Pid : parentId, Nid : nid}
	has, err :=  Engine.Get(&term)
	return has, err
}

// Create tag category for user if there's no tag category
func createTagCategory(uid int64) (int64, error){
	category := Category{Uid : uid, Name : TAGS} 
	if has, err := Engine.Get(&category); !has && err == nil {
		category.Description = "User tags for post"
		if _, e := Engine.Insert(&category); e != nil {
			return 0, err
		}
	} else if err != nil  {
		return 0, err
	}

	return category.Id, nil
}