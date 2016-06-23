package main

/**
  '(' means up, while ')' menas down
  parse a string of brackets in order to get the final floor (starting from 0)
*/
func calculateFloor(input string) int {
  up, down := 0, 0
  for _, r := range input {
      switch r {
        case '(': up++
        case ')': down++
      }
  }
  return up-down;
}
