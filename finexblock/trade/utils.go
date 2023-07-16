package trade

import "fmt"

func getBalanceKey(uuid, currency string) string {
	return fmt.Sprintf("%v:%v:balance", uuid, currency)
}

func getOrderKey(uuid string) string {
	return fmt.Sprintf("order:%v", uuid)
}

func getAccountLockKey(uuid, currency string) string {
	return fmt.Sprintf("%v:%v:%v", accountLockPrefix, uuid, currency)
}

func wrapErr(wrapper, err error) error {
	return fmt.Errorf("%v : %v", wrapper, err)
}
