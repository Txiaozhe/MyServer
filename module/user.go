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

package module

import (
	"MyServer/sqlHelper"
	"github.com/pkg/errors"
)

type User struct {
	Name     string       `json:"name" form:"name" query:"name"`
	Password string       `json:"password" form:"password" query:"password"`
	Mobile   string       `json:"mobile" form:"mobile" query:"mobile"`
}

//添加新用户
func NewUser(name, password, mobile string) {
	sqlHelper.Insert(sqlHelper.Db, "INSERT INTO user (name, password, mobile) VALUES (?, ?, ?)", name, password, mobile)
}

//检查用户是否已存在
func IsUserExisted(name string) (User, error) {
	rowMap, err := sqlHelper.FetchRow(sqlHelper.Db, "SELECT * FROM `user` WHERE `name` = ? LIMIT 1", name)

	if err != nil {
		return User{}, err
	}

	user := User{
		Name:         (*rowMap)["name"],
		Password:     (*rowMap)["password"],
	}

	if user.Name != "" {
		return user, errors.New("User existed!")
	}

	return user, err
}
