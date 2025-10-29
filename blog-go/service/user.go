package service

import (
	"blog-go/global"
	"blog-go/model/appTypes"
	"blog-go/model/database"
	"blog-go/model/other"
	"blog-go/model/request"
	"blog-go/model/response"
	"blog-go/utils"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) Register(user database.User) (database.User, error) {
	if !errors.Is(global.DB.Where("email = ?", user.Email).First(&database.User{}).Error, gorm.ErrRecordNotFound) {
		return database.User{}, errors.New("this email address is already registered, please check the information you filled in, or retrieve your password")
	}

	user.Password = utils.BcryptHash(user.Password)
	user.UUID = uuid.Must(uuid.NewV4())
	user.Avatar = "/image/avatar.jpg"
	user.Register = appTypes.Email

	if err := global.DB.Create(&user).Error; err != nil {
		return database.User{}, err
	}

	return user, nil

}

func (s *UserService) EmailLogin(u database.User) (database.User, error) {
	var user database.User
	err := global.DB.Where("email = ?", u.Email).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return database.User{}, errors.New("incorrect email or password")
		}
		return user, nil
	}
	return database.User{}, err
}

func (s *UserService) ForgotPassword(req request.ForgotPassword) error {
	var user database.User
	if err := global.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return err
	}
	user.Password = utils.BcryptHash(req.NewPassword)
	return global.DB.Save(&user).Error
}

func (s *UserService) UserCard(req request.UserCard) (response.UserCard, error) {
	var user database.User
	if err := global.DB.Where("uuid = ?", req.UUID).Select("uuid", "username", "avatar", "address", "signature").First(&user).Error; err != nil {
		return response.UserCard{}, err
	}
	return response.UserCard{
		UUID:      user.UUID,
		Username:  user.Username,
		Avatar:    user.Avatar,
		Address:   user.Address,
		Signature: user.Signature,
	}, nil
}

func (s *UserService) Logout(c *gin.Context) {
	uuid := utils.GetUUID(c)
	jwtStr := utils.GetRefreshToken(c)
	utils.ClearRefreshToken(c)
	global.Redis.Del(uuid.String())
	_ = ServiceGroupApp.JwtService.JoinInBlacklist(database.JwtBlacklist{
		Jwt: jwtStr,
	})
}
func (s *UserService) UserResetPassword(req request.UserResetPassword) error {
	var user database.User

	if err := global.DB.Take(&user, req.UserID).Error; err != nil {
		return err
	}

	if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
		return errors.New("original password does not match the current account")
	}

	user.Password = req.NewPassword
	return global.DB.Save(&user).Error

}
func (s *UserService) UserInfo(userID uint) (database.User, error) {
	var user database.User

	if err := global.DB.Take(&user, userID).Error; err != nil {
		return database.User{}, err
	}

	return user, nil
}

func (s *UserService) UserChangeInfo(req request.UserChangeInfo) error {
	var user database.User

	if err := global.DB.Take(&user, req.UserID).Error; err != nil {
		return err
	}

	return global.DB.Model(&user).Updates(req).Error
}

func (s *UserService) UserWeather(ip string) (string, error) {
	result, err := global.Redis.Get("weather-" + ip).Result()
	if err != nil {
		// todo
	}
	return result, nil
}

func (s *UserService) UserChart(req request.UserChart) (response.UserChart, error) {
	where := global.DB.Where(fmt.Sprintf("date_sub(curdate(), interval %d day) <= created_at", req.Date))
	var res response.UserChart
	startDate := time.Now().AddDate(0, 0, -req.Date)
	for i := 1; i <= req.Date; i++ {
		res.DateList = append(res.DateList, startDate.AddDate(0, 0, i).Format("2006-01-02"))
	}

	loginCounts := utils.FetchDateCounts(global.DB.Model(&database.Login{}), where)

	registerCounts := utils.FetchDateCounts(global.DB.Model(&database.User{}), where)

	for _, date := range res.DateList {
		loginCount := loginCounts[date]
		registerCounts := registerCounts[date]
		res.LoginData = append(res.LoginData, loginCount)
		res.RegisterData = append(res.RegisterData, registerCounts)
	}

	return res, nil

}

func (s *UserService) UserList(req request.UserList) (interface{}, int64, error) {
	db := global.DB
	if req.UUID != nil {
		db = db.Where("uuid = ?", req.UUID)
	}
	if req.Register != nil {
		db = db.Where("register = ?", req.Register)
	}

	option := other.MySQLOption{
		PageInfo: req.PageInfo,
		Where:    db,
	}
	return utils.MySQLPagination(&database.User{}, option)
}
func (s *UserService) UserFreeze(req request.UserOperation) error {
	var user database.User
	if err := global.DB.Take(&user, req.ID).Update("freeze", true).Error; err != nil {
		return err
	}
	jwtStr, _ := ServiceGroupApp.JwtService.GetRedisJwt(user.UUID)
	if jwtStr != "" {
		_ = ServiceGroupApp.JwtService.JoinInBlacklist(database.JwtBlacklist{Jwt: jwtStr})
	}
	return nil
}
func (s *UserService) UserUnfreeze(req request.UserOperation) error {
	return global.DB.Take(&database.User{}, req.ID).Update("freeze", true).Error
}

func (s *UserService) UserLoginList(req request.UserLoginList) (interface{}, int64, error) {
	db := global.DB

	if req.UUID != nil {
		var userID uint
		if err := global.DB.Model(&database.User{}).Where("uuid = ?", req.UUID).Pluck("id", &userID).Error; err != nil {
			return nil, 0, nil
		}
		db.Where("user_id = ?", userID)
	}

	option := other.MySQLOption{
		PageInfo: req.PageInfo,
		Where:    db,
		Preload:  []string{"User"},
	}

	return utils.MySQLPagination(&database.Login{}, option)

}
