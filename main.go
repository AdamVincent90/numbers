package main

func main() {

	normalNumbers := newNumbers()
	normalNumbers.print()
	normalNumbers.evenOrOdd()

	randomNumbers := newRandomNumbers()
	randomNumbers.print()
	randomNumbers.evenOrOdd()
}
