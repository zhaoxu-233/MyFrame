package domain

import "time"

//User领域对象，是DDD中的聚合根，即领域层面的entity
//也可以交BO（business object）
type User struct {
	Email    string
	PassWord string
	Name     string
	Brithday string
	Info     string
	//不需要confirmpasword
	CreateTime time.Time
}

//按照DDD模式，应该咋子domain中校验加密
