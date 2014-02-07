light := make(chan int)
go func(){
  //Light task
  light <- 1
}()
heavy := make(chan int)
go func() {
  // Heavy task
  heavy <- 1
}()
<- light
<- heavy
//Code executed after both tasks have been finished
