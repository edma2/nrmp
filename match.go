// package nrmp implements the resident matching algorithm
// as described here:
// http://www.nrmp.org/match-process/match-algorithm/
package nrmp

type Program struct {
	name    string
	matches []*Applicant
	ranking map[*Applicant]int
}

type Applicant struct {
	name    string
	ranking []*Program
}

func (a Applicant) String() string {
	return a.name
}

func (p Program) String() string {
	return p.name
}

func NewProgram(name string, size int, ranking []*Applicant) *Program {
	p := Program{
		name:    name,
		matches: make([]*Applicant, size),
		ranking: make(map[*Applicant]int),
	}
	for i, a := range ranking {
		p.ranking[a] = i
	}
	return &p
}

// Match an applicant to this program and return true
// if successful, false otherwise.
func (p Program) Match(a *Applicant) bool {
	var rank int
	var ranked bool
	if rank, ranked = p.ranking[a]; !ranked {
		return false
	}
	for i, m := range p.matches {
		if m == nil {
			p.matches[i] = a
			return true
		} else {
			if rank < p.ranking[m] {
				p.matches[i] = a
				return true
			}
		}
	}
	return false
}

// Run the match algorithm. After it is finished, program slots
// will contain the assignments. If an error is encountered, return
// a non-nil error.
func Match(as []*Applicant) error {
	for _, a := range as {
		for _, p := range a.ranking {
			if p.Match(a) {
				break
			}
		}
	}
	return nil
}
