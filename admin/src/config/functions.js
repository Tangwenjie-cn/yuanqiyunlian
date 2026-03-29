/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
// 导出一个函数，用于生成指定长度的随机字符串
export function randomStr(e){
	// 如果没有传入参数，则默认生成32位随机字符串
	e = e || 32;
	// 定义一个字符串，包含所有可能的字符
	let t = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678",
	// 获取字符串的长度
	a = t.length,
	// 定义一个空字符串，用于存储生成的随机字符串
	n = "";
	// 循环e次，每次从字符串t中随机取一个字符，拼接到n中
	for (let i = 0; i < e; i++) n += t.charAt(Math.floor(Math.random() * a));
	// 返回生成的随机字符串
	return n
}
// 将RFC时间格式转换为时间格式
export function rfcToTime(RFC){
	// 定义一个函数，用于将数字转换为两位数
	let pad = (n) => n.toString().padStart(2, '0');
    // 将RFC时间格式转换为Date对象
    let date= new Date(RFC);
    // 返回转换后的时间格式
    return date.getFullYear() + "-" + pad(date.getMonth() + 1) + "-" + pad(date.getDate()) + " " + pad(date.getHours()) + ":" + pad(date.getMinutes()) + ":" + pad(date.getSeconds());
}