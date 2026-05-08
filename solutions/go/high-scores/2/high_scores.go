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
	top := slices.Clone(scores)

	slices.Sort(top)
	slices.Reverse(top)

	n := min(len(top), size)

	return top[:n]
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	var latest int
	if len(scores) > 0 {
		latest = scores[len(scores)-1]
	}
	top := getTop(3, scores)
	best := 0
	if len(top) > 0 {
		best = top[0]
	}
	return &HighScores{
		scores: scores,
		latest: latest,
		best:   best,
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
