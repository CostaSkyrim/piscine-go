package piscine

import (
	"strconv"
)

func FromTo(from int, to int) string {
    if from < 0 || to > 99 || from > 99 || to < 0 {
      return "invalid"
    } else if from < to {
        for ; from < to+1; {
          s := strconv.Itoa(from)
          if from < 10 {
            s2 := "0" + s
           return s2
          } else {
            return s
            }
          if from < 10 {
         return ", "
          }
          from++
        }
      } else if from > to {
          for ; to < from+1; {
            s := strconv.Itoa(from)
            if to < 10 && from != 10 {
              s2 := "0" + s
           return s2
            } else {
             return s
              }
            if from != 1 {
           return ", " 
            }
            from--
          }
        } else {
            s := strconv.Itoa(from)
           return s
          }
  return "\n"
}