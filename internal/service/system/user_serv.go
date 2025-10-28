package system

import (
	systemmodel "web-scaffold/internal/model/system"
	systemrepository "web-scaffold/internal/repository/system"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/28 14:35
// @Desc:	系统用户服务实现

// CreateUser 创建用户
func CreateUser(username, password, name, email, phone, avatar string) error {
	user := &systemmodel.User{
		Username: username,
		Password: password,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Avatar:   avatar,
		IsActive: true,
	}

	return systemrepository.CreateUser(user)
}

// DeleteUser 删除用户
func DeleteUser(userID uint) error {
	user, err := systemrepository.FindUserByID(userID)
	if err != nil {
		return err
	}

	return systemrepository.SoftDeleteUser(user)
}

// UpdateUser 更新用户
func UpdateUser(userID uint, username, password, name, email, phone, avatar string) error {
	user, err := systemrepository.FindUserByID(userID)
	if err != nil {
		return err
	}

	if username != "" {
		user.Username = username
	}
	if password != "" {
		user.Password = password
	}
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if phone != "" {
		user.Phone = phone
	}
	if avatar != "" {
		user.Avatar = avatar
	}

	return systemrepository.UpdateUser(user)
}

// FindUserByID 根据ID查询用户
func FindUserByID(userID uint) (*systemmodel.User, error) {
	return systemrepository.FindUserByID(userID)
}

// FindUserByUsername 根据Username查询用户
func FindUserByUsername(username string) (*systemmodel.User, error) {
	return systemrepository.FindUserByUsername(username)
}

// FindUserByName 根据Name查询用户
func FindUserByName(name string) ([]systemmodel.User, error) {
	return systemrepository.FindUserByName(name)
}

// GetUserListWithPage 分页查询用户列表
func GetUserListWithPage(page, pageSize int) ([]systemmodel.User, int64, error) {
	return systemrepository.FindUserListWithPage(page, pageSize)
}
