import { PageRequest } from "@/model/base";

export interface QueryRequest extends PageRequest {
  keywords?: string
}


export interface ConfigItem {
  type: string
  name: string
  label: string
  required: boolean
  placeholder: string
  default?: string
}

export interface Configs {
  name: string
  label: string
  configs: ConfigItem[]
}


export interface ChannelItem {
  id?: number
  code: string
  name: string
  vendor_name?: string
  config: Record<string, unknown>
  status: boolean
  createdAt: string
  updatedAt: string
}
