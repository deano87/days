package main

/**
  '(' means up, while ')' menas down
  parse a string of brackets in order to get the final floor (starting from 0)
  The function will return the final floor and the first index in which the
  player reached the basement level (-1)
*/
func calculateFloor(input string) (int, int) {
  up, down, basementIndex := 0, 0, 0
  for _, r := range input {
      switch r {
        case '(': up++
        case ')': down++
      }

      /* check if the floor is -1 and that the basement index
      hasn't been updated yet */
      if -1 == up-down && 0 == basementIndex {
        basementIndex = up+down;
      }
  }
  return up-down, basementIndex;
}
