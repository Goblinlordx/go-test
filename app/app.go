package app

type App struct {
	val string
}

func ProvideApp() App {
	return App{
		val: "test",
	}
}

func (s *App) Test() string {
	return s.val
}
