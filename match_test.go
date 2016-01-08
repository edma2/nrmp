package nrmp

import "testing"

func TestMatch(t *testing.T) {
	var pikesville, memorial, albany *Program
	var dawson, chen, perez, patel *Applicant

	dawson = &Applicant{
		name: "Dawson",
	}
	chen = &Applicant{
		name: "Chen",
	}
	perez = &Applicant{
		name: "Perez",
	}
	patel = &Applicant{
		name: "Patel",
	}

	pikesville = NewProgram("Pikesville", 1, []*Applicant{perez, patel, chen})
	memorial = NewProgram("Memorial", 2, []*Applicant{perez, dawson, patel})
	albany = NewProgram("Albany", 2, []*Applicant{chen, dawson, perez, patel})

	dawson.ranking = []*Program{memorial, albany, pikesville}
	chen.ranking = []*Program{pikesville, albany, memorial}
	perez.ranking = []*Program{pikesville, albany}
	patel.ranking = []*Program{albany}

	expectedMatching := map[*Program][]*Applicant{
		pikesville: []*Applicant{perez},
		memorial:   []*Applicant{dawson},
		albany:     []*Applicant{chen, patel},
	}

	err := Match([]*Applicant{dawson, chen, perez, patel})
	if err != nil {
		t.Error(err)
	}
	for prog, apps := range expectedMatching {
		for i, actual := range prog.matches {
			var expected *Applicant
			if i < len(apps) {
				expected = apps[i]
			} else {
				expected = nil
			}
			if expected != actual {
				t.Errorf("%s: expected %s in position %d, but got %s!\n", prog, expected, i, actual)
			}
		}
	}
}
