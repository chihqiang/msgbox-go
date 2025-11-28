/**
 * 生成指定长度的随机字符串
 * @param length 字符串长度
 */
export function randomString(length: number,type: string = 'alnum'): string {
  const presets: Record<string, string> = {
    number: '0123456789',
    lower: 'abcdefghijklmnopqrstuvwxyz',
    upper: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ',
    alpha: 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ',
    alnum: 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
  };

  const charset = presets[type] || type; // 若不是预设则当作自定义字符集
  let result = '';

  for (let i = 0; i < length; i++) {
    const idx = Math.floor(Math.random() * charset.length);
    result += charset[idx];
  }

  return result;
}
