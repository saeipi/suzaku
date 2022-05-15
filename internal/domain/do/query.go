package do

type MysqlSelect struct {
	Query  string
	Args   []interface{}
	Limit  int
	Offset int
}

func NewMysqlSelect() *MysqlSelect {
	return &MysqlSelect{
		Query: "1=1",
		Args:  make([]interface{}, 0),
	}
}
