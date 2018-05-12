package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type card struct {
	suit  string
	value int
}

type cards []card

func newCards(fd io.Reader) (cards, int) {
	N := 0
	sc := bufio.NewScanner(fd)
	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}
	cards := make(cards, N)
	if sc.Scan() {
		for i, val := range strings.Fields(string(sc.Text())) {
			cards[i].suit = string(val[0])
			cards[i].value, _ = strconv.Atoi(string(val[1]))
		}
	}
	return cards, N
}

func (hand cards) String() string {
	ability := make([]string, len(hand))
	for i, card := range hand {
		ability[i] = card.suit + strconv.Itoa(card.value)
	}
	return strings.Join(ability, " ")
}

func main() {
	cards, N := newCards(os.Stdin)
	bubble := bubbleSort(cards, N)
	fmt.Println(bubble.String())
	fmt.Println("Stable")
	selection := selectionSort(cards, N)
	fmt.Println(selection.String())
	fmt.Println(isStable(selection, bubble))
}

func bubbleSort(hand cards, N int) cards {
	sorted := make(cards, N)
	copy(sorted, hand)
	for i := 0; i < N-1; i++ {
		for j := N - 1; j > i; j-- {
			if sorted[j].value < sorted[j-1].value {
				sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			}
		}
	}
	return sorted
}

func selectionSort(hand cards, N int) cards {
	sorted := make(cards, N)
	copy(sorted, hand)
	for i := 0; i < N-1; i++ {
		minj := i
		for j := i; j < N; j++ {
			if sorted[j].value < sorted[minj].value {
				minj = j
			}
		}
		if minj == i {
			continue
		}
		sorted[i], sorted[minj] = sorted[minj], sorted[i]
	}
	return sorted
}

// isStable reports whether the target slice is sorted keeping the stability
func isStable(target, ref cards) string {
	for i := 0; i < len(target); i++ {
		if target[i].suit != ref[i].suit {
			return "Not stable"
		}
	}
	return "Stable"
}
