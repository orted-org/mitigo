package mitigo

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestADtoBS(t *testing.T) {

	firstEnglish := firstEnglishDate
	out, err := ADtoBS(firstEnglish)
	if err != nil {
		t.Error(err)
	}
	if out != firstNepaliDate {
		t.Error(fmt.Errorf("must be %v, got %v", firstNepaliDate, out))
	}
}

func TestBSToAD(t *testing.T) {
	firstNepali := firstNepaliDate
	out, err := BStoAD(firstNepali)
	if err != nil {
		t.Error(err)
	}
	if out != firstEnglishDate {
		t.Error(fmt.Errorf("must be %v, got %v", firstEnglishDate, out))
	}
}

func TestDaysBetweenBSDates(t *testing.T) {
	days := rand.Intn(100)
	today := time.Now()
	future := today.Add(24 * time.Hour * time.Duration(days))

	todayBS, _ := ADtoBS(today)
	futureBS, _ := ADtoBS(future)

	diff := DaysBetweenBSDates(todayBS, futureBS)
	if days != diff {
		t.Error(fmt.Errorf("must be %d, got %d", days, diff))
	}
}

func TestBSDaysAfter(t *testing.T) {
	days := rand.Intn(100)
	today := time.Now()
	future := today.Add(24 * time.Hour * time.Duration(days))

	todayBS, _ := ADtoBS(today)
	futureBS, _ := ADtoBS(future)

	out, _ := BSDaysAfter(todayBS, days)

	if out != futureBS {
		t.Error(fmt.Errorf("must be %v, got %v", futureBS, out))
	}
}
