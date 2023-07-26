package main

import (
	"errors"
	"fmt"
)

func main() {
	list := []int{1, 2, 3, 4}
	list2 := []string{"1", "2", "3", "4"}
	//list, _ = Remove2(list, 0)
	var err error
	list, err = Add2(list, 0, 100)
	list2, err = Add2(list2, 0, "100")

	list, err = Remove2(list, 0)
	list2, err = Remove2(list2, 0)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
	fmt.Println(list2)
}

// Remove 删除某个下标下的元素
func Remove(list []int, index int) ([]int, error) {
	if index < 0 || index > len(list)-1 {
		return nil, errors.New("下标错误")
	}
	list = append(list[:index], list[index+1:]...)
	return list, nil
}

// Remove2 改造为泛型方法
func Remove2[T any](list []T, index int) ([]T, error) {
	if index < 0 || index > len(list)-1 {
		return nil, errors.New("下标错误")
	}
	list = append(list[:index], list[index+1:]...)
	return list, nil
}

// Add 在切片的指定位置插入新的元素
func Add(slice []int, index int, val int) ([]int, error) {
	if index < 0 || index > len(slice)-1 {
		return nil, errors.New("下标错误")
	}
	buffer := make([]int, len(slice)+1) //由于是向原来的切片中新增一个元素, 所以长度+1, 初始化后所有位置上都是0

	for i := 0; i < index; i++ {
		buffer[i] = slice[i]
	}
	buffer[index] = val

	for i := index + 1; i < len(buffer); i++ {
		buffer[i] = slice[i-1]
	}
	return buffer, nil
}

// Add2 Add 在切片的指定位置插入新的元素
func Add2[T any](slice []T, index int, val T) ([]T, error) {
	if index < 0 || index > len(slice)-1 {
		return nil, errors.New("下标错误")
	}
	buffer := make([]T, len(slice)+1) //由于是向原来的切片中新增一个元素, 所以长度+1, 初始化后所有位置上都是0

	for i := 0; i < index; i++ {
		buffer[i] = slice[i]
	}
	buffer[index] = val

	for i := index + 1; i < len(buffer); i++ {
		buffer[i] = slice[i-1]
	}
	return buffer, nil
}
