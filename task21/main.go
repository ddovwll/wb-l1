package main

type consolePlayer interface {
	playConsole()
}

type console struct{}

func (c *console) playConsole() {
	println("Playing PlayStation 🎮")
}

type gamer struct {
	player consolePlayer
}

func (g *gamer) playGames() {
	g.player.playConsole()
}

type pc struct{}

func (p *pc) playPc() {
	println("Playing PC 💻")
}

type playStationEmulator struct {
	*pc
}

func (p *playStationEmulator) playConsole() {
	p.playPc()
}

// Адаптер позволяет заменять реализацию ожидаемых интерфейсов, не изменяя код потребителя
// и адаптируемого типа

// Из-за злоупотребления адаптерами возрастает сложность кода
// Также есть вероятность "адаптировать неадаптируемое" и нарушить логику проргаммы
func main() {
	// Исходная реализцация без адаптера
	c := console{}
	g := gamer{player: &c}
	g.playGames()

	// Реализация с адаптером
	pc := pc{}
	playStationEmulator := playStationEmulator{&pc}
	g = gamer{player: &playStationEmulator}
	g.playGames()
}
