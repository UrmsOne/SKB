/*
@Author: urmsone urmsone@163.com
@Date: 2/25/21 4:15 PM
@Name: mystrconv.go
*/
package mystrconv

import "strconv"

func ParseUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	return uint(i), err
}
