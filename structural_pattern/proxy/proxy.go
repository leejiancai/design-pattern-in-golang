package proxy

import "fmt"

// 定义接口
type IUserLogin interface {
	login(userName string, passwd string) bool
	logout(userName string)
}

// UserLogin是要被代理的类，它需要加入一些额外的功能，比如缓存、鉴权等非业务核心逻辑相关的功能，可以交给代理类实现
type UserLogin struct {}
func(user *UserLogin) login(userName string, passwd string) string {
	return userName+passwd
}
func(user *UserLogin) logout(userName string) {
	return
}

// UserLoginProxy1 基于继承的方式的代理。由于Go没有继承的，因此可以使用嵌入的方式模拟继承
type UserLoginProxy1 struct {
	UserLogin
}

func(user *UserLoginProxy1) login(userName string, passwd string) string {
	fmt.Println("Before login() called in UserLoginProxy1")
	ret := user.UserLogin.login(userName, passwd)
	fmt.Println("After login() called in UserLoginProxy1")
	return ret+" UserLoginProxy1"
}

func(user *UserLoginProxy1) logout(userName string){
	user.UserLogin.logout(userName)
	return
}

// 基于组合的代理模式
type UserLoginProxy2 struct {
	userLogin *UserLogin
}

func(user *UserLoginProxy2) login(userName string, passwd string) string {
	fmt.Println("Before login() called in UserLoginProxy2")
	ret := user.userLogin.login(userName, passwd)
	fmt.Println("After login() called in UserLoginProxy2")
	return ret+" UserLoginProxy2"
}

func(user *UserLoginProxy2) logout(userName string){
	user.userLogin.logout(userName)
	return
}







