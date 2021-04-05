package examples

import "context"

/*
 * 支持自定义请求路径，每一个自定义路径都将进行匹配。
 */
// @HttpGet("/user")
// @HttpGet("/user/more")
// @HttpPost("/user")
// @HttpPost("/user/info")
func User5() {}

/*
 * 路径模糊匹配也是支持的。被匹配上的参数，将存放在 context 中等待被获取。
 * 该方法不会匹配 /user/
 *
 * 需要注意的是：如果有确定的路径被注册，将优先使用确定的路径。
 * 例如：还有一个方法 UserAdmin，路径 `/user/admin` 被注册。
 * 1. 请求 `/user/admin` 将请求到 UserAdmin 方法。
 * 2. 请求 `/user/zhangsan` 将请求到 UserBlurry 方法。
 */
// @HttpPost("/user/{id}")
func UserBlurry(ctx context.Context) {
	// id := ctx.Value("id") // id 是 string 类型
}

/*
 * 路径模糊匹配不仅支持将参数存放在 context 中，还支持匹配方法中的参数并自动设置。
 * 注意：参数自动匹配后，不再将参数放在 context 中
 */
// @HttpPost("/user/{id}")
func UserBlurry1(ctx context.Context, id string) {
	// id1 := ctx.Value("id") // id1 将是 nil，无法取到值
}

/*
 * 路径模糊匹配也可以自动匹配方法中的参数类型，并自动转换。
 * 如果参数转换错误，将返回错误，不再请求方法。
 */
// @HttpPost("/user/{id}")
func UserBlurry2(ctx context.Context, id int64) {}

/*
 * 程序也支持自动按文件路径进行路由注册。
 * 扫描地址下所有非特定的公开方法都将被注册进路由，方法将改为小写，并以下划线分割。
 * @NotRouter 将避开路由注册扫描。
 *
 * 例如项目结构如下：
 * | ui/
 * |   servelet/
 * |          v1/
 * |            user/
 * |              create.go
 *
 * crete.go 中有一个函数：AutoRegist
 *
 * chero 扫描地址如下
 * chero.Scan("./ui/servelet")
 *
 * 则会自动注册：
 * GET     /v1/user/auto_regist
 * POST    /v1/user/auto_regist
 * PUT     /v1/user/auto_regist
 * PATCH   /v1/user/auto_regist
 * DELETE  /v1/user/auto_regist
 * OPTIONS /v1/user/auto_regist
 */
func AutoRegist() {}
