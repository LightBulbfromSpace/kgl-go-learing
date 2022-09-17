package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

const (
	PlayerPrompt          = "Please enter the number of players: "
	BadPlayersNumberInput = "Bad players' number input, expected to get a number"
	BadWinnerInput        = "Bad winner's name input, expected to get a string '{Name} wins'"
)

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(cli.readLine())
	if err != nil {
		fmt.Fprintln(cli.out, BadPlayersNumberInput)
		return
	}

	cli.game.Start(numberOfPlayers, cli.out)
	userInput := cli.readLine()
	winner, err := extractWinner(userInput)
	if err != nil {
		fmt.Fprintln(cli.out, BadWinnerInput)
		return
	}
	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) (string, error) {
	if strings.HasSuffix(userInput, " wins") {
		return strings.Replace(userInput, " wins", "", 1), nil
	}
	return "", errors.New(BadWinnerInput)
}
