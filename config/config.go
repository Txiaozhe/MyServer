/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package config
//todo:配置端口
const (
	UnKnown          = -1 // 未知错误
	ErrSucceed       = 0  // 成功
	ErrInvalidParam  = 1  // 参数错误
	ErrInvalidPerm   = 2  // 权限错误
	ErrInvalidUser   = 3  // 用户不存在
	ErrInactiveUser  = 4  // 用户被禁用
	ErrOperationPass = 5  // 用户名或密码错误
	ErrUserExists    = 6  // 用户已存在
	ErrInvalidCode   = 7  // 验证码错误
	ErrSendCode      = 8  // 验证码发送失败
	ErrGroupExists   = 9  // 用户色团已存在
	ErrIncorrectPass = 10 // 密码错误
	ErrCache         = 11 // 存入缓存出错
	ErrConv          = 12 // 类型转换出错
	ErrNoChange      = 13 // 未改变
	ErrIsInGroup     = 14 // 已经存在于该色团
	ErrMysqlNotFound = 15 // 数据库没找到
	ErrLogined       = 16 // 用户已登录


	ErrSession       = 400 // Session 错误
	ErrLoginRequired = 401 // 未登录
	ErrActive        = 402 // 用户未激活
	ErrUpgrade       = 403 // ws 连接错误

	ErrMysql         = 500 // MySQL 错误

	ErrMongoDB       = 600 // MongoDB 错误
)
