package setting

import "errors"

func setupEnv() error {
	return errors.New("環境檔無法寫入，請尋找正確的 env.go")
}
