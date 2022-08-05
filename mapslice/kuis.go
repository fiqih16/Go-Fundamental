package mapslice

import "fmt"

func Score() {
	// Hitung nilai rata rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var total int
	for _, score := range scores {
		total = total + score
	}
	length := len(scores)
	average := float64(total) / float64(length)
	fmt.Println("Rata-rata nilai :", average)
}

func Scoretwo(){
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var goodScores []int
	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println("Nilai yang memenuhi kriteria :", goodScores)
}