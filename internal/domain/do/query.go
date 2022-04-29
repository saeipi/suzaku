package do

type MysqlQuery struct {
	Condition string
	Params    []interface{}
	Page      int
	PageSize  int
}

func NewMysqlQuery() *MysqlQuery {
	return &MysqlQuery{
		Condition: "1=1",
		Params:    make([]interface{}, 0),
	}
}
