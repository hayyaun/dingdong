package game

import "github.com/hayyaun/dingdong/utils"

type Chance struct {
	v int // value
	w int // weight
}

func (chance *Chance) Weight() int {
	return chance.w
}

func (chance *Chance) Value() any {
	return chance.v
}

var chances = []Chance{
	{v: 4, w: 1}, {v: 3, w: 1}, {v: 2, w: 2}, {v: 1, w: 100}, {v: 0, w: 200},
}

func chanceToWeightable(chances []Chance) []utils.Weighable {
	weightables := make([]utils.Weighable, len(chances))
	for i := range chances {
		weightables[i] = &chances[i]
	}
	return weightables
}
