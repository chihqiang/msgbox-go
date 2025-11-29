import { Page } from "@/model/base";
import { QueryRequest, RecordItem } from "@/model/record";
import { ApiResponse, get } from "@/utils/request";

export async function listRecords(query: QueryRequest): Promise<ApiResponse<Page<RecordItem>>> {
  return await get<Page<RecordItem>>('/record', {...query})
}
