/**
    @file project_euler_problem_9.cpp

    @author Terence Henriod

    Project Euler Problem 9

    @brief Solves the following problem for the general case of any sum of a 
           Pythagorean Triplet:

  "A Pythagorean triplet is a set of three natural numbers, a < b < c,
   for which,

   a^2 + b^2 = c^2

   For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

   There exists exactly one Pythagorean triplet for which a + b + c = 1000.
   Find the product abc."

   This program implements a solution that will attempt to find a Pythagorean
   triple for a given sum using the fact that all primitive Pythagorean triples
   can be generated from the first primitive Pythgorean triple, (3, 4, 5).
   Using a breadth-first search technique, all new triples are generated from
   previously found triples and then tested for viability.

   The fact that all primitive triples can be generated from parent triples
   comes from the work of B. Berggren in "Pytagoreiska trianglar" (1934).
   Wikipedia provides an overview of the topic at:
   http://en.wikipedia.org/wiki/Pythagorean_triple#Parent.2Fchild_relationships

    @version Original Code 1.00 (12/31/2013) - T. Henriod
*/

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   HEADER FILES / NAMESPACES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
#include <cassert>
#include <cmath>
#include <cstdio>
#include <queue>
using namespace std;


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   GLOBAL CONSTANTS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
unsigned int kSolutionNotFound = -1;


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   USER DEFINED TYPES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
/**
@class PythagoreanTriple

A triple suitable for representing a Pythagorean triple. Also has functionality
for producing the sum and product of the three members of the triple. By
default, the triple is initialized with values that indicate an error (-1).
Only considers the positive versions of Pythagorean triples.

@var a  The first member of a Pythagorean triple.
@var b  The second member of a Pythagorean triple.
@var c  The third member of a Pythagorean triple.

@function PythagoreanTriple
  The default constructor for the struct. Initializes all members to the
  error indicating value, kSolutionNotFound (-1).

@function tripleSum
  Returns the sum of the members of the triple.

@function tripleProduct
  Returns the product of the members of the triple.
*/
class PythagoreanTriple {
 public:
  unsigned int a;
  unsigned int b;
  unsigned int c;
  PythagoreanTriple() : a(kSolutionNotFound), b(kSolutionNotFound),
                        c(kSolutionNotFound) {};

  unsigned int tripleSum() {
    // return the sum of each member
    return (a + b + c);
  };

  unsigned int tripleProduct() {
    // return the product of the triple
    return (a * b * c);
  };
};


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION PROTOTYPES
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
PythagoreanTriple& findPythagoreanTriple(const unsigned int target_sum);
PythagoreanTriple& createFirstChild(const PythagoreanTriple& parent_triple);
PythagoreanTriple& createSecondChild(const PythagoreanTriple& parent_triple);
PythagoreanTriple& createThirdChild(const PythagoreanTriple& parent_triple);


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   MAIN FUNCTION
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
/**
main

Implements the solution to Project Euler Problem 9 as described above.
*/

int main() {
  // variables
  unsigned int sum = 0;
  unsigned int product = 0;
  PythagoreanTriple solution;

  // collect the sought sum from the user
  printf("Enter the desired sum of a Pythagorean triple: ");
  scanf("%u", &sum);
  printf("\n");

  // attempt to find the Pythagorean triple with the given sum
  solution = findPythagoreanTriple(sum);

  // case: the solution doesn't exist
  if (solution.a == kSolutionNotFound) {
    // report that a solution does not exist
    printf("No Pythagorean triple that sums to %u exists...\n", sum);
  }
  // otherwise, find the product and display the solution
  else {
    // get the product of the triple
    product = solution.tripleProduct();

    // state the result for the user
    printf("The Pythagorean triple that sums to %u is: (%u, %u, %u)\n",
           sum, solution.a, solution.b, solution.c);
    printf("The product of the triple is: %u\n", product);
  }

  // return 0 on successful completion
  return 0;
}


/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                   FUNCTION IMPLEMENTATIONS
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
/**
findPythagoreanTriple

Uses a breadth-first search technique to find a Pythagorean triple that sums
to the given value. If the search fails, a triple with values equal to an
error value is returned.

@param target_sum   The sum that the solution triple should sum to.

@var solution             The solution triple to be returned when the function
                          finishes.
@var current_triple       A pointer to the current Pythagorean triple
                          currently being being tested or used to generate new
                          triples.
@var new_triple           A pointer to the newest Pythagorean triple
                          generated.
@var primitive_triples    A queue used to contain pointers to all Pythagorean
                          triples yet to be used in the search.
@var solution_not_found   A boolean indicator that the search should continue
                          when true.
@var test_sum             The sum of the Pythagorean triple being considered.
                          used in determining the solution triple.
@var factor               The multiplicative factor to be applied to a
                          primitive Pythagorean triple to produce a solution.

@return solution   A reference to a triple that will be a solution for the
                   given sum (if one exists). Otherwise, the triple will
                   contain error indicating values.

@pre
-# A positive target sum was provided.

@post
-# A PythagoreanTriple indicating either a solution or the absence thereof
   is returned.

@detail @bAlgorithm
-# The search begins with the first primitive Pythagorean triple, (3, 4, 5).
-# (Pointers to) triples are added to a queue as they are found.
-# Triples are extracted from the queue and tested to see if their sum is a
   factor of the target sum. If so, go on to step 4. Otherwise, go to step 6.
   If no triple can be extracted from the queue because it is empty, then no
   solution exists. Go to the final step.
-# The target sum is then divided by the triple's sum to find the
   multiplicative factor needed to produce the solution triple.
-# The found triple is then multiplied by the found factor to produce the
   solution. This solution is returned in step 7.
-# If the recently extracted triple's sum is less than the target sum,
   use Berggren's finding to produce three new primitive triple's. Add each
   of these to the queue. Start back at step 2. If the sum is greater than
   the target sum, then the triple is simply tossed out.
-# If a solution is found, the triple is returned. If the queue became empty,
   a triple indicating that no solution exists is returned.

@code
  PythagoreanTriple solution;
  unsigned int example_problem_sum = 1000;

  solution = findPythagoreanTriple(example_problem_sum);
  // solution now contains a Pythagorean triple that sums to 1000
@endcode
*/
PythagoreanTriple& findPythagoreanTriple(const unsigned int target_sum) {
  // variables
  PythagoreanTriple solution;
  PythagoreanTriple* current_triple = NULL;
  PythagoreanTriple* new_triple = NULL;
  queue<PythagoreanTriple*> primitive_triples;
  bool solution_not_found = true;
  unsigned int test_sum = 0;
  unsigned int factor = 0;

  // prime the search with the first Pythagorean triple
  new_triple = new PythagoreanTriple;
  new_triple->a = 3;
  new_triple->b = 4;
  new_triple->c = 5;
  primitive_triples.push(new_triple);

  // perform a breadth-first search for a solution
  while (solution_not_found && !primitive_triples.empty()) {
    // get the next element from the queue
    current_triple = primitive_triples.front();
    primitive_triples.pop();

    // get the triple's sum
    test_sum = current_triple->tripleSum();

    // case: the target is a multiple of the current item's sum
    if ((target_sum % test_sum) == 0) {
      // a solution has been found, indicate such
      solution_not_found = false;

      // find the complement factor between the test sum and the target
      factor = target_sum / test_sum;

      // generate the solution triple
      solution.a = factor * current_triple->a;
      solution.b = factor * current_triple->b;
      solution.c = factor * current_triple->c;
    }
    // case: the test sum is less than the target sum
    else if (test_sum < target_sum) {
      // continue the search by adding the three children generated
      // from the current primitive triple to the queue
      // generate the first child
      new_triple = new PythagoreanTriple;
      *new_triple = createFirstChild(*current_triple);
      primitive_triples.push(new_triple);

      // generate the second child
      new_triple = new PythagoreanTriple;
      *new_triple = createSecondChild(*current_triple);
      primitive_triples.push(new_triple);

      // generate the third child
      new_triple = new PythagoreanTriple;
      *new_triple = createThirdChild(*current_triple);
      primitive_triples.push(new_triple);
    }
    // otherwise, do nothing with the triple

    // discard the current triple after taking appropriate action
    delete current_triple;
  }
  // should a solution not be found, the solution contains this information by
  // default due to its constructor

  // return the dynamic memory used
  while (!primitive_triples.empty()) {
    delete primitive_triples.front();
    primitive_triples.pop();
  }

  // return the solution triple
  return solution;
}

/**
createFirstChild

Generates the first primitive Pythagorean triple that can be found from a
given Pythagorean triple. The only guarantee pertaining to the order of
the triple that is returned is that the 'c' member will be in the 3rd
position. Of course, if a valid parent is supplied, a valid child will be
produced.

@param parent_triple   A primitive Pythagorean triple.

@var first_child   The primitive Pythagorean triple to be
                   produced/returned.
@var a_term        The intermediate result pertaining to the 'a' term.
@var b_term        The intermediate result pertaining to the 'b' term.
@var c_term        The intermediate result pertaining to the 'c' term.

@return first_child   The first primitive Pythagorean triple that can be
                      created from a given triple.

@pre
-# For the function to produce a valid result, parent_triple shall be a
   primitive Pythagorean triple. 

@post
-# The triple returned will be a primitive Pythagorean triple.

@detail @bAlgorithm
-# The following linear transformation is applied to the vector of the
   parent triple's members is applied:
   [a b c][ 1  2  2] = [new_a new_b new_c]
          [-2 -1 -2]
          [ 2  2  3]
See: B. Berggren in "Pytagoreiska trianglar" (1934).

@code
  PythagoreanTriple parent;
    parent.a = 3;
    parent.b = 4;
    parent.c = 5;
  PythagoreanTriple newChild;

  newChild = createFirstChild(parent);
  // newChild now contains the triple (5, 12, 13).
@endcode
*/
PythagoreanTriple& createFirstChild(const PythagoreanTriple& parent_triple) {
  // variables
  PythagoreanTriple first_child;
  unsigned int a_term;
  unsigned int b_term;
  unsigned int c_term;

  // compute the child's 'a' ( = a - 2b + 2c)
  a_term = parent_triple.a;
  b_term = 2 * parent_triple.b;
  c_term = 2 * parent_triple.c;
  first_child.a = (a_term + c_term - b_term);

  // compute the child's 'b' ( = 2a - b + 2c)
  a_term = 2 * parent_triple.a;
  b_term = parent_triple.b;
  // the c term remains the same as in the previous step
  first_child.b = (a_term + c_term - b_term);

  // compute the child's 'c' ( = 2a - 2b + 3c)
  // the a term remains the same as in the previous step
  b_term = 2 * parent_triple.b;
  c_term = 3 * parent_triple.c;
  first_child.c = (a_term + c_term - b_term);

  // return the newly computed first child
  return first_child;
}

/**
createSecondChild

Generates the second primitive Pythagorean triple that can be found from
a given Pythagorean triple. The only guarantee pertaining to the order of
the triple that is returned is that the 'c' member will be in the 3rd
position. Of course, if a valid parent is supplied, a valid child will be
produced.

@param parent_triple   A primitive Pythagorean triple.

@var second_child   The primitive Pythagorean triple to be
                    produced/returned.
@var a_term         The intermediate result pertaining to the 'a' term.
@var b_term         The intermediate result pertaining to the 'b' term.
@var c_term         The intermediate result pertaining to the 'c' term.

@return second_child   The second primitive Pythagorean triple that can
                       be created from a given triple.

@pre
-# For the function to produce a valid result, parent_triple shall be a
   primitive Pythagorean triple. 

@post
-# The triple returned will be a primitive Pythagorean triple.

@detail @bAlgorithm
-# The following linear transformation is applied to the vector of the
   parent triple's members is applied:
   [a b c][ 1  2  2] = [new_a new_b new_c]
          [ 2  1  2]
          [ 2  2  3]
See: B. Berggren in "Pytagoreiska trianglar" (1934).

@code
  PythagoreanTriple parent;
    parent.a = 3;
    parent.b = 4;
    parent.c = 5;
  PythagoreanTriple newChild;

  newChild = createSecondChild(parent);
  // newChild now contains the triple (21, 20, 29).
@endcode
*/
PythagoreanTriple& createSecondChild(const PythagoreanTriple& parent_triple) {
  // variables
  PythagoreanTriple second_child;
  unsigned int a_term;
  unsigned int b_term;
  unsigned int c_term;

  // compute the child's 'a' ( = a + 2b + 2c)
  a_term = parent_triple.a;
  b_term = 2 * parent_triple.b;
  c_term = 2 * parent_triple.c;
  second_child.a = (a_term + b_term + c_term);

  // compute the child's 'b' ( = 2a + b + 2c)
  a_term = 2 * parent_triple.a;
  b_term = parent_triple.b;
  // the c term remains the same as in the previous step
  second_child.b = (a_term + b_term + c_term);

  // compute the child's 'c' ( = 2a + 2b + 3c)
  // the a term remains the same as in the previous step
  b_term = 2 * parent_triple.b;
  c_term = 3 * parent_triple.c;
  second_child.c = (a_term + b_term + c_term);

  // return the newly computed second child
  return second_child;
}

/**
createThirdChild

Generates the third primitive Pythagorean triple that can be found from a
given Pythagorean triple. The only guarantee pertaining to the order of
the triple that is returned is that the 'c' member will be in the 3rd
position. Of course, if a valid parent is supplied, a valid child will be
produced.

@param parent_triple   A primitive Pythagorean triple.

@var third_child   The primitive Pythagorean triple to be
                   produced/returned.
@var a_term        The intermediate result pertaining to the 'a' term.
@var b_term        The intermediate result pertaining to the 'b' term.
@var c_term        The intermediate result pertaining to the 'c' term.

@return third_child   The third primitive Pythagorean triple that can be
                      created from a given triple.

@pre
-# For the function to produce a valid result, parent_triple shall be a
   primitive Pythagorean triple. 

@post
-# The triple returned will be a primitive Pythagorean triple.

@detail @bAlgorithm
-# The following linear transformation is applied to the vector of the
   parent triple's members is applied:
   [a b c][-1  2  2] = [new_a new_b new_c]
          [-2  1  2]
          [-2  2  3]
See: B. Berggren in "Pytagoreiska trianglar" (1934).

@code
  PythagoreanTriple parent;
    parent.a = 3;
    parent.b = 4;
    parent.c = 5;
  PythagoreanTriple newChild;

  newChild = createThirdChild(parent);
  // newChild now contains the triple (15, 8, 17).
@endcode
*/
PythagoreanTriple& createThirdChild(const PythagoreanTriple& parent_triple) {
  // variables
  PythagoreanTriple third_child;
  unsigned int a_term;
  unsigned int b_term;
  unsigned int c_term;

  // compute the child's 'a' ( = -a + 2b + 2c)
  a_term = parent_triple.a;
  b_term = 2 * parent_triple.b;
  c_term = 2 * parent_triple.c;
  third_child.a = (b_term + c_term - a_term);

  // compute the child's 'b' ( = -2a + b + 2c)
  a_term = 2 * parent_triple.a;
  b_term = parent_triple.b;
  // the c term remains the same as in the previous step
  third_child.b = (b_term + c_term - a_term);

  // compute the child's 'c' ( = -2a + 2b + 3c)
  // the a term remains the same as in the previous step
  b_term = 2 * parent_triple.b;
  c_term = 3 * parent_triple.c;
  third_child.c = (b_term + c_term - a_term);

  // return the newly computed third child
  return third_child;
}
