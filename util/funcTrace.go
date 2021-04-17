package util

import (
	"log"
	"runtime"
)

// TraceToLog (Português): Retorna os dados da função pai em caso de evento para ser colocado em log
//   file: caminho do arquivo
//   linha: linha do evento
//   funcName: nome da função onde o evento ocorreu
func TraceToLog() (file string, line int, funcName string, ok bool) {
	var pc uintptr

	pc, file, line, ok = runtime.Caller(1)
	if !ok {
		log.Printf("TraceToLog().error: runtime.Caller() unknown error")
		return
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		funcName = ""
		log.Printf("unknownFuncName()[%v]: %v", line, file)
		return
	}

	funcName = fn.Name()
	log.Printf("%v() - line: %v - %v", funcName, line, file)
	return
}
