package main

import "fmt"

type Invoker struct {
	command Command
}

func (i *Invoker) invoke() {
	i.command.execute()
}

type Command interface {
	execute()
}

type Backup struct {
	db Database
}

func (b *Backup) execute() {
	b.db.takeBackup()
}

type Restore struct {
	db Database
}

func (b *Restore) execute() {
	b.db.restoreDB()
}

type Database interface {
	takeBackup()
	restoreDB()
}

type MongoDB struct{}

func (m *MongoDB) takeBackup() {
	fmt.Println("Backup MongoDB database.")
}

func (m *MongoDB) restoreDB() {
	fmt.Println("Restore MongoDB database.")
}

func main() {
	mgDB := &MongoDB{}

	backupCmd := &Backup{
		db: mgDB,
	}

	restoreCmd := &Restore{
		db: mgDB,
	}

	backupInvoker := &Invoker{
		command: backupCmd,
	}

	restoreInvoker := &Invoker{
		command: restoreCmd,
	}
	backupInvoker.invoke()
	restoreInvoker.invoke()
}
