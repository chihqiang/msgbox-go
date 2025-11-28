/**
 * Cookie 操作工具函数 - 基于 js-cookie 的二次封装
 * 提供类型安全的 Cookie 读写操作
 */

import Cookies from 'js-cookie';

// 常用的 Cookie 名称常量
const COOKIE_KEYS = {
  TOKEN: 'token',
} as const;

/**
 * 自定义 Cookie 选项接口，扩展 js-cookie 的选项
 */
export interface CookieOptions extends Cookies.CookieAttributes {
  /**
   * 是否自动进行编码/解码
   * 默认值：true
   */
  encode?: boolean;
}

/**
 * 设置 Cookie 值
 * @param name Cookie 名称
 * @param value Cookie 值
 * @param options Cookie 选项
 * @returns 设置成功返回 true，失败返回 false
 */
export function set(name: string, value: string, options: CookieOptions = {}): boolean {
  try {
    const { encode = true, ...cookieOptions } = options;
    
    // 如果启用了编码，处理特殊字符
    const finalValue = encode ? encodeURIComponent(value) : value;
    
    // 设置默认路径为根路径
    const defaultOptions: Cookies.CookieAttributes = {
      path: '/',
      ...cookieOptions,
    };
    
    // 处理 SameSite 为 'none' 的情况，需要同时设置 secure
    if (defaultOptions.sameSite === 'none' && !defaultOptions.secure) {
      defaultOptions.secure = true;
    }
    
    Cookies.set(name, finalValue, defaultOptions);
    return true;
  } catch (error) {
    console.error('Failed to set cookie:', error);
    return false;
  }
}

/**
 * 获取 Cookie 值
 * @param name Cookie 名称
 * @param options Cookie 选项
 * @returns Cookie 值，如果不存在或发生错误则返回 null
 */
export function get(name: string, options: { encode?: boolean } = {}): string | null {
  try {
    const { encode = true } = options;
    const value = Cookies.get(name);
    
    if (value === undefined) {
      return null;
    }
    
    return encode ? decodeURIComponent(value) : value;
  } catch (error) {
    console.error('Failed to get cookie:', error);
    return null;
  }
}

/**
 * 删除 Cookie
 * @param name Cookie 名称
 * @param options Cookie 选项，需要与设置时的路径和域保持一致
 * @returns 删除成功返回 true，失败返回 false
 */
export function remove(name: string, options: Cookies.CookieAttributes = {}): boolean {
  try {
    // 设置默认路径为根路径
    const defaultOptions: Cookies.CookieAttributes = {
      path: '/',
      ...options,
    };
    
    Cookies.remove(name, defaultOptions);
    // 验证是否真的删除成功
    return get(name) === null;
  } catch (error) {
    console.error('Failed to remove cookie:', error);
    return false;
  }
}

/**
 * 存储对象类型的 Cookie（自动序列化为 JSON）
 * @param name Cookie 名称
 * @param value 要存储的对象
 * @param options Cookie 选项
 * @returns 设置成功返回 true，失败返回 false
 */
export function setJson<T>(name: string, value: T, options: CookieOptions = {}): boolean {
  try {
    const jsonString = JSON.stringify(value);
    return set(name, jsonString, options);
  } catch (error) {
    console.error('Failed to set JSON cookie:', error);
    return false;
  }
}

/**
 * 获取对象类型的 Cookie（自动反序列化为对象）
 * @param name Cookie 名称
 * @returns 解析后的对象，如果不存在或解析失败则返回 null
 */
export function getJson<T>(name: string): T | null {
  try {
    const cookieValue = get(name);
    if (cookieValue === null) {
      return null;
    }
    
    return JSON.parse(cookieValue) as T;
  } catch (error) {
    console.error('Failed to parse JSON cookie:', error);
    return null;
  }
}

/**
 * 设置 Token Cookie
 * @param token 令牌值
 * @param expiresIn 过期时间（秒）
 * @returns 设置成功返回 true，失败返回 false
 */
export function setToken(token: string, expiresIn?: number): boolean {
  const options: CookieOptions = {
    secure: import.meta.env.PROD, // 生产环境使用 secure
    sameSite: 'lax',
  };
  
  if (expiresIn) {
    options.expires = expiresIn
  }
  
  return set(COOKIE_KEYS.TOKEN, token, options);
}

/**
 * 获取 Token Cookie
 * @returns Token 值，如果不存在则返回 null
 */
export function getToken(): string | null {
  return get(COOKIE_KEYS.TOKEN);
}

export function removeToken(): boolean {
  return remove(COOKIE_KEYS.TOKEN);
}
