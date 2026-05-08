package highscores

import (
	"slices"
)

type HighScores struct {
	scores []int
	latest int
	best   int
	top    []int
}

func getTop(size int, scores []int) []int {
	top := make([]int, size)
	tmp := make([]int, size+1)
	for _, score := range scores {
		for i, topScore := range top {
			if score > topScore {
				tmp = slices.Insert(top, i, score)
				n := min(len(tmp), size)
				top = tmp[:n]
				break
			}
		}
	}
	n := min(len(scores), size)
	return top[:n]
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	var latest int
	if len(scores) > 0 {
		latest = scores[len(scores)-1]
	}
	top := getTop(3, scores)
	return &HighScores{
		scores: scores,
		latest: latest,
		best:   top[0],
		top:    top,
	}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.latest
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	return s.best
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	return s.top
}
