package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"syscall"
)

func setUid() {

	go func() {

		runtime.LockOSThread()
		user, err := user.Lookup("mac")
		if err == nil {
			log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)
			uid, _ := strconv.Atoi(user.Uid)
			gid, _ := strconv.Atoi(user.Gid)
			if err := syscall.Setuid(uid); err != nil {
				fmt.Println("syscall.Stuid err", err)
			}
			err := syscall.Setgid(gid)
			if err != nil {
				fmt.Println("syscall.Setgid err", err)
			}
		}
	}()
}

func main() {
	cmd := exec.Command("sudo su ")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	user, err := user.Lookup("song")
	if err == nil {
		log.Printf("uid=%s,gid=%s", user.Uid, user.Gid)

		uid, _ := strconv.Atoi(user.Uid)
		gid, _ := strconv.Atoi(user.Gid)

		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	}

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
