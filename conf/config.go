package conf

import (
	"github.com/spf13/viper"
)

// ViperConfig ...
type ViperConfig struct {
	*viper.Viper
}

// StudyEventGo ...
var StudyEventGo *ViperConfig

func init() {
	StudyEventGo = readConfig(map[string]interface{}{
		"debug_route": false,
		"debug_db":    true,
		"port":        4567,
		"redis_host":  "localhost:6379",
		"broker_host": "redis://localhost:6379",
		"db":          "mysql",
		"db_host":     "localhost",
		"db_port":     3306,
		"db_name":     "lily",
		"db_user":     "root",
		"db_pass":     "",
		"worker_name": "study_event_go_worker",
	})
}

func readConfig(defaults map[string]interface{}) *ViperConfig {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	return &ViperConfig{
		Viper: v,
	}
}
