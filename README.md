# Polynomial-evaluation-using-Go-language
**Objective :** Write a program in Go to evaluate a polynomial over given integer points. You are supposed to use concurrency features of Go for this assignment.

**Input format:**</br>
numtestcases</br>
polydegree</br>
polycoeff</br>
numpoints</br>
Point1 Point2 …</br>

The first line of input is an integer denoting number of test cases on which the program will be run followed by test cases. In each test case, the first line is degree of the polynomial(N), the second line contains N+1 integers specifying the coefficients of the polynomial from higher degree to lower degree. The next line has the number of points(M) on which the polynomial is to be evaluated. The last line of the test case contains M integers on which the polynomial is to be evaluated.

**Example Input:**</br>
2</br>
2</br>
-3 1 0</br>
4</br>
34 400 23 -4</br>
5</br>
34 43 -32 21 0 -9</br>
3</br>
-23 0 4567</br>

Here the first polynomial is -3x^2+x and the second polynomial is 34x^5+43x^4-32x^3+21x^2+9.
The first polynomial is to be evaluated on 4 points (34, 400, 23, -4) and second polynomial is to be evaluated on 3 points (-23, 0, 4567).

**Output format:**</br>
\#</br>
Eval1</br>
Eval2</br>
...</br>
\#</br>
Eval3</br>
...</br>
\#</br>

The output of the program is one evaluation per line with evaluations for different polynomials separated by a line containing ‘#’ without any extra characters. Note that it is guaranteed that the evaluation of the polynomial will not contain more than 1000 digits.

**Example Output:(Given output is not for the input specified earlier.)**</br>
\#</br>
34113492349</br>
223498235923542349</br>
234235964564</br>
\#</br>
235235234353</br>
\#</br>

The given example output contains 3 evaluations on first polynomial and 1 evaluation on second polynomial.
