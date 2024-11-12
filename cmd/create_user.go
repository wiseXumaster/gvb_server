package cmd

import (
	"bufio"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils"
	"os"
	"strings"
)

func CreateUser(permissions string) {
	// 创建用户的逻辑
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)

	// 使用 bufio.NewReader 读取用户输入
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("请输入用户名：")
	userName, _ = reader.ReadString('\n')
	userName = strings.TrimSpace(userName)

	fmt.Printf("请输入昵称（可选）：")
	nickName, _ = reader.ReadString('\n')
	nickName = strings.TrimSpace(nickName)
	if nickName == "" {
		nickName = "默认昵称" // 设置默认昵称
	}

	fmt.Printf("请输入邮箱（可选）：")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email == "" {
		email = "default@example.com" // 设置默认邮箱
	}

	fmt.Printf("请输入密码：")
	password, _ = reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Printf("请再次输入密码：")
	rePassword, _ = reader.ReadString('\n')
	rePassword = strings.TrimSpace(rePassword)

	// 判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		// 存在
		global.Log.Error("用户名已存在，请重新输入")
		return
	}

	// 校验两次密码
	if password != rePassword {
		global.Log.Error("两次密码不一致，请重新输入")
		return
	}

	// 对密码进行 hash
	hashPwd := utils.HashPwd(password)

	// 设定角色
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}

	// 设置头像：默认头像
	avatar := "/uploads/avatar/default.png"

	// 入库
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功!", userName)
}
