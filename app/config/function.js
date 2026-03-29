/*
 * 源启云联
 *
 * @author    twj
 * @email     120722324@qq.com
 * @company   成都源启科技有限公司
 * @copyright © 2026 成都源启科技有限公司
 * @license   Apache-2.0
 */
// 生成随机字符串
export function generateRandomString(length = 6) {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
  let result = ''
  for (let i = 0; i < length; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  return result
}
//对象转换为url查询参数
export function objToQuery(obj) {
  if (!obj) return ''
  return Object.keys(obj)
    .map(key => {
      const value = obj[key]
      if (Array.isArray(value)) {
        return value
          .map(v => `${encodeURIComponent(key)}=${encodeURIComponent(v)}`)
          .join('&')
      }
      return `${encodeURIComponent(key)}=${encodeURIComponent(value)}`
    })
    .join('&')
}
// 秒转换时间
export function formatTime(seconds) {
  seconds = Number(seconds); // 确保是数字
  
  const h = Math.floor(seconds / 3600);
  const m = Math.floor((seconds % 3600) / 60);
  const s = seconds % 60;
  
  // 如果需要补零
  const pad = (num) => num.toString().padStart(2, '0');
  
  if (h > 0) {
    return `${pad(h)}:${pad(m)}:${pad(s)}`;
  } else {
    return `${pad(m)}:${pad(s)}`;
  }
}
/**
 * 格式化ISO时间为本地时间（或指定时区时间）
 * @param {string} isoTime - ISO格式时间字符串（如：2026-01-20T17:52:27+08:00）
 * @param {string} [format='YYYY-MM-DD HH:MM:SS'] - 输出格式，支持：
 *   YYYY-年份，MM-月份(2位)，DD-日期(2位)，
 *   HH-小时(24小时制,2位)，mm-分钟(2位)，ss-秒(2位)
 * @param {string} [timeZone] - 指定时区（如'Asia/Shanghai'，默认使用本地时区）
 * @returns {string} 格式化后的时间字符串
 */
export function formatIsoTime(isoTime, format = 'YYYY-MM-DD HH:mm:ss', timeZone) {
  const date = new Date(isoTime);
  // 若解析失败，返回原始值或提示
  if (isNaN(date.getTime())) {
    console.warn('无效的时间格式:', isoTime);
    return isoTime;
  }

  // 处理时区：若指定时区，使用Intl获取对应时区的时间字段
  let year, month, day, hours, minutes, seconds;
  if (timeZone) {
    const options = {
      timeZone,
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit',
      hour12: false
    };
    const parts = new Intl.DateTimeFormat('zh-CN', options).formatToParts(date);
    const map = {};
    parts.forEach(part => { map[part.type] = part.value; });
    year = map.year;
    month = map.month;
    day = map.day;
    hours = map.hour;
    minutes = map.minute;
    seconds = map.second;
  } else {
    // 本地时区：直接获取本地时间字段
    year = date.getFullYear();
    month = String(date.getMonth() + 1).padStart(2, '0');
    day = String(date.getDate()).padStart(2, '0');
    hours = String(date.getHours()).padStart(2, '0');
    minutes = String(date.getMinutes()).padStart(2, '0');
    seconds = String(date.getSeconds()).padStart(2, '0');
  }

  // 替换格式字符串
  return format
    .replace('YYYY', year)
    .replace('MM', month)
    .replace('DD', day)
    .replace('HH', hours)
    .replace('mm', minutes)
    .replace('ss', seconds);
}