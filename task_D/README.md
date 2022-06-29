D. Подсистема регистрации

Реализуйте часть подсистемы валидации для функциональности регистрации.

Задана последовательность строк — это логины, которые указывали пользователи при регистрации. Для каждого логина вывести YES или NO в зависимости от того валидный это логин или нет.

Надо проверять, что логин — это последовательность из латинских букв в любом регистре, цифр и символов '_' или '-' (подчеркивание и дефис). Логин не должен начинаться с дефиса. Логин должен иметь длину от 2 до 24 символов. Логины, которые отличаются только регистром, считаются одинаковыми.

Если логин задан многократно, то только первое его вхождение проходит валидацию (зарегистрироваться повторно с таким же логином нельзя). Например, если сперва происходит регистрация с логином tester, а затем с логином TesteR, то вторая попытка будет неуспешной (надо вывести NO для неё).

Входные данные
В первой строке записано целое число t (1≤t≤10^4) — количество наборов входных данных в тесте.

Вторая строка содержит целое число n (1≤t≤2*10^5) — количество попыток регистрации в системе. Попытки заданы в хронологическом порядке (от самой ранней до самой поздней).

Далее идёт n строк, каждая содержит данные, которые указал пользователь в виде логина. Длина заданной строки может быть любой от 1 до 255 символов. В качестве символов могут быть любые символы с кодами от 33 (символ '!') до 126 (символ ' ') — такие символы не содержат национальные буквы различных языков и могут быть прочитаны как строка средствами стандартной библиотеки языка обычным способом.

Сумма значений n по всем наборам входных данных не превосходит 2*10^5. Сумма длин логинов по всем наборам входных данных не превосходит 10^6

Выходные данные
Выведите ответы на t наборов входных данных. Каждый ответ должен состоять из n строк, которые содержат YES или NO.

После каждого ответа на набор выходных данных выводите дополнительный перевод строки.

Вы можете выводить YES и NO в любом регистре (например, строки yEs, yes, Yes и YES будут распознаны как положительный ответ).

Пример входных данных
3
10
a
_tester
tester
tester1
~~~!!!
911
-test-me
test_me-93
?
abcdefghijlmnopqkrstuvwxyz
8
1
11
111
a
aa
aaa
_
__
5
one
One
onE
ONE
two

Выходные данные
NO
YES
YES
YES
NO
YES
NO
YES
NO
NO

NO
YES
YES
NO
YES
YES
NO
YES

YES
NO
NO
NO
YES