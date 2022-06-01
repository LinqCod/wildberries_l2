package exercise_7

func or(channels ...<-chan interface{}) <-chan interface{} {
	exit := make(chan interface{})
	//Конкурентно ожидаем закрытия одного из каналов
	//Как только один из них закроется,
	//используем вспомогательный канал exit для оповещения
	//и заканчиваем выполнение функции
	for _, v := range channels {
		go func(c <-chan interface{}) {
			select {
			case <-c:
				exit <- "end"
				close(exit)
			}
		}(v)
	}

	<-exit
	return exit
}
