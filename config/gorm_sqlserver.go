package config

type SqlServer struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *SqlServer) Dsn() string {
	return "sqlserver://" + s.Username + ":" + s.Password + "@" + s.Path + ":" + s.Port + "?database=" + s.Dbname + "&" + s.Config
}

func (s *SqlServer) GetLogMode() string {
	return s.LogMode
}
