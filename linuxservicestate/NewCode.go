package linuxservicestate

import "github.com/alimtvnetwork/core-v8/constants"

func NewCode(code int) ExitCode {
	if code >= constants.MaxUnit8AsInt || code < 0 {
		return InvalidCode
	}

	codeByte := byte(code)

	if codeByte >= BasicEnumImpl.Max() {
		return InvalidCode
	}

	return NewCodeMapping(codeByte)
}
