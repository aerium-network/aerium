package zmq

import (
	"bytes"

	"github.com/aerium-network/aerium/types/block"
	"github.com/aerium-network/aerium/util/logger"
	"github.com/go-zeromq/zmq4"
)

type rawBlockPub struct {
	basePub
}

func newRawBlockPub(socket zmq4.Socket, logger *logger.SubLogger) Publisher {
	return &rawBlockPub{
		basePub: basePub{
			topic:     TopicRawBlock,
			zmqSocket: socket,
			logger:    logger,
		},
	}
}

func (r *rawBlockPub) onNewBlock(blk *block.Block) {
	rawHeader := make([]byte, 0)
	buf := bytes.NewBuffer(rawHeader)

	if err := blk.Header().Encode(buf); err != nil {
		r.logger.Error("failed to encode block header", "err", err, "publisher", r.TopicName())

		return
	}

	rawMsg := r.makeTopicMsg(buf.Bytes(), blk.Height())
	message := zmq4.NewMsg(rawMsg)

	if err := r.zmqSocket.Send(message); err != nil {
		r.logger.Error("zmq publish message error", "err", err, "publisher", r.TopicName())

		return
	}

	r.logger.Debug("ZMQ published the message successfully",
		"publisher", r.TopicName(),
		"block_height", blk.Height())

	r.seqNo++
}
