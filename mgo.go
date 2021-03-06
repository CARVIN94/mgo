package mgo

import (
	"time"

	"github.com/CARVIN94/go-util/log"
	"github.com/globalsign/mgo"
)

var (
	session  *mgo.Session
	database string
)

// Config 数据库基础配置
type Config struct {
	Hosts    string
	Database string
	UserName string
	Password string
	Timeout  time.Duration
}

// Task 查询任务模型
type Task struct {
	*mgo.Collection
}

// Connect 初始化并连接数据库
func Connect(config *Config) {
	defer log.Success("数据库连接成功" + " " + config.Hosts)
	var err error

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{config.Hosts},
		Direct:    false,
		PoolLimit: 4096,
		Database:  config.Database,
		Username:  config.UserName,
		Password:  config.Password,
		Timeout:   config.Timeout,
	}

	session, err = mgo.DialWithInfo(dialInfo)
	log.FailOnError(err, "数据库连接失败")

	database = config.Database

	session.SetMode(mgo.Monotonic, true)
}

//GetDatabase 获取数据库名称
func GetDatabase() string {
	return database
}

// Collection 连接必用前置方法
func Collection(database string, collection string) Task {
	s := session.Copy()
	c := s.DB(database).C(collection)
	return Task{c}
}

// End 连接必用后置方法
func (task *Task) End() {
	task.Database.Session.Close()
}

// EnableNotFound 处理没有搜索到数据的情况
func (task *Task) EnableNotFound() {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			if err.Error() != "not found" {
				panic(err)
			}
		}
	}
}
