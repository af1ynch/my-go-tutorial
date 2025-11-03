package oop

import (
	"fmt"
	"time"
)

// User 结构体定义
type User struct {
	ID       int
	Name     string
	Email    string
	Age      int
	IsActive bool
}

// Account 结构体定义
type Account struct {
	accountNumber string
	balance       float64
	owner         string
}

// NewUser 构造函数
func NewUser(username, email string) *User {
	return &User{
		ID:       generateUserID(),
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
