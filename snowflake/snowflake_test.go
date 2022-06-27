package snowflake

import (
	"testing"
)

func TestSnowflake(t *testing.T) {
	snowflake := new(Snowflake)
	snowflake.timestamp = 165226305000
	snowflake.workerid = 1
	a := snowflake.NextVal()
	snowflake.workerid = 1
	b := snowflake.NextVal()
	//time.Sleep(time.Second * 3)
	snowflake.workerid = 4095
	c := snowflake.NextVal()
	t.Log(a, b, c)
}
