package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in     *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{store,  bufio.NewScanner(in)}
}

func (cli *CLI) PlayPoker() {
	cli.in.Scan()
	player := extractWinner(cli.in.Text())
	cli.store.RecordWin(player)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
