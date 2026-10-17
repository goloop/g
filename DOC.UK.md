# g — довідник

Повний довідник пакета `g`: ментальна модель, усі функції за темами, обмеження
типів і набір практичних рецептів.

Англійська версія: **[DOC.md](DOC.md)**.

## Зміст

- [Ментальна модель](#ментальна-модель)
- [Обмеження типів](#обмеження-типів)
- [Логіка та умови](#логіка-та-умови)
- [Перевірки наявності й значень](#перевірки-наявності-й-значень)
- [Вказівники](#вказівники)
- [Числа й математика](#числа-й-математика)
- [Статистика й випадковість](#статистика-й-випадковість)
- [Колекції](#колекції)
- [Функціональні помічники](#функціональні-помічники)
- [Операції над множинами](#операції-над-множинами)
- [Рядки й парсинг](#рядки-й-парсинг)
- [Дата й час](#дата-й-час)
- [Пошук і ранжування](#пошук-і-ранжування)
- [Діапазони](#діапазони)
- [Помилки](#помилки)
- [Рецепти й поради](#рецепти-й-поради)

## Ментальна модель

`g` — це плаский набір маленьких узагальнених (generic) помічників. Немає стану,
немає конфігурації й немає прихованої конкурентності — кожна функція є чистою,
ощадливою до алокацій операцією над своїми аргументами.

Через увесь пакет проходять три ідеї:

1. **Узагальнення замість дублювання.** Один `Map`, один `Sort`, один `Clamp`,
   що працюють для будь-якого типу елементів, обмежені кількома інтерфейсами
   (`Numerable`, `Verifiable`, `Ordered`).
2. **Делегування стандартній бібліотеці на гарячих шляхах.** Сортування,
   min/max, перевірка входження та випадковість спираються на `slices`, `cmp`
   і `math/rand/v2`, тож ви безкоштовно отримуєте перевірену поведінку
   `O(n log n)` і потокобезпечну випадковість.
3. **Детерміновані крайні випадки.** Порожній ввід, `NaN`, переповнення й
   вироджені діапазони повертають визначене значення, а не панікують. Там, де
   результат може бути відсутній, варіант `…Ok` повідомляє про це булом.

Імпортуйте під коротким іменем:

```go
import g "github.com/goloop/g/v2"
```

## Обмеження типів

Помічники обмежені невеликим набором інтерфейсів:

| Обмеження | Дозволяє |
|-----------|----------|
| `Integer`    | усі знакові й беззнакові цілі типи |
| `Float`      | `float32`, `float64` |
| `Numerable`  | `Integer` ∪ `Float` — усе, над чим можлива арифметика |
| `Ordered`    | `Integer` ∪ `Float` ∪ `~string` — усе з `< <= >= >` |
| `Verifiable` | `Integer` ∪ `Float` ∪ `string` ∪ `rune` — порівнювані класичні типи |
| `comparable` | вбудоване обмеження Go, використовується помічниками множин |

`Pair[T, U]` — це узагальнена структура (`First`, `Second`), тип елемента `Zip`.

## Логіка та умови

```go
func If[L bool | trit.Trit, T any](e L, t, f T, u ...T) T
func IfFunc[T any](cond bool, t, f func() T) T
```

`If` — це тернарний оператор, якого немає в Go. Умова може бути `bool` або
`trit.Trit`; з `Trit` можна передати необов'язкове третє значення `u` для стану
`Unknown`.

```go
max := g.If(a > b, a, b)
label := g.If(t, "yes", "no", "maybe") // t — це trit.Trit
```

`If` **жадібний** — обидві гілки `t` і `f` обчислюються до виклику. Коли гілка
дорога або небезпечна для безумовного обчислення, використовуйте `IfFunc`, що
виконує лише обрану гілку:

```go
cfg := g.IfFunc(hasFile, func() Config { return load() }, func() Config { return Default() })
```

## Перевірки наявності й значень

```go
func All[T any](v ...T) bool          func AllList[T any](v []T) bool
func Any[T any](v ...T) bool          func AnyList[T any](v []T) bool
func Contains[T Verifiable](v T, vs []T) bool
func In[T Verifiable](v T, list ...T) bool
func Index[T comparable](vs []T, v T) int
func IsEmpty[T any](v T) bool
func IsTrue[T any](v T) bool           func IsFalse[T any](v T) bool
func Value[T any](v ...T) T
```

`All`/`Any` повідомляють, чи кожен / чи хоч один аргумент є ненульовим
(«істинним») значенням. Варіанти `…List` беруть зріз замість варіативного списку.

`In` і `Contains` — тести на входження (`In` варіативний, `Contains` бере зріз).
`Index` повертає позицію першого збігу або `-1`.

`IsEmpty` істинний для нульового значення будь-якого типу. `IsTrue`/`IsFalse` —
одноаргументні форми. `Value` повертає перший аргумент або нульове значення,
якщо викликано без аргументів, — зручно для «необов'язкових варіативних» API.

```go
if g.In(status, 200, 201, 204) { /* ... */ }
name := g.Value(names...) // "", якщо names порожній
```

## Вказівники

```go
func Ptr[T any](v ...T) *T
func PtrIf[T any](exp bool, v ...T) *T
```

`Ptr` повертає вказівник на свій аргумент (або на нульове значення без
аргументів) — корисно для побудови структур із полями-вказівниками. `PtrIf`
повертає вказівник, лише коли `exp` істинний, інакше `nil`.

```go
req := Request{Timeout: g.Ptr(30)}       // *int
req.Retry = g.PtrIf(enabled, 3)          // *int або nil
```

## Числа й математика

```go
func Abs[T Numerable](v T) T
func Clamp[T Ordered](v, lo, hi T) T
func Sum[T Numerable](v ...T) T
func SafeSum[T Numerable](v ...T) (T, error)
func Product[T Numerable](v ...T) T
func Min[T Verifiable](v ...T) T        func Max[T Verifiable](v ...T) T
func MinOk[T Verifiable](v ...T) (T, bool)
func MaxOk[T Verifiable](v ...T) (T, bool)
func MinList[T Verifiable](v []T, defaults ...T) T
func MaxList[T Verifiable](v []T, defaults ...T) T
func MinListOk[T Verifiable](v []T, defaults ...T) (T, bool)
func MaxListOk[T Verifiable](v []T, defaults ...T) (T, bool)
func IsEven[T Numerable](v T, f ...bool) bool
func IsOdd[T Numerable](v T, f ...bool) bool
func IsWhole[T Numerable](v T) bool
```

`Sum` переповнюється як оператор `+`. Коли потрібно **виявити** переповнення,
використовуйте `SafeSum`, що повертає помилку з обгорнутим
[`ErrOverflow`](#помилки).

```go
total := g.Sum(1, 2, 3)                  // 6
_, err := g.SafeSum(math.MaxInt64, 1)    // errors.Is(err, g.ErrOverflow)
```

`Clamp` обмежує значення діапазоном `[lo, hi]`:

```go
g.Clamp(15, 0, 10) // 10
g.Clamp(-3, 0, 10) // 0
```

`Min`/`Max` делегують стандартній бібліотеці й детерміновано поширюють `NaN`.
На порожньому вводі повертають нульове значення; варіанти `…Ok` повідомляють, чи
було хоч одне значення. Форми `…List` беруть зріз плюс необов'язкові запасні
значення, що використовуються, коли зріз порожній.

`IsEven`/`IsOdd` обчислюють парність у `float64`, тож лишаються коректними для
значень за межами діапазону `int` (напр. `1e19`). `IsWhole` повідомляє, чи не
має float дробової частини.

## Статистика й випадковість

```go
func Average[T Numerable](v ...T) float64
func Median[T Numerable](v ...T) float64
func Random[T Numerable](v ...T) T
func RandomList[T any](v []T) T
func RandomListPlural[T any](n int, v []T) []T
func RandomMap[K comparable, T any](m map[K]T) T
func RandomMapPlural[K comparable, T any](n int, m map[K]T) []T
```

`Average` і `Median` повертають `float64` для будь-якого числового вводу.
`Median` для парної довжини усереднює два серединні значення.

`Random` ніколи не панікує на виродженому діапазоні: `Random()` повертає нульове
значення, `Random(n)` — значення в `[0, n)`, а `Random(a, b)` — у `[a, b)`;
недодатний чи інвертований діапазон дає нижню межу. `RandomList` / `RandomMap`
обирають один елемент; варіанти `…Plural` обирають `n`. Випадковість спирається
на потокобезпечне джерело `math/rand/v2`, тож ці функції безпечні для виклику з
кількох горутин.

```go
n := g.Random(1, 10)              // int у [1, 10)
pick := g.RandomList(options)     // один елемент
three := g.RandomListPlural(3, options)
```

## Колекції

```go
func Sort[T Verifiable](v []T, inverse ...bool)
func Reverse[T any](v []T)
func Shuffle[T any](v []T)
func Distinct[T comparable](v []T) []T
func Merge[T Verifiable](a []T, b []T, sort ...bool) []T
func Zip[T, U any](a []T, b []U) []Pair[T, U]
func CartesianProduct[T any](a []T, b []T) [][2]T
```

`Sort` змінює зріз на місці за допомогою introsort зі стандартної бібліотеки
(гарантовано `O(n log n)`); передайте `true` для спадного порядку. `Reverse` і
`Shuffle` також змінюють на місці. `Distinct` повертає новий зріз без дублікатів,
зберігаючи порядок першої появи.

`Merge` конкатенує два зрізи (за бажанням сортуючи результат). `Zip` попарно
поєднує два зрізи в `[]Pair`; `CartesianProduct` дає кожну комбінацію `[a, b]`.

```go
g.Sort(nums)                 // за зростанням, на місці
g.Sort(nums, true)           // за спаданням
u := g.Distinct([]int{1,1,2}) // [1 2]
pairs := g.Zip(keys, values) // []Pair[K, V]
```

## Функціональні помічники

```go
func Map[T any, U any](vs []T, f func(T) U) []U
func Filter[T any](vs []T, f func(T) bool) []T
func Reduce[T any, U any](vs []T, f func(U, T) U, init U) U
```

Класична трійця. `Map` перетворює кожен елемент (тип результату може відрізнятись),
`Filter` лишає елементи, що відповідають предикату, `Reduce` згортає зріз в один
акумулятор.

```go
doubled := g.Map(nums, func(n int) int { return n * 2 })
evens := g.Filter(nums, func(n int) bool { return n%2 == 0 })
total := g.Reduce(nums, func(acc, n int) int { return acc + n }, 0)
```

## Операції над множинами

```go
func Union[T comparable](a []T, b []T) []T
func Intersection[T comparable](a []T, b []T) []T
func Difference[T comparable](a []T, b []T) []T        // псевдонім: Diff
func SymmetricDifference[T comparable](a []T, b []T) []T // псевдонім: Sdiff
func Complement[T comparable](a []T, b []T) []T
```

Алгебра множин на зрізах. `Union` повертає унікальні елементи обох зрізів;
`Intersection` — елементи в обох; `Difference` (`Diff`) — елементи `a`, яких
немає в `b`; `SymmetricDifference` (`Sdiff`) — елементи рівно в одному;
`Complement` — елементи `b`, відсутні в `a`.

```go
g.Union([]int{1,2}, []int{2,3})        // [1 2 3]
g.Intersection([]int{1,2}, []int{2,3}) // [2]
g.Difference([]int{1,2}, []int{2,3})   // [1]
```

## Рядки й парсинг

```go
func StringToBool(v string, def ...bool) (bool, error)
func StringToInt(v string, def ...int) (int, error)
func StringToFloat(v string, def ...float64) (float64, error)
func BoolToString(b bool) string
func IntToString[T Integer](v T) string
func FloatToString[T ~float32 | ~float64](v T) string
func Weed(s string, patterns ...string) string
func Preserve(s string, patterns ...string) string
func Trim(s string, patterns ...string) string
```

Парсери `StringToX` приймають необов'язковий дефолт, що повертається при невдачі
парсингу (помилку вони все одно повертають, тож ви вирішуєте, що робити).
Форматери `XToString` — обернені до них.

`Weed` вилучає символи зі списку `patterns`, `Preserve` лишає лише ці символи, а
`Trim` зрізає їх з обох кінців. Є готові константи-набори: `Numbers`, `Letters`,
`Symbols`, `Whitespaces`. Без заданого набору ціллю за замовчуванням є пробіли.

```go
n, _ := g.StringToInt("123")             // 123
n, _ = g.StringToInt("x", 42)            // 42 (+ помилка)
digits := g.Preserve("+1-234-567", g.Numbers) // "1234567"
clean := g.Weed("Hello\t World\n")       // "Hello World"
```

## Дата й час

```go
func StringToDate(s string, patterns ...string) (time.Time, error)
func DateToString(t time.Time, patterns ...string) (string, error)
func DateToStrings(t time.Time, patterns ...string) ([]string, error)
func ChangeTimeZone(t time.Time, timezone string) (time.Time, error)
func SetTimeZone(t time.Time, timezone string) (time.Time, error)
func MoveTimeZone(t time.Time, tz int) time.Time
```

`StringToDate` пробує список макетів (або вбудований набір поширених форматів,
якщо жодного не задано) і повертає помилку з обгорнутим
[`ErrParseDate`](#помилки), коли нічого не збіглося. `DateToString` форматує час
першим придатним макетом; `DateToStrings` повертає по рядку на макет.

Помічники часових поясів різняться за задумом:

- `ChangeTimeZone` — перевести момент в інший пояс (настінний час змінюється,
  момент той самий).
- `SetTimeZone` — переінтерпретувати те саме показання настінного часу в іншому
  поясі (момент змінюється).
- `MoveTimeZone` — зсунути на фіксоване зміщення в годинах.

```go
d, _ := g.StringToDate("2023-12-01")
s, _ := g.DateToString(d, "2006-01-02")
ny, _ := g.ChangeTimeZone(time.Now(), "America/New_York")
```

## Пошук і ранжування

```go
func HLookup[T comparable, U any](v T, lookup []T, result []U, def U) U
func VLookup[T comparable, U any](v T, lookup []T, result []U, def U) U
func Rank[T Verifiable](number T, array []T, ascending ...bool) int
```

Помічники у стилі таблиць. `HLookup`/`VLookup` знаходять `v` у `lookup` і
повертають значення з тим самим індексом у `result` або `def`, якщо збігу немає.
`Rank` повертає 1-базований ранг `number` у `array` (за замовчуванням за
спаданням; передайте `true` для зростання).

```go
label := g.VLookup(code, codes, labels, "unknown")
r := g.Rank(7, []float64{1, 5, 2, 3, 7, 8}) // 2 (7 — друге за величиною)
```

## Діапазони

```go
func Range(a int, opt ...int) []int
func Rangef[T any](fn func(int) T, a int, opt ...int) []T
```

`Range` будує цілочисловий зріз за конвенцією Python `range`: `Range(n)` — це
`[0, n)`, `Range(a, b)` — `[a, b)`, а `Range(a, b, step)` додає крок. `Rangef`
відображає функцію на ті самі індекси, даючи `[]T`. Результат обмежений
`MaxRangeSize` як запобіжник від випадкових величезних алокацій.

```go
g.Range(5)          // [0 1 2 3 4]
g.Range(2, 8, 2)    // [2 4 6]
g.Rangef(func(i int) string { return strconv.Itoa(i) }, 3) // ["0" "1" "2"]
```

## Помилки

```go
var ErrOverflow  = errors.New("g: numeric overflow")
var ErrParseDate = errors.New("g: date and time could not be recognized")
```

Обидва — сентинели, що повертаються **обгорнутими**, тож звіряйте їх через
`errors.Is`:

```go
if _, err := g.SafeSum(a, b); errors.Is(err, g.ErrOverflow) { /* ... */ }
if _, err := g.StringToDate(s); errors.Is(err, g.ErrParseDate) { /* ... */ }
```

## Рецепти й поради

**Дефолт для можливо-порожнього опціонального.** `Value` згортає варіативний
аргумент до одного значення з нульовим запасним:

```go
func greet(name ...string) string {
    return "Hi, " + g.If(g.IsEmpty(g.Value(name...)), "friend", g.Value(name...))
}
```

**Уникнути дорогого дефолту.** Візьміть `IfFunc`/`PtrIf`, щоб невикористана
гілка не виконувалась:

```go
conn := g.IfFunc(cached, func() *Conn { return pool.Get() }, func() *Conn { return dial() })
```

**Безпечне до переповнення підсумовування.** Надавайте перевагу `SafeSum` у коді,
що додає числа ззовні; лишайте `Sum` для завідомо малих входів, де переповнення
прийнятне.

**Детерміноване семплування в тестах.** `Random*` спираються на `math/rand/v2`;
для відтворюваних тестів семплюйте з фіксованого зрізу через `RandomListPlural`
і перевіряйте входження, а не конкретне значення.

**Очищення користувацького вводу.** Поєднуйте `Weed`/`Preserve`/`Trim` з
константами-наборами, щоб нормалізувати телефони, коди й ідентифікатори без
регулярного виразу.

## Реліз v2.3.1

SafeSum виявляє overflow та NaN також для іменованих float-типів. Ptr[any](nil) повертає вказівник на nil-інтерфейс. Intersection зберігає порядок і обходиться однією таблицею пошуку.
