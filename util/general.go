/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package util

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ExitWithCtrlC(waitDuration ...time.Duration) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan

		wait := time.Second * 10
		if len(waitDuration) > 0 {
			wait = waitDuration[0]
		}

		time.Sleep(wait)
		os.Exit(1)
	}()
}

func WaitUntilCtrlC(then ...func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	for _, fn := range then {
		fn()
	}
}
