package system

func TryCatch(fun func(), handler func(err interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}