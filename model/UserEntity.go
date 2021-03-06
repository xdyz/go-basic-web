package model

// 定义UserEntity结构体  设置字段各种属性，如主键，默认值，索引等，配置json格式名称 如果不需要json序列化，可以不设置
// 这里的 json 设置就是 当接口查询时，返回给前端的 对象数据中 的属性 是什么
type UserEntity struct {
	BasicEntity
	Username string `gorm:"type:varchar(50);column:username;not null;unique;comment:'用户名'" json:"username"`
	Password string `gorm:"type:varchar(50);not null;comment:'密码'" json:"password"`
	Nickname string `gorm:"type:varchar(50);column:nickname;not null;comment:'昵称'" json:"nickname"`
	Email    string `gorm:"type:varchar(50);not null;comment:'邮箱'" json:"email"`
	IsRoot   bool   `gorm:"type:tinyint(1);not null;default:0;comment:'是否是管理员'" json:"isRoot"`
}

// 表名
func (UserEntity) TableName() string {
	return "user"
}
