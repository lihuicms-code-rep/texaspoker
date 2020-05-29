package coretest

import (
	"github.com/lihuicms-code-rep/texaspoker/core/coreimplement"
	"testing"
)

func TestNewDeckOfCards(t *testing.T) {
	deckCards := coreimplement.NewDeckOfCards(52)
	deckCards.String()
}
