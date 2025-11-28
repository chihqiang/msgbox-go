import { PageRequest } from "@/model/base";

export interface QueryRequest extends PageRequest {
  keywords?: string
  status?: string
}

// 保持向后兼容
export interface TemplateItem {
  id?: number
  name: string
  channel_id: number | null // 修改为支持null值，以便在创建新模板时显示placeholder
  code: string
  vendor_code: string
  signature: string
  content: string
  status: boolean
  used_count: number
  created_at: string
  updated_at: string
}
