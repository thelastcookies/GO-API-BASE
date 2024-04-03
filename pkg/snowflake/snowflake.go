package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"time"
)

var IDGen *Snow

type Snow struct {
	node *snowflake.Node
}

func New(startTime string, machineID int64) (*Snow, error) {
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return nil, errors.Wrap(err, "[snowflake.New] invalid param 'startTime'")
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err := snowflake.NewNode(machineID)
	if err != nil {
		return nil, errors.Wrap(err, "[snowflake.New] failed to generate snowflake")
	}

	return &Snow{
		node: node,
	}, nil
}

// Snow 生成 64 位雪花 ID
func (s *Snow) Snow() int64 {
	return s.node.Generate().Int64()
}

func init() {
	IDGen, _ = New("2023-01-01", 1)
}
