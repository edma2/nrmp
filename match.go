// package nrmp implements the resident matching algorithm
// as described here:
// http://www.nrmp.org/match-process/match-algorithm/
package nrmp

type applicantSet map[*Applicant]bool

type Program struct {
	name  string
	slots applicantSet
	ranks []*Applicant
}

type Applicant struct {
	name  string
	ranks []*Program
}

// Run the match algorithm. After it is finished, program slots
// will contain the assignments. If an error is encountered, return
// a non-nil error.
func Match(apps []*Applicant) error {
	return nil
}
