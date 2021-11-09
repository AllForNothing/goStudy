package mylogger

type ILog interface {
	Debug()
	Info()
}


type MyLogger struct {
	
}

func NewMyLogger()  MyLogger {
	return MyLogger{}
}

func (l MyLogger) Debug()  {
	fmt.Printf("", )
}