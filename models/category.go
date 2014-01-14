package models

import (
	"fmt"
)

// TODO

const (
	TAGS = "tags"
)

func GetNodeTags(uid int64, nid int64) (*[]CategoryTerm, error) {
	tags := make([]CategoryTerm, 0)
	var (
		cid int64
		err error
	)
	if cid, err = createTagCategory(uid); err == nil {
		if nid != 0 {
			if err = Engine.Join("inner", "node_category_term", "node_category_term.tid=category_term.id").Where("category_term.uid = ? and node_category_term.nid = ? and category_term.cid = ?", uid, nid, cid).Find(&tags); err == nil {
			} else {
				fmt.Println("error")
				fmt.Println(err)
			}
		} else {
			err = Engine.Distinct("Name").Where("category_term.uid = ? and category_term.cid = ?", uid, cid).Join("inner", "node_category_term", "node_category_term.tid=category_term.id").OrderBy("node_category_term.weight").Find(&tags)
		}

	}
	return &tags, err
}

func GetUserTags(uid int64) (*[]CategoryTerm, error) {
	return GetNodeTags(uid, 0)
}

// Add tags to post
func AddTags(uid int64, nid int64, tags []string) error {
	if len(tags) == 0 {
		return nil
	}
	var err error = nil
	var cid int64
	if cid, err = createTagCategory(uid); err == nil {
		nodeTags, _ := GetNodeTags(uid, nid)

		//ids := make([]int64, 0)
		for _, v := range *nodeTags {
			//ids = append(ids, v.Id)
			Engine.Delete(&NodeCategoryTerm{Tid: v.Id})
		}
		// cannot use ids (type []int64) as type []interface {} in function argument
		// Engine.In("tid", ids...).Delete(&NodeCategoryTerm{})

		var weight int64 = 0
		for _, t := range tags {
			if err = createOrUpdateTag(uid, nid, cid, t, weight); err != nil {
				return err
			}
			weight = weight + 1
		}
	}
	return err
}

func createOrUpdateTag(uid int64, nid int64, cid int64, name string, weight int64) error {
	var err error
	var term CategoryTerm
	if term, err = createCategoryTerm(uid, cid, 0, name, 0); err == nil {
		tag := NodeCategoryTerm{}
		if has, e := Engine.Where("Nid = ? AND Tid = ?", nid, term.Id).Get(&tag); has {
			tag.Weight = weight
			_, err = Engine.Id(tag.Id).Cols("weight").Update(&tag)
		} else if e == nil {
			_, err = Engine.Insert(&NodeCategoryTerm{Nid: nid, Tid: term.Id, Weight: weight})
		}
	}
	return err
}

func ExistCategoryTerm(uid int64, parentId int64, name string) (bool, error) {
	term := CategoryTerm{Uid: uid, Name: name, Pid: parentId}
	has, err := Engine.Get(&term)
	return has, err
}

func createCategoryTerm(uid int64, cid int64, parentId int64, name string, weight int64) (CategoryTerm, error) {
	term := CategoryTerm{Uid: uid, Cid: cid, Name: name, Pid: parentId}
	var err error
	if has, _ := Engine.Get(&term); has {
		return term, nil
	} else {
		term.Weight = weight
		if _, err = Engine.Insert(&term); err == nil {
			return term, nil
		}
	}
	return term, err
}

// Create tag category for user if there's no tag category
func createTagCategory(uid int64) (int64, error) {
	category := Category{Uid: uid, Name: TAGS}
	if has, err := Engine.Get(&category); !has && err == nil {
		category.Description = "User tags for post"
		if _, e := Engine.Insert(&category); e != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	return category.Id, nil
}
