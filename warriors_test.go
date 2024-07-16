package wapi_test

import (
	"testing"
	"time"

	wapi "github.com/SabienNguyen/WAPI"
)

func TestGetRoster(t *testing.T) {
	roster, err := wapi.GetRoster()

	// Check if there's an error
	if err != nil {
		t.Fatalf("GetRoster returned an error: %v", err)
	}

	// Check if the roster is not empty
	if len(roster) == 0 {
		t.Error("GetRoster returned an empty roster")
	}

	// Check if each player has the required fields
	for i, player := range roster {
		if player.Name == "" {
			t.Errorf("Player %d has an empty name", i)
		}
		if player.Position == "" {
			t.Errorf("Player %d has an empty position", i)
		}
	}

	// Optional: Print out the first player for manual verification
	if len(roster) > 0 {
		t.Logf("First player: %+v", roster[0])
	}
}

func TestGetSchedule(t *testing.T) {
	schedule, err := wapi.GetSchedule()

	// Check if there's an error
	if err != nil {
		t.Fatalf("GetSchedule returned an error: %v", err)
	}

	// Check if the schedule is not empty
	if len(schedule) == 0 {
		t.Error("GetSchedule returned an empty schedule")
	}

	// Check each game in the schedule
	for i, game := range schedule {
		// Check if the date is not empty
		if game.Date == "" {
			t.Errorf("Game %d has an empty date", i)
		}

		// Try to parse the date to ensure it's in a valid format
		_, err := time.Parse("Jan 2", game.Date)
		if err != nil {
			t.Errorf("Game %d has an invalid date format: %s", i, game.Date)
		}
	}

	// Print out the first few games for manual verification
	for i := 0; i < min(5, len(schedule)); i++ {
		t.Logf("Game %d: %+v", i, schedule[i])
	}
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestGetTeamInfo(t *testing.T) {
	teamInfo, err := wapi.GetTeamInfo()

	// Check if there's an error
	if err != nil {
		t.Fatalf("GetTeamInfo returned an error: %v", err)
	}

	// Check if the record is not empty and in the correct format
	if teamInfo.Record == "" {
		t.Error("Team record is empty")
	}

	// Check if the seed is not empty and in the correct format
	if teamInfo.Seed == "" {
		t.Error("Team seed is empty")
	}

	// Print the team info for manual verification
	t.Logf("Team Info: Record: %s, Seed: %s", teamInfo.Record, teamInfo.Seed)
}
