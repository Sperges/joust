package service

import (
	"context"
	"dice"
	"joust/data"
	"joust/model"
)

// TODO: find a better place for rollers
var KNIGHT_AGE_ROLLER = dice.NewRoller().Dice(4, 6).DropHigh(1).AddModifier("base", 12)
var HORSE_AGE_ROLLER = dice.NewRoller().Dice(4, 6).DropLow(1).AddModifier("base", 2)

type GenerateService struct {
	KnightRepo *data.KnightRepo
}

func (s *GenerateService) Knights(ctx context.Context, amount int) error {
	for i := 0; i < amount; i++ {
		if err := s.Knight(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *GenerateService) Knight(ctx context.Context) error {
	knight, err := s.RandomKnight(ctx)
	if err != nil {
		return err
	}

	if err := s.KnightRepo.Create(ctx, knight); err != nil {
		return err
	}

	return nil
}

func (s *GenerateService) RandomKnight(ctx context.Context) (*model.Knight, error) {
	first_name, err := s.KnightRepo.RandomFirstName(ctx)
	if err != nil {
		return &model.Knight{}, err
	}

	last_name, err := s.KnightRepo.RandomLastName(ctx)
	if err != nil {
		return &model.Knight{}, nil
	}

	region, err := s.KnightRepo.RandomRegion(ctx)
	if err != nil {
		return &model.Knight{}, err
	}

	age := uint(KNIGHT_AGE_ROLLER.Roll())

	stats, err := s.RandomKnightStats(ctx, age)
	if err != nil {
		return &model.Knight{}, nil
	}

	return &model.Knight{
		KnightStats: *stats,
		UserID:      0,
		FirstName:   first_name,
		LastName:    last_name,
		Age:         age,
		Region:      region,
	}, nil
}

func (s *GenerateService) RandomKnightStats(_ctx context.Context, age uint) (*model.KnightStats, error) {

	// TODO: holy fuck, do better than this
	ageModifier := 0
	if age > 17 {
		ageModifier = 1
	} else if age > 22 {
		ageModifier = 2
	} else if age > 26 {
		ageModifier = 1
	} else if age > 30 {
		ageModifier = 0
	}

	statsRoller := dice.NewRoller().Dice(4, 6).DropLow(1).AddModifier("age", ageModifier)

	chivalryRoll := dice.NewRoller().Dice(1, 20).Roll()

	return &model.KnightStats{
		KnightID: 0,

		Power:       statsRoller.Roll(),
		Avoidance:   statsRoller.Roll(),
		Technique:   statsRoller.Roll(),
		Accuracy:    statsRoller.Roll(),
		Balance:     statsRoller.Roll(),
		Riding:      statsRoller.Roll(),
		Deflection:  statsRoller.Roll(),
		Valor:       statsRoller.Roll(),
		Chivalry:    chivalryRoll,
		Skulduggery: 20 - chivalryRoll,
		Showmanship: statsRoller.Roll(),
	}, nil
}
