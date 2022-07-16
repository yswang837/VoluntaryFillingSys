package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"strconv"
	"time"
)

const (
	timeLayout = "2006-01-02"
)

var Node *sf.Node

func Init(startTime string, machineId string) (err error) {
	machineIdNum, err := strconv.ParseInt(machineId, 10, 64)
	if err != nil {
		return err
	}
	var st time.Time
	st, err = time.Parse(timeLayout, startTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	Node, err = sf.NewNode(machineIdNum)
	return err
}

func GenID() string {
	return Node.Generate().String()
}
