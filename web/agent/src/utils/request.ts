/**
 * HTTP请求封装模块
 *
 * 提供基于axios的HTTP请求功能封装，包含：
 * - 统一的请求/响应拦截器
 * - 标准的API响应数据结构
 * - 统一的错误处理机制
 * - 常用HTTP方法的类型安全封装
 */
import axios from 'axios';
import type { AxiosInstance, AxiosRequestConfig, AxiosError, InternalAxiosRequestConfig } from 'axios';
import { getToken } from '@/utils/cookie';
import { message } from 'ant-design-vue';

/**
 * API响应数据接口定义
 *
 * @template T 响应数据的类型，默认为unknown
 * @property code 响应状态码，200表示成功
 * @property data 响应数据内容
 * @property msg 响应消息描述（后端使用msg字段）
 */
export interface ApiResponse<T = unknown> {
  code: number;
  data: T;
  msg: string;
}

/**
 * 创建axios实例
 *
 * 配置基础URL、超时时间和默认请求头
 */
const service: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL || '/api', // 基础URL，优先从环境变量获取
  timeout: 10000, // 请求超时时间：10秒
  headers: {
    'Content-Type': 'application/json;charset=utf-8' // 默认Content-Type
  }
});

/**
 * 请求拦截器
 *
 * 在发送请求前进行统一处理，如添加认证信息、请求头等
 */
service.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // 添加认证token
    const token = getToken();
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error: AxiosError) => {
    console.error('请求错误:', error);
    return Promise.reject(error);
  }
);

/**
 * 响应拦截器
 *
 * 统一处理响应数据和错误
 */
service.interceptors.response.use(
    (response) => {
    // 对响应数据进行处理
    const res = response.data;
    // 检查业务状态码
    // 后端接口规范：{ code: number, data: any, msg: string }
    if (res.code === 0) {
       return res;
    }
    message.error(res.msg || 'Error');
    return Promise.reject(new Error(res.msg || 'Error'))
  },
  (error: AxiosError) => {
    // 处理HTTP响应错误
    let errorMsg = '网络请求失败';
    if (error.response) {
      const { status } = error.response;
      switch (status) {
        case 401:
          errorMsg = '未授权，请重新登录';
          break;
        case 403:
          errorMsg = '拒绝访问';
          break;
        case 404:
          errorMsg = '请求的资源不存在';
          break;
        case 500:
          errorMsg = '服务器内部错误';
          break;
        default:
          errorMsg = `请求失败: ${error.response.statusText || status}`;
      }
    } else if (error.request) {
      // 请求已发送但未收到响应
      errorMsg = '网络连接失败，请检查网络';
    }
    message.error(errorMsg);
    return Promise.reject(new Error(errorMsg));
  }
);

/**
 * 发送GET请求
 *
 * @template T 响应数据类型
 * @param url 请求URL
 * @param params 查询参数
 * @param config 可选的请求配置
 * @returns Promise<ApiResponse<T>> 响应数据
 */
export function get<T = unknown>(url: string, params?: Record<string, unknown>, config?: AxiosRequestConfig<unknown>): Promise<ApiResponse<T>> {
  return service.get(url, { params, ...config });
}

/**
 * 发送POST请求
 *
 * @template T 响应数据类型
 * @param url 请求URL
 * @param data 请求数据
 * @param config 可选的请求配置
 * @returns Promise<ApiResponse<T>> 响应数据
 */
export function post<T = unknown>(url: string, data?: unknown, config?: AxiosRequestConfig<unknown>): Promise<ApiResponse<T>> {
  return service.post(url, data, config);
}

/**
 * 发送PUT请求
 *
 * @template T 响应数据类型
 * @param url 请求URL
 * @param data 请求数据
 * @param config 可选的请求配置
 * @returns Promise<ApiResponse<T>> 响应数据
 */
export function put<T = unknown>(url: string, data?: unknown, config?: AxiosRequestConfig<unknown>): Promise<ApiResponse<T>> {
  return service.put(url, data, config);
}

/**
 * 发送DELETE请求
 *
 * @template T 响应数据类型
 * @param url 请求URL
 * @param params 查询参数
 * @param config 可选的请求配置
 * @returns Promise<ApiResponse<T>> 响应数据
 */
export function del<T = unknown>(url: string, params?: Record<string, unknown>, config?: AxiosRequestConfig<unknown>): Promise<ApiResponse<T>> {
  return service.delete(url, { params, ...config });
}

/**
 * 发送PATCH请求
 *
 * @template T 响应数据类型
 * @param url 请求URL
 * @param data 请求数据
 * @param config 可选的请求配置
 * @returns Promise<ApiResponse<T>> 响应数据
 */
export function patch<T = unknown>(url: string, data?: unknown, config?: AxiosRequestConfig<unknown>): Promise<ApiResponse<T>> {
  return service.patch(url, data, config);
}

/**
 * 上传文件请求
 *
 * @template T 响应数据类型
 * @param url 请求URL
 * @param formData 文件表单数据
 * @param config 可选的请求配置
 * @returns Promise<ApiResponse<T>> 响应数据
 */
export function upload<T = unknown>(url: string, formData: FormData, config?: AxiosRequestConfig<unknown>): Promise<ApiResponse<T>> {
  return service.post(url, formData, {
    headers: {
      'Content-Type': 'multipart/form-data' // 设置为文件上传类型
    },
    ...config
  });
}
