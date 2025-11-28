import { post } from '@/utils/request';
import type { ApiResponse } from '@/utils/request';
import { LoginRequest, LoginResponse, RegisterRequest } from '@/model/auth';

/**
 * 执行用户登录
 *
 * @param data 登录请求数据，包含邮箱和密码
 * @returns Promise<ApiResponse<LoginResponseData>> 登录响应数据
 * @example
 * ```typescript
 * login({ email: 'user@example.com', password: 'password123' })
 *   .then(response => {
 *     console.log('登录成功，用户ID:', response.data.id);
 *     console.log('登录成功，用户名:', response.data.name);
 *     console.log('登录成功，令牌:', response.data.token);
 *     console.log('登录成功，过期时间:', response.data.expires_in);
 *   })
 *   .catch(error => {
 *     console.error('登录失败:', error);
 *   });
 * ```
 */
export async function login(data: LoginRequest): Promise<ApiResponse<LoginResponse>> {
  return await post<LoginResponse>('/login', data);
}

/**
 * 执行用户注册
 *
 * @param data 注册请求数据，包含邮箱、密码和手机号
 * @returns Promise<ApiResponse<null>> 注册响应数据（无返回数据）
 */
export async function register(data: RegisterRequest): Promise<ApiResponse<null>> {
  return await post<null>('/register', data);
}
