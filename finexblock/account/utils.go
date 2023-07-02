package account

import "fmt"

func getKey(uuid, currency string) string {
	return fmt.Sprintf("%v:%v", uuid, currency)
}

func getAccountLockKey(uuid, currency string) string {
	return fmt.Sprintf("%v:%v:%v", accountLockPrefix, uuid, currency)
}

func wrapErr(wrapper, err error) error {
	return fmt.Errorf("%v : %v", wrapper, err)
}
