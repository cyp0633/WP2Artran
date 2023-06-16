package conf

import "testing"

func TestParseConf(t *testing.T) {
	confPath := "conf.yaml"
	ParseConf(confPath)
	t.Log(Conf)
}