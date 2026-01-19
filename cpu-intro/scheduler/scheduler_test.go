package scheduler

import "testing"

func TestLoadProgram(t *testing.T) {
    t.Run("CPU 명령어 1개를 올바르게 파싱해야 한다", func(t *testing.T) {
        program := "c1,i,c2"
        p := LoadProgram(program)

        if len(p.Code) != 4 {
            t.Errorf("Expected code length 4, got %d", len(p.Code))
        }
        if p.Code[1] != "io" {
            t.Errorf("Expected first instruction 'io', got %s", p.Code[0])
        }
    })
}

func TestTickLogic(t *testing.T) {
	t.Run("CPU 명령어를 실행하면 PC가 증가하고, 마지막에 Done이 되어야 한다", func(t *testing.T) {
		s := NewScheduler()
		p := LoadProgram("c1")
		p.State = Running // 테스트를 위해 강제로 Running 설정
		s.Processes = append(s.Processes, p)

		s.Tick()

		if p.PC != 1 {
			t.Errorf("Expected PC 1, got %d", p.PC)
		}
		if p.State != Done {
			t.Errorf("Expected state Done, got %v", p.State)
		}
	})

	t.Run("IO 명령어를 만나면 상태가 Blocked로 변해야 한다", func(t *testing.T) {
		s := NewScheduler()
		p := LoadProgram("i1")
		p.State = Running
		s.Processes = append(s.Processes, p)

		s.Tick()

		if p.State != Blocked {
			t.Errorf("Expected state Blocked, got %v", p.State)
		}
	})
}
