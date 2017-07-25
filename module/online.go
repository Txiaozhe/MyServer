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
	"github.com/pkg/errors"
	"MyServer/utils"
	"sync"
)

type OnlineTable struct {
	Online         map[string]string
	mutex          sync.Mutex
}

var onlineTable = OnlineTable{
	Online:      make(map[string]string),
	mutex:       sync.Mutex{},
}

func PutUser(key, value string) {
	defer onlineTable.mutex.Unlock()

	onlineTable.mutex.Lock()
	onlineTable.Online[key] = value
}

func RemoveUser(key string) error {
	found := CheckIsOnline(key)
	if !found {
		return errors.New("The key: " + key + " is not exist.")
	}
	delete(onlineTable.Online, key)
	return nil
}

func CheckIsOnline(key string) bool {
	found := false

	if _, ok := onlineTable.Online[key]; ok {
		found = true
	}

	return found
}

func LogOnline() {
	for k, v := range onlineTable.Online {
		utils.Log("online.go", 60, k, v)
	}
}
