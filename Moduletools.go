package main

import (
	"math"
)

//1.分页模块 调用函数PagingMethod
//参数1 pageindex 传过来的要到的页面 参数2 PageSize 定义每个页的数据数 参数3 num 所有页数据数量
//返回值1 (pagehead 上一页) 返回值2 (pagenext 下一页) 返回值3 (pageCount 全部页码)
//[]Article数据自己传
//获取到对应数据
//func ShowIndex(PageIndex, PageSize int) ([]Article, error) {
//article, err := models.ShowIndex(page, PageSize)

//测试
//func main() {
//	pagehead, pagenext,pageCount,pageqie:=PagingMethod(1000004,10,1000023)
//	fmt.Println("上一页, ",pagehead)
//	fmt.Println("下一页, ",pagenext)
//	fmt.Println("总页数, ",pageCount)
//	fmt.Println("中间页数",pageqie)
//}
func PagingMethod(pageIndex, PageSize, num int) (pagehead, pagenext, pageCount int, pageqie []int) {
	if pageIndex<=0 {
		pageIndex=1
	}
	pageCount = int(math.Ceil(float64(num) / float64(PageSize)))
	//上一页和下一页实现
	if pageIndex == 1 {
		pagehead = pageIndex
		pagenext = pageIndex + 1
	} else if pageIndex == pageCount {
		pagehead = pageIndex - 1
		pagenext = pageIndex
	} else {
		pagehead = pageIndex - 1
		pagenext = pageIndex + 1
	}
	pageqie = ShowPageIndex(pageCount, pageIndex)
	return
}
func ShowPageIndex(pageCount int, pageindex int) (pageqie []int) {
	//页码数据＜5
	if pageCount < 5 {
		for i := 1; i <= pageCount; i++ {
			pageqie = append(pageqie, i)
		}
		return pageqie
	}
	if pageindex <= 3 {
		for i := 1; i <= 5; i++ {
			pageqie = append(pageqie, i)
		}
		return pageqie
	}
	if pageindex < pageCount-2 {
		for i := pageindex - 2; i <= pageindex+2; i++ {
			pageqie = append(pageqie, i)
		}
		return pageqie
	}
	for i := pageCount - 4; i <= pageCount; i++ {
		pageqie = append(pageqie, i)
	}
	return pageqie
}

