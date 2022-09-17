package poker

import "io"

type Game interface {
	Start(playersNum int, alertsDestination io.Writer)
	Finish(winner string)
}
