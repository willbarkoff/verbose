package verbose

import "fmt"

// Style represents a prefix for a log
type Style struct {
	Foreground int
	Background int
	Text       string
	Level      int
}

func (s *Style) generateString() string {
	if !ColorsEnabled {
		return fmt.Sprintf("%10s ", s.Text)
	}
	return fmt.Sprintf("\033[9%d;4%dm%10s \033[m ", s.Foreground, s.Background, s.Text)
}

// Println prints a log item
func (s *Style) Println(a ...interface{}) {
	if MaximumLogLevel < s.Level {
		return
	}
	fmt.Print(s.generateString())
	fmt.Println(a...)
}
