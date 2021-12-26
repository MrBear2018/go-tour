package model

import "github.com/jinzhu/gorm"

type Tag struct {
	*Model        // TODO  结构体嵌套时， 用指针和用结构体本身的差别？
	Name   string `json:"name"`
	State  uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

/*
新增行为的回调函数
计算当前与当前 tag 的 name 和 state 相同的条数
*/
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 目前这种查询方法等同于语句 SELECT * FROM `blog_tag`  WHERE ...  LIMIT pageSize OFFSET pageOffset；
// 当表格很大的时候，会出现慢查询，因为会做全表扫描,速度会很慢 且 有的数据库结果集返回不稳定 Limit限制的是从结果集的M位置处取出N条输出 ,其余抛弃.
// 参考资料 https://zhuanlan.zhihu.com/p/92552787   https://blog.csdn.net/it_erge/article/details/105416718
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
