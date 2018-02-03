package main

import (
  "fmt"
)

func swap(x, y string) (string, string) {
    return y, x
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func add(x string, y string) string {
  if(len(x) > len(y)) {
    x, y = swap(x, y)
  }

  var str string = ""
  n1 := len(x)
  n2 := len(y)
  diff := n2 - n1

  carry := 0
  for i := n1 - 1; i >= 0; i-- {
    sum := (int(x[i] - '0') + int(y[i+diff] - '0') + carry);
    str = str + string(sum % 10 + '0')
    carry = sum / 10
  }

  for i := n2 - n1 - 1; i >= 0; i-- {
    sum := (int(y[i] - '0') + carry);
    str = str + string(sum % 10 + '0')
    carry = sum / 10
  }

   if(carry > 0) {
      str = str + string(carry + '0')
   }

   str = reverse(str)
   return str
}

func multiplySingleDigit(y byte, x string, count int, c chan string) {
  var str string = ""
  digit := int(y - '0')

  n := len(x)

  carry := 0
  for i := n - 1; i >= 0; i-- {
    sum := (int(x[i] - '0') * digit + carry)
    str = str + string(sum % 10 + '0')
    carry = sum / 10
  }

  if(carry > 0) {
     str = str + string(carry + '0')
  }

  str = reverse(str)

  // adding count number of zeroes to the end
  for i := 0; i < count; i++ {
    str = str + string('0')
  }

  c <- str
}

func multiply(x string, y string) string {
  var str string = ""

  n := len(y)

  c := make(chan string, n)

  for i := n - 1; i >= 0; i-- {
    count := n - 1 - i
    // count gives the number of zeroes to be added at the last of answer
    go multiplySingleDigit(y[i], x, count, c)
    // each digit of y is multiplied to x concurrently where they are written to the channel c
  }

  for i := 0; i < n; i++ {
    temp := <- c
    str = add(str, temp)
  }
  return str
}

func evaluate(x string, a []string, size int) {
  var temp string
  temp = a[0]
  for i := 1; i <= size; i++ {
        temp = add(a[i], multiply(x, temp))
  }

  fmt.Println(temp)
}

func main() {
  var t, n, k int
  fmt.Scanln(&t)
  // t is the number of testcases

  for t > 0 {
    fmt.Scanln(&n)
    // n is the degree of polynomial

    var a []string
    a = make([]string, n + 1)
    for i := 0; i <= n; i++ {
      fmt.Scanln(&a[i])
    }
    // a[i] are the coefficients of polynomial

    fmt.Scanln(&k)
    // k is the number of points we want to evaluate polynimial

    var x []string
    x = make([]string, k)
    for i := 0; i < k; i++ {
      fmt.Scanln(&x[i])
      // x[i] are the points

      go evaluate(x[i], a[:], n)
    }
    t--
  }
  var input string
  fmt.Scanln(&input)
}
