package models

import "github.com/astaxie/beego/orm"

var Ormer orm.Ormer

const (
	USER_STATUS_ON  = 1
	USER_STATUS_OFF = 2

	TODO_STATUS_NEW = 1
	TODO_STATUS_ING = 2
	TODO_STATUS_DONE = 3
	TODO_STATUS_TIMEOUT = 4

	TODO_LEVEL_NORMAL = 1
	TODO_LEVEL_IMPORTANT = 2

	ITYPE_STATUS_ON = 1
	ITYPE_STATUS_OFF = 2

	DEFAULT_PAGE_SIZE = 10
	DEFAULT_PAGE_NUM  = 1

	DEFAULT_ID = 0
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Department), new(Position), new(Item), new(IType), new(AskList), new(TODO))
}

func NewOrm() {
	Ormer = orm.NewOrm()
	Ormer.Using("default") // 默认使用 default，你可以指定为其他数据库
}

type RespSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Desc    string `json:"desc"`
}

type Account struct {
	Id       int
	UserName string
	Password string
	Number   string
	Mail     string
}

type Department struct {
	Id         int
	ChargeId   int
	ChargeName string
	Count      int
	Desc       string
}

type Position struct {
	Id   int
	Desc string
}

type User struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"` // TODO mail,personalPhoneNum,address
	Sex          int     `json:"sex"`
	Age          int     `json:"age"`
	DepartmentId int     `json:"department_id"`
	PositionId   int     `json:"position_id"` // 添加TODO，说明多少人的职位未分配，部门未分配
	BoardDate    string  `json:"board_date"`  //入职时间
	Item         []*Item `json:"item" orm:"reverse(many)"`
	Status       int     `json:"status"` // 1:在职、2:离职
}

type UserRet struct {
	TotalNum int64  `json:"total_num"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Users    []User `json:"users"`
}

type Item struct {
	Id              int
	Name            string
	Number          string  //手机号
	Balance         float64 //话费余额
	User            *User `orm:"rel(fk)"`
	SerialCode      string
	ShoppingCode    string
	SourcePlateForm string
	ShopDate        string
	ITypeId         int    //资产类别
	Desc            string // TODO 确定orm是否支持自动添加时间
}

type IType struct {
	Id       int `json:"id"`
	ParentId int `json:"parent_id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Status int `json:"status"` //1:on、2：off
}

type ITypeRet struct {
	TotalNum int64 `json:"total_num"`
	PageNum int `json:"page_num"`
	PageSize int `json:"page_size"`
	ITypes []IType `json:"i_types"`
}

type AskList struct {
	Id      int
	AType   int // 申领，申还，申换
	WhoId   int
	Who     string
	forWhat string
	Date    string
	Status  int //TODO 申请状态
	Desc    string
}

type TODO struct {
	Id     int `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Status int `json:"status"` //1:新建、2:进行中、3:已完成、4:已超时
	Desc   string `json:"desc"`
	DeadLine string `json:"dead_line"`
	Level int `json:"level"`//1:一般、 2:重要
}

type TODORet struct {
	TotalNum int64  `json:"total_num"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	TODOs    []TODO `json:"todos"`
}
