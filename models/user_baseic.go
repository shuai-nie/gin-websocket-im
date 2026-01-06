package models

type UserBaseic struct {
	gorm.Model
	Name string
	Password string
	Phone string
	Email string
	Identity string
	ClientIp string
	ClientPort string
	LoginTime uint64
	HeartbeatTime uint64
	LogOutTime uint64
	IsLogout bool
	DeviceInfo string
}

func (table *UserBaseic) TableName() string {
	return "user_basic"
}

