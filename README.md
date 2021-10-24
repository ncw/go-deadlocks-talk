# Go Deadlocks Talk

This is sample code for my Go Deadlocks talk.

## Simple

- [simple](simple/simple-deadlock.go) - a super simple deadlock
- [simple2](simple2/simple-deadlock2.go) - the same but with an extra Go deadlock detector defeating go routine

## File and Directory example

- [vfs1](vfs1/vfs.go) - starting out simply with no locking
- [vfs2](vfs2/vfs.go) - adding locking to cause a deadlock
- [vfs3-release-mutex](vfs3-release-mutex/vfs.go) - avoid the deadlock by releasing the mutex
- [vfs4-ro-variable](vfs4-ro-variable/vfs.go) - avoid the deadlock with a read only variable
- [vfs5-go-deadlock](vfs5-go-deadlock/vfs.go) - demonstrate the use of [go-deadlock](https://github.com/sasha-s/go-deadlock)

## Channel deadlock

- [channel-deadlock](channel-deadlock/job.go) - the channel deadlock example
- [channel-deadlock-fixed](channel-deadlock-fixed/job.go) - the fixed example

## Read lock

- [read-lock](read-lock/read-lock.go) - demonstrate recursively taking read locks is a bad idea
