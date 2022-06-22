package setting

import (
	"fmt"
	"os"
)

type pair struct {
	key   string
	value string
}

func setupEnv() error {
	return checkAndSet([]pair{
		{"PING", "PONG"},
		{"PROJECT_NAME", "i18n-town"},

		{"REDIS_HOST", "icache"},
		{"REDIS_PORT", "6379"},
		{"REDIS_DEFAULT_DB", "0"},

		{"PSQL_USERNAME", "postgres"},
		{"PSQL_PASSWORD", "root"},
		{"PSQL_HOST", "idatabase"},
		{"PSQL_PORT", "5432"},

		{"IMSGQUEUE_USERNAME", "guest"},
		{"IMSGQUEUE_PASSWORD", "guest"},
		{"IMSGQUEUE_HOST", "imsg_queue"},
		{"IMSGQUEUE_PORT", "5672"},

		{"IADMIN_HOST", "iadmin"},
		{"IADMIN_PORT", "50051"},
		{"IADMIN_PSQL_USERNAME", "postgres"},
		{"IADMIN_PSQL_PASSWORD", "root"},
		{"IADMIN_PSQL_HOST", "idatabase"},
		{"IADMIN_PSQL_PORT", "5432"},
		{"IADMIN_PSQL_DB_NAME", "iadmin"},

		{"IWORD_HOST", "iword"},
		{"IWORD_PORT", "8080"},
		{"IWORD_PSQL_USERNAME", "postgres"},
		{"IWORD_PSQL_PASSWORD", "root"},
		{"IWORD_PSQL_HOST", "idatabase"},
		{"IWORD_PSQL_PORT", "5432"},
		{"IWORD_PSQL_DB_NAME", "iword"},
	})
}

func checkAndSet(pairs []pair) error {
	for _, p := range pairs {
		key, value := p.key, p.value
		val, ok := os.LookupEnv(key)

		if !ok { // 不存在
			err := os.Setenv(key, value)
			if err != nil {
				return err
			}
		} else if val != value { // 存在且值不同
			return fmt.Errorf("%s已存在值：%s, 無法寫入 %s\n", key, val, value)
		}
	}

	return nil
}
