package system

import (
	"errors"
	"web-scaffold/internal/core/config"
	"web-scaffold/internal/core/constant"
	"web-scaffold/pkg/encrypt"
	"web-scaffold/pkg/logger"

	systemmodel "web-scaffold/internal/model/system"
	systemrepository "web-scaffold/internal/repository/system"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/28 14:35
// @Desc:	系统用户服务实现

// CreateUser 创建用户
func CreateUser(username, password, name, email, phone, avatar, role string) error {
	preUser, err := systemrepository.FindUserByUsername(username)
	if err != nil && err.Error() != "record not found" {
		logger.Errorf("[%d] %s\n", constant.SERVICE_FIND_DATA_ERROR, err.Error())
		return err
	}

	if preUser != nil {
		return errors.New("username already exists")
	}

	password = encrypt.Sha256String(password, config.Config.SecretKey)
	newUser := &systemmodel.User{
		Username: username,
		Password: password,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Avatar:   avatar,
		Role:     role,
		IsActive: true,
	}

	err = systemrepository.CreateUser(newUser)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_CREATE_DATA_ERROR, err.Error())
		return err
	} else {
		return nil
	}
}

// DeleteUser 删除用户
func DeleteUser(userID uint) error {
	user, err := systemrepository.FindUserByID(userID)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_FIND_DATA_ERROR, err.Error())
		return err
	}

	user.IsActive = false
	user.JwtSign = "-"
	err = systemrepository.UpdateUser(user)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_UPDATE_DATA_ERROR, err.Error())
		return err
	}

	err = systemrepository.SoftDeleteUser(user)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_DELETE_DATA_ERROR, err.Error())
		return err
	} else {
		return nil
	}
}

// UpdateUser 更新用户
func UpdateUser(userID uint, username, password, name, email, phone, avatar, role, jwtSign string) error {
	user, err := systemrepository.FindUserByID(userID)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_FIND_DATA_ERROR, err.Error())
		return err
	}

	if username != "" && username != user.Username {
		existUser, _ := systemrepository.FindUserByUsername(username)
		if existUser != nil {
			return errors.New("username already exists")
		}
		user.Username = username
	}
	if password != "" {
		user.Password = encrypt.Sha256String(password, config.Config.SecretKey)
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
	if role != "" {
		user.Role = role
	}
	if jwtSign != "" {
		user.JwtSign = jwtSign
	}

	err = systemrepository.UpdateUser(user)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_UPDATE_DATA_ERROR, err.Error())
		return err
	} else {
		return nil
	}
}

// FindUserByID 根据ID查询用户
func FindUserByID(userID uint) (*systemmodel.User, error) {
	user, err := systemrepository.FindUserByID(userID)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_FIND_DATA_ERROR, err.Error())
		return nil, err
	} else {
		return user, nil
	}
}

// FindUserByUsername 根据Username查询用户
func FindUserByUsername(username string) (*systemmodel.User, error) {
	user, err := systemrepository.FindUserByUsername(username)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_FIND_DATA_ERROR, err.Error())
		return nil, err
	} else {
		return user, nil
	}
}

// FindUserByName 根据Name查询用户
func FindUserByName(name string) ([]systemmodel.User, error) {
	users, err := systemrepository.FindUserByName(name)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_FIND_DATA_ERROR, err.Error())
		return nil, err
	} else {
		return users, nil
	}
}

// FindUserListWithPage 分页查询用户列表
func FindUserListWithPage(page, size int) ([]systemmodel.User, int64, error) {
	users, total, err := systemrepository.FindUserListWithPage(page, size)
	if err != nil {
		logger.Errorf("[%d] %s\n", constant.SERVICE_FIND_DATA_ERROR, err.Error())
		return nil, 0, err
	} else {
		return users, total, nil
	}
}
