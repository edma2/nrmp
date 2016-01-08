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

type queue interface {
	get() *Applicant
	add(a *Applicant)
	empty() bool
}

type sliceQueue struct {
	items []*Applicant
}

func (q *sliceQueue) get() *Applicant {
	a := q.items[0]
	q.items = q.items[1:]
	return a
}

func (q *sliceQueue) add(a *Applicant) {
	q.items = append(q.items, a)
}

func (q *sliceQueue) empty() bool {
	return len(q.items) == 0
}

func newQueue(items []*Applicant) queue {
	return &sliceQueue{items}
}

func (a *Applicant) String() string {
	return a.name
}

func (p *Program) String() string {
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
// if successful, false otherwise. The existing match that
// was kicked out is also returned.
func (p *Program) match(a *Applicant) (*Applicant, bool) {
	var rank int
	var ranked bool
	if rank, ranked = p.ranking[a]; !ranked {
		return nil, false
	}
	for i, m := range p.matches {
		if m == nil {
			p.matches[i] = a
			return nil, true
		} else {
			if rank < p.ranking[m] {
				p.matches[i] = a
				return m, true
			}
		}
	}
	return nil, false
}

// Run the match algorithm. After it is finished, program slots
// will contain the assignments. If an error is encountered, return
// a non-nil error.
func Match(as []*Applicant) error {
	q := newQueue(as)
	for !q.empty() {
		a := q.get()
		for _, p := range a.ranking {
			if old, ok := p.match(a); ok {
				if old != nil {
					q.add(old)
				}
				break
			}
		}
	}
	return nil
}
