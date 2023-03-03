package repository

type UserRepository interface {
	Create()
}

type userRepository struct {
	//DB *gorm.DB
}

func (u *userRepository) Create() {
	//array := [6]int{2, 3, 5, 7, 11, 13}
}
