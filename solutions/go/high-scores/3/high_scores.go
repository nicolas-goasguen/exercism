package highscores

import "slices"

type HighScores struct {
	scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{
		scores: scores,
	}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	n := len(s.scores)

	if n == 0 {
		return 0
	}

	return s.scores[n-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	if len(s.scores) == 0 {
		return 0
	}
	return slices.Max(s.scores)
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	top := slices.Clone(s.scores)

	slices.Sort(top)
	slices.Reverse(top)

	n := min(len(top), 3)

	return top[:n]
}
