/*   1. Crie um diretório para seus fontes, por exemplo goprogs.
     2. Dentro desta pasta voce ira criar seus programas em Go.
     3. Mude para o diretório goprogs (cd gopros)
     4. Entre: "go mod init goprogs".    Este comando criará um arquivo go.mod em goprogs
     5. Entre: "go get github.com/nsf/termbox-go"
     6. Agora voce pode usar este pacote.  veja exemplo abaixo.
	 Instrucoes:
	 Use as setas para andar sobre o tabuleiro.
	 Digite uma letra para gravar a mesma por onde passa.
	 Digite ESC para sair.
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
    "io"
	"io/ioutil"
    "log"

	term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync() // cosmestic purpose
}

type GameState struct {
	xSize int
	ySize int
	mesa  [][]rune
}
type Player struct {
	x  int
	y  int
	ch rune
}

func printState(g GameState) {
	reset()
	//fmt.Println(g)
	for i := 0; i < g.ySize; i++ {
		for j := 0; j < g.xSize; j++ {
			fmt.Print(string(g.mesa[i][j]))
		}
		fmt.Println("")
	}
}

func main() {

	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	fmt.Println("Enter any key to see their ASCII code or press ESC button to quit")

	const Xs = 28
	const Ys = 30

	mesa1 := make([][]rune, Ys)
	for i := 0; i < Ys; i++ {
		mesa1[i] = make([]rune, Xs)
	}

	filename := "map.txt"

	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inputdata := string(filebuffer)

	r := bufio.NewReader(strings.NewReader(inputdata))
	
	i := 0
	j := 0

    for {
        if c, sz, err := r.ReadRune(); err != nil {
            if err == io.EOF {
                break
            } else {
                log.Fatal(err)
            }
        } else {
			if c == '\n' {
					j = 0
					i++
					fmt.Println(i)
				} else {
					mesa1[i][j] = c
					j++
				}
            fmt.Printf("%q [%d] - i: %d, j: %d\n", string(c), sz, i, j)
        }
    }

	g := GameState{xSize: Xs, ySize: Ys, mesa: mesa1}
		printState(g)
	p := Player{x: 0, y: 0, ch: 'k'}
	pOld := Player{x: -1, y: -1, ch: ' '}

	printState(g)

keyPressListenerLoop:
	for {
		if !(p == pOld) {
			g.mesa[p.y][p.x] = p.ch
			printState(g)
			pOld = p
		}
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {

			// -CASOS NAO USADOS AQUI ----------------------------------------------
			case term.KeyF1:
				reset()
				fmt.Println("F1 pressed")
			case term.KeyF2:
				reset()
				fmt.Println("F2 pressed")
			case term.KeyF3:
				reset()
				fmt.Println("F3 pressed")
			case term.KeyF4:
				reset()
				fmt.Println("F4 pressed")
			case term.KeyF5:
				reset()
				fmt.Println("F5 pressed")
			case term.KeyF6:
				reset()
				fmt.Println("F6 pressed")
			case term.KeyF7:
				reset()
				fmt.Println("F7 pressed")
			case term.KeyF8:
				reset()
				fmt.Println("F8 pressed")
			case term.KeyF9:
				reset()
				fmt.Println("F9 pressed")
			case term.KeyF10:
				reset()
				fmt.Println("F10 pressed")
			case term.KeyF11:
				reset()
				fmt.Println("F11 pressed")
			case term.KeyF12:
				reset()
				fmt.Println("F12 pressed")
			case term.KeyInsert:
				reset()
				fmt.Println("Insert pressed")
			case term.KeyDelete:
				reset()
				fmt.Println("Delete pressed")
			case term.KeyHome:
				reset()
				fmt.Println("Home pressed")
			case term.KeyEnd:
				reset()
				fmt.Println("End pressed")
			case term.KeyPgup:
				reset()
				fmt.Println("Page Up pressed")
			case term.KeyPgdn:
				reset()
				fmt.Println("Page Down pressed")
			case term.KeySpace:
				reset()
				fmt.Println("Space pressed")
			case term.KeyBackspace:
				reset()
				fmt.Println("Backspace pressed")
			case term.KeyEnter:
				reset()
				fmt.Println("Enter pressed")
			case term.KeyTab:
				reset()
				fmt.Println("Tab pressed")
				// -ATE AQUI NAO USA AINDA ----------------------------------------------

				// -USAMOS DAQUI PARA BAIXO----------------------------------------------
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyArrowUp:
				reset()
				// fmt.Println("Arrow Up pressed")
				mesa1[p.y][p.x] = '.'
				result := (p.y + Ys - 1) % Ys
				if mesa1[result][p.x] != '#' {
					p.y = result;
				}
			case term.KeyArrowDown:
				// reset()
				fmt.Println("Arrow Down pressed")
				mesa1[p.y][p.x] = '.'
				p.y = (p.y + 1) % Ys
				// mesa1[p.y][p.x] = '.'
				// result := (p.y + 1) % Ys
				// if mesa1[result][p.x] != '#' {
				// 	p.y = result
				// }
			case term.KeyArrowLeft:
				// reset()
				fmt.Println("Arrow Left pressed")
				mesa1[p.y][p.x] = '.'
				result := (p.x + Xs - 1) % Xs
				if mesa1[p.y][result] == '#' {
					p.x = result
				}
			case term.KeyArrowRight:
				// reset()
				fmt.Println("Arrow Right pressed")
				mesa1[p.y][p.x] = '.'
				p.x = (p.x + 1) % Xs
			default:
				// we only want to read a single character or one key pressed event
				// reset()
				p.ch = ev.Ch
				fmt.Println("ASCII : ", ev.Ch)
				// --------------------------------------------------------------------

			}

		case term.EventError:
			panic(ev.Err)
		}
	}
}
