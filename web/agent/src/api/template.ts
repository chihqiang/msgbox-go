import { Page } from "@/model/base"
import { QueryRequest,  TemplateItem } from "@/model/template"
import { ApiResponse, post,get } from "@/utils/request"


// 保持向后兼容的API函数
export async function listTemplates(query: QueryRequest): Promise<ApiResponse<Page<TemplateItem>>> {
  return await get<Page<TemplateItem>>('/template', {...query})
}

export async function createTemplate(template: TemplateItem): Promise<ApiResponse<null>> {
  return await post<null>('/template/create', template)
}

export async function updateTemplate(template: TemplateItem): Promise<ApiResponse<null>> {
  return await post<null>('/template/update', template)
}

export async function deleteTemplate(id: number): Promise<ApiResponse<null>> {
  return await post<null>(`/template/delete`,{"id":id})
}
