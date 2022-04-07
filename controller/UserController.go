package controller

import (
	"fmt"
	"gin_demo2/common"
	"gin_demo2/dto"
	"gin_demo2/model"
	"gin_demo2/response"
	"gin_demo2/utils"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// 注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()

	var requestUser = model.User{}
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)

	// 获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	email := requestUser.Email
	password := requestUser.Password

	// 数据验证 -- 补充验证邮箱
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) <= 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 如果名称没有传就随机给一个10位的随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}

	log.Println(name, telephone, password)
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机已被注册")
		return
	}
	// 判断邮箱是否存在
	if isEmailExist(DB, email) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "邮箱已被注册")
		return
	}
	// 创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
		Email:     email,
	}
	DB.Create(&newUser)
	// 发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")

		log.Printf("token generate error : %v", err)
		return
	}
	// 返回结果
	response.Success(ctx, gin.H{
		"token": token,
	}, "注册成功")
}

// 登录
func Login(ctx *gin.Context) {

	DB := common.GetDB()
	// 获取参数
	requestUser := model.User{}
	ctx.Bind(&requestUser)
	email := requestUser.Email
	password := requestUser.Password
	// 数据验证
	// if len(email) != 11 {
	// 	response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
	// 	return
	// }
	if len(password) <= 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("email=?", email).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	// 发送token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")

		log.Printf("token generate error : %v", err)
		return
	}
	// 返回结果
	response.Success(ctx, gin.H{
		"token": token,
	}, "登录成功")

}

// 用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{
		"user": dto.ToUserDto(user.(model.User)),
	}, "获取信息成功")
}

// 联系方式是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)

	return user.ID != 0

}

// 邮箱是否存在
func isEmailExist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email=?", email).First(&user)
	return user.ID != 0
}

// 图片上传--待完成
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded", file.Filename))
		return
	}

	filepath := path.Join("E:/项目代码/go代码/gin_demo2/resources/image", file.Filename)
	r := gin.Default()
	r.StaticFS("E:/项目代码/go代码/gin_demo2/resources/image", http.Dir("/image"))
	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"uploading": "done", "message": "success", "url": "http://" + c.Request.Host + "/resources/image/" + file.Filename})

}
