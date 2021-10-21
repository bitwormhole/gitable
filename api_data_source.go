package ptable

// DataSource 代表指向DB的数据源
type DataSource interface {
	GetSession() Session
	GetDatabase() Database
}

// Repository 代表某个表的存储库
type Repository interface {
	Init(ds DataSource) error
}
