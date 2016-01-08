package nrmp

import "testing"

func TestMatch(t *testing.T) {
	var pikesville, memorial, albany Program
	var dawson, chen, perez, patel Applicant

	dawson = Applicant{
		name:  "Dawson",
		ranks: []*Program{&memorial, &albany, &pikesville},
	}
	chen = Applicant{
		name:  "Chen",
		ranks: []*Program{&pikesville, &albany, &memorial},
	}
	patel = Applicant{
		name:  "Patel",
		ranks: []*Program{&albany},
	}

	pikesville = Program{
		name:  "Pikesville",
		slots: make(applicantSet, 1),
		ranks: []*Applicant{&perez, &patel, &chen},
	}
	memorial = Program{
		name:  "Memorial",
		slots: make(applicantSet, 2),
		ranks: []*Applicant{&perez, &dawson, &patel},
	}
	albany = Program{
		name:  "Albany",
		slots: make(applicantSet, 2),
		ranks: []*Applicant{&chen, &dawson, &perez, &patel},
	}

	expectedMatching := map[*Program][]*Applicant{
		&pikesville: []*Applicant{&perez},
		&memorial:   []*Applicant{&dawson},
		&albany:     []*Applicant{&chen, &patel},
	}

	err := Match([]*Applicant{&dawson, &chen, &perez, &patel})
	if err != nil {
		t.Error(err)
	}
	for prog, apps := range expectedMatching {
		if len(apps) != len(prog.slots) {
			t.Errorf("Size of %s was expected %d, but was actually %d!\n", prog.name, len(apps), len(prog.slots))
		} else {
			for _, app := range apps {
				if !prog.slots[app] {
					t.Errorf("%s did not contain %s!\n", prog.name, app.name)
				}
			}
		}
	}
}
