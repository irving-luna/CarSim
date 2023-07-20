package parts

import "fmt"

type Motor struct {
	On  bool
	Rpm int
}

func (m *Motor) SetOn(on bool) {
	m.On = on
}

func (m *Motor) IsOn() bool {
	return m.On
}

func (m *Motor) SetRPM(rpm int, sb SeatbeltManipulator) error {
	if rpm < 0 {
		return fmt.Errorf("Negative RPM")
	}
	if rpm > 0 && sb.GetState() == SeatbeltStateEngauged {
		m.Rpm = rpm
	}
	return fmt.Errorf("Seatbelt must be engauged to increase RPM")
}

func (m *Motor) GetRPM() int {
	return m.Rpm
}

func (m *Motor) SetWheelsRPM(wm WheelsManipulator) error {
	wheelRpm :=m.GetRPM()/2

	return wm.SetRPM(wheelRpm)
}

func (m *Motor) Print() {
	fmt.Printf("Motor: \n\tOn: %t, \n\tRPM: %d\n", m.On, m.Rpm)
}
