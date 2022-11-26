package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`

	// gorm
	Mysql     Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	SqlServer SqlServer       `mapstructure:"sqlserver" json:"sqlserver" yaml:"sqlserver"`
	DBList    []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
}
