package types

import "github.com/google/uuid"

type player struct {
	ID   uuid.UUID
	Name string50
}

func (p *player) IsEqual(p1 player) bool {
	return p.ID == p1.ID
}

func NewPlayer(n string) (player, error) {
	name, err := NewString50(n)
	if err != nil {
		return player{}, err
	}

	p := player{
		ID:   uuid.New(),
		Name: name,
	}

	return p, nil
}

func NewPlayerWithId(id, n string) (player, error) {
	tmpId, err := uuid.Parse(id)
	if err != nil {
		return player{}, err
	}

	p, err := NewPlayer(n)
	if err != nil {
		return player{}, err
	}

	p.ID = tmpId

	return p, nil
}
