Задача с тренировочного контеста Яндекса
Дан упорядоченный по неубыванию массив целых 32-разрядных чисел. Требуется удалить из него все повторения.

Желательно получить решение, которое не считывает входной файл целиком в память, т.е., использует лишь константный объем памяти в процессе работы.

Input format
Первая строка входного файла содержит единственное число n, n ≤ 1000000.

На следующих n строк расположены числа — элементы массива, по одному на строку. Числа отсортированы по неубыванию.

Output format
Выходной файл должен содержать следующие в порядке возрастания уникальные элементы входного массива.

Sample 1
Input 	 
5
2
4
8
8
8
Output
2
4
8

Sample 2
Input 	 
5
2
2
2
8
8

Output
2
8
