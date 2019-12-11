package conf

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"time"
)

// 配置文件初始化
func init() {
	viper.SetConfigType("YAML")
	data, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		log.Fatalf("Read 'config.yaml' fail: %v\n", err)
	}
	_ = viper.ReadConfig(bytes.NewBuffer(data))
	_ = viper.UnmarshalKey("server", ServerConf)
	_ = viper.UnmarshalKey("email", EmailConf)
	_ = viper.UnmarshalKey("database", DBConf)
	_ = viper.UnmarshalKey("redis", RedisConf)
	_ = viper.UnmarshalKey("logger", LoggerConf)
	_ = viper.UnmarshalKey("cors", CORSConf)
}

//系统配置
type server struct {
	RunMode         string        `mapstructure:"runMode"`
	Port            int           `mapstructure:"port"`
	ReadTimeout     time.Duration `mapstructure:"readTimeout"`
	WriteTimeout    time.Duration `mapstructure:"writeTimeOut"`
	JWTSecret       string        `mapstructure:"jwtSecret"`
	JWTExpire       int           `mapstructure:"jwtExpire"`
	PrefixURL       string        `mapstructure:"PrefixUrl"`
	StaticRootPath  string        `mapstructure:"staticRootPath"`
	UploadImagePath string        `mapstructure:"uploadImagePath"`
	ImageFormats    []string      `mapstructure:"imageFormats"`
	UploadLimit     float64       `mapstructure:"uploadLimit"`
}

var ServerConf = &server{}

//email
type email struct {
	ServName         string `mapstructure:"servName"`
	UserName         string `mapstructure:"userName"`
	Password         string `mapstructure:"password"`
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	ContentTypeHTML  string `mapsttucture:"contentTypeHTML"`
	ContentTypePlain string `mapsttucture:"contentTypePlain"`
}

var EmailConf = &email{}

//db
type database struct {
	DBType      string `mapstructure:"dbType"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	DBName      string `mapstructure:"dbName"`
	TablePrefix string `mapstructure:"tablePrefix"`
	Debug       bool   `mapstructure:"debug"`
	MaxIdle     int    `mapstructure:"maxIdle"`
	MaxOpen     int    `mapstructure:"maxOpen"`
}

var DBConf = &database{}

// redis
type redis struct {
	Host        string        `mapstructure:"host"`
	Port        int           `mapstructure:"port"`
	Password    string        `mapstructure:"password"`
	DBNum       int           `mapstructure:"db"`
	MaxIdle     int           `mapstructure:"maxIdle"`
	MaxActive   int           `mapstructure:"maxActive"`
	IdleTimeout time.Duration `mapstructure:"idleTimeout"`
}

var RedisConf = &redis{}

// logger 日志
type logger struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

var LoggerConf = &logger{}

// cors 跨域资源共享配置
type cors struct {
	AllowAllOrigins  bool          `mapstructure:"allowAllOrigins"`
	AllowMethods     []string      `mapstructure:"allowMethods"`
	AllowHeaders     []string      `mapstructure:"allowHeaders"`
	ExposeHeaders    []string      `mapstructure:"exposeHeaders"`
	AllowCredentials bool          `mapstructure:"allowCredentials"`
	MaxAge           time.Duration `mapstructure:"maxAge"`
}

var CORSConf = &cors{}
