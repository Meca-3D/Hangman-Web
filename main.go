package main


const PORT = ":8080"

var run = true

func main() {
	initGame()
	for run {
		web()
	}
}
