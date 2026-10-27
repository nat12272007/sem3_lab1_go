Лабораторная работа № 1. Абстрактные структуры данных

выполнила Сергеева Наталья Александрровна АБс-524 АВТФ

взаимодействие со структурами данных (интерфейс) прописанно в main.cpp

команда для запуска программы: go run main.go --repl
для выхода нужно ввести EXIT

список доступных команд:
1) Массив (префикс M, имя по умолчанию myarr).
	MPUSH myarr <значение> — добавить элемент в конец.
	MINSERT myarr <индекс> <значение> — вставить элемент по индексу.
	MGET myarr <индекс> — получить элемент по индексу.
	MSET myarr <индекс> <значение> — заменить элемент по индексу.
	MDEL myarr <индекс> — удалить элемент по индексу.
	MLEN myarr — вывести длину массива.
2) Односвязный список (префикс F, имя по умолчанию mylist).
	FADDHEAD mylist <значение> — добавить в голову.
	FADDTAIL mylist <значение> — добавить в хвост.
	FADDBEFORE mylist <target> <значение> — добавить перед элементом target.
	FADDAFTER mylist <target> <значение> — добавить после элемента target.
	FDELHEAD mylist — удалить из головы.
	FDELTAIL mylist — удалить из хвоста.
	FDELBEFORE mylist <target> — удалить перед элементом target.
	FDELAFTER mylist <target> — удалить после элемента target.
	FDELVAL mylist <значение> — удалить по значению.
	FFIND mylist <значение> — поиск элемента (TRUE/FALSE).
	FGET mylist <индекс> — чтение по индексу.
3) Двусвязный список (префикс L, имя по умолчанию mydlist).
	LADDHEAD mydlist <значение> — добавить в голову.
	LADDTAIL mydlist <значение> — добавить в хвост.
	LADDBEFORE mydlist <target> <значение> — добавить перед элементом target.
	LADDAFTER mydlist <target> <значение> — добавить после элемента target.
	LDELHEAD mydlist — удалить из головы.
	LDELTAIL mydlist — удалить из хвоста.
	LDELBEFORE mydlist <target> — удалить перед элементом target.
	LDELAFTER mydlist <target> — удалить после элемента target.
	LDELVAL mydlist <значение> — удалить по значению.
	LFIND mydlist <значение> — поиск по значению (TRUE/FALSE).
	LGET mydlist <индекс> — чтение по индексу.
	LPRINTBACK mydlist — вывод списка в обратном порядке.
4) Стек (префикс S, имя по умолчанию mystack).
	SPUSH mystack <значение> — добавить элемент (push).
	SPOP mystack — извлечь элемент (pop) с выводом на экран.
	SPEEK mystack — прочитать вершину без удаления.
5) Очередь (префикс Q, имя по умолчанию myqueue).
	QPUSH myqueue <значение> — добавить элемент в очередь (push).
	QPOP myqueue — извлечь элемент (pop) с выводом на экран.
	QPEEK myqueue — прочитать начало очереди без удаления.
