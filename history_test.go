package zabbix_test

import "testing"

func TestHistories(t *testing.T) {
	res, err := getAPI(t).HistoriesGet(map[string]interface{}{"limit": 1})
	if err != nil {
		t.Fatal(err)
	}
	if len(res) == 0 {
		t.Fail()
	}
}
