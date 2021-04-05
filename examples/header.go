package examples

/*
 * 头匹配，只有当包含特定header 的 kv 时，才会进入这个方法，否则考虑其他方法。
 * 注意：头匹配中，大小写不区分
 *
 * 示例：
* ```http
 * POST /header
 * Nick: xiaoming
 *
 * {"id": 2, "name": "zhangsan"}
 * ```
 * 这个方法由于 header 中 content-type 不等于 application/json，将不会调用 HeaderEqual 方法。
*/
// @HeaderEqual("Content-Type=application/json")
// @HttpPost("/header")
func HeaderEqual() {}
