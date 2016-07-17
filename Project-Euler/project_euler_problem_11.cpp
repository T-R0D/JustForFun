/**
    @file project_euler_problem_11.cpp

    @author Terence Henriod

    Project Euler Problem 11

    @brief Solves the following problem for the general case of any square grid
           of two digit numbers (up to 1000 x 1000):

  "In the 20×20 grid below, four numbers along a diagonal line have been marked
   in red:

   08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
   49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
   81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
   52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
   22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
   24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
   32 98 81 28 64 23 67 10#26#38 40 67 59 54 70 66 18 38 64 70
   67 26 20 68 02 62 12 20 95#63#94 39 63 08 40 91 66 49 94 21
   24 55 58 05 66 73 99 26 97 17#78#78 96 83 14 88 34 89 63 72
   21 36 23 09 75 00 76 44 20 45 35#14#00 61 33 97 34 31 33 95
   78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
   16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
   86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
   19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
   04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
   88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
   04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
   20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
   20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
   01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48

   The product of these numbers is 26 × 63 × 78 × 14 = 1788696.

   What is the greatest product of four adjacent numbers in the same direction
   (up, down, left, right, or diagonally) in the 20×20 grid?"

    @version Original Code 1.00 (1/3/2014) - T. Henriod
*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <cmath>
#include <cstdio>
#include <iostream>
#include <fstream>
using namespace std;


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   GLOBAL CONSTANTS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
const int kFirstPrime = 2;

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
unsigned int findLargestProductInFile(const char* file_name,
                                      const unsigned int grid_size,
                                      const unsigned int num_terms);

unsigned int twoDigitCharToInt(const char ones_place, const char tens_place);

unsigned int bestDirectionalProduct(const unsigned int* data,
                                    const unsigned int ndx,
                                    const unsigned int grid_size,
                                    const unsigned int num_terms,
                                    const unsigned int current_largest );

unsigned int getUpwardProduct(const unsigned int* data,
                              const unsigned int data_size,
                              const unsigned int num_terms,
                              const unsigned int row_ndx,
                              const unsigned int column_ndx);

unsigned int getBackwardProduct(const unsigned int* data,
                                const unsigned int data_size,
                                const unsigned int num_terms,
                                const unsigned int row_ndx,
                                const unsigned int column_ndx);

unsigned int getDiagonalProduct(const unsigned int* data,
                                const unsigned int data_size,
                                const unsigned int num_terms,
                                const unsigned int row_ndx,
                                const unsigned int column_ndx);

unsigned int getOtherDiagonalProduct(const unsigned int* data,
                                const unsigned int data_size,
                                const unsigned int num_terms,
                                const unsigned int row_ndx,
                                const unsigned int column_ndx);


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
int main() {
  // variables
  unsigned int largest_product = 0;
  unsigned int grid_size = 20;
  unsigned int num_terms = 4;
  char file_name[40] = "project_euler_problem_11.txt";

  // get the respective data
/*
  printf("Enter the file name: ");
  scanf("%s", file_name);
  printf("Enter the size of the data square: ");
  scanf("%u", &grid_size);
  printf("Enter the number of terms to include in the product: ");
  scanf("%u", &num_terms);
*/

  // get the largest product from the file
  largest_product = findLargestProductInFile(file_name, grid_size, num_terms);

  // report the result to the user
  printf("\nThe largest product in %s is: %8u\n", file_name, largest_product);

  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
unsigned int findLargestProductInFile(const char* file_name,
                                      const unsigned int grid_size,
                                      const unsigned int num_terms) {
  // variables
  unsigned int largest_product = 0;
  unsigned int data_count = (grid_size * grid_size);
  unsigned int data_ndx = 0;
  unsigned int* data = NULL;
  char ones_digit = '0';
  char tens_digit = '0';
  fstream data_file;

  // get dynamic memory
  data = new unsigned int [data_count];

  // open the file
  data_file.clear();
  data_file.open(file_name);
  
  // case: the file opened
  if (data_file.good()) {
    // prime the data reading loop
    data_file >> tens_digit >> ones_digit;

    // read and process data and find product until there is none left to get
    while (data_file.good() && (data_ndx < data_count)) {
      // add the number data to the array
      data[data_ndx] = twoDigitCharToInt(ones_digit, tens_digit);
      
      // find all possible products, keep the largest one
      largest_product = bestDirectionalProduct(data, data_ndx, grid_size,
                                               num_terms, largest_product);

      // count the batch of read data
      data_ndx++;

      // read the next batch of data
      data_file >> tens_digit >> ones_digit;
    }
  }
  // return the dynamic memory
  delete [] data;

  // close the file
  data_file.close();

  // return the found product
  return largest_product;
}

unsigned int twoDigitCharToInt(const char ones_place, const char tens_place) {
  // variables
  int result = 0;

  // add the value of the one's place character to the number
  result += (ones_place - '0');

  // add the value of the ten's place character to the number
  result += (tens_place - '0') * 10;

  // return the result
  return result;
}

unsigned int bestDirectionalProduct(const unsigned int* data,
                                    const unsigned int ndx,
                                    const unsigned int grid_size,
                                    const unsigned int num_terms,
                                    const unsigned int current_largest ) {
  // variables
  unsigned int largest_product = current_largest;
  unsigned int upward_product = 0;
  unsigned int backward_product = 0;
  unsigned int diagonal_product = 0;
  unsigned int other_diagonal_product = 0;
  unsigned int row_ndx = 0;
  unsigned int column_ndx = 0;

  // compute the two dimensional indices
  row_ndx = ndx / grid_size;
  column_ndx = ndx % grid_size;

  // get the upwards product
  upward_product = getUpwardProduct(data, grid_size, num_terms, row_ndx,
                                     column_ndx);

  // case: the upwards product is now the best
  if (upward_product > largest_product) {
    // keep the upwards product
    largest_product = upward_product;

cout << row_ndx << ' ' << column_ndx << ' ' << largest_product << endl;

  }

  // get the backwards product
  backward_product = getBackwardProduct(data, grid_size, num_terms, row_ndx,
                                     column_ndx);

  // case: the backwards product is now the best
  if (backward_product > largest_product) {
    // keep the backwards product
    largest_product = backward_product;

cout << row_ndx << ' ' << column_ndx << ' ' << largest_product << endl;

  }

  // get the diagonal product
  diagonal_product = getDiagonalProduct(data, grid_size, num_terms, row_ndx,
                                        column_ndx);

if (ndx == ((9 * 20) + 11)) {
  cout << diagonal_product << endl;
}

  // case: the diagonal product is now the best
  if (diagonal_product > largest_product) {
    // keep the diagonal product
    largest_product = diagonal_product;

cout << row_ndx << ' ' << column_ndx << ' ' << largest_product << endl;

  }

  // get the other diagonal product
  other_diagonal_product = getOtherDiagonalProduct(data, grid_size, num_terms,
                                                   row_ndx, column_ndx);

if (ndx == ((9 * 20) + 11)) {
  cout << other_diagonal_product << endl;
}

  // case: the diagonal product is now the best
  if (other_diagonal_product > largest_product) {
    // keep the diagonal product
    largest_product = other_diagonal_product;

cout << row_ndx << ' ' << column_ndx << ' ' << largest_product << endl;

  }

  // return the largest product
  return largest_product;
}

unsigned int getUpwardProduct(const unsigned int* data,
                              const unsigned int data_size,
                              const unsigned int num_terms,
                              const unsigned int row_ndx,
                              const unsigned int column_ndx) {
  // variables
  unsigned int product = 1;
  int temp_row = 0;
  int row_stop = (row_ndx - num_terms + 1);

  // case: the supplied indices pass the boundary check
  if (row_stop >= 0) {
    // multiply the numbers in the row
    for (temp_row = row_ndx;
         temp_row >= row_stop;
         temp_row--) {
      // include the number in the product
      product *= data[(temp_row * data_size) + column_ndx];
    }
  }
  // case: the indices do not pass the boundary check
  else {
    // set the result to 0
    product = 0;
  }

  // return the resulting product
  return product;
}

unsigned int getBackwardProduct(const unsigned int* data,
                                const unsigned int data_size,
                                const unsigned int num_terms,
                                const unsigned int row_ndx,
                                const unsigned int column_ndx) {
  // variables
  unsigned int product = 1;
  int temp_column = column_ndx;
  int column_stop = (column_ndx - num_terms + 1);

  // case: the supplied indices pass the boundary check
  if (column_stop >= 0) {
    // multiply the numbers in the row
    while (temp_column >= column_stop) {
      // include the number in the product
      product *= data[(row_ndx * data_size) + temp_column];

      // move the cursor
      temp_column--;
    }
  }
  // case: the indices do not pass the boundary check
  else {
    // set the result to 0
    product = 0;
  }

  // return the resulting product
  return product;
}

unsigned int getDiagonalProduct(const unsigned int* data,
                                const unsigned int data_size,
                                const unsigned int num_terms,
                                const unsigned int row_ndx,
                                const unsigned int column_ndx) {
  // variables
  unsigned int product = 1;
  int temp_row = 0;
  int temp_column = 0;
  int row_stop = (row_ndx - num_terms + 1);
  int column_stop = (column_ndx - num_terms + 1);

  // case: the supplied indices pass the boundary check
  if ((row_stop >= 0) &&
      (column_stop >= 0)) {
    // multiply the numbers in the row
    for (temp_row = row_ndx, temp_column = column_ndx;
         temp_column >= column_stop;
         temp_row--, temp_column--) {
      // include the number in the product
      product *= data[(temp_row * data_size) + temp_column];
    }
  }
  // case: the indices do not pass the boundary check
  else {
    // set the result to 0
    product = 0;
  }

  // return the resulting product
  return product;
}

unsigned int getOtherDiagonalProduct(const unsigned int* data,
                                const unsigned int data_size,
                                const unsigned int num_terms,
                                const unsigned int row_ndx,
                                const unsigned int column_ndx) {
  // variables
  unsigned int product = 1;
  int temp_row = 0;
  int temp_column = 0;
  int row_stop = (row_ndx - num_terms + 1);
  int column_stop = (column_ndx + num_terms);

  // case: the supplied indices pass the boundary check
  if ((row_stop >= 0) &&
      (column_stop >= 0)) {
    // multiply the numbers in the row
    for (temp_row = row_ndx, temp_column = column_ndx;
         temp_column < column_stop;
         temp_row--, temp_column++) {
      // include the number in the product
      product *= data[(temp_row * data_size) + temp_column];
    }
  }
  // case: the indices do not pass the boundary check
  else {
    // set the result to 0
    product = 0;
  }

  // return the resulting product
  return product;
}
