package teamserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/concourse/atc"
	"github.com/concourse/atc/api/present"
)

func (s *Server) ListTeams(w http.ResponseWriter, r *http.Request) {
	hLog := s.logger.Session("list-teams")

	teams, err := s.teamFactory.FindTeams()
	if err != nil {
		hLog.Error("failed-to-get-teams", errors.New("sorry"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	presentedTeams := make([]atc.Team, len(teams))
	for i, team := range teams {
		presentedTeams[i] = present.Team(team)
	}

	json.NewEncoder(w).Encode(presentedTeams)
}
