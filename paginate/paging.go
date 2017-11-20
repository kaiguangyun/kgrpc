package paginate

import (
	"reflect"
	"github.com/kaiguangyun/kgrpc/helper"
	"strings"
)

const (
	PagingSize      = 10   // 默认页码
	PagingByPrimary = iota // 主键分页
	PagingByNumber         // 页码分页
)

const (
	PagingDefaultSortField     = "id"    // 默认排序字段
	PagingDefaultSortOrderAsc  = "asc"   // 升序排序
	PagingDefaultSortOrderDesc = "desc"  // 倒序排序
	PagingSortFieldSuffix      = "_sort" // 排序后缀
)

// 获取分页选项
func GetPagingOptions(in *PageOptions, PagingMode int) (offset, limit int32) {
	// 设置分页默认值
	SetPagingDefaultOptions(in)
	// 获取分页偏移
	switch PagingMode {
	case PagingByPrimary:
		offset, limit = GetPagingModeByPrimaryOptions(in)
	case PagingByNumber:
		offset, limit = GetPagingModeByNumberOptions(in)
	}
	return offset, limit
}

// 设置 : 默认每页 10 条，页码 第 1 页
func SetPagingDefaultOptions(in *PageOptions) {
	// set default pageSize ：
	if in.PageSize < 1 {
		in.PageSize = int32(PagingSize)
	}
	// set default first page : 1
	if in.PageNumber < 1 {
		in.PageNumber = 1
	}
	// 设置默认查询字段、排序
	in.SortField, in.SortFieldTo = SetPagingModeByPrimarySelectFieldAndSort(in.SortField, in.SortFieldTo)
}

// 默认排序
func SetPagingModeByPrimarySelectFieldAndSort(sortField, sortFieldTo string) (columnName, sortOrder string) {
	// 排序字段
	columnName = helper.ToSnakeString(helper.TrimSpace(sortField))
	// 默认排序字段
	if columnName == "" {
		columnName = PagingDefaultSortField
	}
	// 添加排序后缀
	if columnName != PagingDefaultSortField && ! strings.HasSuffix(columnName, PagingSortFieldSuffix) {
		columnName += PagingSortFieldSuffix
	}
	// 排序顺序
	sortOrder = PagingDefaultSortOrderDesc
	sortFieldTo = strings.ToLower(helper.TrimSpace(sortFieldTo))
	if sortFieldTo == PagingDefaultSortOrderAsc {
		sortOrder = PagingDefaultSortOrderAsc
	}
	return columnName, sortOrder
}

// structPointer 必须是 struct 的 指针
func PagingOptionsFieldNameIsValid(structPointer interface{}, fieldName string) bool {
	sElem := reflect.ValueOf(structPointer).Elem()
	return sElem.FieldByName(helper.ToCamelString(fieldName)).IsValid()
}

// 页码分页模式选项
func GetPagingModeByNumberOptions(in *PageOptions) (offset, limit int32) {

	offset = in.PageSize * (in.PageNumber - 1)
	limit = in.PageSize

	return offset, limit
}

// 主键分页模式选项
func GetPagingModeByPrimaryOptions(in *PageOptions) (offset, limit int32) {

	offset = 0
	limit = in.PageSize

	return offset, limit
}

//panic if s is not a struct pointer
func GetSortValue(s interface{}, sortField string) int64 {
	sortField = helper.ToCamelString(sortField)
	sElem := reflect.ValueOf(s).Elem()
	// 是否存在字段
	if ! sElem.FieldByName(sortField).IsValid() {
		return 0
	}
	// 字段值
	value := sElem.FieldByName(sortField).Interface()
	// 判断字段值
	switch value.(type) {
	case int64:
		return value.(int64)
	case int32:
		return int64(value.(int32))
	case int:
		return int64(value.(int))
	default:
		return 0
	}
	return 0
}

// Set Paging Result
func SetPagingResult(in *PageOptions, TotalRecords int32, SortValue int64) (paginate PageResult) {
	paginate.TotalRecords = TotalRecords
	// 总页数
	if paginate.TotalRecords%in.PageSize == 0 {
		paginate.TotalPages = paginate.TotalRecords / in.PageSize
	} else {
		paginate.TotalPages = paginate.TotalRecords/in.PageSize + 1
	}
	paginate.PageSize = in.PageSize     // 显示条数
	paginate.PageNumber = in.PageNumber // 当前页码
	paginate.SortValue = SortValue      // 排序字段

	return
}
