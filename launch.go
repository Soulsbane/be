package main

func main() {
	var commands Commands

	commands.addCommand("minion", "java -jar /home/soulsbane/bin/minion/Minion-jfx.jar")
	commands.dump()
	commands.run("minion")
	commands.run("what")
}
