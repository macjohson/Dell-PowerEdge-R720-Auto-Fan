package utils

type Account struct {
	Username string
	Password string
	IP       string
}

func (a *Account) Prefix() []string {
	return []string{"-I", "lanplus", "-H", a.IP, "-U", a.Username, "-P", a.Password}
}

func (a *Account) Command(args []string) []string {
	commands := []string{"ipmitool"}

	commands = append(commands, a.Prefix()...)

	for _, arg := range args {
		commands = append(commands, arg)
	}

	return commands
}
