package interfaces

import (
	"strconv"
	"testing"
)

func TestFindMatchLinksAmount(t *testing.T) {
	doc := loadDoc("begegnungen.html")
	expectedMatchLinkCount := 14

	matchLinks := FindMatchLinks(doc)
	if expectedMatchLinkCount != len(matchLinks) {
		t.Errorf("False amount of match links. expected: %d, result: %d", expectedMatchLinkCount, len(matchLinks))
	}
}

func TestFindMatchLinksCheckFirstLink(t *testing.T) {
	doc := loadDoc("begegnungen.html")
	expectedMatchLink := "http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe?task=begegnung_spielplan&veranstaltungid=64&id=3815"

	matchLinks := FindMatchLinks(doc)
	if expectedMatchLink != matchLinks[0] {
		t.Errorf("Parsing first match link failed. expected: %s, result: %s", expectedMatchLink, matchLinks[0])
		t.Errorf(expectedMatchLink)
		t.Errorf(matchLinks[0])
	}
}

func TestFindLigaLinks(t *testing.T) {
	doc := loadDoc("uebersicht.html")
	expectedLinkCount := 5

	links := FindLigaLinks(doc)
	linkCount := len(links)

	if expectedLinkCount != linkCount {
		t.Errorf("False amount of links. expected: %d, result: %d", expectedLinkCount, linkCount)
	}
}

func TestFindSeasonsCount(t *testing.T) {
	doc := loadDoc("uebersicht.html")
	expectedseasonIdsCount := 5

	seasonIds := FindSeasons(doc)
	seasonIdsCount := len(seasonIds)

	if expectedseasonIdsCount != seasonIdsCount {
		t.Errorf("False amount of links. expected: %d, result: %d", expectedseasonIdsCount, seasonIdsCount)
	}
}

func TestFindSeasonsFirstId(t *testing.T) {
	doc := loadDoc("uebersicht.html")
	expectedseasonId := 7

	seasonIds := FindSeasons(doc)
	seasonId, _ := strconv.Atoi(seasonIds[0])

	if expectedseasonId != seasonId {
		t.Errorf("False amount of links. expected: %d, result: %d", expectedseasonId, seasonId)
	}
}
