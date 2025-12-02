export interface LoginRequest {
  /** 用户邮箱地址 */
  email: string;
  /** 用户密码 */
  password: string;
}

export interface LoginResponse {
  /** 用户ID */
  id: number;
  /** 用户名 */
  name: string;
  /** 认证令牌 */
  token: string;
  /** 令牌过期时间（秒） */
  expires_in: number;
}


export interface RegisterRequest {
  email: string;
  password: string;
  phone: string;
  code: string;
}

