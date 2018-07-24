package models

type Data struct {
	ID        int
	ServerIP  string
	ClientIP  string
	Metadata  map[string]string
	Msg       string
	tableName struct{} `sql:"myservice.data"`
}
