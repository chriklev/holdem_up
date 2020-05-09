package main

import "fmt"

func findWinrate(x1, x2 int, count *int) {
	// For each possible opponent hand
	for y1 := 0; y1 < 52; y1++ {
		if y1 != x1 && y1 != x2 {
			for y2 := 0; y2 < y1; y2++ {
				if y2 != x1 && y2 != x2 {
					// For each possible table
					for a := 0; a < 52; a++ {
						if a != x1 && a != x2 && a != y1 && a != y2 {
							for b := 0; b < a; b++ {
								if b != x1 && b != x2 && b != y1 && b != y2 {
									for c := 0; c < b; c++ {
										if c != x1 && c != x2 && c != y1 && c != y2 {
											for d := 0; d < c; d++ {
												if d != x1 && d != x2 && d != y1 && d != y2 {
													for e := 0; e < d; e++ {
														if e != x1 && e != x2 && e != y1 && e != y2 {
															*count++
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
}

func main() {
	count := 0
	//for each possible starting hand
	for f := 0; f < 2; f++ {
		for i := 0; i < 13; i++ {
			for j := 13 * f; j < i+14*f; j++ {
				findWinrate(i, j, &count)
			}
		}
	}
	fmt.Println(count)
}
