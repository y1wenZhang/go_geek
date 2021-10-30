package main

import (
	"database/sql"
	"fmt"
	"pkg/errors"
)

// Q: dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// A: 在自己的应用程序中调用了其他包的函数，应该将这个error向上抛，交给上层处理。
// 处理error的两种方法，1. 在当前步骤，自己降级处理；2.将错误向上抛，交给上层处理。
// 保持只处理错误一次

func Dao(query string) error {
	err := mockError()

	if err == sql.ErrNoRows {
		return errors.Wrapf(err, fmt.Sprintf("data not found: %s", query))
	}
}
