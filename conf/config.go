package conf

import (
	"time"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var ArchiveConf *ArchiveConfig

func LoadConfig(f string) error {
	content, e := ioutil.ReadFile(f)
	if e != nil {
		return e
	}

	if _, err := toml.Decode(string(content), &ArchiveConf); err != nil {
		return err
	}

	return nil
}

type ArchiveConfig struct {
	ArchiveRPC *ArchiveRPConfig `toml:"archive_rpc_conf"`
}

type ArchiveRPConfig struct {
	Addrs   string   `toml:"addrs"`
	Timeout Duration `toml:"time_out"`
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
