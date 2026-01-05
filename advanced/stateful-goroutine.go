package main

type StatrfulWorker struct {
	count int
	ch    chan int
}

func (worker *StatrfulWorker) start() {
	go func() {
		for {
			select {
			case value := <-worker.ch:
				worker.count += value
			default:
				return
			}
		}
	}()

}
