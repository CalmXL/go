package config

// 常量配置法
// 没有命名空间
// const (
// 	DBHost     = "127.0.0.1"
// 	DBPort     = "3306"
// 	DBUser     = "root"
// 	DBPassword = "xL1210..."
// 	DBName     = "study"
// )
//
// const (
// 	GinHost = "127.0.0.1"
// 	GinPort = "8080"
// )

// 映射表配置法

// var DBConfig = map[string]string{
// 	"HOST":     "127.0.0.1",
// 	"PORT":     "3306",
// 	"USER":     "root",
// 	"PASSWORD": "xL1210...",
// 	"NAME":     "study",
// }
//
// var GINConfig = map[string]string{
// 	"HOST": "127.0.0.1",
// 	"PORT": "8080",
// }

// flag 配置法

// var DBHost = flag.String("DB_HOST", "127.0.0.1", "Database host")
// var DBPort = flag.String("DB_PORT", "3306", "Database port")
// var DBUser = flag.String("DB_USER", "root", "Database username")
// var DBPassword = flag.String("DB_PASSWORD", "xL1210...", "Database password")
// var DBName = flag.String("DB_NAME", "study", "Database name")
//
// var GinHost = flag.String("GIN_HOST", "127.0.0.1", "GIN host")
// var GinPort = flag.String("GIN_PORT", "8080", "GIN port")

// 结构体配置法

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type GinConfig struct {
	Host string
	Port string
}

var DB = DBConfig{
	Host:     "127.0.0.1",
	Port:     "3306",
	User:     "root",
	Password: "xL1210...",
	Name:     "study",
}

var Gin = GinConfig{
	Host: "127.0.0.1",
	Port: "8080",
}
