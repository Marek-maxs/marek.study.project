package amqp

import "strings"

func checkTopic(pub, sub string) bool {
	if sub == "" {
		return true
	}

	p := strings.Split(pub, ".")

	for i, seg := range strings.Split(sub, ".") {
		if seg == "#" {
			return true
		}

		if seg != "*" && seg != p[i] {
			return false
		}
	}

	return true
}

func explanLifeCycle(queue Queue) (durable, autoDelete, exclusive bool) {
	switch queue.Lifecycle {
	case Temporary:
		autoDelete = true
	case Exclusive:
		autoDelete = true
		exclusive = true
	case Durable:
		durable = true
	}

	return
}