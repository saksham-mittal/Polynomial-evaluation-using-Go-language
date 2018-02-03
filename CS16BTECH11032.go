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

func isSmaller(x string, y string) bool {
    // Calculate lengths of both string
    n1 := len(x)
    n2 := len(y)

    if (n1 < n2) {
       return true
    }
    if (n2 < n1) {
       return false
    }

    for i := 0; i < n1; i++ {
       if (x[i] < y[i]) {
          return true
      } else if (x[i] > y[i]) {
          return false
      }
    }

    return false
}

func subtract(y string, x string) string {
  n1 := len(y)
  n2 := len(x)

  var str string = ""
  y = reverse(y)
  x = reverse(x)

  carry := 0
  for i := 0; i < n2; i++ {
    sub := (int(y[i] - '0') - int(x[i] - '0') - carry)
    if (sub < 0) {
      sub = sub + 10;
      carry = 1;
    } else {
      carry = 0;
    }

    str = str + string(sub + '0')
  }

  for i := n2; i < n1; i++ {
    sub := (int(y[i] - '0') - carry)
    carry = 0
    str = str + string(sub + '0')
  }

  str = reverse(str)

  return str
}

func add(x string, y string) string {
  var sign int = 0
  var str string = ""

  if(x == "") {
    if(y == "") {
      return str
    } else {
      return y
    }
  } else {
    if(y == "") {
      return x
    }
  }

  if(y[0] == '-') {
    y = y[1 : ]
    sign = 3
  }

  if(x[0] == '-') {
    x = x[1 : ]
    if(sign == 3) {
      sign = 2
    } else {
      sign = 1
    }
  }

  // sign = 0 - both are positive
  //        1 - only x is negative
  //        2 - both x and y is negative
  //        3 - only y is negative

  if(isSmaller(y, x)) {
    x, y = swap(x, y)

    if(sign == 1) {
      sign = 3
    } else if(sign == 3) {
      sign = 1
    }
  }

  if(sign == 0 || sign == 2) {
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

     if(sign == 2) {
       str = str + string('-')
     }

     str = reverse(str)
   } else {
     str = subtract(y, x);

     if(sign == 3) {
       str = reverse(str)
       str = str + string('-')
       str = reverse(str)
     }
   }
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
  var sign int

  if(y[0] == '-') {
    y = y[1 : ]
    // removing - sign if 1st character is '-'
    sign = 1
  }
  if(x[0] == '-') {
    x = x[1 : ]
    if(sign == 1) {
      sign = 0
    } else {
      sign = 1
    }
  }
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

  if(sign == 1) {
    // if the product is negative then add - sign at beginning
    str = reverse(str)
    str = str + string('-')
    str = reverse(str)
  }
  return str
}

func evaluate(x string, a []string, size int) {
  var temp string
  temp = a[0]
  // fmt.Println("temp =", temp)
  for i := 1; i <= size; i++ {
    // fmt.Println("x * temp =", multiply(x, temp))
    // fmt.Println("a[i] =", a[i])
    temp = add(a[i], multiply(x, temp))
    // fmt.Println("a[i] + (x * temp) =", temp)
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
