package main

//
//import "sync"
//
////func doSomething() error {
////	var mu sync.Mutex
////	mu.Lock()
////
////	r1, err := OpenResource1()
////	if err != nil {
////		mu.Unlock()
////		return err
////	}
////
////	r2, err := OpenResource2()
////	if err != nil {
////		r1.Close()
////		mu.Unlock()
////		return err
////	}
////
////	r3, err := OpenResource3()
////	if err != nil {
////		r2.Close()
////		r1.Close()
////		mu.Unlock()
////		return err
////	}
////
////	// 使用r1，r2, r3
////	err = doWithResources()
////	if err != nil {
////		r3.Close()
////		r2.Close()
////		r1.Close()
////		mu.Unlock()
////		return err
////	}
////
////	r3.Close()
////	r2.Close()
////	r1.Close()
////	mu.Unlock()
////	return nil
////}
//
//func doSomething() error {
//	var mu sync.Mutex
//	mu.Lock()
//	defer mu.Unlock()
//
//	r1, err := OpenResource1()
//	if err != nil {
//		return err
//	}
//	defer r1.Close()
//
//	r2, err := OpenResource2()
//	if err != nil {
//		return err
//	}
//	defer r2.Close()
//
//	r3, err := OpenResource3()
//	if err != nil {
//		return err
//	}
//	defer r3.Close()
//
//	// 使用r1，r2, r3
//	return doWithResources()
//}
