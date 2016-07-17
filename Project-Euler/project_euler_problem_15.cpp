/**
    @file project_euler_problem_15.cpp

    @author Terence Henriod

    Project Euler Problem 15

    @brief Solves the following problem for the general case of any number of
           size of square/rectangular grid:

  "Starting in the top left corner of a 2×2 grid, and only being able to move to
   the right and down, there are exactly 6 routes to the bottom right corner.

   How many such routes are there through a 20×20 grid?"

  Note: the problem is more of a combinatorics problem than a programming one.

    @version Original Code 1.00 (1/4/2014) - T. Henriod
*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <cmath>
#include <cstdio>
#include <iostream>
#include <iomanip>
#include <fstream>
using namespace std;


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   GLOBAL CONSTANTS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION PROTOTYPES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
/**
FunctionName

A short description

@param

@return

@pre
-# 

@post
-# 

@detail @bAlgorithm
-# 

@exception

@code
@endcode
*/
double nCr(const unsigned int n, const unsigned int r);


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main() {
  // variables
  unsigned int horizontal_size = 20;
  unsigned int vertical_size = 20;
  unsigned int total_moves_needed = 0;
  double number_of_solutions = 0;

/*
  // get the dimensions of the grid
  printf("Enter the horiontal dimension of the grid: ");
  scanf("%u", &horizontal_size);
  printf("Enter the vertical dimension of the grid: ");
  scanf("%u", &vertical_size);
*/

  // compute the total number of moves needed
  total_moves_needed = horizontal_size + vertical_size;

  // The number of strategies for moving from the upper left corner of the grid
  // to the lower left is like choosing which sequence members (1st move, 4th
  // move, etc...) must be downward moves. Simply use counting theory to compute
  // this: C( # total moves, # downward moves ). Choosing spots in the move
  // order for rightward moves will produce the same result due to the symmetry
  // of the combination function.
  number_of_solutions = nCr(total_moves_needed, vertical_size);

  // report the result to the user
  printf("\nThe number of solutions to traverse a %u x %u grid is: ",
         horizontal_size, vertical_size);
  cout << setprecision(15) << number_of_solutions << endl;

  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
double nCr(const unsigned int n, const unsigned int r) {
  // variables
  double result = 0;
  double numerator = 0;
  double denominator = 0;
  unsigned int stopping_point = n - r;
  unsigned int temp = 0;

  // case: r is larger than (n-r)
  if (r > stopping_point) {
    // compute the (n-r)! for the denominator value
    for (temp = 2, denominator = 1; temp <= stopping_point; temp++) {
      // include the term in the product
      denominator *= temp;
    }

    // change the stopping point
    stopping_point = r;
  }
  // case: (n-r) is larger than r
  else {
    // compute the r! for the denominator value
    for (temp = 2, denominator = 1; temp <= r; temp++) {
      // include the term in the product
      denominator *= temp;
    }
  }

  // compute the n! numerator part
  for (temp = stopping_point + 1, numerator = 1; temp <= n; temp++) {
    // include the term in the product
    numerator *= temp;
  }

  // compute the whole solution
  result = numerator / denominator;

  // return the resulting product
  return result;
}
