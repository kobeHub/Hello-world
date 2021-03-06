package main


import (
  "fmt"
)

type student struct {
  firstName string
  lastName string
  grade string
  country string
}


// filter all the students which make the `f` return true
func filter(s []student, f func(student) bool) []student {
  var res []student
  for _, v := range s {
    if f(v) {
      res = append(res, v)
    }
  }
  return res
}

// impl iMap
func iMap(s []*student, f func(*student)) {
  for _, v := range s {
    f(v)
  }
}

func main() {
  s1 := student {
    "Nikofl",
    "Inno",
    "B",
    "Japan",
  }
  s2 := student{
    "James",
    "Leborn",
    "A",
    "America",
  }
  s3 := student {
    "Kiturl",
    "Deropmerl",
    "B",
    "Greek",
  }

  students := []student{s1, s2, s3}
  set := filter(students, func(s student) bool {
    return s.grade == "B"
  })
  fmt.Println("All the students whose grade is B\n", set)

  sets := []*student{&s1, &s2, &s3}
  fmt.Println("Now every students' grade should be upgrade")
  iMap(sets, func(s *student) {
    (*s).grade += "+"
  })

  for _, v := range sets {
    fmt.Println(*v)
  }
}
