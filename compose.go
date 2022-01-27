package HandlerFramework

func ComposeHandlers(handlers ...HandlerFunc) *HandlerApp {
	return newHandlerApp(handlers)
}
