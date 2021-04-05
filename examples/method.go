package examples

// 以下注释中，多行注释表示对方法的描述，单行注释表示项目的使用方式

/*
 * 以 Get 开头的方法将默认只支持 get 请求
 */
func GetUser(id int64) {}

/*
 * 使用 `@HttpGet` 表示该方法支持 get 请求
 */
// @HttpGet
func User1() {}

/*
 * 使用 `@HttpGet` 与使用 `@HttpPost` 等并不冲突，这样表示该方法支持 get 和 post，
 * 但其他方式（例如 delete 等）不予支持。
 */
// @HttpGet
// @HttpPost
func User2() {}

/*
 * 不使用任何 HTTP method 说明，代表将允许默认的方法。
 * 可以在 chero.SetDefaultMethod 方法中设置，
 * 如果不设置，默认允许所有的请求方式。
 */
func User3() {}

/*
 * 允许多种请求方式，支持 get、post、put、patch、delete、options，可自由组合。
 */
// @HttpGet
// @HttpPost
// @HttpPut
// @HttpPatch
// @HttpDelete
// @HttpOptions
func User4() {}
