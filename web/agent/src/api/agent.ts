import { get, post, type ApiResponse } from "@/utils/request"
import { AgentInfo, resetAgentSecret } from "@/model/agent"

export async function getAgentInfo(): Promise<ApiResponse<AgentInfo>> {
  return await get<AgentInfo>('/info')
}

export async function resetSecret(): Promise<ApiResponse<resetAgentSecret>> {
  return await post<resetAgentSecret>('/reset/agent/secret')
}
