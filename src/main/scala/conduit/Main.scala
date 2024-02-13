object Main {
  def main(args: Array[String]): Unit = {
    println("Hello, world")

    val greeter = new Greeter("scala")
    println(greeter.greet())
  }
}


class Greeter(name: String) {
  def greet(): String = s"Hello, $name!"
}
