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

package handler

import (
	"github.com/labstack/echo"
	"MyServer/module"
	"MyServer/utils"
	"MyServer/config"
	"net/http"
)

func ReviseInfo(c echo.Context) error {
	status := config.ErrSucceed
	data := module.UserInfo{}

	sess := utils.GlobalSessions.SessionStart(c.Response().Writer, c.Request())

	name := sess.Get("login")

	if name != nil {
		data, _ = module.GetUserInfoFromSql(name.(string))
	} else {
		status = config.ErrLoginRequired
	}

	status_json := module.Err{
		Status: status,
		Data: data,
	}

	return c.JSONPretty(http.StatusOK, status_json, " ")
}
