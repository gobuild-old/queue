package queue_test

import (
	"testing"

	"github.com/gobuild/queue"
)

func TestHello(t *testing.T) {
	q := &queue.Thunder{}
	q.Fetch()
}
