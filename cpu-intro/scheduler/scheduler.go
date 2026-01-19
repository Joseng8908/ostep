package scheduler

import(
	"strings"
	"strconv"
)

type State int

const (
	Ready State = iota
	Running
	Blocked
	Done
)

type Process struct {
	ID int
	Code []string
	PC int	//program counter, 현재 몇 번째 명령어 수행중인지
	State State //현재 프로세스의 상태
}

type Scheduler struct {
	Processes []*Process
	Current int
	Ticks int
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Processes: []*Process{},
		Current: 0,
		Ticks: 0,
	}
}

func LoadProgram(program string) *Process {
	p := &Process{Code: []string{}}
	parts := strings.Split(program, ",")
	for _, part := range parts{
		if len(part) == 0 {continue}
		
		opcode := part[0]

		switch opcode {
		case 'c': {
				num, _ := strconv.Atoi(part[1:])
				for i := 0; i < num; i++ {
					p.Code = append(p.Code, "cpu")
				}
			}

		case 'i': {
				p.Code = append(p.Code, "io")
			}
		}
	}
	return p 
}

func (s *Scheduler) Tick() {
	if len(s.Processes) == 0 {
		return
	}

	p := s.Processes[s.Current] 

	if p.State != Running || p.PC >= len(p.Code) {
		return
	}

	inst := p.Code[p.PC]

	switch inst {
	case "cpu" :
		p.PC++
	case "io" :
		p.State = Blocked
		p.PC++
		return
	}

	if p.PC >= len(p.Code) {
		p.State = Done
	}
	s.Ticks++
}
