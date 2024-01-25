package service

import (
	"context"
	"joust/data"
	"joust/model"
	"math/rand"
)

const (
	NO_HIT      = 0
	LANCE_HIT   = 1
	LANCE_BREAK = 5
	UNHORSE     = 10
)

type MatchService struct {
	KnightRepo *data.KnightRepo
	HorseRepo  *data.HorseRepo
	MatchRepo  *data.MatchRepo
}

func (s *MatchService) SimMatch(ctx context.Context, homeID, awayID uint) (*model.Match, error) {
	homeKnight, err := s.KnightRepo.ReadById(ctx, homeID)
	if err != nil {
		return &model.Match{}, err
	}

	awayKnight, err := s.KnightRepo.ReadById(ctx, awayID)
	if err != nil {
		return &model.Match{}, err
	}

	homeScores := []model.BoxScore{}
	awayScores := []model.BoxScore{}

	for i := 1; i <= 5; i++ {
		homeScore, awayScore := s.Scores(homeKnight.KnightStats, awayKnight.KnightStats)
		homeScore.PassNumber = i
		awayScore.PassNumber = i
		homeScores = append(homeScores, *homeScore)
		awayScores = append(awayScores, *awayScore)
	}

	homeTotal := uint(s.SumBoxScores(homeScores))
	awayTotal := uint(s.SumBoxScores(awayScores))

	match := &model.Match{
		BoxScores: append(homeScores, awayScores...),
		HomeId:    homeID,
		HomeTotal: homeTotal,
		AwayId:    awayID,
		AwayTotal: awayTotal,
		WinnerId: func() uint {
			if homeTotal < awayTotal {
				return awayID
			} else {
				return homeID
			}
		}(),
	}

	if err := s.MatchRepo.Create(ctx, match); err != nil {
		return &model.Match{}, err
	}

	return match, nil
}

func (s *MatchService) Scores(homeStats, awayStats model.KnightStats) (*model.BoxScore, *model.BoxScore) {
	return s.Score(homeStats, awayStats), s.Score(awayStats, homeStats)
}

func (s *MatchService) Score(my, their model.KnightStats) *model.BoxScore {
	boxScore := &model.BoxScore{
		KnightID: my.ID,

		AccuracyRoll:   0,
		AvoidanceRoll:  0,
		PowerRoll:      0,
		DeflectionRoll: 0,
		TechniqueRoll:  0,
		BalanceRoll:    0,
	}

	s.ToHit(boxScore, &my, &their)

	return boxScore
}

func (s *MatchService) ToHit(boxScore *model.BoxScore, my, their *model.KnightStats) *model.BoxScore {
	boxScore.AccuracyRoll = s.Roll(my.Accuracy) + 5
	boxScore.AvoidanceRoll = s.Roll(their.Avoidance)

	if boxScore.AccuracyRoll > boxScore.AvoidanceRoll {
		s.Hit(boxScore, my, their)
	} else {
		boxScore.Score = NO_HIT
	}

	return boxScore
}

func (s *MatchService) Hit(boxScore *model.BoxScore, my, their *model.KnightStats) *model.BoxScore {
	if s.Unhorse(boxScore, my, their) {
		boxScore.Score = UNHORSE
	} else if s.LanceBreak(boxScore, my, their) {
		boxScore.Score = LANCE_BREAK
	} else {
		boxScore.Score = LANCE_HIT
	}
	return boxScore
}

func (s *MatchService) LanceBreak(boxScore *model.BoxScore, my, their *model.KnightStats) bool {
	boxScore.PowerRoll = s.Roll(my.Power) + 5
	boxScore.DeflectionRoll = s.Roll(their.Deflection)
	return boxScore.PowerRoll > boxScore.DeflectionRoll
}

func (s *MatchService) Unhorse(boxScore *model.BoxScore, my, their *model.KnightStats) bool {
	boxScore.TechniqueRoll = s.Roll(my.Technique) - 5
	boxScore.BalanceRoll = s.Roll(their.Balance)
	return boxScore.TechniqueRoll > boxScore.BalanceRoll
}

func (s *MatchService) Roll(score int) int {
	return rand.Intn(20) + 1 + score
}

func (s *MatchService) SumBoxScores(arr []model.BoxScore) int {
	sum := 0
	for _, boxScore := range arr {
		sum += boxScore.Score
	}
	return sum
}
