package conf

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     int32
	Name     string
}

// DbMasterList 系统中所有mysql主库 root:root@tcp(127.0.0.1:3306)/lottery?charset=utf-8
var DbMasterList = []DbConfig{
	{
		User:     "root",
		Password: "Aa_5469826",
		Host:     "localhost",
		Port:     3306,
		Name:     "lottery",
	},
}

var DbMaster DbConfig = DbMasterList[0]
