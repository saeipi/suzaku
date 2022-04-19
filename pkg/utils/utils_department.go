package utils

type DepartmentInfo struct {
	Department   string `json:"department"`    // 部门
	DepartmentId int    `json:"department_id"` // 部门id
	ParentId     int    `json:"parent_id"`     // 父部门id
	Deleted      int    `json:"deleted"`       // 1:删除
}

type Department struct {
	Maps map[int]string
}

func (e Department) DepartmentPath(departmentId int, maps map[int]*DepartmentInfo) (path string) {
	if e.Maps == nil {
		return
	}
	if departmentId == 0 {
		return
	}
	if _, ok := e.Maps[departmentId]; ok {
		path = e.Maps[departmentId]
		return
	}
	if _, ok := maps[departmentId]; ok {
		path = path + maps[departmentId].Department
		parentPath := e.DepartmentPath(maps[departmentId].ParentId, maps)
		if parentPath != "" {
			path = parentPath + "/" + path
		}
	}
	e.Maps[departmentId] = path
	return
}
