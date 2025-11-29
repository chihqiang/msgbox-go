import { ApiResponse, get, post } from "@/utils/request"
import { ChannelItem, Configs, QueryRequest } from "@/model/channel"
import { Page } from "@/model/base"

export async function getChannelConfigs(): Promise<ApiResponse<Configs[]>> {
  return await get<Configs[]>('/channel/configs')
}


export async function listChannels(query: QueryRequest): Promise<ApiResponse<Page<ChannelItem>>> {
  return await get<Page<ChannelItem>>('/channel', {...query})
}

export async function createChannel(channel: ChannelItem): Promise<ApiResponse<null>> {
  return await post<null>('/channel/create', channel)
}

export async function updateChannel(channel: ChannelItem): Promise<ApiResponse<null>> {
  return await post<null>('/channel/update', channel)
}

export async function deleteChannel(id: number): Promise<ApiResponse<null>> {
  return await post<null>(`/channel/delete`,{"id":id})
}
