package oop

import (
	"fmt"
	"time"
)

// BaseModel 结构体定义
type BaseModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Address 结构体定义
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// User 结构体定义
type User struct {
	BaseModel // 匿名字段
	Name      string
	Email     string
	Age       int
	Address   Address
	IsActive  bool
}

// Account 结构体定义
type Account struct {
	accountNumber string
	balance       float64
	owner         string
}

// Document 结构体定义
type Document struct {
	title   string
	content string
	size    int
}

// Course 结构体定义
type Course struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Price       float64  `json:"price"`
	IsPublished bool     `json:"is_published"`
	Tags        []string `json:"tags"`
	Author      string   `json:"-"` // 序列化时忽略此字段
}

// NewUser 构造函数
func NewUser(username, email string) *User {
	return &User{
		BaseModel: BaseModel{
			ID:        generateUserID(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
		Name:     username,
		Email:    email,
		IsActive: true,
	}
}

// NewAccount 构造函数
func NewAccount(accountNumber string, owner string) *Account {
	return &Account{
		accountNumber: accountNumber,
		balance:       0.0,
		owner:         owner,
	}
}

// 模拟生成用户ID
func generateUserID() int {
	// 模拟生成用户ID的逻辑
	return int(time.Now().Unix())
}

// GetBalance 查询账户余额
func (a Account) GetBalance() float64 {
	return a.balance
}

// GetInfo 查询账户信息
func (a Account) GetInfo() string {
	return fmt.Sprintf("账户：%s, 户主：%s，余额：%.2f", a.accountNumber, a.owner, a.balance)
}

// Deposit 存款（值接收者，不会修改原始数据
func (a Account) Deposit(amount float64) float64 {
	a.balance += amount
	return a.balance
}

// CreateDocument 创建文档（构造函数）
func CreateDocument(title string) *Document {
	return &Document{
		title:   title,
		content: "",
		size:    0,
	}
}

// 值接收者 vs 指针接收者

// GetTitle 获取文档标题（值接收者）
func (d Document) GetTitle() string {
	return d.title
}

// GetSize 获取文档大小（值接收者）
func (d Document) GetSize() int {
	return d.size
}

// 指针接收者：适用于需要修改接收者的方法

// SetTitle 设置文档标题（指针接收者）
func (d Document) SetTitle(title string) {
	d.title = title
}

// AppendContent 追加文档内容（指针接收者）
func (d *Document) AppendContent(content string) {
	d.content += content
	d.size += len(content)
}

// Clear 清空文档内容（指针接收者）
func (d *Document) Clear() {
	d.content = ""
	d.size = 0
}
