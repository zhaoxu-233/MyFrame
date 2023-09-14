package web

import (
	"exercise_code/webook/internal/domain"
	"exercise_code/webook/internal/repository"
	"exercise_code/webook/internal/service"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"net/http"
)

//正则的变量声明
const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`

	userIdKey       = "userId"
	dateRegexPatter = "^([0-9]{4})-([0-9]{2})-([0-9]{2})$"
)

//在这个上边定义所以和user有关的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	dateExp     *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc:         svc,
		emailExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		dateExp:     regexp.MustCompile(dateRegexPatter, regexp.None),
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	//分组路由
	ug := server.Group("/users")
	ug.POST("signup", u.SignUp)
	ug.POST("login", u.LoginJWT)
	//ug.POST("login", u.Login)
	ug.POST("edit", u.Edit)
	ug.POST("profile", u.Profile)
	//server.POST("/user/signup", u.SingUp)
	//server.POST("/user/login", u.Login)
	//server.POST("/user/edit", u.Edit)
	//server.GET("/user/profile", u.Profile)
}

func (u *UserHandler) SignUp(c *gin.Context) {
	//内部结构体，不想被别的方法使用
	//定义结构体来接受前端传过来的数据
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
		//Context         string `json:"context"`
	}
	var req SignUpReq
	//bind方法会根据content-type来解析你的数据到req里面
	//解析错误，就会协会一个400错误
	//使用bind方法接受请求
	if err := c.Bind(&req); err != nil {
		return
	}
	//预编译
	//emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		//记录log
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "邮箱格式错误")
		return
	}
	if req.Password != req.ConfirmPassword {
		c.String(http.StatusOK, "两次输入密码不一致")
		return
	}
	//passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "密码必须大于8位数")
		return
	}
	err = u.svc.SignUp(c, domain.User{
		Email:    req.Email,
		PassWord: req.Password,
	})
	//最佳实践
	//errors.Is(err,service.ErrUserDuplicateEmail)
	if err == repository.ErrUserDuplicateEmail {
		c.String(http.StatusOK, "邮箱/密码冲突")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统异常")
		return
	}
	c.String(http.StatusOK, "注册成功")
	fmt.Printf("%v", req)
	//数据库操作

}

//使用jwt实现login方法
func (u *UserHandler) LoginJWT(c *gin.Context) {
	//定义结构体接受前端参数
	type LoginReq struct {
		Email    string `json:"email"`
		PassWord string `json:"passWord"`
	}
	var req LoginReq
	err := c.Bind(&req)
	if err != nil {
		return
	}
	//获取完前端信息，之后应该调用service中的业务逻辑代码
	user, err := u.svc.Login(c, req.Email, req.PassWord)
	if err == service.ErrInvalidUserorPassWord {
		c.String(http.StatusOK, "用户名或密码错误")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	//使用jwt设置登录态
	//生成一个jwt token
	token := jwt.New(jwt.SigningMethodHS512)
	tokenStr, err := token.SignedString([]byte("MM653MID5HDZ3LLAG57SB294YBHS76UU"))
	if err != nil {
		c.String(http.StatusInternalServerError, "系统错误")
		return
	}
	c.Header("x-jwt-token", tokenStr)
	fmt.Println(tokenStr)
	fmt.Println(user)

	c.String(http.StatusOK, "login method 登录成功")
	return
}

func (u *UserHandler) Login(c *gin.Context) {
	//定义结构体接受前端参数
	type LoginReq struct {
		Email    string `json:"email"`
		PassWord string `json:"passWord"`
	}
	var req LoginReq
	err := c.Bind(&req)
	if err != nil {
		return
	}
	//获取完前端信息，之后应该调用service中的业务逻辑代码
	user, err := u.svc.Login(c, req.Email, req.PassWord)
	if err == service.ErrInvalidUserorPassWord {
		c.String(http.StatusOK, "用户名或密码错误")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	//登录成功后设置session
	//初始化,创建session
	sess := sessions.Default(c)
	//设置session的值,刚在session里面的值
	sess.Set("userId", user.Id)
	//可以设置一些options，options控制的事cookie的内容
	sess.Options(sessions.Options{
		//生产环境设置secure和httponly就可以
		//Secure: true,//需要使用hhtps协议
		HttpOnly: true,
		MaxAge:   30, //控制cookie的过期时间
	})
	//保存session
	sess.Save()
	c.String(http.StatusOK, "login method 登录成功")
	return
}
func (u *UserHandler) Edit(c *gin.Context) {
	type EditReq struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Brithday string `json:"brithday"`
		Info     string `json:"info"`
	}
	var req EditReq
	err := c.Bind(&req)
	if err != nil {
		return
	}
	if len(req.Name) > 20 {
		c.String(http.StatusOK, "昵称长度超过限制范围")
		return
	}
	err = u.svc.Edit(c, req.Email, domain.User{
		Info:     req.Info,
		Name:     req.Name,
		Brithday: req.Brithday,
	})
	if err != nil {
		c.String(http.StatusOK, "编辑信息失败")
		return
	}
	ok, err := u.dateExp.MatchString(req.Brithday)
	if err != nil {
		c.String(http.StatusOK, "系统异常")
		return
	}
	if !ok {
		c.String(http.StatusOK, "生日信息格式错误")
		return
	}
	if len(req.Info) > 300 {
		c.String(http.StatusOK, "内容超出限制范围")
		return
	}
	c.String(http.StatusOK, "edit method 编辑成功")
	return
}
func (u *UserHandler) Profile(c *gin.Context) {
	type ProfileReq struct {
		Email string `json:"email"`
	}
	var req ProfileReq
	err := c.Bind(&req)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	user, err := u.svc.Profile(c, req.Email)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{"message": "查询成功", "data": user})

}

func (u *UserHandler) LogOut(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Options(sessions.Options{
		MaxAge: -1,
	})
	sess.Save()
	c.String(http.StatusOK, "退出登录")
}
