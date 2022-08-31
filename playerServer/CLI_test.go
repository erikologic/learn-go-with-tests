package poker_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	poker "github.com/mrenrich84/learn-go-with-tests/playerServer"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}
		stdout := &bytes.Buffer{}

		in := userSends("3", "Chris")
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		game := &poker.GameSpy{}

		in := userSends("8", "Cleo")
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		game := &poker.GameSpy{}

		stdout := &bytes.Buffer{}
		in := userSends("pies", "")

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func assertFinishCalledWith(t *testing.T, game *poker.GameSpy, winnerName string) {
	t.Helper()
	if game.FinishedWith != winnerName {
		t.Errorf("wanted the game to finish with %s, got %s", winnerName, game.FinishedWith)
	}
}

func assertGameStartedWith(t *testing.T, game *poker.GameSpy, numberOfPlayers int) {
	t.Helper()
	if game.StartedWith != numberOfPlayers {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayers, game.StartedWith)
	}
}

func assertGameNotStarted(t *testing.T, game *poker.GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func userSends(numberOfPlayers, winnerName string) io.Reader {
	return strings.NewReader(fmt.Sprintf("%s\n%s wins\n", numberOfPlayers, winnerName))
}



func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
