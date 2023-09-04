
## Wire: Automated Initialization in Go

https://github.com/google/wire

Wire is a code generation tool that automates connecting components using dependency injection. Dependencies between components are represented in Wire as function parameters, encouraging explicit initialization instead of global variables. Because Wire operates without runtime state or reflection, code written to be used with Wire is useful even for hand-written initialization.

Wire 是一种代码生成工具，它使用依赖注入自动连接组件。 组件之间的依赖关系在 Wire 中表示为函数参数，鼓励显式初始化而不是全局变量。 由于 Wire 在没有运行时状态或反射的情况下运行，因此编写用于 Wire 的代码即使对于手写初始化也很有用。

https://go.dev/blog/wire

## Installing

```bash
go install github.com/google/wire/cmd/wire@latest
```

依赖关系注入是一种标准技术，用于生成灵活且松散耦合的代码，方法是显式地为组件提供它们工作所需的所有依赖项。

在 Go 中，这通常采用将依赖项传递给构造函数的形式：

```go
// NewUserStore returns a UserStore that uses cfg and db as dependencies.
func NewUserStore(cfg *Config, db *mysql.DB) (*UserStore, error) {...}
```

```go
// NewUserStore is the same function we saw above; it is a provider for UserStore,
// with dependencies on *Config and *mysql.DB.
func NewUserStore(cfg *Config, db *mysql.DB) (*UserStore, error) {...}

// NewDefaultConfig is a provider for *Config, with no dependencies.
func NewDefaultConfig() *Config {...}

// NewDB is a provider for *mysql.DB based on some connection info.
func NewDB(info *ConnectionInfo) (*mysql.DB, error) {...}
```
