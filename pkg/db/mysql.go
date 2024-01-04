package db

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

//=====
// Mysql配置对象
//=====

// MySQL mysql配置
type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`

	Dsn string
	// 因为使用的是 Mysql的连接池，需要对连接池做一些规划配置
	// 控制当前程序的 Mysql打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 控制 Mysql复用，比如 5， 最多运行5个复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 一个连接的生命周期，这个和 Mysql Server配置有关系，必须小于 Server 配置
	// 比如一个链接用 12 h 换一个 conn，保证一点的可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// Idle 连接 最多允许存货多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
	// 作为私有变量，用于控制DetDB
	lock sync.Mutex
}

type Option func(*MySQL)

func WithMaxOpenConn(maxOpenConn int) Option {
	return func(m *MySQL) {
		m.MaxOpenConn = maxOpenConn
	}
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return func(m *MySQL) {
		m.MaxIdleConn = maxIdleConn
	}
}

func WithMaxLifeTime(maxLifeTime int) Option {
	return func(m *MySQL) {
		m.MaxLifeTime = maxLifeTime
	}
}

func WithMaxIdleTime(maxIdleTime int) Option {
	return func(m *MySQL) {
		m.MaxIdleTime = maxIdleTime
	}
}

const (
	DefaultMaxOpenConn = 10
	DefaultMaxIdleConn = 5
	DefaultMaxLifeTime = 7200
	DefaultMaxIdleTime = 14400
)

func NewMySQL(host, port, username, password, database string, opts ...Option) (*MySQL, error) {
	if host == "" || port == "" || username == "" || password == "" || database == "" {
		return nil, fmt.Errorf("required fields are missing")
	}

	c := &MySQL{
		Host:     host,
		Port:     port,
		UserName: username,
		Password: password,
		Database: database,
	}

	for _, opt := range opts {
		opt(c)
	}

	// Set default values if not provided in options
	if c.MaxOpenConn == 0 {
		c.MaxOpenConn = DefaultMaxOpenConn
	}
	if c.MaxIdleConn == 0 {
		c.MaxIdleConn = DefaultMaxIdleConn
	}
	if c.MaxLifeTime == 0 {
		c.MaxLifeTime = DefaultMaxLifeTime
	}
	if c.MaxIdleTime == 0 {
		c.MaxIdleTime = DefaultMaxIdleTime
	}

	return c, nil
}

func NewMySQLFromDSN(dsn string, opts ...Option) *MySQL {
	if dsn == "" {
		return nil
	}

	c := &MySQL{
		Dsn: dsn,
	}
	for _, opt := range opts {
		opt(c)
	}

	// Set default values if not provided in options
	if c.MaxOpenConn == 0 {
		c.MaxOpenConn = DefaultMaxOpenConn
	}
	if c.MaxIdleConn == 0 {
		c.MaxIdleConn = DefaultMaxIdleConn
	}
	if c.MaxLifeTime == 0 {
		c.MaxLifeTime = DefaultMaxLifeTime
	}
	if c.MaxIdleTime == 0 {
		c.MaxIdleTime = DefaultMaxIdleTime
	}

	return c
}

var mysqlDb *gorm.DB

func (m *MySQL) GetConnection() (*gorm.DB, error) {
	m.lock.Lock() // 锁住临界区，保证线程安全
	defer m.lock.Unlock()

	if mysqlDb == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		mysqlDb = conn
	}
	return mysqlDb, nil
}

// gorm获取数据库连接
func (m *MySQL) getDBConn() (*gorm.DB, error) {
	var dsn string
	if m.Dsn != "" {
		dsn = m.Dsn
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
			m.UserName, m.Password, m.Host, m.Port, m.Database)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("连接Mysql：%s，error：%s", dsn, err.Error())
	}

	// 维护连接池
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(m.MaxOpenConn)
	sqlDB.SetMaxIdleConns(m.MaxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))

	// 用于测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql %s，error：%s", dsn, err.Error())
	}
	return db, nil
}