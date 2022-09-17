package poker_test

import (
	"bytes"
	"fmt"
	poker "github.com/LightBulbfromSpace/build_an_application"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
	"time"
)

type GameSpy struct {
	StartCalled bool
	StartedWith int
	BlindAlert  []byte

	FinishCalled bool
	FinishedWith string
}

func (g *GameSpy) Start(playersNum int, to io.Writer) {
	g.StartCalled = true
	g.StartedWith = playersNum
	to.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalled = true
	g.FinishedWith = winner
}

type SpyBlindAlerter struct {
	alerts []sheduledAlert
}

func (a *SpyBlindAlerter) SheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	a.alerts = append(a.alerts, sheduledAlert{sheduledAt: at, amount: amount})
}

type sheduledAlert struct {
	sheduledAt time.Duration
	amount     int
}

func (a sheduledAlert) String() string {
	return fmt.Sprintf("%d crips at %v", a.amount, a.sheduledAt)
}

var (
	dummyBlindAlerter = &SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	//dummyStdIn        = &bytes.Buffer{}
	dummyStdOut = &bytes.Buffer{}
	dummyGame   = &GameSpy{}
)

func TestCLI(t *testing.T) {
	t.Run("start game for 6 and record Alex's name as winner", func(t *testing.T) {
		in := userSends("6\nAlex wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 6)
		assertFinishCalledWith(t, game, "Alex")
	})
	t.Run("it prompts the user to enter the number of players, start game for 9 and record Lima's name as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("9\nLima wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, game.StartedWith)
		assertFinishCalledWith(t, game, game.FinishedWith)
	})
	t.Run("it prompts the user to enter the number of players and tell about wrong input", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("Flaw\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayersNumberInput, "\n")
	})
	t.Run("bad winner input", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("5\nLloyd is a killer\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertMessagesToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInput, "\n")
	})
}

func userSends(input string) io.Reader {
	return strings.NewReader(input)
}

func assertAlert(t *testing.T, actualAlert, expectedAlert sheduledAlert) {
	t.Helper()
	amountGot := actualAlert.amount
	gotScheduledTime := actualAlert.sheduledAt
	if amountGot != expectedAlert.amount {
		t.Errorf("got amount %d, want %s", amountGot, expectedAlert.String())
	}
	if gotScheduledTime != expectedAlert.sheduledAt {
		t.Errorf("got scheduled time of %v, want %s", gotScheduledTime, expectedAlert.String())
	}
}

func assertMessagesToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	assert.Equal(t, want, got)
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	if game.StartCalled {
		t.Error("game should not have started")
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, numberOfplayersWanted int) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == numberOfplayersWanted
	})

	if !passed {
		t.Errorf("expected game to be sheduled for %d players, got %d players", numberOfplayersWanted, game.StartedWith)
	}
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishedWith == winner
	})

	if !passed {
		t.Errorf("expected to get winner %s, got %s", winner, game.FinishedWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}
