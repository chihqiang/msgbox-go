import { PageRequest } from "@/model/base";

export interface QueryRequest extends PageRequest {
  keywords?: string
}

export interface RecordItem {
  id: number;
  receiver: string;
  channel_name: string;
  channel_config: Record<string, undefined>;
  vendor_name: string;
  vendor_code: string;
  signature: string;
  title: string;
  content: string;
  variables?: Record<string, undefined>;
  extra?: Record<string, undefined>;
  status: number;
  status_msg?: string;
  send_time: string;
  error?: string;
  response?: string;
  delivery_time?: string;
  delivery_raw?: string;
  created_at?: string;
  updated_at?: string;
}

