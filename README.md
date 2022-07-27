# golang-learning

[Learn go with tests](https://quii.gitbook.io/learn-go-with-tests/) exercises repository template.

# About me

Student of mathematical and information direction.

# Learned lessons

### github.com/LightBulbfromSpace/kld-go-learning/arrays_and_slices/sum - 100.0%

<details>
  <summary><code>func Sum(numbers []int) int</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/arrays_and_slices/sumAll - 100.0%

<details>
  <summary><code>func SumAll(nums ...[]int) (sums []int)</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/arrays_and_slices/sumTails - 100.0%

<details>
  <summary><code>func SumTails(numsToSum ...[]int) []int</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/concurrency/faster_version_v2 - 100.0%

<details>
  <summary><code>func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool</code></summary>

    TYPES
    type WebsiteChecker func(string) bool
</details>

### github.com/LightBulbfromSpace/kld-go-learning/concurrency/slow_version_v1 - 100.0%

<details>
  <summary><code>func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool</code></summary>

    TYPES
    type WebsiteChecker func(string) bool
</details>

### github.com/LightBulbfromSpace/kld-go-learning/context/v1 - 100.0%

<details>
  <summary><code>func Server(store Store) http.HandlerFunc</code></summary>

    TYPES
    type Store interface {
    Fetch() string
    Cancel()
    }
    type StubStore struct {
    // Has unexported fields.
    }
</details>

<details>
  <summary><code>func (s *StubStore) Cancel()</code></summary>

</details>

<details>
  <summary><code>func (s *StubStore) Fetch() string</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/context/v2 - 100.0%

<details>
  <summary><code>func Server(store Store) http.HandlerFunc</code></summary>

    TYPES
    type Store interface {
    Fetch(ctx context.Context) (string, error)
    }
</details>

### github.com/LightBulbfromSpace/kld-go-learning/dependency_injection/greet - 50.0%

<details>
  <summary><code>func Greet(writer io.Writer, name string)</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/hello - 92.9%

<details>
  <summary><code>func DefaultReceiver(language string) (receiver string)</code></summary>

</details>

<details>
  <summary><code>func GreetingPrefix(language string) (prefix string)</code></summary>

</details>

<details>
  <summary><code>func Hello(name string, language string) string</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/hellogo - 100.0%
Package hellogo contains first steps in language.
<details>
  <summary><code>func Hello() string</code></summary>

    Hello is first function.
</details>

### github.com/LightBulbfromSpace/kld-go-learning/integers - 100.0%

<details>
  <summary><code>func Add(a, b int) int</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/intro_to_propely_based_tests/v1 - 100.0%

<details>
  <summary><code>func ConvertToRoman(arabic int) string</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/intro_to_propely_based_tests/v2 - 93.3%

<details>
  <summary><code>func ConvertToArabic(roman string) int</code></summary>

</details>

<details>
  <summary><code>func ConvertToRoman(arabic int) string</code></summary>

    TYPES
    type RomanNum struct {
    // Has unexported fields.
    }
</details>

### github.com/LightBulbfromSpace/kld-go-learning/intro_to_propely_based_tests/v3_to_arabic_new - 96.6%
VARIABLES

var ErrArabicOverflow = errors.New("unacceptable number (more than 3999)")
<details>
  <summary><code>func ConvertToArabic(roman string) (result uint16)</code></summary>

</details>

<details>
  <summary><code>func ConvertToRoman(arabic uint16) (string, error)</code></summary>

    TYPES
    type RomanNumeral struct {
    // Has unexported fields.
    }
    type RomanNumerals []RomanNumeral
</details>

<details>
  <summary><code>func (r RomanNumerals) Exists(symbols ...byte) bool</code></summary>

</details>

<details>
  <summary><code>func (r RomanNumerals) ValueOf(symbols ...byte) uint16</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/iteration - 100.0%

<details>
  <summary><code>func Repeat(str string, num int) string</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/maps - 83.3%
CONSTANTS

const (
	ErrNotFound         = DictionaryErr("can't find the word")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it doesn't exist")
)

TYPES

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error

func (d Dictionary) Delete(word string)

func (d Dictionary) Search(word string) (string, error)

func (d Dictionary) Update(word, newDefinition string) error

type DictionaryErr string

func (e DictionaryErr) Error() string
### github.com/LightBulbfromSpace/kld-go-learning/math - 100.0%
TYPES

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point
### github.com/LightBulbfromSpace/kld-go-learning/mocking/configurable_sleeper - 81.8%

<details>
  <summary><code>func Countdown(out io.Writer, slp Sleeper)</code></summary>

    TYPES
    type ConfigurableSleeper struct {
    // Has unexported fields.
    }
</details>

<details>
  <summary><code>func (c *ConfigurableSleeper) Sleep()</code></summary>

    type Sleeper interface {
    Sleep()
    }
    type SpyCountdownOperations struct {
    Calls []string
    }
</details>

<details>
  <summary><code>func (s *SpyCountdownOperations) Sleep()</code></summary>

</details>

<details>
  <summary><code>func (s *SpyCountdownOperations) Write(p []byte) (n int, err error)</code></summary>

    type SpyTime struct {
    // Has unexported fields.
    }
</details>

<details>
  <summary><code>func (st *SpyTime) Sleep(duration time.Duration)</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/mocking/countdown_v1 - 62.5%

<details>
  <summary><code>func Countdown(out io.Writer, slp Sleeper)</code></summary>

    TYPES
    type RealSleeper struct{}
</details>

<details>
  <summary><code>func (rs *RealSleeper) Sleep()</code></summary>

    type Sleeper interface {
    Sleep()
    }
    type SpySleeper struct {
    Calls int
    }
</details>

<details>
  <summary><code>func (s *SpySleeper) Sleep()</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/pointers_and_errors - 83.3%
VARIABLES

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
var ErrNegativeValue = errors.New("amount of ETH must be positive")

TYPES

type Ethereum float64

func (e Ethereum) String() string

type Stringer interface {
	String() string
}

type Wallet struct {
	// Has unexported fields.
}

func (w *Wallet) Balance() Ethereum

func (w *Wallet) Deposit(sum Ethereum) error

func (w *Wallet) Withdraw(sum Ethereum) error
### github.com/LightBulbfromSpace/kld-go-learning/readme - 94.3%

<details>
  <summary><code>func MakeContent(packages []Package) string</code></summary>

    TYPES
    type Function struct {
    Interface string
    DocLines  []string
    }
    type Package struct {
    Package      string
    TestCoverage string     // TestCoverage from go test -cover
    Description  string     // Description from go doc comment above package
    Functions    []Function // Functions and their docs above
    }
    Package represents exercises directory.
</details>

<details>
  <summary><code>func GetPackages(commandRunner func(cmd string, args ...string) string) (result []Package)</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/reflection/v1 - 100.0%

### github.com/LightBulbfromSpace/kld-go-learning/reflection/v2 - 100.0%

### github.com/LightBulbfromSpace/kld-go-learning/reflection/v3 - 100.0%

### github.com/LightBulbfromSpace/kld-go-learning/select/concurrency_check - 90.0%

<details>
  <summary><code>func ConfigurableWebsiteRacer(a, b string, timeout time.Duration) (winner string, err error)</code></summary>

</details>

<details>
  <summary><code>func WebsiteRacer(a, b string) (winner string, err error)</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/select/server_mocking - 87.5%

<details>
  <summary><code>func WebsiteRacer(a, b string) string</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/select/test_external_sites - 88.9%

<details>
  <summary><code>func WebsiteRacer(a, b string) (winner string)</code></summary>

</details>

### github.com/LightBulbfromSpace/kld-go-learning/structs_methods_interfaces - 100.0%
TYPES

type Circle struct {
	// Has unexported fields.
}

func (c Circle) Area() float64

func (c Circle) Perimeter() float64

type Rectangle struct {
	// Has unexported fields.
}

func (rect Rectangle) Area() float64

func (rect Rectangle) Perimeter() float64

type RightTriangle struct {
	// Has unexported fields.
}

func (t RightTriangle) Area() float64

func (t RightTriangle) Perimeter() float64

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Square struct {
	// Has unexported fields.
}

func (s Square) Area() float64

func (s Square) Perimeter() float64
### github.com/LightBulbfromSpace/kld-go-learning/sync - 100.0%
TYPES

type Counter struct {
	// Has unexported fields.
}

func NewCounter() *Counter

func (c *Counter) Inc()

func (c *Counter) Value() int