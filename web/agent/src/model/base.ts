import { ChannelItem } from "./channel"

export interface Page<T> {
  map(arg0: (item: ChannelItem) => { label: string; value: string }): () => ArrayIterator<{ label: string; value: string }>
  total: number
  data: T[]
}

export interface PageRequest {
  page: number
  size : number
}

export interface SelectOption {
  label: string
  value: number | string
}
