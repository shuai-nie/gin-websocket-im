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

func (u UserBaseic) List() ([]*UserBaseic, error) {
	userList := make([]*UserBaseic, 10)
	err := common.DB.Model(UserBaseic{}).Where("id <> 0").Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func (u UserBaseic)Create(user UserBaseic) error {
	return common.DB.MOdel(u).Create(&user).Error
}

func (u UserBaseic) Delete() error {
	return common.DB.Model(u).Delete(u).Error
}

func (u UserBaseic) Update(user UserBaseic) error {
	return common.DB.Model(u).Update(&user).Error
}