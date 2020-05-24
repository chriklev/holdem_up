package startinghandprobs

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/spiritofsim/phe"
)

func findWinrate(x1, x2 int) float32 {
	count := 0
	winCount := 0
	drawCount := 0
	// For each possible opponent hand
	for y1 := 1; y1 < 53; y1++ {
		if y1 != x1 && y1 != x2 {
			for y2 := 1; y2 < y1; y2++ {
				if y2 != x1 && y2 != x2 {
					// For each possible table
					for a := 5; a < 53; a++ {
						if a != x1 && a != x2 && a != y1 && a != y2 {
							for b := 4; b < a; b++ {
								if b != x1 && b != x2 && b != y1 && b != y2 {
									for c := 3; c < b; c++ {
										if c != x1 && c != x2 && c != y1 && c != y2 {
											for d := 2; d < c; d++ {
												if d != x1 && d != x2 && d != y1 && d != y2 {
													for e := 1; e < d; e++ {
														if e != x1 && e != x2 && e != y1 && e != y2 {
															playerScore, _ := phe.Eval(phe.Card(a), phe.Card(b), phe.Card(c), phe.Card(d), phe.Card(e), phe.Card(x1), phe.Card(x2))
															opponentScore, _ := phe.Eval(phe.Card(a), phe.Card(b), phe.Card(c), phe.Card(d), phe.Card(e), phe.Card(y1), phe.Card(y2))

															result := playerScore.Compare(opponentScore)
															if result == 1 {
																winCount++
															} else if result == 0 {
																drawCount++
															}
															count++
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return (float32(winCount) + float32(drawCount)/2) / float32(count)
}

func worker(hands <-chan [2]int, winRates *[13][13]float32, count *int) {
	for hand := range hands {
		x1 := hand[0]
		x2 := hand[1]
		wr := findWinrate(x1, x2)

		// Add winrate to table
		x1--
		x2--
		if x1%4 == x2%4 {
			(*winRates)[x1/4][x2/4] = wr
		} else {
			(*winRates)[x2/4][x1/4] = wr
		}
		// Increment and print progress counter
		*count++
		fmt.Printf("\r%v/169", *count)
	}
}

func writeToCSV(winrates [13][13]float32, path string) {
	csvfile, err := os.Create(path)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)

	for _, wr := range winrates {
		strwr := make([]string, 13)
		for j := range wr {
			strwr[j] = strconv.FormatFloat(float64(wr[j]), 'f', 4, 32)
		}
		_ = csvwriter.Write(strwr)
	}

	csvwriter.Flush()
	csvfile.Close()
	fmt.Printf("Data written to: %v\n", path)
}

func main() {
	t0 := time.Now()
	fmt.Printf("Tree search started at %v\n", t0)

	winRates := [13][13]float32{}
	count := 0

	// Channel for workers to recieve hands
	hands := make(chan [2]int, 169)

	nWorkers := 8
	var wg sync.WaitGroup
	wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func() {
			worker(hands, &winRates, &count)
			wg.Done()
		}()
	}

	//for each possible starting hand
	for f := 0; f < 2; f++ {
		for i := 1 + f; i < 53; i += 4 {
			for j := 1; j < i; j += 4 {
				hands <- [2]int{i, j}
			}
		}
	}
	close(hands)

	wg.Wait()
	fmt.Println("\n----Done----")

	fileName := "starting_hand_win_rates.csv"
	writeToCSV(winRates, fileName)

	t1 := time.Now()
	fmt.Printf("Time elapsed: %v\n", t1.Sub(t0))
}
