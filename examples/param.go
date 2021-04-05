package examples

import (
	"context"
	"io"

	"github.com/SmallTianTian/chero/examples/model"
)

/*
 * http query params 默认将放置在 context 中
 *
 * 例如：
 * 	请求 `/query?id=123`，在 context 获取 id，将得到字符串 123
 */
// @HttpGet("/query")
func Query(ctx context.Context) {
	// id := ctx.Value("id") // id == "123"
}

/*
 * http query params 也可以直接赋值到方法接收的参数中，这时候 context 中无内容。
 */
// @HttpGet("/query")
func Query1(ctx context.Context, id string) {
	// id := ctx.Value("id") // id == nil
}

/*
 * http query params 赋值的同时，也可以按接收类型进行转换。
 */
// @HttpGet("/query")
func Query2(ctx context.Context, id int64) {}

/*
 * http query params 的来源可以有多种，优先级从高到低依次是：
 * path > body > query > header > cookie
 *
 * 例如：
 * ```http
 * POST /user/1?id=3&name=lisi&age=15
 * Content-Type: application/json
 * Id: 4
 * Name: wangwu
 * Age: 100
 * Nick: xiaoming
 *
 * {"id": 2, "name": "zhangsan"}
 * ```
 *
 * 参数获取结果如下：
 * id: 1
 * name: zhangsan
 * age: 15
 * nick: xiaoming
 */
// @HttpPost("/user/{id}")
func MultipartParam(ctx context.Context) {
	// id := ctx.Value("id")     // id == 1
	// name := ctx.Value("name") // name == zhangsan
	// age := ctx.Value("age")   // age == 15
	// nick := ctx.Value("nick") // nick == xiaoming
}

/*
 * http query params 除了正常的优先级之外，你也可以强制指定使用来源。
 * 相同 param 被多次强制指定，以最后指定为准。
 * 有如下几种强制指定
 * @PathParam
 * @BodyParam
 * @QueryParam
 * @HeaderParam
 * @CookieParam
 *
 * 例如：
 * ```http
 * POST /user/1?id=3&name=lisi
 * Content-Type: application/json
 * Id: 4
 * Name: wangwu
 * ```
 *
 * 参数获取结果如下：
 * id: 4
 * name: wangwu
 */
// @HttpPost("/user/{id}")
// @PathParam("id")
// @HeaderParam("id")
// @HeaderParam("name")
func MultipartParam1(ctx context.Context) {
	// id := ctx.Value("id")     // id == 4
	// name := ctx.Value("name") // name == wangwu
}

/*
 * http query params 也会自动依据 Content-Type 进行解析，然后注册到 context 中。
 *
 * 例如:
 * ```http
 * POST /user
 * Content-Type: application/xml
 * Id: 4
 * Name: wangwu
 * Age: 100
 * Nick: xiaoming
 *
 * <xml>
 *  <id>2</id>
 * 	<name>zhangsan</name>
 * </xml>
 * ```
 */
// @HttpPost("/user")
func MultipartBody(ctx context.Context) {
	// id := ctx.Value("id")     // id == 1
	// name := ctx.Value("name") // name == zhangsan
}

/*
 * http query params 一样也可以自动按类型绑定到方法的入参中，将不会再出现在 context 中。
 *
 * 例如:
 * ```http
 * POST /user
 * Content-Type: application/xml
 * Id: 4
 * Name: wangwu
 * Age: 100
 * Nick: xiaoming
 *
 * <xml>
 *  <id>2</id>
 * 	<name>zhangsan</name>
 * </xml>
 * ```
 */
// @HttpPost("/user")
func MultipartBody1(ctx context.Context, id int64) {
	// id := ctx.Value("id")     // id == nil
	// name := ctx.Value("name") // name == zhangsan
}

/*
 * 表单上传文件中，文件将是 `io.ReadSeekCloser` 类型。
 */
func FormUploadFile(ctx context.Context) {
	// fileName := ctx.Value("file_name") // fileName == "README.md"
	// file := ctx.Value("file") // file is a `io.ReadSeekCloser`
}

/*
 * 表单上传文件中，文件绑定入参的时候，不仅能是 `io.ReadSeekCloser` 类型，
 * 还能是 `[]byte` 类型，其他类型不予支持。
 */
func FormUploadFile1(ctx context.Context, file1 io.ReadSeekCloser, file2 []byte) {}

/*
 * 函数入参除了支持基础类型外，也支持 map 和用户自定义 strcut。
 * 注意：
 * 1. map 的值必须是 interface 类型。
 * 2. 绑定顺序：基础类型 = 用户自定义 struct > map > context
 *         即：基础类型的变量名和用户自定义 struct 重复时，两个都将被赋值，
 *            再考虑塞入值为 interface 类型的 map 中，最后考虑 context。
 *
 * ```http
 * POST /user/1?id=3&name=lisi&age=15
 * Content-Type: application/json
 * Id: 4
 * Name: wangwu
 * Age: 100
 * Nick: xiaoming
 *
 * {"id": 2, "name": "zhangsan"}
 * ```
 *
 * 参数获取结果如下：
 * id: 1
 * u.ID: 1
 * u.Name: zhangsan
 * m: {"age": "15", "nick": "xiaoming"}
 * ctx: nil
 */
// @HttpPost(/user/{id})
func UserStruct(ctx context.Context, id int64, u model.User, m map[string]interface{}) {
	// id == u.ID // id 和 u.ID 值一样，都是 1
	// u.Name == "zhangsan"
	// m["age"] == "15" // 15 为字符串类型
	// m["nick"] == "xiaoming"
	// m["id"] == nil
	// ctx["id"] == nil
}
